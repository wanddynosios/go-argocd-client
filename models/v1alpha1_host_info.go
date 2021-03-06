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

// V1alpha1HostInfo HostInfo holds host name and resources metrics
// TODO: describe purpose of this type
// TODO: describe members of this type
//
// swagger:model v1alpha1HostInfo
type V1alpha1HostInfo struct {

	// name
	Name string `json:"name,omitempty"`

	// resources info
	ResourcesInfo []*V1alpha1HostResourceInfo `json:"resourcesInfo"`

	// system info
	SystemInfo *V1NodeSystemInfo `json:"systemInfo,omitempty"`
}

// Validate validates this v1alpha1 host info
func (m *V1alpha1HostInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourcesInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1HostInfo) validateResourcesInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourcesInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourcesInfo); i++ {
		if swag.IsZero(m.ResourcesInfo[i]) { // not required
			continue
		}

		if m.ResourcesInfo[i] != nil {
			if err := m.ResourcesInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourcesInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourcesInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1HostInfo) validateSystemInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemInfo) { // not required
		return nil
	}

	if m.SystemInfo != nil {
		if err := m.SystemInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("systemInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha1 host info based on the context it is used
func (m *V1alpha1HostInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourcesInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1HostInfo) contextValidateResourcesInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourcesInfo); i++ {

		if m.ResourcesInfo[i] != nil {
			if err := m.ResourcesInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourcesInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourcesInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1HostInfo) contextValidateSystemInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemInfo != nil {
		if err := m.SystemInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("systemInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1HostInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1HostInfo) UnmarshalBinary(b []byte) error {
	var res V1alpha1HostInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
