// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestoreUnit restore unit
//
// swagger:model RestoreUnit
type RestoreUnit struct {

	// keyspace
	Keyspace string `json:"keyspace,omitempty"`

	// tables
	Tables []*RestoreTable `json:"tables"`
}

// Validate validates this restore unit
func (m *RestoreUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreUnit) validateTables(formats strfmt.Registry) error {

	if swag.IsZero(m.Tables) { // not required
		return nil
	}

	for i := 0; i < len(m.Tables); i++ {
		if swag.IsZero(m.Tables[i]) { // not required
			continue
		}

		if m.Tables[i] != nil {
			if err := m.Tables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreUnit) UnmarshalBinary(b []byte) error {
	var res RestoreUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
