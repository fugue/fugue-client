// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProviderOptionsUpdateInput Mutable provider options.
//
// swagger:model ProviderOptionsUpdateInput
type ProviderOptionsUpdateInput struct {

	// aws
	Aws *ProviderOptionsAwsUpdateInput `json:"aws,omitempty"`

	// aws govcloud
	AwsGovcloud *ProviderOptionsAwsUpdateInput `json:"aws_govcloud,omitempty"`

	// azure
	Azure *ProviderOptionsAzureUpdateInput `json:"azure,omitempty"`
}

// Validate validates this provider options update input
func (m *ProviderOptionsUpdateInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsGovcloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProviderOptionsUpdateInput) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *ProviderOptionsUpdateInput) validateAwsGovcloud(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsGovcloud) { // not required
		return nil
	}

	if m.AwsGovcloud != nil {
		if err := m.AwsGovcloud.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_govcloud")
			}
			return err
		}
	}

	return nil
}

func (m *ProviderOptionsUpdateInput) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProviderOptionsUpdateInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderOptionsUpdateInput) UnmarshalBinary(b []byte) error {
	var res ProviderOptionsUpdateInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
