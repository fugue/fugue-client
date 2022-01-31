package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createRepositoryEnvironmentOptions struct {
	Name               string
	ComplianceFamilies []string
	URL                string
	Branch             string
}

// NewCreateRepositoryEnvironmentCommand returns a command that creates a repository environment
func NewCreateRepositoryEnvironmentCommand() *cobra.Command {

	var opts createRepositoryEnvironmentOptions

	cmd := &cobra.Command{
		Use:     "environment",
		Short:   "Create a Repository environment",
		Aliases: []string{"env"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := environments.NewCreateEnvironmentParams()
			scanScheduleEnabled := false
			params.Environment = &models.CreateEnvironmentInput{
				ComplianceFamilies:  opts.ComplianceFamilies,
				Name:                opts.Name,
				Provider:            "repository",
				ScanScheduleEnabled: &scanScheduleEnabled,
				ProviderOptions: &models.ProviderOptions{
					Repository: &models.ProviderOptionsRepository{
						Branch: opts.Branch,
						URL:    opts.URL,
					},
				},
			}

			fmt.Fprintf(os.Stderr, "Sending request...\n")
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

			families := strings.Join(env.ComplianceFamilies, ", ")
			if families == "" {
				families = "-"
			}

			items := []interface{}{
				Item{"ENVIRONMENT_ID", env.ID},
				Item{"NAME", env.Name},
				Item{"URL", env.ProviderOptions.Repository.URL},
				Item{"BRANCH", env.ProviderOptions.Repository.Branch},
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
	cmd.Flags().StringVar(&opts.URL, "url", "", "URL to repository")
	cmd.Flags().StringVar(&opts.Branch, "branch", "", "Branch in repository to use")
	cmd.Flags().StringSliceVar(&opts.ComplianceFamilies, "compliance-families", []string{}, "Compliance families")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("url")
	cmd.MarkFlagRequired("branch")

	return cmd
}

func init() {
	repositoryCmd.AddCommand(NewCreateRepositoryEnvironmentCommand())
}
