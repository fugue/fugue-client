package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/client/rule_waivers"
	"github.com/fugue/fugue-client/format"
	"github.com/fugue/fugue-client/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type createRuleWaiverOptions struct {
	Name              string
	Comment           string
	EnvironmentID     string
	RuleID            string
	ResourceID        string
	ResourceType      string
	ResourceProvider  string
	ResourceTag       string
	ExpiresAt         string
	ExpiresAtDuration string
}

// (?i) case insensitive
// Make the P and T optional (this is not standard!!!)
var durationPattern = regexp.MustCompile(`(?i)^P?((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<week>\d+)W)?((?P<day>\d+)D)?(T?((?P<hour>\d+)H)?((?P<minute>\d+)M)?((?P<second>\d+)S)?)?$`)

func parseDuration(duration string) (*models.Duration, error) {

	if duration == "" {
		return nil, nil
	}

	var match []string
	var d models.Duration

	if durationPattern.MatchString(duration) {
		match = durationPattern.FindStringSubmatch(duration)
	} else {
		return nil, errors.New("could not parse duration string")
	}

	for i, name := range durationPattern.SubexpNames() {
		part := match[i]
		if i == 0 || name == "" || part == "" {
			continue
		}

		val, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		v := int64(val)
		switch name {
		case "year":
			d.Years = v
		case "month":
			d.Months = v
		case "week":
			d.Weeks = v
		case "day":
			d.Days = v
		case "hour":
			d.Hours = v
		// case "minute":
		// case "second":
		default:
			return nil, fmt.Errorf("unknown field %s", name)
		}
	}
	return &d, nil
}

func parseExpiresAt(expiresAt string) (*time.Time, error) {

	if expiresAt == "" {
		return nil, nil
	}
	// try timestamp first
	expiresAtInt, err := strconv.ParseInt(expiresAt, 10, 64)
	if err != nil {
		// try RFC3339 format
		t, err2 := time.Parse(time.RFC3339, expiresAt)
		if err2 != nil {
			err = errors.Wrap(err, err2.Error())
			return nil, err
		}
		return &t, nil
	}
	t := time.Unix(expiresAtInt, 0)
	return &t, nil
}

// NewCreateRuleWaiverCommand returns a command that creates a custom rule
func NewCreateRuleWaiverCommand() *cobra.Command {

	var opts createRuleWaiverOptions

	cmd := &cobra.Command{
		Use:     "rule-waiver",
		Short:   "Create a rule waiver",
		Aliases: []string{"waiver", "rule_waiver", "rule-waivers", "waivers", "rule_waivers"},
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			// Mutually exclusive flags
			if opts.ExpiresAt != "" && opts.ExpiresAtDuration != "" {
				err := errors.New("cannot specify both --expires-at and --expires-at-duration")
				CheckErr(err)
			}

			duration, err := parseDuration(opts.ExpiresAtDuration)
			CheckErr(err)

			expiresAtPtr, err := parseExpiresAt(opts.ExpiresAt)
			CheckErr(err)
			var expiresAt int64 = 0
			if expiresAtPtr != nil {
				expiresAt = expiresAtPtr.Unix()
			}

			params := rule_waivers.NewCreateRuleWaiverParams()
			params.Input = &models.CreateRuleWaiverInput{
				Name:              &opts.Name,
				Comment:           opts.Comment,
				EnvironmentID:     &opts.EnvironmentID,
				RuleID:            &opts.RuleID,
				ResourceID:        &opts.ResourceID,
				ResourceType:      &opts.ResourceType,
				ResourceProvider:  &opts.ResourceProvider,
				ResourceTag:       opts.ResourceTag,
				ExpiresAt:         expiresAt,
				ExpiresAtDuration: duration,
			}

			resp, err := client.RuleWaivers.CreateRuleWaiver(params, auth)
			if err != nil {
				switch respError := err.(type) {
				case *custom_rules.CreateCustomRuleInternalServerError:
					Fatal(respError.Payload.Message, DefaultErrorExitCode)
				default:
					CheckErr(err)
				}
			}

			waiver := resp.Payload

			var itemTag Item
			if waiver.ResourceTag != "" {
				itemTag = Item{"RESOURCE_TAG", waiver.ResourceTag}
			} else {
				itemTag = Item{"RESOURCE_TAG", "-"}
			}

			var itemTime Item
			if waiver.ExpiresAt != 0 {
				t := time.Unix(waiver.ExpiresAt, 0)
				itemTime = Item{"EXPIRES_AT", t.Format(time.RFC3339)}
			} else {
				itemTime = Item{"EXPIRES_AT", "-"}
			}

			items := []interface{}{
				Item{"RULE_WAIVER_ID", *waiver.ID},
				Item{"NAME", *waiver.Name},
				Item{"COMMENT", waiver.Comment},
				Item{"ENVIRONMENT_ID", *waiver.EnvironmentID},
				Item{"ENVIRONMENT_NAME", waiver.EnvironmentName},
				Item{"RULE_ID", *waiver.RuleID},
				Item{"RESOURCE_ID", *waiver.ResourceID},
				Item{"RESOURCE_TYPE", *waiver.ResourceType},
				Item{"RESOURCE_PROVIDER", *waiver.ResourceProvider},
				itemTag,
				itemTime,
				Item{"CREATED_AT", format.Unix(waiver.CreatedAt)},
				Item{"CREATED_BY", waiver.CreatedBy},
				Item{"CREATED_BY_DISPLAY_NAME", waiver.CreatedByDisplayName},
				Item{"UPDATED_AT", format.Unix(waiver.UpdatedAt)},
				Item{"UPDATED_BY", waiver.UpdatedBy},
				Item{"UPDATED_BY_DISPLAY_NAME", waiver.UpdatedByDisplayName},
			}

			table, err := format.Table(format.TableOpts{
				Rows:       items,
				Columns:    []string{"Attribute", "Value"},
				ShowHeader: true,
			})
			CheckErr(err)

			for _, tableRow := range table {
				fmt.Println(tableRow)
			}
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "Waiver name")
	cmd.Flags().StringVar(&opts.Comment, "comment", "", "Comment describing the rule waiver purpose")
	cmd.Flags().StringVar(&opts.RuleID, "rule-id", "", "Rule ID (e.g. FG_R00217, <UUID Custom Rule ID>)")
	cmd.Flags().StringVar(&opts.EnvironmentID, "environment-id", "", "Environment ID")
	cmd.Flags().StringVar(&opts.ResourceID, "resource-id", "", "Resource ID (e.g. resource-123, 'resource-*')")
	cmd.Flags().StringVar(&opts.ResourceType, "resource-type", "", "Resource Type (e.g. AWS.S3.Bucket, '*')")
	cmd.Flags().StringVar(&opts.ResourceProvider, "resource-provider", "", "Resource Provider (e.g. aws.us-east-1, azure, '*')")
	// resource-tag is optional in the API: if resource-tag == "", the CLI is not posting the resource-tag json field
	cmd.Flags().StringVar(&opts.ResourceTag, "resource-tag", "", "Resource tag (e.g. 'env:prod', 'env:*', '*')")

	cmd.Flags().StringVar(&opts.ExpiresAt, "expires-at", "", "Expires at in RFC3339 representation or Unix timestamp (e.g. '2020-01-01T00:00:00Z' or '1577836800')")
	cmd.Flags().StringVar(&opts.ExpiresAtDuration, "expires-at-duration", "", "Expires at duration in ISO 8601 format (e.g. 'P3Y6M4DT12H') or '4d', 1d12h, etc.")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("rule-id")
	cmd.MarkFlagRequired("environment-id")
	cmd.MarkFlagRequired("resource-id")
	cmd.MarkFlagRequired("resource-type")
	cmd.MarkFlagRequired("resource-provider")

	return cmd
}

func init() {
	createCmd.AddCommand(NewCreateRuleWaiverCommand())
}
