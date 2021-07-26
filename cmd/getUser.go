package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/users"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

// NewGetUserCommand returns a command that retrieves user details
func NewGetUserCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user [user_id]",
		Short: "Retrieve details for a user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// When printing json gets the first (pos 0) API call printed
			jsonPositionToShow = 0
			client, auth := getClient()

			userID := args[0]
			params := users.NewGetUserByIDParams()
			params.UserID = userID

			resp, err := client.Users.GetUserByID(params, auth)

			if err != nil {
				switch respError := err.(type) {
				case *users.GetUserByIDNotFound:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			user := resp.Payload

			var groups []string
			for key, value := range user.Groups {
				groups = append(groups, fmt.Sprintf("%s:%s", key, value))
			}

			items := []interface{}{
				Item{"USER_ID", *user.ID},
				Item{"EMAIL", *user.Email},
				Item{"FIRST_NAME", user.FirstName},
				Item{"LAST_NAME", user.LastName},
				Item{"OWNER", user.Owner},
				Item{"GROUPS", strings.Join(groups, ", ")},
				Item{"STATUS", *user.Status},
				Item{"RESOURCE_TYPE", user.ResourceType},
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

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetUserCommand())
}
