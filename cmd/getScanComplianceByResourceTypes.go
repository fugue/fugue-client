package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type getScanComplianceByResourceTypesOptions struct {
	Offset        int64
	MaxItems      int64
	ResourceTypes []string
	Families      []string
	Columns       []string
}

// NewGetScanComplianceByResourceTypesCommand returns a command that retrives
// compliance by resource types
func NewGetScanComplianceByResourceTypesCommand() *cobra.Command {

	var opts getScanComplianceByResourceTypesOptions

	cmd := &cobra.Command{
		Use:   "compliance-by-resource-types [scan_id]",
		Short: "Show compliance results by resource type",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := scans.NewGetComplianceByResourceTypesParams()
			params.ScanID = args[0]

			if opts.Offset > 0 {
				params.Offset = &opts.Offset
			}
			if opts.MaxItems > 0 {
				params.MaxItems = &opts.MaxItems
			}
			if len(opts.ResourceTypes) > 0 {
				params.ResourceType = opts.ResourceTypes
			}
			if len(opts.Families) > 0 {
				params.Family = opts.Families
			}

			resp, err := client.Scans.GetComplianceByResourceTypes(params, auth)
			CheckErr(err)

			rtypes := resp.Payload.Items
			rows := make([]interface{}, len(rtypes))
			for i, rtype := range rtypes {
				rows[i] = rtype
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
	cmd.Flags().StringSliceVar(&opts.ResourceTypes, "resource-type", nil, "Resource type filter")
	cmd.Flags().StringSliceVar(&opts.Families, "family", nil, "Compliance family filter")
	cmd.Flags().StringSliceVar(&opts.Columns, "columns", []string{"ResourceType", "Compliant", "Total"}, "columns to show")

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetScanComplianceByResourceTypesCommand())
}
