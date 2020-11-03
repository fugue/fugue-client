package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

// downloadResult describes the scheme that we download from the
// `links["output"]` key, which is slightly different than what we usually
// expect, so we need to do some conversions here.
//
// Note that these types are reused in `cmd/getRuleInput.go`.
type downloadResult struct {
	Input  *models.TestCustomRuleInputScan `json:"input,omitempty"`
	Errors []*downloadError                `json:"errors,omitempty"`
	Report *downloadReport                 `json:"report,omitempty"`
}

type downloadReport struct {
	Resources []*downloadResource `json:"resources"`
}

type downloadError struct {
	Text     string `json:"_text,omitempty"`
	Severity string `json:"severity,omitempty"`
}

type downloadResource struct {
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
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			CheckErr(testRule(opts, args[0]))
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
