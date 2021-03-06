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

// V1alpha1ApplicationSourceHelm ApplicationSourceHelm holds helm specific options
//
// swagger:model v1alpha1ApplicationSourceHelm
type V1alpha1ApplicationSourceHelm struct {

	// FileParameters are file parameters to the helm template
	FileParameters []*V1alpha1HelmFileParameter `json:"fileParameters"`

	// IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values
	IgnoreMissingValueFiles bool `json:"ignoreMissingValueFiles,omitempty"`

	// Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation
	Parameters []*V1alpha1HelmParameter `json:"parameters"`

	// PassCredentials pass credentials to all domains (Helm's --pass-credentials)
	PassCredentials bool `json:"passCredentials,omitempty"`

	// ReleaseName is the Helm release name to use. If omitted it will use the application name
	ReleaseName string `json:"releaseName,omitempty"`

	// SkipCrds skips custom resource definition installation step (Helm's --skip-crds)
	SkipCrds bool `json:"skipCrds,omitempty"`

	// ValuesFiles is a list of Helm value files to use when generating a template
	ValueFiles []string `json:"valueFiles"`

	// Values specifies Helm values to be passed to helm template, typically defined as a block
	Values string `json:"values,omitempty"`

	// Version is the Helm version to use for templating ("3")
	Version string `json:"version,omitempty"`
}

// Validate validates this v1alpha1 application source helm
func (m *V1alpha1ApplicationSourceHelm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationSourceHelm) validateFileParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.FileParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.FileParameters); i++ {
		if swag.IsZero(m.FileParameters[i]) { // not required
			continue
		}

		if m.FileParameters[i] != nil {
			if err := m.FileParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationSourceHelm) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1alpha1 application source helm based on the context it is used
func (m *V1alpha1ApplicationSourceHelm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationSourceHelm) contextValidateFileParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileParameters); i++ {

		if m.FileParameters[i] != nil {
			if err := m.FileParameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationSourceHelm) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ApplicationSourceHelm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ApplicationSourceHelm) UnmarshalBinary(b []byte) error {
	var res V1alpha1ApplicationSourceHelm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
