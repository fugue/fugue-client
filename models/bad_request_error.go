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

// BadRequestError Error returned when the API is presented with a bad request.
// swagger:model BadRequestError
type BadRequestError struct {

	// HTTP status code for the error.
	Code int64 `json:"code"`

	// Detailed human-readable message about the bad request.
	Message string `json:"message"`

	// Type of bad request.
	// Enum: [BadRequest AlreadyAttachedToDifferentTenantError AlreadyAttachedToTenantError AlreadyInvitedError InvalidCredential InvalidJSON InvalidParameterValue MissingParameter RoleNotAssumable WorkAlreadyStartedException]
	Type string `json:"type"`
}

// Validate validates this bad request error
func (m *BadRequestError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var badRequestErrorTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BadRequest","AlreadyAttachedToDifferentTenantError","AlreadyAttachedToTenantError","AlreadyInvitedError","InvalidCredential","InvalidJSON","InvalidParameterValue","MissingParameter","RoleNotAssumable","WorkAlreadyStartedException"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		badRequestErrorTypeTypePropEnum = append(badRequestErrorTypeTypePropEnum, v)
	}
}

const (

	// BadRequestErrorTypeBadRequest captures enum value "BadRequest"
	BadRequestErrorTypeBadRequest string = "BadRequest"

	// BadRequestErrorTypeAlreadyAttachedToDifferentTenantError captures enum value "AlreadyAttachedToDifferentTenantError"
	BadRequestErrorTypeAlreadyAttachedToDifferentTenantError string = "AlreadyAttachedToDifferentTenantError"

	// BadRequestErrorTypeAlreadyAttachedToTenantError captures enum value "AlreadyAttachedToTenantError"
	BadRequestErrorTypeAlreadyAttachedToTenantError string = "AlreadyAttachedToTenantError"

	// BadRequestErrorTypeAlreadyInvitedError captures enum value "AlreadyInvitedError"
	BadRequestErrorTypeAlreadyInvitedError string = "AlreadyInvitedError"

	// BadRequestErrorTypeInvalidCredential captures enum value "InvalidCredential"
	BadRequestErrorTypeInvalidCredential string = "InvalidCredential"

	// BadRequestErrorTypeInvalidJSON captures enum value "InvalidJSON"
	BadRequestErrorTypeInvalidJSON string = "InvalidJSON"

	// BadRequestErrorTypeInvalidParameterValue captures enum value "InvalidParameterValue"
	BadRequestErrorTypeInvalidParameterValue string = "InvalidParameterValue"

	// BadRequestErrorTypeMissingParameter captures enum value "MissingParameter"
	BadRequestErrorTypeMissingParameter string = "MissingParameter"

	// BadRequestErrorTypeRoleNotAssumable captures enum value "RoleNotAssumable"
	BadRequestErrorTypeRoleNotAssumable string = "RoleNotAssumable"

	// BadRequestErrorTypeWorkAlreadyStartedException captures enum value "WorkAlreadyStartedException"
	BadRequestErrorTypeWorkAlreadyStartedException string = "WorkAlreadyStartedException"
)

// prop value enum
func (m *BadRequestError) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, badRequestErrorTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BadRequestError) validateType(formats strfmt.Registry) error {

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
func (m *BadRequestError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BadRequestError) UnmarshalBinary(b []byte) error {
	var res BadRequestError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
