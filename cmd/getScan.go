package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

// NewGetScanCommand returns a command that retrives details of a single scan
func NewGetScanCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "scan [scan_id]",
		Short: "Get scan details",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := scans.NewGetScanParams()
			params.ScanID = args[0]

			resp, err := client.Scans.GetScan(params, auth)
			if err != nil {
				switch respError := err.(type) {
				case *scans.GetScanNotFound:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			scan := resp.Payload

			var summary models.ResourceSummary
			if resp.Payload.ResourceSummary != nil {
				summary = *resp.Payload.ResourceSummary
			}

			message := "-"
			if scan.Message != "" {
				message = scan.Message
			}

			items := []interface{}{
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
	getCmd.AddCommand(NewGetScanCommand())
}
