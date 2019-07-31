package cmd

import (
	"github.com/fugue/fugue-client/client/environments"
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
				case "remediation":
					params.Environment.Remediation = opts.Remediation
				case "scan-schedule-enabled":
					params.Environment.ScanScheduleEnabled = opts.ScanScheduleEnabled
				}
			})

			resp, err := client.Environments.UpdateEnvironment(params, auth)
			CheckErr(err)

			showResponse(resp.Payload)
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "Environment name")
	cmd.Flags().StringVar(&opts.BaselineID, "baseline-id", "", "Baseline scan ID")
	cmd.Flags().Int64Var(&opts.ScanInterval, "scan-interval", 0, "Scan interval (seconds)")
	cmd.Flags().StringSliceVar(&opts.ComplianceFamilies, "compliance-families", nil, "Compliance families")
	cmd.Flags().StringSliceVar(&opts.RemediateResourceTypes, "remediate-resource-types", nil, "Remediation resource types")
	cmd.Flags().StringSliceVar(&opts.SurveyResourceTypes, "survey-resource-types", nil, "Survey resource types")
	cmd.Flags().BoolVar(&opts.Remediation, "remediation", false, "Remediation enabled")
	cmd.Flags().BoolVar(&opts.ScanScheduleEnabled, "scan-schedule-enabled", false, "Scan schedule enabled")

	return cmd
}

func init() {
	updateCmd.AddCommand(NewUpdateEnvironmentCommand())
}
