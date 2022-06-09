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

// V1alpha1ResourceAction TODO: describe this type
// TODO: describe members of this type
//
// swagger:model v1alpha1ResourceAction
type V1alpha1ResourceAction struct {

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// params
	Params []*V1alpha1ResourceActionParam `json:"params"`
}

// Validate validates this v1alpha1 resource action
func (m *V1alpha1ResourceAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ResourceAction) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(m.Params) { // not required
		return nil
	}

	for i := 0; i < len(m.Params); i++ {
		if swag.IsZero(m.Params[i]) { // not required
			continue
		}

		if m.Params[i] != nil {
			if err := m.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1alpha1 resource action based on the context it is used
func (m *V1alpha1ResourceAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ResourceAction) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Params); i++ {

		if m.Params[i] != nil {
			if err := m.Params[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ResourceAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ResourceAction) UnmarshalBinary(b []byte) error {
	var res V1alpha1ResourceAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
