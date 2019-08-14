package cmd

import (
	"fmt"
	"time"

	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
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
			CheckErr(err)

			scan := resp.Payload
			summary := resp.Payload.ResourceSummary

			createdAt := time.Unix(scan.CreatedAt, 0)
			finishedAt := time.Unix(scan.FinishedAt, 0)

			message := "-"
			if scan.Message != "" {
				message = scan.Message
			}

			items := []interface{}{
				Item{"ID", scan.ID},
				Item{"CREATED_AT", createdAt.Format(time.RFC3339)},
				Item{"FINISHED_AT", finishedAt.Format(time.RFC3339)},
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
