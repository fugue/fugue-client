package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/users"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type listUsersOptions struct {
	Offset         int64
	MaxItems       int64
	OrderDirection string
	FetchAll       bool
	Email          string
}

type listUsersViewItem struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Status    string
	Groups    string
	Owner     bool
}

// NewListUsersCommand returns a command that lists users in Fugue
func NewListUsersCommand() *cobra.Command {

	var opts listUsersOptions

	cmd := &cobra.Command{
		Use:   "users",
		Short: "Lists details for multiple users",
		Long:  `Lists details for multiple users`,
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			if opts.MaxItems < 1 {
				opts.MaxItems = 1
			}
			if opts.MaxItems > 100 || opts.FetchAll {
				opts.MaxItems = 100
			}
			if opts.Offset < 0 {
				opts.Offset = 0
			}

			var usersList []*models.User
			offset := opts.Offset
			for {
				params := users.NewListUsersParams()
				params.Offset = &offset
				params.MaxItems = &opts.MaxItems
				if opts.OrderDirection != "" {
					params.OrderDirection = &opts.OrderDirection
				}
				params.Email = &opts.Email
				resp, err := client.Users.ListUsers(params, auth)
				CheckErr(err)
				for _, user := range resp.Payload.Items {
					usersList = append(usersList, user)
				}
				if opts.FetchAll && resp.Payload.IsTruncated {
					offset = resp.Payload.NextOffset
					continue
				}
				break
			}

			var rows []interface{}
			for _, user := range usersList {

				numGroups := len(user.Groups)
				var groups string
				if numGroups == 1 {
					var groupNames []string
					for _, name := range user.Groups {
						groupNames = append(groupNames, name)
					}

					groups = groupNames[0]
				} else {
					groups = fmt.Sprintf("%v groups", numGroups)
				}

				rows = append(rows, listUsersViewItem{
					ID:        *user.ID,
					Email:     *user.Email,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Groups:    groups,
					Owner:     user.Owner,
				})
			}

			table, err := format.Table(format.TableOpts{
				Rows:       rows,
				Columns:    []string{"ID", "Email", "FirstName", "LastName", "Groups", "Owner"},
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "Offset into results")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 100, "Max items to return")
	cmd.Flags().StringVar(&opts.OrderDirection, "order-direction", "asc", "Order by direction [asc | desc]")
	cmd.Flags().StringVar(&opts.Email, "email", "", "Retrieve users with a provided email address")
	cmd.Flags().BoolVar(&opts.FetchAll, "all", false, "Retrieve all users")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListUsersCommand())
}
