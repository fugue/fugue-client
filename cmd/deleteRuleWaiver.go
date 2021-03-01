package cmd

import (
	"github.com/fugue/fugue-client/client/rule_waivers"
	"github.com/spf13/cobra"
)

// NewDeleteRuleWaiverCommand returns a command that deletes a custom rule
func NewDeleteRuleWaiverCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "rule-waiver [rule_waiver_id]",
		Short:   "Deletes a rule waiver",
		Aliases: []string{"waiver", "rule_waiver"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := rule_waivers.NewDeleteRuleWaiverParams()
			params.RuleWaiverID = args[0]

			_, err := client.RuleWaivers.DeleteRuleWaiver(params, auth)
			CheckErr(err)
		},
	}
	return cmd
}

func init() {
	deleteCmd.AddCommand(NewDeleteRuleWaiverCommand())
}
