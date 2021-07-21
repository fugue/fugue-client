package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/families"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type listFamiliesOptions struct {
	Columns        []string
	Offset         int64
	MaxItems       int64
	OrderBy        string
	OrderDirection string
	FetchAll       bool
}

type listFamiliesViewItem struct {
	ID          string
	Name        string
	Source      string
	Description string
	Providers   string
	Recommended bool
	CreatedAt   string
	CreatedBy   string
	UpdatedAt   string
	UpdatedBy   string
}

// NewListFamiliesCommand returns a command that lists families in Fugue
func NewListFamiliesCommand() *cobra.Command {

	var opts listFamiliesOptions

	cmd := &cobra.Command{
		Use:     "families",
		Short:   "Lists details for multiple families",
		Long:    `Lists details for multiple families`,
		Aliases: []string{"fams", "fam"},
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

			searchParams := []string{}

			var fams []*models.FamilyList
			offset := opts.Offset
			for {

				params := families.NewListFamiliesParams()
				params.Offset = &offset
				params.MaxItems = &opts.MaxItems
				if opts.OrderBy != "" {
					params.OrderBy = &opts.OrderBy
				}
				if opts.OrderDirection != "" {
					params.OrderDirection = &opts.OrderDirection
				}

				if len(searchParams) > 0 {
					paramsJSON, _ := json.Marshal(searchParams)
					jsonString := string(paramsJSON)
					params.Query = &jsonString
				}
				resp, err := client.Families.ListFamilies(params, auth)
				CheckErr(err)

				fams = append(fams, resp.Payload.Items...)

				if opts.FetchAll && resp.Payload.IsTruncated {
					offset = resp.Payload.NextOffset
					continue
				}
				break
			}

			var rows []interface{}
			for _, family := range fams {

				description := family.Description
				if len(description) > 32 {
					description = description[:29] + "..."
				}
				providers := strings.Join(family.Providers[:], ",")
				if len(providers) > 32 {
					providers = providers[:29] + "..."
				}

				rows = append(rows, listFamiliesViewItem{

					ID:          family.ID,
					Name:        family.Name,
					Source:      family.Source,
					Description: description,
					Providers:   providers,
					Recommended: family.Recommended,
					CreatedAt:   format.Unix(family.CreatedAt),
					CreatedBy:   family.CreatedBy,
					UpdatedAt:   format.Unix(family.UpdatedAt),
					UpdatedBy:   family.UpdatedBy,
				})

			}

			table, err := format.Table(format.TableOpts{
				Rows:       rows,
				Columns:    opts.Columns,
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	defaultCols := []string{
		"ID",
		"Name",
		"Source",
		"Description",
		"Providers",
		"Recommended",
	}

	cmd.Flags().StringSliceVar(&opts.Columns, "columns", defaultCols, "Columns to show")
	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "Offset into results")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 100, "Max items to return")
	cmd.Flags().StringVar(&opts.OrderBy, "order-by", "name", "Order by attribute")
	cmd.Flags().StringVar(&opts.OrderDirection, "order-direction", "asc", "Order by direction [asc | desc]")
	cmd.Flags().BoolVar(&opts.FetchAll, "all", false, "Retrieve all families")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListFamiliesCommand())
}
