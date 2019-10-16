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

var regoResourceTypeRegex = regexp.MustCompile(`([\w]+)[.]([\w]+)[.]([\w]+)`)

func pathToRuleName(path string) string {
	baseName := filepath.Base(path)
	extension := strings.ToLower(filepath.Ext(path))
	if extension != ".rego" {
		return ""
	}
	return baseName[:len(baseName)-len(extension)]
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

	rego := regoFile{
		Text: string(contents),
		Name: baseName[:len(baseName)-len(extension)],
	}

	for _, line := range strings.Split(rego.Text, "\n") {
		// We will extract the rule description and resource type from
		// comment lines in the rego file. Ignore non-comment lines here.
		line = strings.TrimSpace(line)
		if len(line) == 0 || !strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(strings.Trim(line, "#"))
		// Look for a resource type in the form of "Provider.Service.Type"
		if rego.ResourceType == "" {
			match := regoResourceTypeRegex.FindStringSubmatch(line)
			if len(match) == 4 {
				rego.ResourceType = strings.Join(match[1:4], ".")
				rego.Provider = strings.ToUpper(match[1])
				continue
			}
			lineUpper := strings.ToUpper(line)
			if lineUpper == "AWS" || lineUpper == "AWS_GOVCLOUD" || lineUpper == "AZURE" {
				rego.Provider = lineUpper
			}
		}
		// Comment lines are otherwise considered part of the rule description
		if rego.Description == "" {
			rego.Description = line
		}
	}
	if rego.ResourceType == "" {
		rego.ResourceType = "MULTIPLE"
	}
	if rego.Description == "" {
		return nil, errors.New("Expected a rego comment to use as the " +
			"rule description.")
	}
	return &rego, nil
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

				ruleName := pathToRuleName(path)
				if ruleName == "" {
					return nil
				}

				rego, err := loadRego(path)
				if err != nil {
					fmt.Println("WARN:", err)
					return nil
				}
				if rego == nil {
					return nil
				}

				existingRule := getRuleByName(ruleName)
				if existingRule == nil {
					fmt.Println("Creating rule", ruleName)
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
					fmt.Println("Updating rule", ruleName)
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