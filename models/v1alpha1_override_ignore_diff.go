// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1OverrideIgnoreDiff OverrideIgnoreDiff contains configurations about how fields should be ignored during diffs between
// the desired state and live state
//
// swagger:model v1alpha1OverrideIgnoreDiff
type V1alpha1OverrideIgnoreDiff struct {

	// JSONPointers is a JSON path list following the format defined in RFC4627 (https://datatracker.ietf.org/doc/html/rfc6902#section-3)
	JSONPointers []string `json:"jSONPointers"`

	// JQPathExpressions is a JQ path list that will be evaludated during the diff process
	JqPathExpressions []string `json:"jqPathExpressions"`

	// ManagedFieldsManagers is a list of trusted managers. Fields mutated by those managers will take precedence over the
	// desired state defined in the SCM and won't be displayed in diffs
	ManagedFieldsManagers []string `json:"managedFieldsManagers"`
}

// Validate validates this v1alpha1 override ignore diff
func (m *V1alpha1OverrideIgnoreDiff) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha1 override ignore diff based on context it is used
func (m *V1alpha1OverrideIgnoreDiff) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1OverrideIgnoreDiff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1OverrideIgnoreDiff) UnmarshalBinary(b []byte) error {
	var res V1alpha1OverrideIgnoreDiff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}