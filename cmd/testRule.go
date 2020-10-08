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

	viaDownload := true
	params := custom_rules.NewTestCustomRuleParams()
	params.ViaDownload = &viaDownload
	params.Rule = &models.TestCustomRuleInput{
		ScanID:       &opts.ScanID,
		ResourceType: opts.ResourceType,
		RuleText:     &regoText,
	}

	resp, err := client.CustomRules.TestCustomRule(params, auth)
	if err != nil {
		return err
	}

	// Download temporary file and perform conversions.
	downloadResult, err := getDownloadResult(resp.Payload.Links)
	if err != nil {
		return err
	}

	if len(downloadResult.Errors) > 0 {
		for _, regoError := range downloadResult.Errors {
			fmt.Println(regoError.Text)
		}
		return nil
	}

	var rows []interface{}
	if downloadResult.Report != nil {
		for _, resource := range downloadResult.Report.Resources {
			rows = append(rows, testRuleResource{
				ID:     resource.ID,
				Result: validToString(resource.Valid),
				Type:   resource.Type,
			})
		}
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

	viaDownload := true
	params := custom_rules.NewTestCustomRuleInputParams()
	params.ScanID = opts.ScanID
	params.ViaDownload = &viaDownload

	resp, err := client.CustomRules.TestCustomRuleInput(params, auth)
	if err != nil {
		return err
	}

	downloadResult, err := getDownloadResult(resp.Payload.Links)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(downloadResult.Input, "", "    ")
	if err != nil {
		return err
	}
	os.Stdout.Write(bytes)
	fmt.Println("") // Add final newline after JSON.
	return nil
}

// downloadResult describes the scheme that we download from the
// `links["output"]` key, which is slightly different than what we usually
// expect, so we need to do some conversions here.
type downloadResult struct {
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

// fromDownloadResource convers a tmpResource to a models.TestCustomRuleOutputResource
func validToString(valid *bool) string {
	if valid == nil {
		return models.TestCustomRuleOutputResourceResultUNKNOWN
	} else if *valid {
		return models.TestCustomRuleOutputResourceResultPASS
	} else {
		return models.TestCustomRuleOutputResultFAIL
	}
}

func getDownloadResult(links map[string]string) (*downloadResult, error) {
	url, ok := links["output"]
	if !ok {
		return nil, fmt.Errorf("missing 'output' link in result")
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result downloadResult
	bytes, err := ioutil.ReadAll(resp.Body)
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
