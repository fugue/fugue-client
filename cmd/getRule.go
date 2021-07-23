package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type getRuleOptions struct {
	ShowText bool
}

// NewGetRuleCommand returns a command that retrieves custom rule details
func NewGetRuleCommand() *cobra.Command {

	var opts getRuleOptions

	cmd := &cobra.Command{
		Use:   "rule [rule_id]",
		Short: "Retrieve details for a custom rule",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := custom_rules.NewGetCustomRuleParams()
			params.RuleID = args[0]

			resp, err := client.CustomRules.GetCustomRule(params, auth)
			CheckErr(err)

			rule := resp.Payload

			if opts.ShowText {
				fmt.Println(rule.RuleText)
				return
			}

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

	cmd.Flags().BoolVar(&opts.ShowText, "text", false, "Show rule text")

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetRuleCommand())
}
