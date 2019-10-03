package cmd

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync files to your account",
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
