package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a resource",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
