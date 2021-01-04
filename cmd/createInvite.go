package cmd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/invites"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createInviteOptions struct {
	Email                     string
	GroupIds					 []string
	Expires                   bool
}


// NewCreateInviteCommand returns a command that creates an invite
func NewCreateInviteCommand() *cobra.Command {

	var opts createInviteOptions

	cmd := &cobra.Command{
		Use:     "invite",
		Short:   "Create an invite",
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := invites.NewCreateInviteParams()

			params.Invite = &models.CreateInviteInput{
				Email: &opts.Email,
				GroupIds: opts.GroupIds,
				Expires: &opts.Expires,
			}

			resp, err := client.Invites.CreateInvite(params, auth)

			buf := new(bytes.Buffer)
			fmt.Println(buf)
			if err != nil {
				switch respError := err.(type) {
				case *invites.CreateInviteInternalServerError:
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
				Rows:       items,
				Columns:    []string{"Attribute", "Value"},
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().StringVar(&opts.Email, "email", "", "Email")
	cmd.Flags().StringSliceVar(&opts.GroupIds, "group-ids", []string{}, "Groups to assign the user once they accept the issued invitation")
	cmd.Flags().BoolVar(&opts.Expires, "expires", true, "Indicates if the invite should expire")

	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("group-ids")

	return cmd
}

func init() {
	createCmd.AddCommand(NewCreateInviteCommand())
}
