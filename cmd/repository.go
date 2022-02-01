package cmd

import (
	"github.com/spf13/cobra"
)

var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Repository subcommands",
}

func init() {
	createCmd.AddCommand(repositoryCmd)
}
