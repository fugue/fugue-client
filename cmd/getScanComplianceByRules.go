package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type getScanComplianceByRulesOptions struct {
	Offset   int64
	MaxItems int64
	Results  []string
	Families []string
	Columns  []string
}

// NewGetScanComplianceByRulesCommand returns a command that retrives
// compliance by rule
func NewGetScanComplianceByRulesCommand() *cobra.Command {

	var opts getScanComplianceByRulesOptions

	cmd := &cobra.Command{
		Use:   "compliance-by-rules [scan_id]",
		Short: "Show compliance results by rule",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := scans.NewGetComplianceByRulesParams()
			params.ScanID = args[0]

			if opts.Offset > 0 {
				params.Offset = &opts.Offset
			}
			if opts.MaxItems > 0 {
				params.MaxItems = &opts.MaxItems
			}
			if len(opts.Results) > 0 {
				params.Result = format.NormalizeStrings(opts.Results)
			}
			if len(opts.Families) > 0 {
				params.Family = format.NormalizeStrings(opts.Families)
			}

			resp, err := client.Scans.GetComplianceByRules(params, auth)
			CheckErr(err)

			rules := resp.Payload.Items
			rows := make([]interface{}, len(rules))
			for i, rule := range rules {
				rows[i] = rule
			}

			table, err := format.Table(format.TableOpts{
				Rows:       rows,
				Columns:    opts.Columns,
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "Offset")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 0, "Max items")
	cmd.Flags().StringSliceVar(&opts.Results, "result", nil, "Rule result filter")
	cmd.Flags().StringSliceVar(&opts.Families, "family", nil, "Compliance family filter")
	cmd.Flags().StringSliceVar(&opts.Columns, "columns", []string{"Family", "Rule", "Result"}, "columns to show")

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetScanComplianceByRulesCommand())
}
