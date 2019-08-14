package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
	"github.com/go-openapi/runtime"
	"github.com/spf13/cobra"
)

// NewTriggerScanCommand returns a command that scans a specified environment
func NewTriggerScanCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "scan [environment_id]",
		Short: "Trigger a scan",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			createParams := scans.NewCreateScanParams()
			createParams.EnvironmentID = args[0]

			createResp, err := client.Scans.CreateScan(createParams, auth)
			if err != nil {
				switch respError := err.(type) {
				case *scans.CreateScanBadRequest:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				case *runtime.APIError:
					if respError.Code == 404 {
						Fatal("Environment not found", DefaultErrorExitCode)
					}
					CheckErr(err)
				default:
					CheckErr(err)
				}
			}

			scanID := createResp.Payload.ID

			params := scans.NewGetScanParams()
			params.ScanID = scanID
			resp, err := client.Scans.GetScan(params, auth)
			CheckErr(err)

			scan := resp.Payload

			items := []interface{}{
				Item{"SCAN_ID", scan.ID},
				Item{"CREATED_AT", format.Unix(scan.CreatedAt)},
				Item{"STATUS", scan.Status},
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
	rootCmd.AddCommand(NewTriggerScanCommand())
}
