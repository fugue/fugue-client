package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createGoogleEnvironmentOptions struct {
	Name                string
	ServiceAccountEmail string
	ProjectID           string
	ScanInterval        int64
	ComplianceFamilies  []string
}

// NewCreateGoogleEnvironmentCommand returns a command that creates an environment
func NewCreateGoogleEnvironmentCommand() *cobra.Command {

	var opts createGoogleEnvironmentOptions

	cmd := &cobra.Command{
		Use:     "environment",
		Short:   "Create an Google environment",
		Aliases: []string{"env"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			scanScheduleEnabled := opts.ScanInterval != 0
			var scanIntervalPtr *int64
			if scanScheduleEnabled {
				scanIntervalPtr = &opts.ScanInterval
			}

			params := environments.NewCreateEnvironmentParams()
			params.Environment = &models.CreateEnvironmentInput{
				ComplianceFamilies:     opts.ComplianceFamilies,
				Name:                   opts.Name,
				Provider:               "google",
				ScanInterval:           scanIntervalPtr,
				SurveyResourceTypes:    []string{},
				RemediateResourceTypes: []string{},
				ScanScheduleEnabled:    &scanScheduleEnabled,

				ProviderOptions: &models.ProviderOptions{
					Google: &models.ProviderOptionsGoogle{
						ServiceAccountEmail: opts.ServiceAccountEmail,
						ProjectID:           opts.ProjectID,
					},
				},
			}

			resp, err := client.Environments.CreateEnvironment(params, auth)
			CheckErr(err)
			env := resp.Payload

			families := strings.Join(env.ComplianceFamilies, ",")

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
	cmd.Flags().StringVar(&opts.ServiceAccountEmail, "email", "", "Google Service Account Email")
	cmd.Flags().StringVar(&opts.ProjectID, "project-id", "", "Google Project ID (if not given, the project_id is extracted from the service acccount email)")

	cmd.Flags().Int64Var(&opts.ScanInterval, "scan-interval", 86400, "Scan interval (seconds)")
	cmd.Flags().StringSliceVar(&opts.ComplianceFamilies, "compliance-families", []string{}, "Compliance families")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("survey-resource-groups")

	return cmd
}

func init() {
	googleCmd.AddCommand(NewCreateGoogleEnvironmentCommand())
}
