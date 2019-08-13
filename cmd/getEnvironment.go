package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type Item struct {
	Attribute string
	Value     interface{}
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

			maxItems := int64(1)
			listScansParams := scans.NewListScansParams()
			listScansParams.EnvironmentID = args[0]
			listScansParams.MaxItems = &maxItems
			lsResp, err := client.Scans.ListScans(listScansParams, auth)
			CheckErr(err)

			var lastScanID string
			if len(lsResp.Payload.Items) > 0 {
				lastScanID = lsResp.Payload.Items[0].ID
			}

			families := strings.Join(env.ComplianceFamilies, ",")
			nextScanAt := time.Unix(env.NextScanAt, 0)
			lastScanAt := time.Unix(env.LastScanAt, 0)

			items := []interface{}{
				Item{"ID", env.ID},
				Item{"Name", env.Name},
				Item{"Provider", env.Provider},
				Item{"ScanInterval", env.ScanInterval},
				Item{"LastScanID", lastScanID},
				Item{"LastScanAt", lastScanAt.Format(time.RFC3339)},
				Item{"NextScanAt", nextScanAt.Format(time.RFC3339)},
				Item{"ScanStatus", env.ScanStatus},
				Item{"ComplianceFamilies", families},
				Item{"Drift", env.Drift},
				Item{"Remediation", env.Remediation},
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
	getCmd.AddCommand(NewGetEnvironmentCommand())
}
