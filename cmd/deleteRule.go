package cmd

import (
	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/spf13/cobra"
)

// NewDeleteRuleCommand returns a command that deletes a custom rule
func NewDeleteRuleCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "rule [rule_id]",
		Short: "Deletes a custom rule",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := custom_rules.NewDeleteCustomRuleParams()
			params.RuleID = args[0]

			_, err := client.CustomRules.DeleteCustomRule(params, auth)
			CheckErr(err)
		},
	}
	return cmd
}

func init() {
	deleteCmd.AddCommand(NewDeleteRuleCommand())
}
