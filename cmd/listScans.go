package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type listScansOptions struct {
	Offset         int64
	MaxItems       int64
	OrderBy        string
	OrderDirection string
	Status         []string
	RangeFrom      int64
	RangeTo        int64
}

type listScansViewItem struct {
	ScanID        string
	EnvironmentID string
	CreatedAt     string
	FinishedAt    string
	Status        string
	Message       string
}

// NewListScansCommand returns a command that lists scans in Fugue
func NewListScansCommand() *cobra.Command {

	var opts listScansOptions

	cmd := &cobra.Command{
		Use:     "scans [environment_id]",
		Short:   "List scans belonging to an environment",
		Aliases: []string{"scan"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := scans.NewListScansParams()
			params.EnvironmentID = args[0]

			if opts.Offset > 0 {
				params.Offset = &opts.Offset
			}
			if opts.MaxItems > 0 {
				params.MaxItems = &opts.MaxItems
			}
			if opts.RangeFrom > 0 {
				params.RangeFrom = &opts.RangeFrom
			}
			if opts.RangeTo > 0 {
				params.RangeTo = &opts.RangeTo
			}
			if opts.OrderBy != "" {
				params.OrderBy = &opts.OrderBy
			}
			if opts.OrderDirection != "" {
				params.OrderDirection = &opts.OrderDirection
			}
			if len(opts.Status) > 0 {
				params.Status = opts.Status
			}

			resp, err := client.Scans.ListScans(params, auth)
			CheckErr(err)

			scans := resp.Payload.Items

			rows := make([]interface{}, len(scans))
			for i, scan := range scans {
				rows[i] = listScansViewItem{
					ScanID:        scan.ID,
					EnvironmentID: scan.EnvironmentID,
					CreatedAt:     format.Unix(scan.CreatedAt),
					FinishedAt:    format.Unix(scan.FinishedAt),
					Status:        scan.Status,
					Message:       scan.Message,
				}
			}

			table, err := format.Table(format.TableOpts{
				Rows:       rows,
				Columns:    []string{"ScanID", "CreatedAt", "FinishedAt", "Status"},
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "offset into results")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 20, "max items to return")
	cmd.Flags().StringVar(&opts.OrderBy, "order-by", "", "order by attribute")
	cmd.Flags().StringVar(&opts.OrderDirection, "order-direction", "", "order by direction [asc | desc]")
	cmd.Flags().StringSliceVar(&opts.Status, "status", nil, "Scan status filter [IN-PROGRESS | SUCCESS | ERROR]")
	cmd.Flags().Int64Var(&opts.RangeFrom, "range-from", 0, "Range from time filter")
	cmd.Flags().Int64Var(&opts.RangeTo, "range-to", 0, "Range to time filter")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListScansCommand())
}
