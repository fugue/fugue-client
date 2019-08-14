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

// UpdateEnvironmentInput A managed environment.
// swagger:model UpdateEnvironmentInput
type UpdateEnvironmentInput struct {

	// Scan ID of the baseline if baseline is enabled.
	BaselineID string `json:"baseline_id,omitempty"`

	// List of compliance families validated against the environment.
	ComplianceFamilies []string `json:"compliance_families"`

	// Name of the environment.
	Name string `json:"name,omitempty"`

	// Name of the cloud service provider for the environment.
	// Enum: [aws aws_govcloud azure]
	Provider string `json:"provider,omitempty"`

	// provider options
	ProviderOptions *ProviderOptionsUpdateInput `json:"provider_options,omitempty"`

	// List of resource types remediated for the environment if remediation is enabled.
	RemediateResourceTypes []string `json:"remediate_resource_types"`

	// Indicates whether remediation is enabled for the environment.
	Remediation bool `json:"remediation,omitempty"`

	// Time in seconds between the end of one scan to the start of the next. Must also set scan_schedule_enabled to true.
	// Minimum: 300
	ScanInterval int64 `json:"scan_interval,omitempty"`

	// Indicates whether an environment is scanned on a schedule.
	ScanScheduleEnabled bool `json:"scan_schedule_enabled,omitempty"`

	// List of resource types surveyed for the environment.
	SurveyResourceTypes []string `json:"survey_resource_types"`
}

// Validate validates this update environment input
func (m *UpdateEnvironmentInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateEnvironmentInputTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aws","aws_govcloud","azure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateEnvironmentInputTypeProviderPropEnum = append(updateEnvironmentInputTypeProviderPropEnum, v)
	}
}

const (

	// UpdateEnvironmentInputProviderAws captures enum value "aws"
	UpdateEnvironmentInputProviderAws string = "aws"

	// UpdateEnvironmentInputProviderAwsGovcloud captures enum value "aws_govcloud"
	UpdateEnvironmentInputProviderAwsGovcloud string = "aws_govcloud"

	// UpdateEnvironmentInputProviderAzure captures enum value "azure"
	UpdateEnvironmentInputProviderAzure string = "azure"
)

// prop value enum
func (m *UpdateEnvironmentInput) validateProviderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateEnvironmentInputTypeProviderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateEnvironmentInput) validateProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEnvironmentInput) validateProviderOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderOptions) { // not required
		return nil
	}

	if m.ProviderOptions != nil {
		if err := m.ProviderOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider_options")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateEnvironmentInput) validateScanInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.ScanInterval) { // not required
		return nil
	}

	if err := validate.MinimumInt("scan_interval", "body", int64(m.ScanInterval), 300, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateEnvironmentInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateEnvironmentInput) UnmarshalBinary(b []byte) error {
	var res UpdateEnvironmentInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
