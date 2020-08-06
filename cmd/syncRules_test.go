package cmd

import (
	"testing"
)

func TestParseRego(t *testing.T) {
	tests := []struct {
		name        string
		rego        *regoFile
		expProvider string
		expResource string
		expSeverity string
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
			"AWS",
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
			"AWS",
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
			"AWS_GOVCLOUD",
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
			"AWS_GOVCLOUD",
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
			"Azure",
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
			"Azure",
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
			"AWS_GOVCLOUD",
			"AWS.S3.Bucket",
			"Medium",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rego.ParseText()

			if err != nil {
				t.Errorf("Error in parseRego %s", err.Error())
			}
			if tt.rego.Provider != tt.expProvider {
				t.Errorf("Expected %s, actual %s", tt.expProvider, tt.rego.Provider)
			}
			if tt.rego.ResourceType != tt.expResource {
				t.Errorf("Expected %s, actual %s", tt.expResource, tt.rego.ResourceType)
			}
		})
	}
}
