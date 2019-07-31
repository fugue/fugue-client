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

// NotFoundError Error returned when the API request references a non-existent resource.
// swagger:model NotFoundError
type NotFoundError struct {

	// HTTP status code for the error.
	Code int64 `json:"code"`

	// Detailed human-readable message about the not found error.
	Message string `json:"message"`

	// Type of not found error.
	// Enum: [NotFound]
	Type string `json:"type"`
}

// Validate validates this not found error
func (m *NotFoundError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var notFoundErrorTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotFound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notFoundErrorTypeTypePropEnum = append(notFoundErrorTypeTypePropEnum, v)
	}
}

const (

	// NotFoundErrorTypeNotFound captures enum value "NotFound"
	NotFoundErrorTypeNotFound string = "NotFound"
)

// prop value enum
func (m *NotFoundError) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, notFoundErrorTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NotFoundError) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotFoundError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotFoundError) UnmarshalBinary(b []byte) error {
	var res NotFoundError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
