package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type Item struct {
	Key   string
	Value interface{}
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
			CheckErr(err)

			env := resp.Payload

			items := []interface{}{
				Item{"ID", env.ID},
				Item{"Name", env.Name},
				Item{"Provider", env.Provider},
				Item{"ScanInterval", env.ScanInterval},
				Item{"LastScanAt", env.LastScanAt},
				Item{"NextScanAt", env.NextScanAt},
				Item{"ScanStatus", env.ScanStatus},
				Item{"Compliance Families", env.ComplianceFamilies},
				Item{"Drift", env.Drift},
				Item{"Remediation", env.Remediation},
			}

			table, err := format.Table(format.TableOpts{
				Rows:       items,
				Columns:    []string{"Key", "Value"},
				ShowHeader: false,
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
