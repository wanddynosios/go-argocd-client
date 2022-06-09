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

// V1alpha1Operation Operation contains information about a requested or running operation
//
// swagger:model v1alpha1Operation
type V1alpha1Operation struct {

	// Info is a list of informational items for this operation
	Info []*V1alpha1Info `json:"info"`

	// initiated by
	InitiatedBy *V1alpha1OperationInitiator `json:"initiatedBy,omitempty"`

	// retry
	Retry *V1alpha1RetryStrategy `json:"retry,omitempty"`

	// sync
	Sync *V1alpha1SyncOperation `json:"sync,omitempty"`
}

// Validate validates this v1alpha1 operation
func (m *V1alpha1Operation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSync(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1Operation) validateInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Info) { // not required
		return nil
	}

	for i := 0; i < len(m.Info); i++ {
		if swag.IsZero(m.Info[i]) { // not required
			continue
		}

		if m.Info[i] != nil {
			if err := m.Info[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1Operation) validateInitiatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.InitiatedBy) { // not required
		return nil
	}

	if m.InitiatedBy != nil {
		if err := m.InitiatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiatedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("initiatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Operation) validateRetry(formats strfmt.Registry) error {
	if swag.IsZero(m.Retry) { // not required
		return nil
	}

	if m.Retry != nil {
		if err := m.Retry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retry")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Operation) validateSync(formats strfmt.Registry) error {
	if swag.IsZero(m.Sync) { // not required
		return nil
	}

	if m.Sync != nil {
		if err := m.Sync.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha1 operation based on the context it is used
func (m *V1alpha1Operation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRetry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSync(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1Operation) contextValidateInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Info); i++ {

		if m.Info[i] != nil {
			if err := m.Info[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1Operation) contextValidateInitiatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.InitiatedBy != nil {
		if err := m.InitiatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiatedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("initiatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Operation) contextValidateRetry(ctx context.Context, formats strfmt.Registry) error {

	if m.Retry != nil {
		if err := m.Retry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retry")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Operation) contextValidateSync(ctx context.Context, formats strfmt.Registry) error {

	if m.Sync != nil {
		if err := m.Sync.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1Operation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1Operation) UnmarshalBinary(b []byte) error {
	var res V1alpha1Operation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
