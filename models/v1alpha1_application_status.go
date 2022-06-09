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

// V1alpha1ApplicationStatus ApplicationStatus contains status information for the application
//
// swagger:model v1alpha1ApplicationStatus
type V1alpha1ApplicationStatus struct {

	// Conditions is a list of currently observed application conditions
	Conditions []*V1alpha1ApplicationCondition `json:"conditions"`

	// health
	Health *V1alpha1HealthStatus `json:"health,omitempty"`

	// History contains information about the application's sync history
	History []*V1alpha1RevisionHistory `json:"history"`

	// observed at
	ObservedAt *V1Time `json:"observedAt,omitempty"`

	// operation state
	OperationState *V1alpha1OperationState `json:"operationState,omitempty"`

	// reconciled at
	ReconciledAt *V1Time `json:"reconciledAt,omitempty"`

	// Resources is a list of Kubernetes resources managed by this application
	Resources []*V1alpha1ResourceStatus `json:"resources"`

	// SourceType specifies the type of this application
	SourceType string `json:"sourceType,omitempty"`

	// summary
	Summary *V1alpha1ApplicationSummary `json:"summary,omitempty"`

	// sync
	Sync *V1alpha1SyncStatus `json:"sync,omitempty"`
}

// Validate validates this v1alpha1 application status
func (m *V1alpha1ApplicationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObservedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReconciledAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
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

func (m *V1alpha1ApplicationStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.History) { // not required
		return nil
	}

	for i := 0; i < len(m.History); i++ {
		if swag.IsZero(m.History[i]) { // not required
			continue
		}

		if m.History[i] != nil {
			if err := m.History[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("history" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateObservedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ObservedAt) { // not required
		return nil
	}

	if m.ObservedAt != nil {
		if err := m.ObservedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("observedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("observedAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateOperationState(formats strfmt.Registry) error {
	if swag.IsZero(m.OperationState) { // not required
		return nil
	}

	if m.OperationState != nil {
		if err := m.OperationState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operationState")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateReconciledAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ReconciledAt) { // not required
		return nil
	}

	if m.ReconciledAt != nil {
		if err := m.ReconciledAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reconciledAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reconciledAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateSync(formats strfmt.Registry) error {
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

// ContextValidate validate this v1alpha1 application status based on the context it is used
func (m *V1alpha1ApplicationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObservedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperationState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReconciledAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummary(ctx, formats); err != nil {
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

func (m *V1alpha1ApplicationStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {
		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) contextValidateHistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.History); i++ {

		if m.History[i] != nil {
			if err := m.History[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("history" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) contextValidateObservedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.ObservedAt != nil {
		if err := m.ObservedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("observedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("observedAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) contextValidateOperationState(ctx context.Context, formats strfmt.Registry) error {

	if m.OperationState != nil {
		if err := m.OperationState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operationState")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) contextValidateReconciledAt(ctx context.Context, formats strfmt.Registry) error {

	if m.ReconciledAt != nil {
		if err := m.ReconciledAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reconciledAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reconciledAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {
			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) contextValidateSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.Summary != nil {
		if err := m.Summary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) contextValidateSync(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1alpha1ApplicationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ApplicationStatus) UnmarshalBinary(b []byte) error {
	var res V1alpha1ApplicationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}