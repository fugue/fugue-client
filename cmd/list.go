package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List a collection of resources",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
