package cmd

import (
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS subcommands",
}

func init() {
	createCmd.AddCommand(awsCmd)
}
