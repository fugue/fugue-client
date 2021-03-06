// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomRuleError An error for a custom rule
//
// swagger:model CustomRuleError
type CustomRuleError struct {

	// Severity of the error.
	// Enum: [error warning]
	Severity string `json:"severity,omitempty"`

	// Text describing the error
	Text string `json:"text,omitempty"`
}

// Validate validates this custom rule error
func (m *CustomRuleError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customRuleErrorTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["error","warning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customRuleErrorTypeSeverityPropEnum = append(customRuleErrorTypeSeverityPropEnum, v)
	}
}

const (

	// CustomRuleErrorSeverityError captures enum value "error"
	CustomRuleErrorSeverityError string = "error"

	// CustomRuleErrorSeverityWarning captures enum value "warning"
	CustomRuleErrorSeverityWarning string = "warning"
)

// prop value enum
func (m *CustomRuleError) validateSeverityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customRuleErrorTypeSeverityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomRuleError) validateSeverity(formats strfmt.Registry) error {

	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomRuleError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomRuleError) UnmarshalBinary(b []byte) error {
	var res CustomRuleError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
