package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRego(t *testing.T) {
	tests := []struct {
		name         string
		rego         *regoFile
		expProviders []string
		expResource  string
		expSeverity  string
	}{
		{
			"single-aws",
			&regoFile{
				Text: `
# Provider: AWS
# Resource-Type: AWS.EC2.Instance
# Severity: Low
# Description: fake
deny{}`,
			},
			[]string{"AWS"},
			"AWS.EC2.Instance",
			"Low",
		},
		{
			"multiple-aws",
			&regoFile{
				Text: `
# Provider: AWS
# Resource-Type: MULTIPLE
# Description: fake
# Severity: Medium
deny{}`,
			},
			[]string{"AWS"},
			"MULTIPLE",
			"Medium",
		},
		{
			"single-aws-govcloud",
			&regoFile{
				Text: `
# Provider: AWS_GOVCLOUD
# Resource-Type: AWS.EC2.Instance
# Description: fake
# Severity: High
deny{}`,
			},
			[]string{"AWS_GOVCLOUD"},
			"AWS.EC2.Instance",
			"High",
		},
		{
			"multiple-aws-govcloud",
			&regoFile{
				Text: `
# Provider: AWS_GOVCLOUD
# Resource-Type: MULTIPLE
# Description: fake
# Severity: Critical
deny{}`,
			},
			[]string{"AWS_GOVCLOUD"},
			"MULTIPLE",
			"Critical",
		},
		{
			"single-azure",
			&regoFile{
				Text: `
# Provider: Azure
# Resource-Type: Azure.Compute.VirtualMachine
# Description: fake
# Severity: Informational
deny{}`,
			},
			[]string{"Azure"},
			"Azure.Compute.VirtualMachine",
			"Informational",
		},
		{
			"multiple-azure",
			&regoFile{
				Text: `
# Provider: Azure
# Resource-Type: MULTIPLE
# Description: fake
# Severity: Low
deny{}`,
			},
			[]string{"Azure"},
			"MULTIPLE",
			"Low",
		},
		{
			"becki1-aws-s3-bucket-sse",
			&regoFile{
				Text: `
# Provider: AWS_GOVCLOUD
# Resource-Type: AWS.S3.Bucket
# Description: SSE encryption should be enabled for S3 buckets (AES-256 or KMS).
# Severity: Medium
allow {
  input.server_side_encryption_configuration[_].rule[_][_][_].sse_algorithm = _
}`,
			},
			[]string{"AWS_GOVCLOUD"},
			"AWS.S3.Bucket",
			"Medium",
		},
		{
			"single-aws-no-severity",
			&regoFile{
				Text: `
# Provider: AWS
# Resource-Type: AWS.EC2.Instance
# Description: fake
deny{}`,
			},
			[]string{"AWS"},
			"AWS.EC2.Instance",
			"High",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rego.ParseText()

			if err != nil {
				t.Errorf("Error in parseRego %s", err.Error())
			}
			assert.Equal(t, tt.expProviders, tt.rego.Providers)
			assert.Equal(t, tt.expResource, tt.rego.ResourceType)
			assert.Equal(t, tt.expSeverity, tt.rego.Severity)
		})
	}
}
