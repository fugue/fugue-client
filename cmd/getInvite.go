package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/invites"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

// NewGetInviteCommand returns a command that retrieves invite details
func NewGetInviteCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "invite [invite_id]",
		Short: "Retrieve details for a invite",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// When printing json gets the first (pos 0) API call printed
			jsonPositionToShow = 0
			client, auth := getClient()

			inviteID := args[0]

			params := invites.NewGetInviteByIDParams()
			params.InviteID = inviteID

			resp, err := client.Invites.GetInviteByID(params, auth)

			if err != nil {
				switch respError := err.(type) {
				case *invites.GetInviteByIDNotFound:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			invite := resp.Payload

			var groups []string
			for key, value := range invite.Groups {
				groups = append(groups, fmt.Sprintf("%s:%s", key, value))
			}

			var createdAt int64
			if invite.CreatedAt != nil {
				createdAt = *invite.CreatedAt
			}

			var expiresAt int64
			if invite.ExpiresAt != nil {
				expiresAt = *invite.ExpiresAt
			}

			items := []interface{}{
				Item{"INVITE_ID", *invite.ID},
				Item{"EMAIL", *invite.Email},
				Item{"GROUPS", strings.Join(groups, ", ")},
				Item{"STATUS", *invite.Status},
				Item{"CREATED_AT", format.Unix(createdAt)},
				Item{"UPDATED_AT", format.Unix(invite.UpdatedAt)},
				Item{"EXPIRES_AT", format.Unix(expiresAt)},
				Item{"RESOURCE_TYPE", invite.ResourceType},
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
	getCmd.AddCommand(NewGetInviteCommand())
}
