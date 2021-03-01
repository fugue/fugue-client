package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/rule_waivers"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

// NewGetRuleWaiverCommand returns a command that retrieves rule waiver details
func NewGetRuleWaiverCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "rule-waiver [rule_waiver_id]",
		Short:   "Retrieve details for a rule waiver",
		Aliases: []string{"waiver", "rule_waiver"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := rule_waivers.NewGetRuleWaiverParams()
			params.RuleWaiverID = args[0]

			resp, err := client.RuleWaivers.GetRuleWaiver(params, auth)
			CheckErr(err)

			waiver := resp.Payload

			var controls []string
			if waiver.RuleComplianceMapping != nil {
				for _, mapping := range waiver.RuleComplianceMapping.(map[string]interface{}) {
					family := mapping.(map[string]interface{})
					familyControls := family["controls"].([]interface{})

					for _, control := range familyControls {
						controls = append(controls, control.(string))
					}

				}
			}

			items := []interface{}{
				Item{"RULE_WAIVER_ID", *waiver.ID},
				Item{"NAME", *waiver.Name},
				Item{"COMMENT", waiver.Comment},
				Item{"ENVIRONMENT_ID", *waiver.EnvironmentID},
				Item{"ENVIRONMENT_NAME", waiver.EnvironmentName},
				Item{"RULE_ID", *waiver.RuleID},
				Item{"RULE_DESCRIPTION", waiver.RuleDescription},
				Item{"RULE_COMPLIANCE_MAPPING", strings.Join(controls, ",")},
				Item{"RESOURCE_ID", *waiver.ResourceID},
				Item{"RESOURCE_TYPE", *waiver.ResourceType},
				Item{"RESOURCE_PROVIDER", *waiver.ResourceProvider},
				Item{"CREATED_AT", format.Unix(waiver.CreatedAt)},
				Item{"CREATED_BY", waiver.CreatedBy},
				Item{"CREATED_BY_DISPLAY_NAME", waiver.CreatedByDisplayName},
				Item{"UPDATED_AT", format.Unix(waiver.UpdatedAt)},
				Item{"UPDATED_BY", waiver.UpdatedBy},
				Item{"UPDATED_BY_DISPLAY_NAME", waiver.UpdatedByDisplayName},
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

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetRuleWaiverCommand())
}
