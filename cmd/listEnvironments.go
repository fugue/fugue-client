package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type listEnvironmentsOptions struct {
	Columns        []string
	Provider       string
	NameFilter     string
	Offset         int64
	MaxItems       int64
	OrderBy        string
	OrderDirection string
	FetchAll       bool
}

type listEnvironmentsViewItem struct {
	ID                 string
	Name               string
	Provider           string
	Region             string
	Regions            string
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

			if opts.MaxItems < 1 {
				opts.MaxItems = 1
			}
			if opts.MaxItems > 100 || opts.FetchAll {
				opts.MaxItems = 100
			}
			if opts.Offset < 0 {
				opts.Offset = 0
			}

			var envs []*models.Environment
			offset := opts.Offset
			for {
				params := environments.NewListEnvironmentsParams()
				params.Offset = &offset
				params.MaxItems = &opts.MaxItems
				if opts.OrderBy != "" {
					params.OrderBy = &opts.OrderBy
				}
				if opts.OrderDirection != "" {
					params.OrderDirection = &opts.OrderDirection
				}
				resp, err := client.Environments.ListEnvironments(params, auth)
				CheckErr(err)
				for _, env := range resp.Payload.Items {
					envs = append(envs, env)
				}
				if opts.FetchAll && resp.Payload.IsTruncated {
					offset = resp.Payload.NextOffset
					continue
				}
				break
			}

			nameFilter := strings.ToLower(opts.NameFilter)

			var rows []interface{}
			for _, env := range envs {

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
				var regionsTmp []string
				if env.Provider == "aws" {
					region = env.ProviderOptions.Aws.Region
					regionsTmp = env.ProviderOptions.Aws.Regions
				} else if env.Provider == "aws_govcloud" {
					region = env.ProviderOptions.AwsGovcloud.Region
					regionsTmp = env.ProviderOptions.AwsGovcloud.Regions
				}
				var regions string
				if len(regionsTmp) == 0 {
					regions = region
				} else if len(regionsTmp) > 2 {
					regions = strings.Join(regionsTmp[:2], ",") + "..."
				} else {
					regions = strings.Join(regionsTmp, ",")
				}

				rows = append(rows, listEnvironmentsViewItem{
					ID:                 env.ID,
					Name:               env.Name,
					HasBaseline:        env.BaselineID != "",
					Provider:           env.Provider,
					Region:             region,
					Regions:            regions,
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
		"Regions",
		"HasBaseline",
		"ScanInterval",
		"ScanStatus",
	}

	cmd.Flags().StringVar(&opts.Provider, "provider", "", "Provider filter")
	cmd.Flags().StringVar(&opts.NameFilter, "name", "", "Name filter (substring match, case insensitive)")
	cmd.Flags().StringSliceVar(&opts.Columns, "columns", defaultCols, "Columns to show")
	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "Offset into results")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 100, "Max items to return")
	cmd.Flags().StringVar(&opts.OrderBy, "order-by", "name", "Order by attribute")
	cmd.Flags().StringVar(&opts.OrderDirection, "order-direction", "asc", "Order by direction [asc | desc]")
	cmd.Flags().BoolVar(&opts.FetchAll, "all", false, "Retrieve all environments")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListEnvironmentsCommand())
}
