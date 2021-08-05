package cmd

import (
	"github.com/fugue/fugue-client/client/families"
	"github.com/spf13/cobra"
)

// NewDeleteFamilyCommand returns a command that deletes a family
func NewDeleteFamilyCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "family [family_id]",
		Short: "Deletes a family",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := families.NewDeleteFamilyParams()
			params.FamilyID = args[0]

			_, err := client.Families.DeleteFamily(params, auth)
			CheckErr(err)
		},
	}
	return cmd
}

func init() {
	deleteCmd.AddCommand(NewDeleteFamilyCommand())
}
