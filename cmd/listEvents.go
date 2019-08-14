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

type listEventsViewItem struct {
	EventID      string
	EventType    string
	CreatedAt    string
	ResourceID   string
	ResourceType string
	Change       string
	OldState     string
	NewState     string
	Error        string
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

			empty := "-"

			rows := make([]interface{}, len(events))
			for i, event := range events {

				item := listEventsViewItem{
					EventID:      event.ID,
					EventType:    event.EventType,
					CreatedAt:    format.Unix(event.CreatedAt),
					Error:        event.Error,
					ResourceID:   empty,
					ResourceType: empty,
					OldState:     empty,
					NewState:     empty,
					Change:       empty,
				}

				if event.ComplianceDiff != nil {
					diff := event.ComplianceDiff
					item.ResourceID = diff.ResourceID
					item.ResourceType = diff.ResourceType
					item.OldState = diff.OldState
					item.NewState = diff.NewState

					// This will happen when a resource type is "UNKNOWN"
					if item.ResourceID == "" {
						item.ResourceID = empty
					}
				}

				if event.ResourceDiff != nil {
					diff := event.ResourceDiff
					item.ResourceID = diff.ResourceID
					item.ResourceType = diff.ResourceType
					item.Change = diff.Change
				}

				// Resource IDs can be extremely long.
				// Truncate at a max length for now.
				idLen := len(item.ResourceID)
				if idLen > 50 {
					item.ResourceID = "..." + item.ResourceID[idLen-50:]
				}

				rows[i] = item
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

	defaultCols := []string{
		// "EventID",
		"EventType",
		"CreatedAt",
		"ResourceID",
		"ResourceType",
		"Change",
		"OldState",
		"NewState",
	}

	cmd.Flags().Int64Var(&opts.Offset, "offset", 0, "Offset")
	cmd.Flags().Int64Var(&opts.MaxItems, "max-items", 20, "Max items")
	cmd.Flags().Int64Var(&opts.RangeFrom, "range-from", 0, "Range from")
	cmd.Flags().Int64Var(&opts.RangeTo, "range-to", 0, "Range to")
	cmd.Flags().StringSliceVar(&opts.EventType, "event-type", nil, "Event types")
	cmd.Flags().StringSliceVar(&opts.Change, "change", nil, "Change")
	cmd.Flags().StringSliceVar(&opts.Remediated, "remediated", nil, "Remediated")
	cmd.Flags().StringSliceVar(&opts.ResourceType, "resource-type", nil, "Resource types")
	cmd.Flags().StringSliceVar(&opts.Columns, "columns", defaultCols, "columns to show")

	return cmd
}

func init() {
	listCmd.AddCommand(NewListEventsCommand())
}
