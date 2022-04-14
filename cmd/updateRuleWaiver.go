package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fugue/fugue-client/client/rule_waivers"

	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type updateRuleWaiverOptions struct {
	ID        string
	Name      string
	Comment   string
	ExpiresAt string
}

// NewUpdateRuleWaiverCommand returns a command that updates a rule waiver
func NewUpdateRuleWaiverCommand() *cobra.Command {

	var opts updateRuleWaiverOptions

	cmd := &cobra.Command{
		Use:     "rule-waiver [rule_waiver_id]",
		Short:   "Update rule waiver settings",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"waiver", "rule_waiver"},
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
				case "expires-at":
					expiresAtPtr, duration, err := parseBothExpiresAt(opts.ExpiresAt)
					CheckErr(err)
					if expiresAtPtr != nil {
						params.Input.ExpiresAt = expiresAtPtr.Unix()
					}
					params.Input.ExpiresAtDuration = duration
				}
			})

			if flagCount == 0 {
				os.Exit(0)
			}

			resp, err := client.RuleWaivers.UpdateRuleWaiver(params, auth)
			CheckErr(err)

			waiver := resp.Payload

			var itemTag Item
			if waiver.ResourceTag != "" {
				itemTag = Item{"RESOURCE_TAG", waiver.ResourceTag}
			} else {
				itemTag = Item{"RESOURCE_TAG", "-"}
			}

			var itemTime Item
			if waiver.ExpiresAt != 0 {
				t := time.Unix(waiver.ExpiresAt, 0)
				itemTime = Item{"EXPIRES_AT", t.Format(time.RFC3339)}
			} else {
				itemTime = Item{"EXPIRES_AT", "-"}
			}

			items := []interface{}{
				Item{"RULE_WAIVER_ID", *waiver.ID},
				Item{"NAME", *waiver.Name},
				Item{"COMMENT", waiver.Comment},
				Item{"ENVIRONMENT_ID", *waiver.EnvironmentID},
				Item{"ENVIRONMENT_NAME", waiver.EnvironmentName},
				Item{"RULE_ID", *waiver.RuleID},
				Item{"RESOURCE_ID", *waiver.ResourceID},
				Item{"RESOURCE_TYPE", *waiver.ResourceType},
				Item{"RESOURCE_PROVIDER", *waiver.ResourceProvider},
				itemTag,
				itemTime,
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

	// Add to the documents:
	// use ISO 8601 format for the duration (e.g. P1Y2M3DT4H5M6S) up to hours (e.g. PT1H)
	// They can also drop the P and T (e.g. 1Y2M3DT4H) and it's case insensitive
	cmd.Flags().StringVar(&opts.ExpiresAt, "expires-at", "",
		"Expires at in RFC3339 representation, Unix timestamp (e.g. '2020-01-01T00:00:00Z' or '1577836800') or at duration in ISO 8601 format (e.g. 'P3Y6M4DT12H') or '4d', 1d12h, etc.")

	return cmd
}

func init() {
	updateCmd.AddCommand(NewUpdateRuleWaiverCommand())
}
