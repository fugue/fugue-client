package cmd

import (
	"fmt"
	"github.com/fugue/fugue-client/client/groups"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type updateUsersGroupsOptions struct {
	UserIds					 []string
	GroupIds					 []string
}


// NewUpdateUsersGroups returns a command that allows updating groups for multiple users
func NewUpdateUsersGroups() *cobra.Command {

	var opts updateUsersGroupsOptions

	cmd := &cobra.Command{
		Use:     "users_groups",
		Short:   "Batch update users group assignments",
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := groups.NewEditUsersGroupAssignmentsParams()
			
			params.EditGroupAssignments = &models.EditUsersGroupAssignmentsInput{
				UserIds: opts.UserIds,
				GroupIds: opts.GroupIds,
			}

			_, err := client.Groups.EditUsersGroupAssignments(params, auth)


			if err != nil {
				switch respError := err.(type) {
				case *groups.EditUsersGroupAssignmentsInternalServerError:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			fmt.Println("Successfully updated user(s) group assignments")
		},
	}

	cmd.Flags().StringSliceVar(&opts.UserIds, "user-ids", []string{}, "Users to update")
	cmd.Flags().StringSliceVar(&opts.GroupIds, "group-ids", []string{}, "Groups to assign to provided users")

	cmd.MarkFlagRequired("group-ids")
	cmd.MarkFlagRequired("group-ids")

	return cmd
}

func init() {
	updateCmd.AddCommand(NewUpdateUsersGroups())
}
