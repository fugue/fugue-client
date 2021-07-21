package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/invites"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type listInvitesOptions struct {
	Offset         int64
	MaxItems       int64
	OrderDirection string
	FetchAll       bool
	Email          string
}

type listInvitesViewItem struct {
	ID        string
	Email     string
	Status    string
	CreatedAt string
	ExpiresAt string
	Groups    string
}

// NewListInvitesCommand returns a command that lists invites in Fugue
func NewListInvitesCommand() *cobra.Command {

	var opts listInvitesOptions

	cmd := &cobra.Command{
		Use:   "invites",
		Short: "Lists details for multiple invites",
		Long:  `Lists details for multiple invites`,
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

			var invitesList []*models.Invite
			offset := opts.Offset
			for {
				params := invites.NewListInvitesParams()
				params.Offset = &offset
				params.MaxItems = &opts.MaxItems
				if opts.OrderDirection != "" {
					params.OrderDirection = &opts.OrderDirection
				}
				params.Email = &opts.Email
				resp, err := client.Invites.ListInvites(params, auth)
				CheckErr(err)
				invitesList = append(invitesList, resp.Payload.Items...)
				if opts.FetchAll && resp.Payload.IsTruncated {
					offset = resp.Payload.NextOffset
					continue
				}
				break
			}

			var rows []interface{}
			for _, invite := range invitesList {

				numGroups := len(invite.Groups)
				var groups string
				if numGroups == 1 {
					var groupNames []string
					for _, name := range invite.Groups {
						groupNames = append(groupNames, name)
					}

					groups = groupNames[0]
				} else {
					groups = fmt.Sprintf("%v groups", numGroups)
				}

				var createdAt int64
				if invite.CreatedAt != nil {
					createdAt = *invite.CreatedAt
				}

				var expiresAt int64
				if invite.ExpiresAt != nil {
					expiresAt = *invite.ExpiresAt
				}

				rows = append(rows, listInvitesViewItem{
					ID:        *invite.ID,
					Email:     *invite.Email,
					Status:    *invite.Status,
					CreatedAt: format.Unix(createdAt),
					ExpiresAt: format.Unix(expiresAt),
					Groups:    groups,
				})
			}

			table, err := format.Table(format.TableOpts{
				Rows:       rows,
				Columns:    []string{"ID", "Email", "Groups", "Status", "CreatedAt", "ExpiresAt"},
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
	cmd.Flags().StringVar(&opts.Email, "email", "", "Retrieve invites with a provided email address")
	cmd.Flags().BoolVar(&opts.FetchAll, "all", false, "Retrieve all invites")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListInvitesCommand())
}
