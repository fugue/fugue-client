package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createRuleOptions struct {
	Name         string
	Description  string
	Provider     string
	Severity     string
	ResourceType string
	RuleText     string
}

// NewCreateRuleCommand returns a command that creates a custom rule
func NewCreateRuleCommand() *cobra.Command {

	var opts createRuleOptions

	cmd := &cobra.Command{
		Use:   "rule",
		Short: "Create a custom rule",
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := custom_rules.NewCreateCustomRuleParams()
			params.Rule = &models.CreateCustomRuleInput{
				Name:         opts.Name,
				Description:  opts.Description,
				ResourceType: opts.ResourceType,
				Severity:     opts.Severity,
				Provider:     opts.Provider,
				RuleText:     opts.RuleText,
			}

			resp, err := client.CustomRules.CreateCustomRule(params, auth)
			if err != nil {
				switch respError := err.(type) {
				case *custom_rules.CreateCustomRuleInternalServerError:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			rule := resp.Payload

			items := []interface{}{
				Item{"RULE_ID", rule.ID},
				Item{"NAME", rule.Name},
				Item{"DESCRIPTION", rule.Description},
				Item{"PROVIDER", rule.Provider},
				Item{"RESOURCE_TYPE", rule.ResourceType},
				Item{"SEVERITY", rule.Severity},
				Item{"STATUS", rule.Status},
				Item{"FAMILIES", "-"}, // Family is always created without families
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
	cmd.Flags().StringVar(&opts.Provider, "provider", "", "Provider")
	cmd.Flags().StringVar(&opts.Severity, "severity", "", "Severity")
	cmd.Flags().StringVar(&opts.ResourceType, "resource-type", "", "Resource type")
	cmd.Flags().StringVar(&opts.RuleText, "text", "", "Rule text")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("provider")
	cmd.MarkFlagRequired("resource-type")
	cmd.MarkFlagRequired("text")

	return cmd
}

func init() {
	createCmd.AddCommand(NewCreateRuleCommand())
}
