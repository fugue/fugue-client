package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fugue/fugue-client/client/families"

	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type updateFamilyOptions struct {
	Name        string
	Description string
	Recommended bool
	RuleIDs     []string
}

// NewUpdateFamilyCommand returns a command that updates a custom family
func NewUpdateFamilyCommand() *cobra.Command {

	var opts updateFamilyOptions

	cmd := &cobra.Command{
		Use:   "family [family_id]",
		Short: "Update family settings",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := families.NewUpdateFamilyParams()
			params.FamilyID = args[0]
			params.Family = &models.UpdateFamilyInput{}

			flagCount := 0

			// Using Visit here allows us to process only flags that were set
			cmd.Flags().Visit(func(f *pflag.Flag) {
				flagCount++
				switch f.Name {
				case "name":
					params.Family.Name = opts.Name
				case "description":
					params.Family.Description = opts.Description
				case "recommended":
					params.Family.Recommended = &opts.Recommended
				case "rule-ids":
					params.Family.RuleIds = opts.RuleIDs
				}
			})

			if flagCount == 0 {
				os.Exit(0)
			}

			resp, err := client.Families.UpdateFamily(params, auth)
			CheckErr(err)

			family := resp.Payload

			var itemRules Item
			if len(family.RuleIds) > 0 {
				itemRules = Item{"RULE_IDS", strings.Join(family.RuleIds[:], ", ")}
			} else {
				itemRules = Item{"RULE_IDS", "-"}
			}
			var itemProviders Item
			if len(family.Providers) > 0 {
				itemProviders = Item{"PROVIDERS", strings.Join(family.Providers[:], ", ")}
			} else {
				itemProviders = Item{"PROVIDERS", "-"}
			}

			items := []interface{}{
				Item{"NAME", family.Name},
				Item{"DESCRIPTION", family.Description},
				itemProviders,
				Item{"RECOMMENDED", family.Recommended},
				itemRules,
				Item{"CREATED_AT", format.Unix(family.CreatedAt)},
				Item{"CREATED_BY", family.CreatedBy},
				Item{"CREATED_BY_DISPLAY_NAME", family.CreatedByDisplayName},
				Item{"UPDATED_AT", format.Unix(family.UpdatedAt)},
				Item{"UPDATED_BY", family.UpdatedBy},
				Item{"UPDATED_BY_DISPLAY_NAME", family.UpdatedByDisplayName},
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

	cmd.Flags().StringVar(&opts.Name, "name", "", "Family name")
	cmd.Flags().StringVar(&opts.Description, "description", "", "Description")
	cmd.Flags().BoolVar(&opts.Recommended, "recommended", true, "If the family is recommended for all new environments")
	cmd.Flags().StringSliceVar(&opts.RuleIDs, "rule-ids", []string{}, "List of rule IDs to associate with the family (e.g. FG_R00217,<UUID Custom Rule ID>)")

	return cmd
}

func init() {
	updateCmd.AddCommand(NewUpdateFamilyCommand())
}
