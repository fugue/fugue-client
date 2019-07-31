// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ResourceSummaryFamiliesItems Compliance summary for the compliance family run against resources for the scan.
// swagger:model resourceSummaryFamiliesItems
type ResourceSummaryFamiliesItems struct {

	// Number of compliant resources in this family.
	Compliant int64 `json:"compliant"`

	// Name of the compliance family.
	Family string `json:"family"`

	// Number of noncompliant resources in this family.
	Noncompliant int64 `json:"noncompliant"`

	// Number of compliance rules failed in this family.
	RulesFailed int64 `json:"rules_failed"`

	// Number of compliance rules passed in this family.
	RulesPassed int64 `json:"rules_passed"`
}

// Validate validates this resource summary families items
func (m *ResourceSummaryFamiliesItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceSummaryFamiliesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceSummaryFamiliesItems) UnmarshalBinary(b []byte) error {
	var res ResourceSummaryFamiliesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
