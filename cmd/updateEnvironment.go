package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client"
	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/go-openapi/runtime"
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
	Regions                []string
	ServiceAccountEmail    string
	URL                    string
	Branch                 string
}

func getEnvironmentToUpdate(client *client.Fugue, auth runtime.ClientAuthInfoWriter, environmentID string) *models.EnvironmentWithSummary {
	paramsGet := environments.NewGetEnvironmentParams()
	paramsGet.EnvironmentID = environmentID
	resp, err := client.Environments.GetEnvironment(paramsGet, auth)
	if err != nil {
		switch respError := err.(type) {
		case *environments.GetEnvironmentNotFound:
			Fatal(respError.Payload.Message, DefaultErrorExitCode)
		default:
			CheckErr(err)
		}
	}
	return resp.Payload
}

// NewUpdateEnvironmentCommand returns a command that updates an environment
func NewUpdateEnvironmentCommand() *cobra.Command {

	var opts updateEnvironmentOptions

	cmd := &cobra.Command{
		Use:     "environment [environment_id]",
		Short:   "Update environment settings",
		Aliases: []string{"env"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := environments.NewUpdateEnvironmentParams()
			params.EnvironmentID = args[0]
			params.Environment = &models.UpdateEnvironmentInput{}

			if len(opts.Regions) > 0 {
				// trying to update the regions. See if this environment has regions already
				env := getEnvironmentToUpdate(client, auth, params.EnvironmentID)

				if env.Provider == "aws" && len(env.ProviderOptions.Aws.Regions) > 0 {
					params.Environment.ProviderOptions = &models.ProviderOptionsUpdateInput{}
					params.Environment.ProviderOptions.Aws = &models.ProviderOptionsAwsUpdateInput{Regions: opts.Regions}
				} else if env.Provider == "aws_govcloud" && len(env.ProviderOptions.AwsGovcloud.Regions) > 0 {
					params.Environment.ProviderOptions = &models.ProviderOptionsUpdateInput{}
					params.Environment.ProviderOptions.AwsGovcloud = &models.ProviderOptionsAwsUpdateInput{Regions: opts.Regions}
				}
			}

			if opts.ServiceAccountEmail != "" {
				env := getEnvironmentToUpdate(client, auth, params.EnvironmentID)
				if env.Provider == "google" {
					params.Environment.ProviderOptions = &models.ProviderOptionsUpdateInput{}
					params.Environment.ProviderOptions.Google = &models.ProviderOptionsGoogleUpdateInput{ServiceAccountEmail: opts.ServiceAccountEmail}
				}
			}

			if opts.URL != "" || opts.Branch != "" {
				env := getEnvironmentToUpdate(client, auth, params.EnvironmentID)
				if env.Provider == "repository" {
					params.Environment.ProviderOptions = &models.ProviderOptionsUpdateInput{}
					params.Environment.ProviderOptions.Repository = &models.ProviderOptionsRepositoryUpdateInput{
						URL:    env.ProviderOptions.Repository.URL,
						Branch: env.ProviderOptions.Repository.Branch,
					}
				} else {
					Fatal("Only repository environments support --url and --branch", 400)
				}
			}

			// Using Visit here allows us to process only flags that were set
			//
			// Note that the generated Go models have `omitempty` set.  This
			// means that any booleans that are `false` are simply dropped from
			// the JSON.  We work around this questionable design decision
			// by using pointers to booleans for `ScanScheduleEnabled` and
			// `Remediation`.
			cmd.Flags().Visit(func(f *pflag.Flag) {
				switch f.Name {
				case "name":
					params.Environment.Name = opts.Name
				case "baseline-id":
					params.Environment.BaselineID = &opts.BaselineID
				case "scan-interval":
					params.Environment.ScanInterval = opts.ScanInterval
				case "compliance-families":
					params.Environment.ComplianceFamilies = opts.ComplianceFamilies
				case "survey-resource-types":
					params.Environment.SurveyResourceTypes = opts.SurveyResourceTypes
				case "remediate-resource-types":
					params.Environment.RemediateResourceTypes = opts.RemediateResourceTypes
				case "scan-schedule-enabled":
					params.Environment.ScanScheduleEnabled = &opts.ScanScheduleEnabled
				case "remediation":
					params.Environment.Remediation = &opts.Remediation
				case "url":
					params.Environment.ProviderOptions.Repository.URL = opts.URL
				case "branch":
					params.Environment.ProviderOptions.Repository.Branch = opts.Branch
				}
			})

			resp, err := client.Environments.UpdateEnvironment(params, auth)
			if err != nil {
				switch respError := err.(type) {
				case *environments.UpdateEnvironmentBadRequest:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			env := resp.Payload

			families := strings.Join(env.ComplianceFamilies, ", ")
			if families == "" {
				families = "-"
			}

			baselineID := env.BaselineID
			if baselineID == "" {
				baselineID = "-"
			}

			items := []interface{}{
				Item{"ENVIRONMENT_ID", env.ID},
				Item{"NAME", env.Name},
				Item{"PROVIDER", env.Provider},
				Item{"SCAN_INTERVAL", env.ScanInterval},
				Item{"BASELINE_ID", baselineID},
				Item{"LAST_SCAN_AT", format.Unix(env.LastScanAt)},
				Item{"NEXT_SCAN_AT", format.Unix(env.NextScanAt)},
				Item{"SCAN_STATUS", env.ScanStatus},
				Item{"COMPLIANCE_FAMILIES", families},
				Item{"DRIFT", env.Drift},
				Item{"REMEDIATION", env.Remediation},
			}

			switch env.Provider {
			case "aws":
				items = append(items, Item{"ROLE_ARN", env.ProviderOptions.Aws.RoleArn})
				if env.ProviderOptions.Aws.Region != "" {
					items = append(items, Item{"REGION", env.ProviderOptions.Aws.Region})
				} else if len(env.ProviderOptions.Aws.Regions) > 0 {
					items = append(items, Item{"REGIONS", strings.Join(env.ProviderOptions.Aws.Regions, ", ")})
				}
			case "aws_govcloud":
				items = append(items, Item{"ROLE_ARN", env.ProviderOptions.AwsGovcloud.RoleArn})
				if env.ProviderOptions.AwsGovcloud.Region != "" {
					items = append(items, Item{"REGION", env.ProviderOptions.AwsGovcloud.Region})
				} else if len(env.ProviderOptions.AwsGovcloud.Regions) > 0 {
					items = append(items, Item{"REGIONS", strings.Join(env.ProviderOptions.AwsGovcloud.Regions, ", ")})
				}
			case "azure":
				items = append(items, Item{"SUBSCRIPTION_ID", env.ProviderOptions.Azure.SubscriptionID})
				items = append(items, Item{"APPLICATION_ID", env.ProviderOptions.Azure.ApplicationID})
			case "google":
				items = append(items, Item{"PROJECT_ID", env.ProviderOptions.Google.ProjectID})
				items = append(items, Item{"SERVICE_ACCOUNT_EMAIL", env.ProviderOptions.Google.ServiceAccountEmail})
			case "repository":
				items = append(items, Item{"URL", env.ProviderOptions.Repository.URL})
				items = append(items, Item{"BRANCH", env.ProviderOptions.Repository.Branch})
			}

			table, err := format.Table(format.TableOpts{
				Rows:         items,
				Columns:      []string{"Attribute", "Value"},
				ShowHeader:   true,
				MaxCellWidth: 70,
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
	cmd.Flags().BoolVar(&opts.ScanScheduleEnabled, "scan-schedule-enabled", true, "Enable automatic scanning schedule")
	cmd.Flags().StringSliceVar(&opts.ComplianceFamilies, "compliance-families", nil, "Compliance families")
	cmd.Flags().BoolVar(&opts.Remediation, "remediation", false, "Enable automatic remediation")
	cmd.Flags().StringSliceVar(&opts.RemediateResourceTypes, "remediate-resource-types", nil, "Remediation resource types")
	cmd.Flags().StringSliceVar(&opts.SurveyResourceTypes, "survey-resource-types", nil, "Survey resource types")
	cmd.Flags().StringSliceVar(&opts.Regions, "regions", nil, "AWS regions")
	cmd.Flags().StringVar(&opts.ServiceAccountEmail, "service-account-email", "", "Google service account email")
	cmd.Flags().StringVar(&opts.URL, "url", "", "URL for repository environment")
	cmd.Flags().StringVar(&opts.Branch, "branch", "", "Repository environment branch")

	return cmd
}

func init() {
	updateCmd.AddCommand(NewUpdateEnvironmentCommand())
}
