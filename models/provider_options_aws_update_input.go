// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProviderOptionsAwsUpdateInput Mutable provider options for AWS.
// swagger:model ProviderOptionsAwsUpdateInput
type ProviderOptionsAwsUpdateInput struct {

	// AWS IAM Role ARN that will be assumed to scan and remediate infrastructure.
	RoleArn string `json:"role_arn"`
}

// Validate validates this provider options aws update input
func (m *ProviderOptionsAwsUpdateInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProviderOptionsAwsUpdateInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderOptionsAwsUpdateInput) UnmarshalBinary(b []byte) error {
	var res ProviderOptionsAwsUpdateInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
