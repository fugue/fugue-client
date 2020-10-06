package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	if url, ok := testOutput.Links["output"]; ok {
		// Download temporary file and perform conversions.
		if tmpResult, err := getTmpResult(url); err == nil {
			testOutput.Errors = make([]*models.CustomRuleError, 0)
			for _, tmpError := range tmpResult.Errors {
				testOutput.Errors = append(testOutput.Errors, fromTmpError(tmpError))
			}
			if tmpResult.Report != nil {
				testOutput.Resources = make([]*models.TestCustomRuleOutputResource, 0)
				for _, tmpResource := range tmpResult.Report.Resources {
					testOutput.Resources = append(testOutput.Resources, fromTmpResource(tmpResource))
				}
			}
		} else {
			return err
		}
	}

	// Check if we need to do a trip through `links["output"]`.
	// if url, ok := resp.Payload.Links["output"]; ok {
	// result, err := getTmpResult(url)
	// }

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

	if url, ok := resp.Payload.Links["output"]; ok {
		// If the response is too big, it is provided as a download link.
		if downloadResult, err := getTmpResult(url); err == nil {
			resp.Payload = downloadResult.Input
		} else {
			return err
		}
	}

	bytes, err := json.MarshalIndent(resp.Payload, "", "    ")
	if err != nil {
		return err
	}
	os.Stdout.Write(bytes)
	fmt.Println("") // Add final newline after JSON.
	return nil
}

// tmpResult describes the scheme that we download from the
// `links["output"]` key, which is slightly different than what we usually
// expect, so we need to do some conversions here.
type tmpResult struct {
	Input  *models.TestCustomRuleInputScan `json:"input,omitempty"`
	Errors []*tmpError                     `json:"errors,omitempty"`
	Report *tmpReport                      `json:"report,omitempty"`
}

type tmpReport struct {
	Resources []*tmpResource `json:"resources"`
}

type tmpError struct {
	Text     string `json:"_text,omitempty"`
	Severity string `json:"severity,omitempty"`
}

type tmpResource struct {
	ID    string `json:"id,omitempty"`
	Valid *bool  `json:"valid,omitempty"`
	Type  string `json:"type,omitempty"`
}

// fromTmpError converts a tmpError to a models.CustomRuleError
func fromTmpError(tmp *tmpError) *models.CustomRuleError {
	return &models.CustomRuleError{
		Text:     tmp.Text,
		Severity: tmp.Severity,
	}
}

// fromTmpResource convers a tmpResource to a models.TestCustomRuleOutputResource
func fromTmpResource(tmp *tmpResource) *models.TestCustomRuleOutputResource {
	resource := models.TestCustomRuleOutputResource{
		ID:   tmp.ID,
		Type: tmp.Type,
	}
	if tmp.Valid == nil {
		resource.Result = models.TestCustomRuleOutputResourceResultUNKNOWN
	} else if *tmp.Valid {
		resource.Result = models.TestCustomRuleOutputResourceResultPASS
	} else {
		resource.Result = models.TestCustomRuleOutputResultFAIL
	}
	return &resource
}

func getTmpResult(url string) (*tmpResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result tmpResult
	bytes, err := ioutil.ReadAll(resp.Body)
	print(string(bytes))
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &result)
	return &result, nil
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
