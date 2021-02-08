package cmd

import (
	"fmt"
	"os"

	"github.com/fugue/fugue-client/client/rule_waivers"

	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type updateRuleWaiverOptions struct {
	ID      string
	Name    string
	Comment string
}

// NewUpdateRuleWaiverCommand returns a command that updates a rule waiver
func NewUpdateRuleWaiverCommand() *cobra.Command {

	var opts updateRuleWaiverOptions

	cmd := &cobra.Command{
		Use:     "rule_waiver [rule_waiver_id]",
		Short:   "Update rule waiver settings",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"waiver"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := rule_waivers.NewUpdateRuleWaiverParams()
			params.RuleWaiverID = args[0]
			params.Input = &models.UpdateRuleWaiverInput{}

			flagCount := 0

			// Using Visit here allows us to process only flags that were set
			cmd.Flags().Visit(func(f *pflag.Flag) {
				flagCount++
				switch f.Name {
				case "name":
					params.Input.Name = opts.Name
				case "comment":
					params.Input.Comment = opts.Comment
				}
			})

			if flagCount == 0 {
				os.Exit(0)
			}

			resp, err := client.RuleWaivers.UpdateRuleWaiver(params, auth)
			CheckErr(err)

			waiver := resp.Payload

			items := []interface{}{
				Item{"NAME", *waiver.Name},
				Item{"COMMENT", waiver.Comment},
				Item{"ENVIRONMENT_ID", *waiver.EnvironmentID},
				Item{"ENVIRONMENT_NAME", waiver.EnvironmentName},
				Item{"RULE_ID", *waiver.RuleID},
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
				Rows:       items,
				Columns:    []string{"Attribute", "Value"},
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "Waiver name")
	cmd.Flags().StringVar(&opts.Comment, "comment", "", "Waiver comment")

	return cmd
}

func init() {
	updateCmd.AddCommand(NewUpdateRuleWaiverCommand())
}
