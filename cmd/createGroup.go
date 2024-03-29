package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/groups"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createGroupOptions struct {
	Name            string
	EnvironmentIds  []string
	Policy          string
	AllEnvironments bool
}

// NewCreateGroupCommand returns a command that creates an group
func NewCreateGroupCommand() *cobra.Command {

	var opts createGroupOptions

	cmd := &cobra.Command{
		Use:   "group",
		Short: "Create a group",
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := groups.NewCreateGroupParams()

			var environmentIDs []string

			if opts.AllEnvironments {
				environmentIDs = []string{"*"}

			} else {
				environmentIDs = opts.EnvironmentIds
			}

			params.Group = &models.CreateGroupInput{
				Name:           opts.Name,
				EnvironmentIds: environmentIDs,
				Policy:         opts.Policy,
			}

			resp, err := client.Groups.CreateGroup(params, auth)

			if err != nil {
				switch respError := err.(type) {
				case *groups.CreateGroupInternalServerError:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			group := resp.Payload

			var environments []string
			for key, value := range group.Environments {
				environments = append(environments, fmt.Sprintf("%s:%s", key, value))
			}

			items := []interface{}{
				Item{"GROUP_ID", group.ID},
				Item{"NAME", group.Name},
				Item{"POLICY", group.Policy},
				Item{"ENVIRONMENTS", strings.Join(environments, ", ")},
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

	cmd.Flags().StringVar(&opts.Name, "name", "", "Group name")
	cmd.Flags().StringVar(&opts.Policy, "policy", "", "Fugue policy to use for the group")
	cmd.Flags().StringSliceVar(&opts.EnvironmentIds, "environment-ids", []string{}, "Environments which this group should be able to access using the provided policy")
	cmd.Flags().BoolVar(&opts.AllEnvironments, "all-environments", false, "Indicates that the group should be created with all environments (current and future) attached")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("policy")

	return cmd
}

func init() {
	createCmd.AddCommand(NewCreateGroupCommand())
}
