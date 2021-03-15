package cmd

import (
	"github.com/spf13/cobra"
)

var googleCmd = &cobra.Command{
	Use:   "google",
	Short: "Google subcommands",
}

func init() {
	createCmd.AddCommand(googleCmd)
}
