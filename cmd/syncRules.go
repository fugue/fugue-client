package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type regoFile struct {
	Name         string
	Provider     string
	ResourceType string
	Description  string
	Text         string
}

func (rego *regoFile) ParseText() error {
	// Extract the header block from the first comment block.
	headerText := ""
	inFirstCommentBlock := false
	for _, line := range strings.Split(rego.Text, "\n") {
		if strings.HasPrefix(line, "#") {
			if !inFirstCommentBlock && headerText == "" {
				inFirstCommentBlock = true
			}
			if inFirstCommentBlock {
				headerText += strings.TrimSpace(strings.TrimPrefix(line, "#"))
				headerText += "\r\n"
			}
		} else {
			inFirstCommentBlock = false
		}
	}

	// Parse the HTTP headers in `headerText`.
	reader := bufio.NewReader(strings.NewReader(headerText + "\r\n"))
	tp := textproto.NewReader(reader)
	headers := make(map[string][]string)
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		log.Fatal(err)
	} else {
		headers = http.Header(mimeHeader)
	}

	// Helper to obtain a specific header.
	getHeader := func(name string) string {
		if arr, ok := headers[name]; ok {
			if len(arr) > 0 {
				return arr[0]
			}
		}
		return ""
	}

	rego.Description = getHeader("Description")
	rego.Provider = getHeader("Provider")

	rt := getHeader("Resource-Type")
	if strings.EqualFold(rt, "MULTIPLE") {
		rego.ResourceType = rt
	} else if strings.EqualFold(rego.Provider[0:3], "AWS") {
		rego.ResourceType = "AWS." + rt
	} else {
		rego.ResourceType = rego.Provider + "." + rt
	}

	// Throw errors if things are missing.
	if rego.ResourceType == "" {
		return errors.New("expected a resource type by the header \"Resource-Type\"")
	}
	if rego.Description == "" {
		return errors.New("expected a description by the header \"Description\"")
	}
	if rego.Provider == "" {
		return errors.New("expected a provider by the header \"Provider\"")
	}

	return nil
}

func loadRego(path string) (*regoFile, error) {

	baseName := filepath.Base(path)
	extension := strings.ToLower(filepath.Ext(path))
	if extension != ".rego" {
		return nil, nil
	}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(contents) == 0 {
		return nil, nil
	}

	name := baseName[:len(baseName)-len(extension)]
	name = strings.ReplaceAll(name, "_", "-")

	rego := regoFile{
		Text: string(contents),
		Name: name,
	}

	err = rego.ParseText()

	return &rego, err
}

// NewSyncRulesCommand returns a command that watches a directory for changes
// to rego files
func NewSyncRulesCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "rules [directory]",
		Short: "Sync rules to the organization",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			getRuleByName := func(name string) *models.CustomRule {
				params := custom_rules.NewListCustomRulesParams()
				resp, err := client.CustomRules.ListCustomRules(params, auth)
				CheckErr(err)
				for _, rule := range resp.Payload.Items {
					if rule.Name == name {
						return rule
					}
				}
				return nil
			}

			updateRule := func(path string) error {

				rego, err := loadRego(path)
				if err != nil {
					fmt.Println("WARN:", err)
					return nil
				}
				if rego == nil {
					return nil
				}

				existingRule := getRuleByName(rego.Name)
				if existingRule == nil {
					fmt.Println("Creating rule", rego.Name)
					params := custom_rules.NewCreateCustomRuleParams()
					params.Rule = &models.CreateCustomRuleInput{
						Name:         rego.Name,
						Description:  rego.Description,
						ResourceType: rego.ResourceType,
						Provider:     rego.Provider,
						RuleText:     rego.Text,
					}
					client.CustomRules.CreateCustomRule(params, auth)
				} else {
					fmt.Println("Updating rule", rego.Name)
					params := custom_rules.NewUpdateCustomRuleParams()
					params.RuleID = existingRule.ID
					params.Rule = &models.UpdateCustomRuleInput{
						Description:  rego.Description,
						ResourceType: rego.ResourceType,
						RuleText:     rego.Text,
					}
					client.CustomRules.UpdateCustomRule(params, auth)
				}
				return nil
			}

			files, err := ioutil.ReadDir(args[0])
			CheckErr(err)

			for _, file := range files {
				if filepath.Ext(file.Name()) == ".rego" {
					updateRule(filepath.Join(args[0], file.Name()))
				}
			}
		},
	}
	return cmd
}

func init() {
	syncCmd.AddCommand(NewSyncRulesCommand())
}
