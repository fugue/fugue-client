package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type testRuleOptions struct {
	ScanID       string
	ResourceType string
}

type testRuleResource struct {
	ID     string
	Result string
	Type   string
}

func NewTestRuleCommand() *cobra.Command {
	var opts testRuleOptions

	cmd := &cobra.Command{
		Use:   "rule [rego file]",
		Short: "Test a custom rule",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			regoFile := args[0]
			regoBytes, err := ioutil.ReadFile(regoFile)
			if err != nil {
				CheckErr(err)
			}
			regoText := string(regoBytes)

			params := custom_rules.NewTestCustomRuleParams()
			params.Rule = &models.TestCustomRuleInput{
				ScanID:       &opts.ScanID,
				ResourceType: opts.ResourceType,
				RuleText:     &regoText,
			}

			resp, err := client.CustomRules.TestCustomRule(params, auth)
			if err != nil {
				CheckErr(err)
			}

			testOutput := resp.Payload
			if len(testOutput.Errors) > 0 {
				for _, regoError := range testOutput.Errors {
					fmt.Println(regoError.Text)
				}
				return
			}

			var rows []interface{}
			for _, resource := range testOutput.Resources {
				rows = append(rows, testRuleResource{
					ID:     resource.ID,
					Result: resource.Result,
					Type:   resource.Type,
				})
			}

			table, err := format.Table(format.TableOpts{
				Rows:       rows,
				Columns:    []string{"ID", "Result", "Type"},
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().StringVar(&opts.ScanID, "scan", "", "Scan ID")
	cmd.Flags().StringVar(&opts.ResourceType, "resource-type", "", "Resource type")

	cmd.MarkFlagRequired("scan")
	cmd.MarkFlagRequired("resource-type")

	return cmd
}

func init() {
	testCmd.AddCommand(NewTestRuleCommand())
}
