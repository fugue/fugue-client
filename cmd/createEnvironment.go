package cmd

import (
	"log"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createEnvironmentOptions struct {
	Name                     string
	Region                   string
	Provider                 string
	Role                     string
	ScanInterval             int64
	ComplianceFamilies       []string
	RemediationResourceTypes []string
	SurveyResourceTypes      []string
	ScanScheduleEnabled      bool
}

// NewCreateEnvironmentCommand returns a command that creates an environment
func NewCreateEnvironmentCommand() *cobra.Command {

	var opts createEnvironmentOptions

	cmd := &cobra.Command{
		Use:     "environment",
		Short:   "Create an environment",
		Aliases: []string{"env"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := environments.NewCreateEnvironmentParams()
			params.Environment = &models.CreateEnvironmentInput{
				ComplianceFamilies: opts.ComplianceFamilies,
				Name:               opts.Name,
				Provider:           opts.Provider,
				ProviderOptions: &models.ProviderOptions{
					Aws: &models.ProviderOptionsAws{
						Region:  opts.Region,
						RoleArn: opts.Role,
					},
				},
				ScanInterval:           opts.ScanInterval,
				SurveyResourceTypes:    opts.SurveyResourceTypes,
				RemediateResourceTypes: opts.RemediationResourceTypes,
				ScanScheduleEnabled:    opts.ScanScheduleEnabled,
			}

			resp, err := client.Environments.CreateEnvironment(params, auth)
			if err != nil {
				log.Fatal(err)
			}
			showResponse(resp.Payload)
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "Environment name")
	cmd.Flags().StringVar(&opts.Region, "region", "", "AWS region")
	cmd.Flags().StringVar(&opts.Provider, "provider", "", "Provider: aws | aws_govcloud")
	cmd.Flags().StringVar(&opts.Role, "role", "", "AWS IAM role arn")
	cmd.Flags().Int64Var(&opts.ScanInterval, "scan-interval", 0, "Scan interval (seconds)")
	cmd.Flags().StringSliceVar(&opts.ComplianceFamilies, "compliance-families", nil, "Compliance families")
	cmd.Flags().StringSliceVar(&opts.RemediationResourceTypes, "remediation-resource-types", nil, "Remediation resource types")
	cmd.Flags().StringSliceVar(&opts.SurveyResourceTypes, "survey-resource-types", nil, "Survey resource types")
	cmd.Flags().BoolVar(&opts.ScanScheduleEnabled, "scan-schedule-enabled", false, "Scan schedule enabled")

	return cmd
}

func init() {
	createCmd.AddCommand(NewCreateEnvironmentCommand())
}
