package cmd

import (
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch directory for changes",
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
