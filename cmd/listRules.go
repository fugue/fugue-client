package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type listRulesOptions struct {
	Columns []string
}

type listRulesViewItem struct {
	ID           string
	Name         string
	Description  string
	Provider     string
	ResourceType string
	RuleText     string
	Status       string
	CreatedAt    string
	CreatedBy    string
	UpdatedAt    string
	UpdatedBy    string
}

// NewListRulesCommand returns a command that lists custom rules in Fugue
func NewListRulesCommand() *cobra.Command {

	var opts listRulesOptions

	cmd := &cobra.Command{
		Use:   "rules",
		Short: "List rules in the organization",
		Long:  `List rules in the organization`,
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := custom_rules.NewListCustomRulesParams()
			resp, err := client.CustomRules.ListCustomRules(params, auth)
			CheckErr(err)

			rules := resp.Payload.Items

			var rows []interface{}
			for _, rule := range rules {
				description := rule.Description
				if len(description) > 32 {
					description = description[:29] + "..."
				}
				rows = append(rows, listRulesViewItem{
					ID:           rule.ID,
					Name:         rule.Name,
					Description:  description,
					Provider:     rule.Provider,
					ResourceType: rule.ResourceType,
					RuleText:     rule.RuleText,
					Status:       rule.Status,
					CreatedAt:    format.Unix(rule.CreatedAt),
					CreatedBy:    rule.CreatedBy,
					UpdatedAt:    format.Unix(rule.UpdatedAt),
					UpdatedBy:    rule.UpdatedBy,
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
		"ResourceType",
		"Status",
		"Description",
	}

	cmd.Flags().StringSliceVar(&opts.Columns, "columns", defaultCols, "Columns to show")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListRulesCommand())
}
