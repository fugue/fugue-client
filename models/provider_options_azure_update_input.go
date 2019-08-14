// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProviderOptionsAzureUpdateInput Mutable provider options for Azure.
// swagger:model ProviderOptionsAzureUpdateInput
type ProviderOptionsAzureUpdateInput struct {

	// The application ID/client ID of the service principal to be used
	ApplicationID string `json:"application_id,omitempty"`

	// The client secret of the service principal to be used
	ClientSecret string `json:"client_secret,omitempty"`

	// The resource groups to be remediated
	RemediateResourceGroups []string `json:"remediate_resource_groups"`

	// The resource groups to be surveyed
	SurveyResourceGroups []string `json:"survey_resource_groups"`
}

// Validate validates this provider options azure update input
func (m *ProviderOptionsAzureUpdateInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProviderOptionsAzureUpdateInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderOptionsAzureUpdateInput) UnmarshalBinary(b []byte) error {
	var res ProviderOptionsAzureUpdateInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
