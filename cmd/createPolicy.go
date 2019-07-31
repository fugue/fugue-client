package cmd

import (
	"github.com/fugue/fugue-client/client/metadata"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createPolicyOptions struct {
	Provider                 string
	RemediationResourceTypes []string
	SurveyResourceTypes      []string
}

// NewCreatePolicyCommand returns a command that creates an IAM policy
// that can be used to allow Fugue to scan an environment
func NewCreatePolicyCommand() *cobra.Command {

	var opts createPolicyOptions

	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Get an AWS IAM policy for survey and remediation",
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := metadata.NewCreatePolicyParams()
			params.Provider = opts.Provider
			params.Input = &models.CreatePolicyInput{
				SurveyResourceTypes:    opts.SurveyResourceTypes,
				RemediateResourceTypes: opts.RemediationResourceTypes,
			}

			resp, err := client.Metadata.CreatePolicy(params, auth)
			CheckErr(err)

			showResponse(resp.Payload)
		},
	}

	cmd.Flags().StringVar(&opts.Provider, "provider", "aws", "Cloud provider [aws | aws_govcloud]")
	cmd.Flags().StringSliceVar(&opts.RemediationResourceTypes, "remediation-types", nil, "Remediation resource types")
	cmd.Flags().StringSliceVar(&opts.SurveyResourceTypes, "survey-types", nil, "Survey resource types")

	return cmd
}

func init() {
	getCmd.AddCommand(NewCreatePolicyCommand())
}
