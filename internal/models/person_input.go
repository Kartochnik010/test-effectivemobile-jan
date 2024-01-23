// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PersonInput person input
//
// swagger:model PersonInput
type PersonInput struct {

	// name
	Name string `json:"name,omitempty"`

	// patronymic
	Patronymic string `json:"patronymic,omitempty"`

	// surname
	Surname string `json:"surname,omitempty"`
}

// Validate validates this person input
func (m *PersonInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PersonInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonInput) UnmarshalBinary(b []byte) error {
	var res PersonInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}