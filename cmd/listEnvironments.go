package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type listEnvironmentsOptions struct {
	Offset         int64
	MaxItems       int64
	OrderBy        string
	OrderDirection string
	Columns        []string
	Provider       string
	NameFilter     string
}

type listEnvironmentsViewItem struct {
	ID                 string
	Name               string
	Provider           string
	Region             string
	ScanInterval       string
	ScanStatus         string
	HasBaseline        bool
	ComplianceFamilies string
}

// NewListEnvironmentsCommand returns a command that lists environments in Fugue
func NewListEnvironmentsCommand() *cobra.Command {

	var opts listEnvironmentsOptions

	cmd := &cobra.Command{
		Use:     "environments",
		Short:   "Lists details for multiple environments",
		Long:    `Lists details for multiple environments`,
		Aliases: []string{"envs", "env"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := environments.NewListEnvironmentsParams()
			if opts.Offset > 0 {
				params.Offset = &opts.Offset
			}
			if opts.MaxItems > 0 {
				params.MaxItems = &opts.MaxItems
			}
			if opts.OrderBy != "" {
				params.OrderBy = &opts.OrderBy
			}
			if opts.OrderDirection != "" {
				params.OrderDirection = &opts.OrderDirection
			}

			resp, err := client.Environments.ListEnvironments(params, auth)
			CheckErr(err)

			environments := resp.Payload.Items

			sort.Slice(environments, func(i, j int) bool {
				return environments[i].Name < environments[j].Name
			})

			nameFilter := strings.ToLower(opts.NameFilter)

			var rows []interface{}
			for _, env := range environments {

				// Optionally filter by provider
				if opts.Provider != "" {
					if env.Provider != opts.Provider {
						continue
					}
				}

				// Optionally filter by name (substring match, case insensitive)
				if nameFilter != "" {
					if !strings.Contains(strings.ToLower(env.Name), nameFilter) {
						continue
					}
				}

				region := "-"
				if env.Provider == "aws" {
					region = env.ProviderOptions.Aws.Region
				} else if env.Provider == "aws_govcloud" {
					region = env.ProviderOptions.AwsGovcloud.Region
				}

				rows = append(rows, listEnvironmentsViewItem{
					ID:                 env.ID,
					Name:               env.Name,
					HasBaseline:        env.BaselineID != "",
					Provider:           env.Provider,
					Region:             region,
					ScanInterval:       strconv.FormatInt(env.ScanInterval, 10),
					ScanStatus:         env.ScanStatus,
					ComplianceFamilies: strings.Join(env.ComplianceFamilies, ","),
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
		"Provider",
		"Region",
		"HasBaseline",
		"ScanInterval",
		"ScanStatus",
	}

	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "offset into results")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 0, "max items to return")
	cmd.Flags().StringVar(&opts.OrderBy, "order-by", "", "order by attribute")
	cmd.Flags().StringVar(&opts.OrderDirection, "order-direction", "", "order by direction [asc | desc]")
	cmd.Flags().StringVar(&opts.Provider, "provider", "", "Provider filter")
	cmd.Flags().StringVar(&opts.NameFilter, "name", "", "Name filter (substring match, case insensitive)")
	cmd.Flags().StringSliceVar(&opts.Columns, "columns", defaultCols, "columns to show")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListEnvironmentsCommand())
}
