package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createAzureEnvironmentOptions struct {
	Name                    string
	ApplicationID           string
	ClientSecret            string
	SubscriptionID          string
	TenantID                string
	ScanInterval            int64
	ComplianceFamilies      []string
	SurveyResourceGroups    []string
	RemediateResourceGroups []string
}

// NewCreateAzureEnvironmentCommand returns a command that creates an environment
func NewCreateAzureEnvironmentCommand() *cobra.Command {

	var opts createAzureEnvironmentOptions

	cmd := &cobra.Command{
		Use:     "environment",
		Short:   "Create an Azure environment",
		Aliases: []string{"env"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			scanInterval := opts.ScanInterval
			var scanIntervalPtr *int64
			if opts.ScanInterval > 0 {
				scanIntervalPtr = &scanInterval
			}
			scanScheduleEnabled := scanInterval != 0

			params := environments.NewCreateEnvironmentParams()
			params.Environment = &models.CreateEnvironmentInput{
				ComplianceFamilies:     opts.ComplianceFamilies,
				Name:                   opts.Name,
				Provider:               "azure",
				ScanInterval:           scanIntervalPtr,
				SurveyResourceTypes:    []string{},
				RemediateResourceTypes: []string{},
				ScanScheduleEnabled:    scanScheduleEnabled,
				ProviderOptions: &models.ProviderOptions{
					Azure: &models.ProviderOptionsAzure{
						ApplicationID:           opts.ApplicationID,
						ClientSecret:            opts.ClientSecret,
						SubscriptionID:          opts.SubscriptionID,
						TenantID:                opts.TenantID,
						SurveyResourceGroups:    opts.SurveyResourceGroups,
						RemediateResourceGroups: opts.RemediateResourceGroups,
					},
				},
			}

			resp, err := client.Environments.CreateEnvironment(params, auth)
			CheckErr(err)
			env := resp.Payload

			families := strings.Join(env.ComplianceFamilies, ",")

			items := []interface{}{
				Item{"ENVIRONMENT_ID", env.ID},
				Item{"NAME", env.Name},
				Item{"PROVIDER", env.Provider},
				Item{"SCAN_INTERVAL", env.ScanInterval},
				Item{"LAST_SCAN_AT", format.Unix(env.LastScanAt)},
				Item{"NEXT_SCAN_AT", format.Unix(env.NextScanAt)},
				Item{"SCAN_STATUS", env.ScanStatus},
				Item{"COMPLIANCE_FAMILIES", families},
				Item{"DRIFT", env.Drift},
				Item{"REMEDIATION", env.Remediation},
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

	cmd.Flags().StringVar(&opts.Name, "name", "", "Environment name")
	cmd.Flags().StringVar(&opts.ApplicationID, "app", "", "Azure Application ID")
	cmd.Flags().StringVar(&opts.ClientSecret, "secret", "", "Azure Client Secret")
	cmd.Flags().StringVar(&opts.SubscriptionID, "sub", "", "Azure Subscription ID")
	cmd.Flags().StringVar(&opts.TenantID, "tenant", "", "Azure Tenant ID")

	cmd.Flags().Int64Var(&opts.ScanInterval, "scan-interval", 0, "Scan interval (seconds)")
	cmd.Flags().StringSliceVar(&opts.ComplianceFamilies, "compliance-families", []string{}, "Compliance families")
	cmd.Flags().StringSliceVar(&opts.RemediateResourceGroups, "remediation-resource-groups", []string{}, "Remediation resource groups")
	cmd.Flags().StringSliceVar(&opts.SurveyResourceGroups, "survey-resource-groups", nil, "Survey resource groups")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("app")
	cmd.MarkFlagRequired("secret")
	cmd.MarkFlagRequired("sub")
	cmd.MarkFlagRequired("tenant")
	cmd.MarkFlagRequired("survey-resource-groups")

	return cmd
}

func init() {
	azureCmd.AddCommand(NewCreateAzureEnvironmentCommand())
}
