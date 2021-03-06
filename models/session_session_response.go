// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SessionSessionResponse SessionResponse wraps the created token or returns an empty string if deleted.
//
// swagger:model sessionSessionResponse
type SessionSessionResponse struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this session session response
func (m *SessionSessionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this session session response based on context it is used
func (m *SessionSessionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SessionSessionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SessionSessionResponse) UnmarshalBinary(b []byte) error {
	var res SessionSessionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
