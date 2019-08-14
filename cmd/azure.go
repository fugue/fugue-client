package cmd

import (
	"github.com/spf13/cobra"
)

var azureCmd = &cobra.Command{
	Use:   "azure",
	Short: "Azure subcommands",
}

func init() {
	createCmd.AddCommand(azureCmd)
}
