// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ComplianceByRuleFailedResourceTypesItems Resource type that failed to satisfy the rule due to a required resource being omitted and associated error messages.
// swagger:model complianceByRuleFailedResourceTypesItems
type ComplianceByRuleFailedResourceTypesItems struct {

	// Messages why the rule failed.
	Messages []string `json:"messages"`

	// Resource type that failed to satisfy the rule.
	ResourceType string `json:"resource_type"`
}

// Validate validates this compliance by rule failed resource types items
func (m *ComplianceByRuleFailedResourceTypesItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComplianceByRuleFailedResourceTypesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplianceByRuleFailedResourceTypesItems) UnmarshalBinary(b []byte) error {
	var res ComplianceByRuleFailedResourceTypesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
