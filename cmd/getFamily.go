package cmd

import (
	"fmt"
	"strings"

	"github.com/fugue/fugue-client/client/families"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

// NewGetFamilyCommand returns a command that retrieves custom family details
func NewGetFamilyCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "family [family_id]",
		Short: "Retrieve details for a family",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := families.NewGetFamilyParams()
			params.FamilyID = args[0]

			resp, err := client.Families.GetFamily(params, auth)
			CheckErr(err)

			family := resp.Payload

			providers := strings.Join(family.Providers[:], ",")
			ruleIDs := strings.Join(family.RuleIds[:], ",")

			items := []interface{}{
				Item{"NAME", family.Name},
				Item{"SOURCE", family.Source},
				Item{"DESCRIPTION", family.Description},
				Item{"PROVIDERS", providers},
				Item{"RECOMMENDED", family.Recommended},
				Item{"RULE IDS", ruleIDs},
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

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetFamilyCommand())
}
