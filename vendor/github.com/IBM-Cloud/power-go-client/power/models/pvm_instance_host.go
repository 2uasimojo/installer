// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PvmInstanceHost pvm instance host
//
// swagger:model PvmInstanceHost
type PvmInstanceHost struct {

	// ID of the dedicated host where the PVM Instance is running, if applicable
	DedicatedHostID string `json:"dedicatedHostID,omitempty"`

	// The PVM Instance Host ID (Internal Use Only)
	ID int64 `json:"id,omitempty"`
}

// Validate validates this pvm instance host
func (m *PvmInstanceHost) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pvm instance host based on context it is used
func (m *PvmInstanceHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PvmInstanceHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PvmInstanceHost) UnmarshalBinary(b []byte) error {
	var res PvmInstanceHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
