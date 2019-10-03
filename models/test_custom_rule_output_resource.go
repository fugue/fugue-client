// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestCustomRuleOutputResource Test results from testing a custom rule on a single resource.
// swagger:model TestCustomRuleOutputResource
type TestCustomRuleOutputResource struct {

	// ID of the resource.
	ID string `json:"id,omitempty"`

	// Whether or not this single resource is compliant.
	// Enum: [PASS FAIL UNKNOWN]
	Result string `json:"result,omitempty"`

	// Type of the resource.
	Type string `json:"type,omitempty"`
}

// Validate validates this test custom rule output resource
func (m *TestCustomRuleOutputResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var testCustomRuleOutputResourceTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PASS","FAIL","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		testCustomRuleOutputResourceTypeResultPropEnum = append(testCustomRuleOutputResourceTypeResultPropEnum, v)
	}
}

const (

	// TestCustomRuleOutputResourceResultPASS captures enum value "PASS"
	TestCustomRuleOutputResourceResultPASS string = "PASS"

	// TestCustomRuleOutputResourceResultFAIL captures enum value "FAIL"
	TestCustomRuleOutputResourceResultFAIL string = "FAIL"

	// TestCustomRuleOutputResourceResultUNKNOWN captures enum value "UNKNOWN"
	TestCustomRuleOutputResourceResultUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *TestCustomRuleOutputResource) validateResultEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, testCustomRuleOutputResourceTypeResultPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TestCustomRuleOutputResource) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCustomRuleOutputResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCustomRuleOutputResource) UnmarshalBinary(b []byte) error {
	var res TestCustomRuleOutputResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
