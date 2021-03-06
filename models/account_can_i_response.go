// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountCanIResponse account can i response
//
// swagger:model accountCanIResponse
type AccountCanIResponse struct {

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this account can i response
func (m *AccountCanIResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account can i response based on context it is used
func (m *AccountCanIResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountCanIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCanIResponse) UnmarshalBinary(b []byte) error {
	var res AccountCanIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
