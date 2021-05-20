package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/go-openapi/runtime"
	"github.com/spf13/cobra"
)

// NewTriggerScanCommand returns a command that scans a specified environment
func NewTriggerScanCommand() *cobra.Command {

	var wait bool
	var scanFailureExitCode int

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

			var scan *models.ScanWithSummary
			var summary models.ResourceSummary
			for {
				resp, err := client.Scans.GetScan(params, auth)
				CheckErr(err)
				scan = resp.Payload
				if resp.Payload.ResourceSummary != nil {
					summary = *resp.Payload.ResourceSummary
				}
				if scan.Status != "IN_PROGRESS" || !wait {
					break
				}
				time.Sleep(time.Second * 30)
			}

			var items []interface{}

			if !wait {
				items = []interface{}{
					Item{"SCAN_ID", scan.ID},
					Item{"CREATED_AT", format.Unix(scan.CreatedAt)},
					Item{"STATUS", scan.Status},
				}
			} else {
				message := "-"
				if scan.Message != "" {
					message = scan.Message
				}
				items = []interface{}{
					Item{"SCAN_ID", scan.ID},
					Item{"CREATED_AT", format.Unix(scan.CreatedAt)},
					Item{"FINISHED_AT", format.Unix(scan.FinishedAt)},
					Item{"STATUS", scan.Status},
					Item{"MESSAGE", message},
					Item{"RESOURCE_COUNT", summary.Total},
					Item{"RESOURCE_TYPES", summary.ResourceTypes},
					Item{"COMPLIANT", summary.Compliant},
					Item{"NONCOMPLIANT", summary.Noncompliant},
					Item{"RULES_PASSED", summary.RulesPassed},
					Item{"RULES_FAILED", summary.RulesFailed},
				}
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

			if(wait && scan.Status == "ERROR") {
				os.Exit(int(scanFailureExitCode))
			}
		},
	}

	cmd.Flags().BoolVar(&wait, "wait", false, "Wait for scan to complete")
	cmd.Flags().IntVar(&scanFailureExitCode, "scan-failure-exit-code", 0, "Sets the exit code to raise when a scan fails. Default is 0. Used with the wait flag")

	return cmd
}

func init() {
	rootCmd.AddCommand(NewTriggerScanCommand())
}
