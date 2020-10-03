package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/spf13/cobra"
)

type testRuleOptions struct {
	ScanID       string
	ResourceType string
	Input        bool
}

type testRuleResource struct {
	ID     string
	Result string
	Type   string
}

// testRule does the actual work of testing the rule.  The command will delegate
// to either this or `getRuleInput`.
func testRule(opts testRuleOptions, regoFile string) error {
	client, auth := getClient()

	regoBytes, err := ioutil.ReadFile(regoFile)
	if err != nil {
		return err
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
		return err
	}

	testOutput := resp.Payload
	if len(testOutput.Errors) > 0 {
		for _, regoError := range testOutput.Errors {
			fmt.Println(regoError.Text)
		}
		return nil
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
	if err != nil {
		return err
	}

	for _, tableRow := range table {
		fmt.Println(tableRow)
	}
	return nil
}

// getRuleInput retrieves the rule input rather than running the rule.
func getRuleInput(opts testRuleOptions) error {
	client, auth := getClient()

	params := custom_rules.NewTestCustomRuleInputParams()
	params.ScanID = opts.ScanID

	resp, err := client.CustomRules.TestCustomRuleInput(params, auth)
	if err != nil {
		return err
	}

	// TODO: we need to take the new `links` feature into account here.
	bytes, err := json.MarshalIndent(resp.Payload, "", "    ")
	if err != nil {
		return err
	}

	os.Stdout.Write(bytes)
	fmt.Println("") // Add final newline after JSON.
	return nil
}

func NewTestRuleCommand() *cobra.Command {
	var opts testRuleOptions

	cmd := &cobra.Command{
		Use:   "rule [rego file]",
		Short: "Test a custom rule",
		Args: func(cmd *cobra.Command, args []string) error {
			if opts.Input {
				return cobra.ExactArgs(0)(cmd, args)
			} else {
				return cobra.ExactArgs(1)(cmd, args)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if opts.Input {
				CheckErr(getRuleInput(opts))
			} else {
				CheckErr(testRule(opts, args[0]))
			}
		},
	}

	cmd.Flags().StringVar(&opts.ScanID, "scan", "", "Scan ID")
	cmd.Flags().StringVar(&opts.ResourceType, "resource-type", "", "Resource type")
	cmd.Flags().BoolVar(&opts.Input, "input", false, "Retrieve rule input")

	cmd.MarkFlagRequired("scan")

	return cmd
}

func init() {
	testCmd.AddCommand(NewTestRuleCommand())
}
