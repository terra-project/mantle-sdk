// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyConstraints policy constraints
//
// swagger:model PolicyConstraints
type PolicyConstraints struct {

	// cap
	Cap *Coin `json:"cap,omitempty"`

	// 0.025%
	ChangeMax string `json:"change_max,omitempty"`

	// 1%
	RateMax string `json:"rate_max,omitempty"`

	// 0.05%
	RateMin string `json:"rate_min,omitempty"`
}

// Validate validates this policy constraints
func (m *PolicyConstraints) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyConstraints) validateCap(formats strfmt.Registry) error {

	if swag.IsZero(m.Cap) { // not required
		return nil
	}

	if m.Cap != nil {
		if err := m.Cap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyConstraints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyConstraints) UnmarshalBinary(b []byte) error {
	var res PolicyConstraints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
