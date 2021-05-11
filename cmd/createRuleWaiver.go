package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/client/rule_waivers"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createRuleWaiverOptions struct {
	Name             string
	Comment          string
	EnvironmentID    string
	RuleID           string
	ResourceID       string
	ResourceType     string
	ResourceProvider string
	ResourceTag      string
	WildcardMode     bool
}

// NewCreateRuleWaiverCommand returns a command that creates a custom rule
func NewCreateRuleWaiverCommand() *cobra.Command {

	var opts createRuleWaiverOptions

	cmd := &cobra.Command{
		Use:     "rule-waiver",
		Short:   "Create a rule waiver",
		Aliases: []string{"waiver", "rule_waiver", "rule-waivers", "waivers", "rule_waivers"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := rule_waivers.NewCreateRuleWaiverParams()
			params.Input = &models.CreateRuleWaiverInput{
				Name:             &opts.Name,
				Comment:          opts.Comment,
				EnvironmentID:    &opts.EnvironmentID,
				RuleID:           &opts.RuleID,
				ResourceID:       &opts.ResourceID,
				ResourceType:     &opts.ResourceType,
				ResourceProvider: &opts.ResourceProvider,
				ResourceTag:      opts.ResourceTag,
				WildcardMode:     opts.WildcardMode,
			}

			resp, err := client.RuleWaivers.CreateRuleWaiver(params, auth)
			if err != nil {
				switch respError := err.(type) {
				case *custom_rules.CreateCustomRuleInternalServerError:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			waiver := resp.Payload

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
				Item{"RESOURCE_TAG", *waiver.ResourceTag},
				Item{"WILDCARD_MODE", *waiver.WildcardMode},
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
	cmd.Flags().StringVar(&opts.Comment, "comment", "", "Comment describing the rule waiver purpose")
	cmd.Flags().StringVar(&opts.RuleID, "rule-id", "", "Rule ID (e.g. FG_R00217, <UUID Custom Rule ID>)")
	cmd.Flags().StringVar(&opts.EnvironmentID, "environment-id", "", "Environment ID")
	cmd.Flags().StringVar(&opts.ResourceID, "resource-id", "*", "Resource ID (e.g. resource-123, 'resource-*')")
	cmd.Flags().StringVar(&opts.ResourceType, "resource-type", "*", "Resource Type (e.g. AWS.S3.Bucket, '*')")
	cmd.Flags().StringVar(&opts.ResourceProvider, "resource-provider", "*", "Resource Provider (e.g. aws.us-east-1, azure, '*')")
	cmd.Flags().StringVar(&opts.ResourceTag, "resource-tag", "*", "Resource tag (e.g. 'env:prod', 'env:*', '*')")
	cmd.Flags().BoolVar(&opts.WildcardMode, "wildcard-mode", true, "Controls whether glob-style wildcard characters are expanded")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("rule-id")
	cmd.MarkFlagRequired("environment-id")

	return cmd
}

func init() {
	createCmd.AddCommand(NewCreateRuleWaiverCommand())
}
