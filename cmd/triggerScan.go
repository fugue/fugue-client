package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/scans"
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

			params := scans.NewCreateScanParams()
			params.EnvironmentID = args[0]

			resp, err := client.Scans.CreateScan(params, auth)
			CheckErr(err)

			fmt.Println(resp.Payload.ID)
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(NewTriggerScanCommand())
}
