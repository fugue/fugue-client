package cmd

import (
	"fmt"

	"github.com/fugue/fugue-client/client/events"
	"github.com/fugue/fugue-client/format"
	"github.com/spf13/cobra"
)

type listEventsOptions struct {
	Offset       int64
	MaxItems     int64
	RangeFrom    int64
	RangeTo      int64
	EventType    []string
	Change       []string
	Remediated   []string
	ResourceType []string
	Columns      []string
}

// NewListEventsCommand returns a command that lists events in an environment
func NewListEventsCommand() *cobra.Command {

	var opts listEventsOptions

	cmd := &cobra.Command{
		Use:   "events [environment_id]",
		Short: "List environment events",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			params := events.NewListEventsParams()
			params.EnvironmentID = args[0]

			if opts.Offset > 0 {
				params.Offset = &opts.Offset
			}
			if opts.MaxItems > 0 {
				params.MaxItems = &opts.MaxItems
			}
			if opts.RangeFrom > 0 {
				params.RangeFrom = &opts.RangeFrom
			}
			if opts.RangeTo > 0 {
				params.RangeTo = &opts.RangeTo
			}
			if len(opts.EventType) > 0 {
				params.EventType = opts.EventType
			}
			if len(opts.Change) > 0 {
				params.Change = opts.Change
			}
			if len(opts.Remediated) > 0 {
				params.Remediated = opts.Remediated
			}
			if len(opts.ResourceType) > 0 {
				params.ResourceType = opts.ResourceType
			}

			resp, err := client.Events.ListEvents(params, auth)
			CheckErr(err)

			events := resp.Payload.Items

			rows := make([]interface{}, len(events))
			for i, event := range events {
				rows[i] = event
			}

			table, err := format.Table(format.TableOpts{
				Rows:       rows,
				Columns:    opts.Columns,
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().String("environment-id", "", "Environment ID (required)")
	cmd.Flags().Int64("offset", 0, "Offset")
	cmd.Flags().Int64("max-items", 0, "Max items")
	cmd.Flags().Int64("range-from", 0, "Range from")
	cmd.Flags().Int64("range-to", 0, "Range to")
	cmd.Flags().StringSlice("event-type", nil, "Event types")
	cmd.Flags().StringSlice("change", nil, "Change")
	cmd.Flags().StringSlice("remediated", nil, "Remediated")
	cmd.Flags().StringSlice("resource-type", nil, "Resource types")
	cmd.Flags().StringSliceVar(&opts.Columns, "columns", []string{"ID", "EventType", "CreatedAt"}, "columns to show")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListEventsCommand())
}
