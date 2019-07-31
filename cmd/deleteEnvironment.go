package cmd

import (
	"github.com/fugue/fugue-client/client/environments"
	"github.com/spf13/cobra"
)

// NewDeleteEnvironmentCommand returns a command that deletes an environment
func NewDeleteEnvironmentCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "environment [environment_id]",
		Aliases: []string{"env"},
		Short:   "Deletes an environment",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := environments.NewDeleteEnvironmentParams()
			params.EnvironmentID = args[0]

			_, err := client.Environments.DeleteEnvironment(params, auth)
			CheckErr(err)
		},
	}
	return cmd
}

func init() {
	deleteCmd.AddCommand(NewDeleteEnvironmentCommand())
}
