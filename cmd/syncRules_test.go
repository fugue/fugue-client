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
	}{
		{
			"single-aws",
			&regoFile{
				Text: `
# Provider: AWS
# Resource-Type: EC2.Instance
# Description: fake
deny{}`,
			},
			"AWS",
			"AWS.EC2.Instance",
		},
		{
			"multiple-aws",
			&regoFile{
				Text: `
# Provider: AWS
# Resource-Type: MULTIPLE
# Description: fake
deny{}`,
			},
			"AWS",
			"MULTIPLE",
		},
		{
			"single-aws-govcloud",
			&regoFile{
				Text: `
# Provider: AWS_GOVCLOUD
# Resource-Type: EC2.Instance
# Description: fake

deny{}`,
			},
			"AWS_GOVCLOUD",
			"AWS.EC2.Instance",
		},
		{
			"multiple-aws-govcloud",
			&regoFile{
				Text: `
# Provider: AWS_GOVCLOUD
# Resource-Type: MULTIPLE
# Description: fake

deny{}`,
			},
			"AWS_GOVCLOUD",
			"MULTIPLE",
		},
		{
			"single-azure",
			&regoFile{
				Text: `
# Provider: Azure
# Resource-Type: Compute.VirtualMachine
# Description: fake
deny{}`,
			},
			"Azure",
			"Azure.Compute.VirtualMachine",
		},
		{
			"multiple-azure",
			&regoFile{
				Text: `
# Provider: Azure
# Resource-Type: MULTIPLE
# Description: fake
deny{}`,
			},
			"Azure",
			"MULTIPLE",
		},
		{
			"becki1-aws-s3-bucket-sse",
			&regoFile{
				Text: `
# Provider: AWS_GOVCLOUD
# Resource-Type: S3.Bucket
# Description: SSE encryption should be enabled for S3 buckets (AES-256 or KMS).

allow {
  input.server_side_encryption_configuration[_].rule[_][_][_].sse_algorithm = _
}`,
			},
			"AWS_GOVCLOUD",
			"AWS.S3.Bucket",
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
