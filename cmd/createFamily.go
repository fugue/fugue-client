package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/families"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type createFamilyOptions struct {
	Name        string
	Description string
	Recommended bool
	RuleIDs     []string
}

// NewCreateFamilyCommand returns a command that creates a family
func NewCreateFamilyCommand() *cobra.Command {

	var opts createFamilyOptions

	cmd := &cobra.Command{
		Use:   "family",
		Short: "Create a family",
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := families.NewCreateFamilyParams()
			params.Family = &models.CreateFamilyInput{
				Name:        opts.Name,
				Description: opts.Description,
				Recommended: opts.Recommended,
				RuleIds:     opts.RuleIDs,
			}

			resp, err := client.Families.CreateFamily(params, auth)
			if err != nil {
				switch respError := err.(type) {
				case *families.CreateFamilyInternalServerError:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			family := resp.Payload

			var item Item
			if len(family.RuleIds) > 0 {
				item = Item{"RULE_IDS", family.RuleIds}
			} else {
				item = Item{"RULE_IDS", "-"}
			}

			items := []interface{}{
				Item{"FAMILY_ID", family.ID},
				Item{"NAME", family.Name},
				Item{"DESCRIPTION", family.Description},
				Item{"RECOMMENDED", family.Recommended},
				item,
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

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("description")

	return cmd
}

func init() {
	createCmd.AddCommand(NewCreateFamilyCommand())
}
