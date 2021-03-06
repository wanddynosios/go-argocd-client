// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountAccount account account
//
// swagger:model accountAccount
type AccountAccount struct {

	// capabilities
	Capabilities []string `json:"capabilities"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tokens
	Tokens []*AccountToken `json:"tokens"`
}

// Validate validates this account account
func (m *AccountAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTokens(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAccount) validateTokens(formats strfmt.Registry) error {
	if swag.IsZero(m.Tokens) { // not required
		return nil
	}

	for i := 0; i < len(m.Tokens); i++ {
		if swag.IsZero(m.Tokens[i]) { // not required
			continue
		}

		if m.Tokens[i] != nil {
			if err := m.Tokens[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tokens" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this account account based on the context it is used
func (m *AccountAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTokens(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAccount) contextValidateTokens(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tokens); i++ {

		if m.Tokens[i] != nil {
			if err := m.Tokens[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tokens" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAccount) UnmarshalBinary(b []byte) error {
	var res AccountAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
