// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Applicationv1alpha1EnvEntry EnvEntry represents an entry in the application's environment
//
// swagger:model applicationv1alpha1EnvEntry
type Applicationv1alpha1EnvEntry struct {

	// Name is the name of the variable, usually expressed in uppercase
	Name string `json:"name,omitempty"`

	// Value is the value of the variable
	Value string `json:"value,omitempty"`
}

// Validate validates this applicationv1alpha1 env entry
func (m *Applicationv1alpha1EnvEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this applicationv1alpha1 env entry based on context it is used
func (m *Applicationv1alpha1EnvEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Applicationv1alpha1EnvEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Applicationv1alpha1EnvEntry) UnmarshalBinary(b []byte) error {
	var res Applicationv1alpha1EnvEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
