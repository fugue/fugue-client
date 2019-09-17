// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RuleValidationResult Result from validating a rule against a single resource
// swagger:model RuleValidationResult
type RuleValidationResult struct {

	// Resource ID to which the result applies
	ResourceID string `json:"resource_id,omitempty"`

	// Resource type to which the result applies
	ResourceType string `json:"resource_type,omitempty"`

	// Result
	Result string `json:"result,omitempty"`
}

// Validate validates this rule validation result
func (m *RuleValidationResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuleValidationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleValidationResult) UnmarshalBinary(b []byte) error {
	var res RuleValidationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
