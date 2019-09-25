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

			families := strings.Join(env.ComplianceFamilies, ",")
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
				items = append(items, Item{"ROLE", env.ProviderOptions.Aws.RoleArn})
				items = append(items, Item{"REGION", env.ProviderOptions.Aws.Region})
			case "aws_govcloud":
				items = append(items, Item{"ROLE", env.ProviderOptions.AwsGovcloud.RoleArn})
				items = append(items, Item{"REGION", env.ProviderOptions.AwsGovcloud.Region})
			case "azure":
				items = append(items, Item{"SUBSCRIPTION_ID", env.ProviderOptions.Azure.SubscriptionID})
				items = append(items, Item{"APPLICATION_ID", env.ProviderOptions.Azure.ApplicationID})
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

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetEnvironmentCommand())
}
