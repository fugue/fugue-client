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

// CustomRule A custom rule
// swagger:model CustomRule
type CustomRule struct {

	// Compliance controls to which the custom rule belongs
	ComplianceControls []string `json:"compliance_controls"`

	// The date and time the rule was created.
	CreatedAt int64 `json:"created_at,omitempty"`

	// Principal that created the rule.
	CreatedBy string `json:"created_by,omitempty"`

	// Description of the custom rule
	Description string `json:"description,omitempty"`

	// ID of the custom rule
	ID string `json:"id,omitempty"`

	// Human readable name of the custom rule
	Name string `json:"name,omitempty"`

	// Provider of the custom rule
	// Enum: [AWS AWS_GOVCLOUD AZURE]
	Provider string `json:"provider,omitempty"`

	// Resource type to which the custom rule applies
	ResourceType string `json:"resource_type,omitempty"`

	// The rego source code for the rule
	RuleText string `json:"rule_text,omitempty"`

	// The origin of this rule
	// Enum: [CUSTOM]
	Source string `json:"source,omitempty"`

	// The current status of the rule
	// Enum: [ENABLED DISABLED INVALID]
	Status string `json:"status,omitempty"`

	// The date and time the rule was last updated.
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// Principal that last updated the rule.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this custom rule
func (m *CustomRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customRuleTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","AWS_GOVCLOUD","AZURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customRuleTypeProviderPropEnum = append(customRuleTypeProviderPropEnum, v)
	}
}

const (

	// CustomRuleProviderAWS captures enum value "AWS"
	CustomRuleProviderAWS string = "AWS"

	// CustomRuleProviderAWSGOVCLOUD captures enum value "AWS_GOVCLOUD"
	CustomRuleProviderAWSGOVCLOUD string = "AWS_GOVCLOUD"

	// CustomRuleProviderAZURE captures enum value "AZURE"
	CustomRuleProviderAZURE string = "AZURE"
)

// prop value enum
func (m *CustomRule) validateProviderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customRuleTypeProviderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomRule) validateProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

var customRuleTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customRuleTypeSourcePropEnum = append(customRuleTypeSourcePropEnum, v)
	}
}

const (

	// CustomRuleSourceCUSTOM captures enum value "CUSTOM"
	CustomRuleSourceCUSTOM string = "CUSTOM"
)

// prop value enum
func (m *CustomRule) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customRuleTypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomRule) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

var customRuleTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","INVALID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customRuleTypeStatusPropEnum = append(customRuleTypeStatusPropEnum, v)
	}
}

const (

	// CustomRuleStatusENABLED captures enum value "ENABLED"
	CustomRuleStatusENABLED string = "ENABLED"

	// CustomRuleStatusDISABLED captures enum value "DISABLED"
	CustomRuleStatusDISABLED string = "DISABLED"

	// CustomRuleStatusINVALID captures enum value "INVALID"
	CustomRuleStatusINVALID string = "INVALID"
)

// prop value enum
func (m *CustomRule) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customRuleTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomRule) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomRule) UnmarshalBinary(b []byte) error {
	var res CustomRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
