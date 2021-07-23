package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fugue/fugue-client/client/custom_rules"

	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type updateRuleOptions struct {
	ID           string
	Name         string
	Description  string
	Severity     string
	ResourceType string
	RuleText     string
	Families     []string
}

// NewUpdateRuleCommand returns a command that updates a custom rule
func NewUpdateRuleCommand() *cobra.Command {

	var opts updateRuleOptions

	cmd := &cobra.Command{
		Use:   "rule [rule_id]",
		Short: "Update rule settings",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := custom_rules.NewUpdateCustomRuleParams()
			params.RuleID = args[0]
			params.Rule = &models.UpdateCustomRuleInput{}

			flagCount := 0

			// Using Visit here allows us to process only flags that were set
			cmd.Flags().Visit(func(f *pflag.Flag) {
				flagCount++
				switch f.Name {
				case "name":
					params.Rule.Name = opts.Name
				case "description":
					params.Rule.Description = opts.Description
				case "severity":
					params.Rule.Severity = opts.Severity
				case "resource-type":
					params.Rule.ResourceType = opts.ResourceType
				case "text":
					params.Rule.RuleText = opts.RuleText
				case "families":
					params.Rule.Families = opts.Families
				}
			})

			if flagCount == 0 {
				os.Exit(0)
			}

			resp, err := client.CustomRules.UpdateCustomRule(params, auth)
			CheckErr(err)

			rule := resp.Payload

			items := []interface{}{
				Item{"NAME", rule.Name},
				Item{"DESCRIPTION", rule.Description},
				Item{"PROVIDER", rule.Provider},
				Item{"SEVERITY", rule.Severity},
				Item{"RESOURCE_TYPE", rule.ResourceType},
				Item{"STATUS", rule.Status},
				Item{"FAMILIES", strings.Join(rule.Families[:], ",")},
				Item{"CREATED_AT", format.Unix(rule.CreatedAt)},
				Item{"CREATED_BY", rule.CreatedBy},
				Item{"CREATED_BY_DISPLAY_NAME", rule.CreatedByDisplayName},
				Item{"UPDATED_AT", format.Unix(rule.UpdatedAt)},
				Item{"UPDATED_BY", rule.UpdatedBy},
				Item{"UPDATED_BY_DISPLAY_NAME", rule.UpdatedByDisplayName},
			}

			table, err := format.Table(format.TableOpts{
				Rows:         items,
				Columns:      []string{"Attribute", "Value"},
				ShowHeader:   true,
				MaxCellWidth: 70,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "Rule name")
	cmd.Flags().StringVar(&opts.Description, "description", "", "Description")
	cmd.Flags().StringVar(&opts.Severity, "severity", "", "Severity")
	cmd.Flags().StringVar(&opts.ResourceType, "resource-type", "", "Resource type")
	cmd.Flags().StringVar(&opts.RuleText, "text", "", "Rule text")
	cmd.Flags().StringSliceVar(&opts.Families, "families", nil, "Families")

	return cmd
}

func init() {
	updateCmd.AddCommand(NewUpdateRuleCommand())
}
