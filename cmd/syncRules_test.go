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
				Text: "# Resource-Type: AWS.EC2.Instance\n# Description: fake\n\ndeny{}",
			},
			"AWS",
			"AWS.EC2.Instance",
		},
		{
			"multiple-aws",
			&regoFile{
				Text: "# Resource-Type: AWS.MULTIPLE\n# Description: fake\n\ndeny{}",
			},
			"AWS",
			"MULTIPLE",
		},
		{
			"multiple-aws-mixed",
			&regoFile{
				Text: "# Resource-Type: AWS.mulTiPLe\n# Description: fake\n\ndeny{}",
			},
			"AWS",
			"MULTIPLE",
		},
		{
			"single-aws-govcloud",
			&regoFile{
				Text: "# Resource-Type: AWS_GOVCLOUD.EC2.Instance\n# Description: fake\n\ndeny{}",
			},
			"AWS_GOVCLOUD",
			"AWS_GOVCLOUD.EC2.Instance",
		},
		{
			"multiple-aws-govcloud",
			&regoFile{
				Text: "# Resource-Type: AWS_GOVCLOUD.MULTIPLE\n# Description: fake\n\ndeny{}",
			},
			"AWS_GOVCLOUD",
			"MULTIPLE",
		},
		{
			"single-azure",
			&regoFile{
				Text: "# Resource-Type: Azure.Compute.VirtualMachine\n# Description: fake\n\ndeny{}",
			},
			"Azure",
			"Azure.Compute.VirtualMachine",
		},
		{
			"multiple-azure",
			&regoFile{
				Text: "# Resource-Type: Azure.MULTIPLE\n# Description: fake\n\ndeny{}",
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
