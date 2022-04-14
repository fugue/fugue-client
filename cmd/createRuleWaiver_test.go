package cmd

import (
	"testing"
	"time"

	"github.com/fugue/fugue-client/models"
	"github.com/stretchr/testify/assert"
)

func Test_parseDuration(t *testing.T) {

	tests := []struct {
		input string
		want  *models.Duration
		err   bool
	}{
		{
			input: "p4y3m2w1d",
			want:  &models.Duration{Years: 4, Months: 3, Weeks: 2, Days: 1},
			err:   false,
		},
		{
			input: "4y3m2w1d",
			want:  &models.Duration{Years: 4, Months: 3, Weeks: 2, Days: 1},
			err:   false,
		},
		{input: "10w1d", want: &models.Duration{Weeks: 10, Days: 1}, err: false},
		{input: "pt24h", want: &models.Duration{Hours: 24}, err: false},
		{input: "t24h", want: &models.Duration{Hours: 24}, err: false},
		{input: "p24h", want: &models.Duration{Hours: 24}, err: false},
		{input: "24h", want: &models.Duration{Hours: 24}, err: false},
		{input: "xpto", want: nil, err: true},
		{input: "", want: nil, err: false},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseDuration(tc.input)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_parseExpiresAt(t *testing.T) {

	tests := []struct {
		input string
		want  func() *time.Time
		err   bool
	}{
		{
			input: "2019-01-01T00:00:00Z",
			want: func() *time.Time {
				t := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
				return &t
			},
			err: false,
		},
		{
			input: "1649285903",
			want: func() *time.Time {
				t := time.Date(2022, time.April, 6, 22, 58, 23, 0, time.UTC)
				return &t
			},
			err: false,
		},
		{
			input: "2022-04-06T19:05:16-04:00",
			want: func() *time.Time {
				t := time.Date(2022, time.April, 6, 23, 5, 16, 0, time.UTC)
				return &t
			},
			err: false,
		},
		{
			input: "invalid",
			want:  func() *time.Time { return nil },
			err:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseExpiresAt(tc.input)
			if tc.err {
				assert.Error(t, err)
				assert.Equal(t, tc.want(), got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want().Unix(), got.Unix())
			}
		})
	}

}

func Test_parseBothExpiresAt(t *testing.T) {

	tests := []struct {
		input         string
		wantTime      func() *time.Time
		wantDuraction *models.Duration
		err           bool
	}{
		{
			input: "2019-01-01T00:00:00Z",
			wantTime: func() *time.Time {
				t := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
				return &t
			},
			err: false,
		},
		{
			input: "1649285903",
			wantTime: func() *time.Time {
				t := time.Date(2022, time.April, 6, 22, 58, 23, 0, time.UTC)
				return &t
			},
			err: false,
		},
		{
			input:         "p4y3m2w1d",
			wantDuraction: &models.Duration{Years: 4, Months: 3, Weeks: 2, Days: 1},
			err:           false,
		},
		{
			input:         "4y3m2w1d",
			wantDuraction: &models.Duration{Years: 4, Months: 3, Weeks: 2, Days: 1},
			err:           false,
		},
		{
			input:         "invalid",
			wantTime:      func() *time.Time { return nil },
			wantDuraction: nil,
			err:           true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			gotTime, gotDuration, err := parseBothExpiresAt(tc.input)
			if gotTime != nil {
				assert.Equal(t, tc.wantTime().Unix(), gotTime.Unix())
			}
			if gotDuration != nil {
				assert.Equal(t, tc.wantDuraction, gotDuration)
			}
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
