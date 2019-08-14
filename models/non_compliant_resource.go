// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NonCompliantResource Describes the rules violated by a resource.
// swagger:model NonCompliantResource
type NonCompliantResource struct {

	// List of rules and messages the resource violates.
	FailedRules []*NonCompliantResourceFailedRulesItems0 `json:"failed_rules"`

	// ID of the failing resource.
	ResourceID string `json:"resource_id,omitempty"`
}

// Validate validates this non compliant resource
func (m *NonCompliantResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailedRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NonCompliantResource) validateFailedRules(formats strfmt.Registry) error {

	if swag.IsZero(m.FailedRules) { // not required
		return nil
	}

	for i := 0; i < len(m.FailedRules); i++ {
		if swag.IsZero(m.FailedRules[i]) { // not required
			continue
		}

		if m.FailedRules[i] != nil {
			if err := m.FailedRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failed_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NonCompliantResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NonCompliantResource) UnmarshalBinary(b []byte) error {
	var res NonCompliantResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NonCompliantResourceFailedRulesItems0 non compliant resource failed rules items0
// swagger:model NonCompliantResourceFailedRulesItems0
type NonCompliantResourceFailedRulesItems0 struct {

	// Compliance family the violated rule belongs to.
	Family string `json:"family,omitempty"`

	// Reasons the resource was found in violation of a rule.
	Messages []string `json:"messages"`

	// ID of the violated rule.
	Rule string `json:"rule,omitempty"`
}

// Validate validates this non compliant resource failed rules items0
func (m *NonCompliantResourceFailedRulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NonCompliantResourceFailedRulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NonCompliantResourceFailedRulesItems0) UnmarshalBinary(b []byte) error {
	var res NonCompliantResourceFailedRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
