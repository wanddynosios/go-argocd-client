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

// V1alpha1ProjectRole ProjectRole represents a role that has access to a project
//
// swagger:model v1alpha1ProjectRole
type V1alpha1ProjectRole struct {

	// Description is a description of the role
	Description string `json:"description,omitempty"`

	// Groups are a list of OIDC group claims bound to this role
	Groups []string `json:"groups"`

	// JWTTokens are a list of generated JWT tokens bound to this role
	JwtTokens []*V1alpha1JWTToken `json:"jwtTokens"`

	// Name is a name for this role
	Name string `json:"name,omitempty"`

	// Policies Stores a list of casbin formatted strings that define access policies for the role in the project
	Policies []string `json:"policies"`
}

// Validate validates this v1alpha1 project role
func (m *V1alpha1ProjectRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJwtTokens(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ProjectRole) validateJwtTokens(formats strfmt.Registry) error {
	if swag.IsZero(m.JwtTokens) { // not required
		return nil
	}

	for i := 0; i < len(m.JwtTokens); i++ {
		if swag.IsZero(m.JwtTokens[i]) { // not required
			continue
		}

		if m.JwtTokens[i] != nil {
			if err := m.JwtTokens[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jwtTokens" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jwtTokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1alpha1 project role based on the context it is used
func (m *V1alpha1ProjectRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJwtTokens(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ProjectRole) contextValidateJwtTokens(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.JwtTokens); i++ {

		if m.JwtTokens[i] != nil {
			if err := m.JwtTokens[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jwtTokens" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jwtTokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ProjectRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ProjectRole) UnmarshalBinary(b []byte) error {
	var res V1alpha1ProjectRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
