// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GroupKind GroupKind specifies a Group and a Kind, but does not force a version.  This is useful for identifying
// concepts during lookup stages without having partially valid types
//
// +protobuf.options.(gogoproto.goproto_stringer)=false
//
// swagger:model v1GroupKind
type V1GroupKind struct {

	// group
	Group string `json:"group,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`
}

// Validate validates this v1 group kind
func (m *V1GroupKind) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 group kind based on context it is used
func (m *V1GroupKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GroupKind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GroupKind) UnmarshalBinary(b []byte) error {
	var res V1GroupKind
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
