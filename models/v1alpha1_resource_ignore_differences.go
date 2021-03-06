// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1ResourceIgnoreDifferences ResourceIgnoreDifferences contains resource filter and list of json paths which should be ignored during comparison with live state.
//
// swagger:model v1alpha1ResourceIgnoreDifferences
type V1alpha1ResourceIgnoreDifferences struct {

	// group
	Group string `json:"group,omitempty"`

	// jq path expressions
	JqPathExpressions []string `json:"jqPathExpressions"`

	// json pointers
	JSONPointers []string `json:"jsonPointers"`

	// kind
	Kind string `json:"kind,omitempty"`

	// ManagedFieldsManagers is a list of trusted managers. Fields mutated by those managers will take precedence over the
	// desired state defined in the SCM and won't be displayed in diffs
	ManagedFieldsManagers []string `json:"managedFieldsManagers"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this v1alpha1 resource ignore differences
func (m *V1alpha1ResourceIgnoreDifferences) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha1 resource ignore differences based on context it is used
func (m *V1alpha1ResourceIgnoreDifferences) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ResourceIgnoreDifferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ResourceIgnoreDifferences) UnmarshalBinary(b []byte) error {
	var res V1alpha1ResourceIgnoreDifferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
