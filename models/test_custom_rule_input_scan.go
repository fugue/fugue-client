// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TestCustomRuleInputScan Scan used as input to a custom rule.
//
// swagger:model TestCustomRuleInputScan
type TestCustomRuleInputScan struct {

	// resources
	Resources []interface{} `json:"resources"`
}

// Validate validates this test custom rule input scan
func (m *TestCustomRuleInputScan) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestCustomRuleInputScan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCustomRuleInputScan) UnmarshalBinary(b []byte) error {
	var res TestCustomRuleInputScan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
