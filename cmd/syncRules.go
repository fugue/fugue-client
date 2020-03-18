package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
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

	for _, line := range strings.Split(rego.Text, "\n") {
		// We will extract the rule description and resource type from
		// comment lines in the rego file. Ignore non-comment lines here.
		line = strings.TrimSpace(line)
		if len(line) == 0 || !strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(strings.Trim(line, "#"))
		// Look for a resource type type denoted by the correct header
		if rego.ResourceType == "" {
			match := regoResourceTypeHeader.FindStringSubmatch(line)
			if len(match) == 6 {
				rego.Provider = match[2]
				if match[4] == "" {
					rego.ResourceType = strings.Join(match[2:4], ".")
				} else {
					rego.ResourceType = strings.ToUpper(match[3])
				}
				continue
			} else {
				return errors.New("unexpected resource type definition")
			}
		}
		// Look for a description type denoted by the correct header
		if rego.Description == "" {
			match := regoDescriptionHeader.FindStringSubmatch(line)
			if len(match) == 3 {
				rego.Description = match[2]
				continue
			} else {
				return errors.New("unexpected description definition")
			}
		}
	}
	if rego.ResourceType == "" {
		return errors.New("expected a resource type by the header \"Resource-Type\"")
	}
	if rego.Description == "" {
		return errors.New("expected a description by the header \"Description\"")
	}
	return nil
}

var regoResourceTypeHeader = regexp.MustCompile(`([rR]esource-[tT]ype\:[\t ]*?)([\w]+)[.]((MULTIPLE)|([\w]+[.][\w]+))`)
var regoDescriptionHeader = regexp.MustCompile(`([dD]escription\:[\t ]*?)(.*)`)

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
