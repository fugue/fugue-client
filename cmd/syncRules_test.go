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
# Resource-Type: AWS.EC2.Instance
# Description: fake
# Provider: AWS
deny{}`,
			},
			"AWS",
			"AWS.EC2.Instance",
		},
		{
			"multiple-aws",
			&regoFile{
				Text: `
# Resource-Type: MULTIPLE
# Description: fake
# Provider: AWS
deny{}`,
			},
			"AWS",
			"MULTIPLE",
		},
		{
			"single-aws-govcloud",
			&regoFile{
				Text: `
# Resource-Type: AWS.EC2.Instance
# Provider: AWS_GOVCLOUD
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
# Resource-Type: MULTIPLE
# Description: fake
# Provider: AWS_GOVCLOUD

deny{}`,
			},
			"AWS_GOVCLOUD",
			"MULTIPLE",
		},
		{
			"single-azure",
			&regoFile{
				Text: `
# Resource-Type: Azure.Compute.VirtualMachine
# Description: fake
# Provider: Azure
deny{}`,
			},
			"Azure",
			"Azure.Compute.VirtualMachine",
		},
		{
			"multiple-azure",
			&regoFile{
				Text: `
# Resource-Type: MULTIPLE
# Description: fake
# Provider: Azure
deny{}`,
			},
			"Azure",
			"MULTIPLE",
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
