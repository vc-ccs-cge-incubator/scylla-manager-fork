// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CopyPathsOptions copy paths options
//
// swagger:model CopyPathsOptions
type CopyPathsOptions struct {

	// A remote name string eg. drive: for the destination
	DstFs string `json:"dstFs,omitempty"`

	// A directory within that remote eg. files/ for the destination
	DstRemote string `json:"dstRemote,omitempty"`

	// Paths relative to srcRemote/dstRemote eg. file.txt for both source and destination
	Paths []string `json:"paths"`

	// A remote name string eg. drive: for the source
	SrcFs string `json:"srcFs,omitempty"`

	// A directory within that remote eg. files/ for the source
	SrcRemote string `json:"srcRemote,omitempty"`
}

// Validate validates this copy paths options
func (m *CopyPathsOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CopyPathsOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyPathsOptions) UnmarshalBinary(b []byte) error {
	var res CopyPathsOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}