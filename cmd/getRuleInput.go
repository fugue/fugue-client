package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/spf13/cobra"
)

type getRuleInputOptions struct {
	ScanID string
}

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

func NewGetRuleInputCommand() *cobra.Command {
	var opts testRuleOptions

	cmd := &cobra.Command{
		Use:   "rule-input",
		Short: "Retrieve rule input",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			CheckErr(getRuleInput(opts))
		},
	}

	cmd.Flags().StringVar(&opts.ScanID, "scan", "", "Scan ID")
	cmd.MarkFlagRequired("scan")

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetRuleInputCommand())
}
