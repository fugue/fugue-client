package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type updateEnvironmentOptions struct {
	Name                   string
	BaselineID             string
	ScanInterval           int64
	ComplianceFamilies     []string
	RemediateResourceTypes []string
	SurveyResourceTypes    []string
	Remediation            bool
	ScanScheduleEnabled    bool
}

// NewUpdateEnvironmentCommand returns a command that updates an environment
func NewUpdateEnvironmentCommand() *cobra.Command {

	var opts updateEnvironmentOptions

	cmd := &cobra.Command{
		Use:     "environment",
		Short:   "Update environment settings",
		Aliases: []string{"env"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := environments.NewUpdateEnvironmentParams()
			params.EnvironmentID = args[0]
			params.Environment = &models.UpdateEnvironmentInput{}

			// The generated Go models have omitempty set on boolean flags.
			// This means we can't send "false" values for these fields:
			// * remediation
			// * scan_schedule_enabled
			// For now we won't support setting these flags in the CLI.

			// Using Visit here allows us to process only flags that were set
			cmd.Flags().Visit(func(f *pflag.Flag) {
				switch f.Name {
				case "name":
					params.Environment.Name = opts.Name
				case "baseline-id":
					params.Environment.BaselineID = opts.BaselineID
				case "scan-interval":
					params.Environment.ScanInterval = opts.ScanInterval
				case "compliance-families":
					params.Environment.ComplianceFamilies = opts.ComplianceFamilies
				case "survey-resource-types":
					params.Environment.SurveyResourceTypes = opts.SurveyResourceTypes
				case "remediate-resource-types":
					params.Environment.RemediateResourceTypes = opts.RemediateResourceTypes
				}
			})

			resp, err := client.Environments.UpdateEnvironment(params, auth)
			CheckErr(err)
			env := resp.Payload

			families := strings.Join(env.ComplianceFamilies, ",")
			if families == "" {
				families = "-"
			}

			items := []interface{}{
				Item{"ENVIRONMENT_ID", env.ID},
				Item{"NAME", env.Name},
				Item{"PROVIDER", env.Provider},
				Item{"SCAN_INTERVAL", env.ScanInterval},
				Item{"LAST_SCAN_AT", format.Unix(env.LastScanAt)},
				Item{"NEXT_SCAN_AT", format.Unix(env.NextScanAt)},
				Item{"SCAN_STATUS", env.ScanStatus},
				Item{"COMPLIANCE_FAMILIES", families},
				Item{"DRIFT", env.Drift},
				Item{"REMEDIATION", env.Remediation},
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

	cmd.Flags().StringVar(&opts.Name, "name", "", "Environment name")
	cmd.Flags().StringVar(&opts.BaselineID, "baseline-id", "", "Baseline scan ID")
	cmd.Flags().Int64Var(&opts.ScanInterval, "scan-interval", 0, "Scan interval (seconds)")
	cmd.Flags().StringSliceVar(&opts.ComplianceFamilies, "compliance-families", nil, "Compliance families")
	cmd.Flags().StringSliceVar(&opts.RemediateResourceTypes, "remediate-resource-types", nil, "Remediation resource types")
	cmd.Flags().StringSliceVar(&opts.SurveyResourceTypes, "survey-resource-types", nil, "Survey resource types")

	return cmd
}

func init() {
	updateCmd.AddCommand(NewUpdateEnvironmentCommand())
}
