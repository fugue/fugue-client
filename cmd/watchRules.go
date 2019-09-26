package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fsnotify/fsnotify"
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
		return nil, fmt.Errorf("Unexpected file type: %s", extension)
	}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
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
		}
		// Comment lines are otherwise considered part of the rule description
		if rego.Description == "" {
			rego.Description = line
		} else {
			rego.Description += "\n" + line
		}
	}
	if rego.ResourceType == "" {
		return nil, errors.New("No resource type detected in rego comments. " +
			"Expected type like AWS.EC2.SecurityGroup.")
	}
	if rego.Description == "" {
		return nil, errors.New("Expected a rego comment to use as the " +
			"rule description.")
	}
	fmt.Printf("REGO: %+v\n", rego)
	return &rego, nil
}

// NewWatchRulesCommand returns a command that watches a directory for changes
// to rego files
func NewWatchRulesCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "rules [directory]",
		Short: "Update rule settings",
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
				CheckErr(err)

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

			watcher, err := fsnotify.NewWatcher()
			CheckErr(err)
			defer watcher.Close()

			err = watcher.Add(args[0])
			CheckErr(err)

			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						updateRule(event.Name)
					} else if event.Op&fsnotify.Create == fsnotify.Create {
						updateRule(event.Name)
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					CheckErr(err)
				}
			}
		},
	}
	return cmd
}

func init() {
	watchCmd.AddCommand(NewWatchRulesCommand())
}
