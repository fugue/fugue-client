package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test custom rules",
}

func init() {
	rootCmd.AddCommand(testCmd)
}
