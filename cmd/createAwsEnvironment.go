package cmd

import (
	"bytes"
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
	Regions                  []string
	Provider                 string
	GovCloud                 bool
	Role                     string
	ScanInterval             int64
	ComplianceFamilies       []string
	RemediationResourceTypes []string
	SurveyResourceTypes      []string
}

func all(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func stringContainsGov(elem string) bool {
	return strings.Contains(elem, "gov")
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

			surveyTypes := opts.SurveyResourceTypes
			if len(surveyTypes) == 0 {
				// Default to all available types
				getTypesParams := metadata.NewGetResourceTypesParams()
				getTypesParams.Provider = opts.Provider
				getTypesParams.Region = &opts.Region // if opts.Regions, opts.Region is ""
				resp, err := client.Metadata.GetResourceTypes(getTypesParams, auth)
				CheckErr(err)
				surveyTypes = resp.Payload.ResourceTypes
			}

			params := environments.NewCreateEnvironmentParams()
			params.Environment = &models.CreateEnvironmentInput{
				ComplianceFamilies:     opts.ComplianceFamilies,
				Name:                   opts.Name,
				Provider:               opts.Provider,
				ScanInterval:           opts.ScanInterval,
				SurveyResourceTypes:    surveyTypes,
				RemediateResourceTypes: opts.RemediationResourceTypes,
				ScanScheduleEnabled:    true,
			}

			providerOpts := &models.ProviderOptionsAws{
				Region:  opts.Region,
				Regions: opts.Regions,
				RoleArn: opts.Role,
			}

			if opts.Provider == "aws_govcloud" {
				params.Environment.ProviderOptions = &models.ProviderOptions{AwsGovcloud: providerOpts}
			} else {
				params.Environment.ProviderOptions = &models.ProviderOptions{Aws: providerOpts}
			}

			resp, err := client.Environments.CreateEnvironment(params, auth)

			buf := new(bytes.Buffer)
			fmt.Println(buf)
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
		Args: func(cmd *cobra.Command, args []string) error {

			if opts.Region == "" && len(opts.Regions) == 0 {
				if opts.Provider != "aws" && opts.Provider != "aws_govcloud" {
					return fmt.Errorf("Regions not specified. Please specify a provider: aws or aws_govcloud")
				}
				opts.Regions = []string{"*"}
			} else if opts.Region != "" {
				opts.Region = strings.ToLower(opts.Region)
				opts.Provider = "aws"
				if strings.Contains(opts.Region, "gov") {
					opts.Provider = "aws_govcloud"
				}
			} else if len(opts.Regions) > 0 {
				opts.Provider = "aws"
				var regions []string
				for _, thisRegion := range opts.Regions {
					regions = append(regions, strings.ToLower(thisRegion))
				}
				if all(regions, stringContainsGov) {
					opts.Provider = "aws_govcloud"
				}
				opts.Regions = regions
			} else {
				return fmt.Errorf("Unknown error: %s", args)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "Environment name")
	cmd.Flags().StringVar(&opts.Region, "region", "", "AWS region (deprecated)")
	cmd.Flags().StringSliceVar(&opts.Regions, "regions", []string{}, "AWS regions (default all regions)")
	cmd.Flags().StringVar(&opts.Provider, "provider", "", "Provider if cannot be resolved from regions")
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
