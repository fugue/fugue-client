package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/fugue/fugue-client/client/rule_waivers"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type listRuleWaiversOptions struct {
	Columns                []string
	SearchQuery            string
	IDFilter               string
	EnvironmentIDFilter    string
	NameFilter             string
	RuleIDFilter           string
	ResourceIDFilter       string
	ResourceProviderFilter string
	ResourceTypeFilter     string
	ResourceTagFilter      string
	Offset                 int64
	MaxItems               int64
	OrderBy                string
	OrderDirection         string
	FetchAll               bool
}

type listRuleWaiversViewItem struct {
	ID                   string
	Name                 string
	Comment              string
	EnvironmentID        string
	EnvironmentName      string
	RuleID               string
	ResourceID           string
	ResourceProvider     string
	ResourceType         string
	ResourceTag          string
	CreatedAt            string
	CreatedBy            string
	CreatedByDisplayName string
	UpdatedAt            string
	UpdatedBy            string
	UpdatedByDisplayName string
}

// NewListRuleWaiversCommand returns a command that lists rule waivers in Fugue
func NewListRuleWaiversCommand() *cobra.Command {

	var opts listRuleWaiversOptions

	cmd := &cobra.Command{
		Use:     "rule-waivers",
		Short:   "Lists details for multiple rule waivers",
		Long:    `Lists details for multiple rule waivers`,
		Aliases: []string{"waivers", "rule_waivers"},
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

			// Search query filter
			if opts.SearchQuery != "" {
				searchParams = append(searchParams, opts.SearchQuery)
			}

			// Optionally filter by id (substring match, case insensitive)
			if opts.IDFilter != "" {
				searchParams = append(searchParams, fmt.Sprintf("id:%s", opts.IDFilter))
			}

			// Optionally filter by environment id (substring match, case insensitive)
			if opts.EnvironmentIDFilter != "" {
				searchParams = append(searchParams, fmt.Sprintf("environment_id:%s", opts.EnvironmentIDFilter))
			}

			// Optionally filter by name (substring match, case insensitive)
			if opts.NameFilter != "" {
				searchParams = append(searchParams, fmt.Sprintf("name:%s", opts.NameFilter))
			}

			// Optionally filter by rule id (substring match, case insensitive)
			if opts.RuleIDFilter != "" {
				searchParams = append(searchParams, fmt.Sprintf("rule_id:%s", opts.RuleIDFilter))
			}

			// Optionally filter by resource id (substring match, case insensitive)
			if opts.ResourceIDFilter != "" {
				searchParams = append(searchParams, fmt.Sprintf("resource_id:%s", opts.ResourceIDFilter))
			}

			// Optionally filter by resource type (substring match, case insensitive)
			if opts.ResourceTypeFilter != "" {
				searchParams = append(searchParams, fmt.Sprintf("resource_type:%s", opts.ResourceTypeFilter))
			}

			// Optionally filter by resource provider (substring match, case insensitive)
			if opts.ResourceProviderFilter != "" {
				searchParams = append(searchParams, fmt.Sprintf("resource_provider:%s", opts.ResourceProviderFilter))
			}

			var waivers []*models.RuleWaiver
			offset := opts.Offset
			for {
				params := rule_waivers.NewListRuleWaiversParams()
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
				resp, err := client.RuleWaivers.ListRuleWaivers(params, auth)
				CheckErr(err)
				waivers = append(waivers, resp.Payload.Items...)
				if opts.FetchAll && resp.Payload.IsTruncated {
					offset = resp.Payload.NextOffset
					continue
				}
				break
			}

			var rows []interface{}
			for _, waiver := range waivers {

				resourceTag := "-"
				if waiver.ResourceTag != "" {
					resourceTag = waiver.ResourceTag
				}
				row := listRuleWaiversViewItem{
					ID:                   *waiver.ID,
					Name:                 *waiver.Name,
					Comment:              waiver.Comment,
					EnvironmentID:        *waiver.EnvironmentID,
					EnvironmentName:      waiver.EnvironmentName,
					RuleID:               *waiver.RuleID,
					ResourceID:           *waiver.ResourceID,
					ResourceType:         *waiver.ResourceType,
					ResourceProvider:     *waiver.ResourceProvider,
					ResourceTag:          resourceTag,
					CreatedAt:            format.Unix(waiver.CreatedAt),
					CreatedBy:            waiver.CreatedBy,
					CreatedByDisplayName: waiver.CreatedByDisplayName,
					UpdatedAt:            format.Unix(waiver.UpdatedAt),
					UpdatedBy:            waiver.UpdatedBy,
					UpdatedByDisplayName: waiver.UpdatedByDisplayName,
				}
				rows = append(rows, row)

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
		"EnvironmentID",
		"EnvironmentName",
		"RuleID",
		"ResourceID",
		"ResourceType",
		"ResourceProvider",
		"ResourceTag",
	}

	cmd.Flags().StringVar(&opts.SearchQuery, "search", "", "Combined filter for ID, Name, and Rule ID")
	cmd.Flags().StringVar(&opts.IDFilter, "id", "", "ID filter (substring match, including provider account identifiers)")
	cmd.Flags().StringVar(&opts.NameFilter, "name", "", "Name filter (substring match, case insensitive)")
	cmd.Flags().StringVar(&opts.EnvironmentIDFilter, "environment-id", "", "Environment ID filter (substring match)")
	cmd.Flags().StringVar(&opts.RuleIDFilter, "rule-id", "", "Rule ID filter (substring match)")
	cmd.Flags().StringVar(&opts.ResourceIDFilter, "resource-id", "", "Resource ID filter (substring match)")
	cmd.Flags().StringVar(&opts.ResourceTypeFilter, "resource-type", "", "Resource Type filter (substring match)")
	cmd.Flags().StringVar(&opts.ResourceProviderFilter, "resource-provider", "", "Resource Provider filter (substring match)")
	cmd.Flags().StringSliceVar(&opts.Columns, "columns", defaultCols, "Columns to show")
	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "Offset into results")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 100, "Max items to return")
	cmd.Flags().StringVar(&opts.OrderBy, "order-by", "name", "Order by attribute")
	cmd.Flags().StringVar(&opts.OrderDirection, "order-direction", "asc", "Order by direction [asc | desc]")
	cmd.Flags().BoolVar(&opts.FetchAll, "all", false, "Retrieve all environments")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListRuleWaiversCommand())
}
