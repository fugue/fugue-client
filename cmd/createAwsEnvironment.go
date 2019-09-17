package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/client/metadata"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createAwsEnvironmentOptions struct {
	Name                     string
	Region                   string
	GovCloud                 bool
	Role                     string
	ScanInterval             int64
	ComplianceFamilies       []string
	RemediationResourceTypes []string
	SurveyResourceTypes      []string
}

// NewCreateAwsEnvironmentCommand returns a command that creates an environment
func NewCreateAwsEnvironmentCommand() *cobra.Command {

	var opts createAwsEnvironmentOptions

	cmd := &cobra.Command{
		Use:     "environment",
		Short:   "Create an AWS environment",
		Aliases: []string{"env"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			provider := "aws"
			if opts.GovCloud {
				provider = "aws_govcloud"
			}

			surveyTypes := opts.SurveyResourceTypes
			if len(surveyTypes) == 0 {
				// Default to all available types
				getTypesParams := metadata.NewGetResourceTypesParams()
				getTypesParams.Provider = provider
				getTypesParams.Region = &opts.Region
				resp, err := client.Metadata.GetResourceTypes(getTypesParams, auth)
				CheckErr(err)
				surveyTypes = resp.Payload.ResourceTypes
			}

			params := environments.NewCreateEnvironmentParams()
			params.Environment = &models.CreateEnvironmentInput{
				ComplianceFamilies:     opts.ComplianceFamilies,
				Name:                   opts.Name,
				Provider:               provider,
				ScanInterval:           opts.ScanInterval,
				SurveyResourceTypes:    surveyTypes,
				RemediateResourceTypes: opts.RemediationResourceTypes,
				ScanScheduleEnabled:    true,
			}

			providerOpts := &models.ProviderOptionsAws{
				Region:  opts.Region,
				RoleArn: opts.Role,
			}

			if opts.GovCloud {
				params.Environment.ProviderOptions = &models.ProviderOptions{AwsGovcloud: providerOpts}
			} else {
				params.Environment.ProviderOptions = &models.ProviderOptions{Aws: providerOpts}
			}

			resp, err := client.Environments.CreateEnvironment(params, auth)
			if err != nil {
				switch respError := err.(type) {
				case *environments.CreateEnvironmentInternalServerError:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

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
	cmd.Flags().StringVar(&opts.Region, "region", "us-east-1", "AWS region")
	cmd.Flags().BoolVar(&opts.GovCloud, "govcloud", false, "Is GovCloud?")
	cmd.Flags().StringVar(&opts.Role, "role", "", "AWS IAM role arn")
	cmd.Flags().Int64Var(&opts.ScanInterval, "scan-interval", 86400, "Scan interval (seconds)")
	cmd.Flags().StringSliceVar(&opts.ComplianceFamilies, "compliance-families", []string{}, "Compliance families")
	cmd.Flags().StringSliceVar(&opts.RemediationResourceTypes, "remediation-resource-types", []string{}, "Remediation resource types")
	cmd.Flags().StringSliceVar(&opts.SurveyResourceTypes, "survey-resource-types", nil, "Survey resource types (defaults to all available types)")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("role")

	return cmd
}

func init() {
	awsCmd.AddCommand(NewCreateAwsEnvironmentCommand())
}
