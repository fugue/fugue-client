// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScanWithSummaryAllOf1ResourceTypeErrorsItems scan with summary all of1 resource type errors items
// swagger:model scanWithSummaryAllOf1ResourceTypeErrorsItems
type ScanWithSummaryAllOf1ResourceTypeErrorsItems struct {

	// error message
	// Required: true
	ErrorMessage *string `json:"error_message"`

	// resource type
	// Required: true
	ResourceType *string `json:"resource_type"`
}

// Validate validates this scan with summary all of1 resource type errors items
func (m *ScanWithSummaryAllOf1ResourceTypeErrorsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScanWithSummaryAllOf1ResourceTypeErrorsItems) validateErrorMessage(formats strfmt.Registry) error {

	if err := validate.Required("error_message", "body", m.ErrorMessage); err != nil {
		return err
	}

	return nil
}

func (m *ScanWithSummaryAllOf1ResourceTypeErrorsItems) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScanWithSummaryAllOf1ResourceTypeErrorsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScanWithSummaryAllOf1ResourceTypeErrorsItems) UnmarshalBinary(b []byte) error {
	var res ScanWithSummaryAllOf1ResourceTypeErrorsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
