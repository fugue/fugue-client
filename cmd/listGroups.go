package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/groups"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type listGroupsOptions struct {
	Offset         int64
	MaxItems       int64
	OrderBy        string
	OrderDirection string
	FetchAll       bool
}

type listGroupsViewItem struct {
	ID                 string
	Name               string
	Environments       string
	Users      		   int
	Policy             string
}

// NewListGroupsCommand returns a command that lists environments in Fugue
func NewListGroupsCommand() *cobra.Command {

	var opts listEnvironmentsOptions

	cmd := &cobra.Command{
		Use:     "groups",
		Short:   "Lists details for multiple groups",
		Long:    `Lists details for multiple groups`,
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			if opts.MaxItems < 1 {
				opts.MaxItems = 1
			}
			if opts.MaxItems > 10 || opts.FetchAll {
				opts.MaxItems = 10
			}
			if opts.Offset < 0 {
				opts.Offset = 0
			}

			var groupsList []*models.Group
			offset := opts.Offset
			for {
				params := groups.NewListGroupsParams()
				params.Offset = &offset
				params.MaxItems = &opts.MaxItems
				if opts.OrderBy != "" {
					params.OrderBy = &opts.OrderBy
				}
				if opts.OrderDirection != "" {
					params.OrderDirection = &opts.OrderDirection
				}
				resp, err := client.Groups.ListGroups(params, auth)
				CheckErr(err)
				for _, group := range resp.Payload.Items {
					groupsList = append(groupsList, group)
				}
				if opts.FetchAll && resp.Payload.IsTruncated {
					offset = resp.Payload.NextOffset
					continue
				}
				break
			}

			var rows []interface{}
			for _, group := range groupsList {
				numEnvironments := len(group.Environments)
				environments := "All"
				if(numEnvironments > 0) {
					environments = fmt.Sprintf("%v", numEnvironments);
				}
				rows = append(rows, listGroupsViewItem{
					ID:                 group.ID,
					Name:              	group.Name,
					Policy:				group.Policy,
					Environments: 		environments,
					Users:				len(group.Users),
				})
			}

			table, err := format.Table(format.TableOpts{
				Rows:       rows,
				Columns:    []string{"ID", "Name", "Policy", "Environments", "Users"},
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "Offset into results")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 10, "Max items to return")
	cmd.Flags().StringVar(&opts.OrderBy, "order-by", "name", "Order by attribute")
	cmd.Flags().StringVar(&opts.OrderDirection, "order-direction", "asc", "Order by direction [asc | desc]")
	cmd.Flags().BoolVar(&opts.FetchAll, "all", false, "Retrieve all groups")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListGroupsCommand())
}
