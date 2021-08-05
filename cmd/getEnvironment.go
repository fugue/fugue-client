package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

// Item is used to display an attribute and its value to the user
type Item struct {
	Attribute string
	Value     interface{}
}

// NewGetEnvironmentCommand returns a command that retrieves environment details
func NewGetEnvironmentCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "environment [environment_id]",
		Aliases: []string{"env"},
		Short:   "Retrieve details for an environment",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// When printing json gets the first (pos 0) API call printed
			jsonPositionToShow = 0
			client, auth := getClient()

			params := environments.NewGetEnvironmentParams()
			params.EnvironmentID = args[0]

			resp, err := client.Environments.GetEnvironment(params, auth)
			if err != nil {
				switch respError := err.(type) {
				case *environments.GetEnvironmentNotFound:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			env := resp.Payload

			maxItems := int64(1)
			listScansParams := scans.NewListScansParams()
			listScansParams.EnvironmentID = args[0]
			listScansParams.MaxItems = &maxItems
			lsResp, err := client.Scans.ListScans(listScansParams, auth)
			CheckErr(err)

			var lastScanID string
			if len(lsResp.Payload.Items) > 0 {
				lastScanID = lsResp.Payload.Items[0].ID
			}

			families := strings.Join(env.ComplianceFamilies, ", ")
			if families == "" {
				families = "-"
			}

			baselineID := env.BaselineID
			if baselineID == "" {
				baselineID = "-"
			}

			items := []interface{}{
				Item{"ENVIRONMENT_ID", env.ID},
				Item{"NAME", env.Name},
				Item{"PROVIDER", env.Provider},
				Item{"SCAN_INTERVAL", env.ScanInterval},
				Item{"BASELINE_ID", baselineID},
				Item{"LAST_SCAN_ID", lastScanID},
				Item{"LAST_SCAN_AT", format.Unix(env.LastScanAt)},
				Item{"NEXT_SCAN_AT", format.Unix(env.NextScanAt)},
				Item{"SCAN_STATUS", env.ScanStatus},
				Item{"COMPLIANCE_FAMILIES", families},
				Item{"DRIFT", env.Drift},
				Item{"REMEDIATION", env.Remediation},
			}

			switch env.Provider {
			case "aws":
				items = append(items, Item{"ROLE_ARN", env.ProviderOptions.Aws.RoleArn})
				if env.ProviderOptions.Aws.Region != "" {
					items = append(items, Item{"REGION", env.ProviderOptions.Aws.Region})
				} else if len(env.ProviderOptions.Aws.Regions) > 0 {
					items = append(items, Item{"REGIONS", strings.Join(env.ProviderOptions.Aws.Regions, ", ")})
				}
			case "aws_govcloud":
				items = append(items, Item{"ROLE_ARN", env.ProviderOptions.AwsGovcloud.RoleArn})
				if env.ProviderOptions.AwsGovcloud.Region != "" {
					items = append(items, Item{"REGION", env.ProviderOptions.AwsGovcloud.Region})
				} else if len(env.ProviderOptions.AwsGovcloud.Regions) > 0 {
					items = append(items, Item{"REGIONS", strings.Join(env.ProviderOptions.AwsGovcloud.Regions, ", ")})
				}
			case "azure":
				items = append(items, Item{"SUBSCRIPTION_ID", env.ProviderOptions.Azure.SubscriptionID})
				items = append(items, Item{"APPLICATION_ID", env.ProviderOptions.Azure.ApplicationID})
			case "google":
				items = append(items, Item{"PROJECT_ID", env.ProviderOptions.Google.ProjectID})
				items = append(items, Item{"SERVICE_ACCOUNT_EMAIL", env.ProviderOptions.Google.ServiceAccountEmail})
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
	getCmd.AddCommand(NewGetEnvironmentCommand())
}
