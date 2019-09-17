package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/metadata"
	"github.com/spf13/cobra"
)

type getResourceTypesOptions struct {
	Provider string
	Region   string
}

// NewGetResourceTypesCommand returns a command that retrives available
// resource types for the given provider and region
func NewGetResourceTypesCommand() *cobra.Command {

	var opts getResourceTypesOptions

	cmd := &cobra.Command{
		Use:     "types",
		Short:   "List supported resource types",
		Aliases: []string{"resource-types"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := metadata.NewGetResourceTypesParams()
			params.Provider = opts.Provider

			if opts.Provider == "aws" || opts.Provider == "aws_govcloud" {
				if opts.Region == "" {
					Fatal("Must specify a region", 1)
				}
				params.Region = &opts.Region
			}

			resp, err := client.Metadata.GetResourceTypes(params, auth)
			CheckErr(err)

			for _, rtype := range resp.Payload.ResourceTypes {
				fmt.Println(rtype)
			}
		},
	}

	cmd.Flags().StringVar(&opts.Provider, "provider", "aws", "Cloud provider [aws | aws_govcloud | azure]")
	cmd.Flags().StringVar(&opts.Region, "region", "", "Region")
	cmd.MarkFlagRequired("provider")

	return cmd
}

func init() {
	getCmd.AddCommand(NewGetResourceTypesCommand())
}
