package cmd

import (
	"github.com/fugue/fugue-client/client/scans"
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

			showResponse(resp.Payload)
		},
	}

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetScanCommand())
}
