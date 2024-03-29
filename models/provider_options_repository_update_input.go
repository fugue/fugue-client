// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProviderOptionsRepositoryUpdateInput Mutable provider options for repository environments.
//
// swagger:model ProviderOptionsRepositoryUpdateInput
type ProviderOptionsRepositoryUpdateInput struct {

	// The branch associated with this environment (e.g. 'main')
	Branch string `json:"branch,omitempty"`

	// The URL of the repository (e.g. 'https://github.com/fugue/regula.git')
	URL string `json:"url,omitempty"`
}

// Validate validates this provider options repository update input
func (m *ProviderOptionsRepositoryUpdateInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProviderOptionsRepositoryUpdateInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderOptionsRepositoryUpdateInput) UnmarshalBinary(b []byte) error {
	var res ProviderOptionsRepositoryUpdateInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
