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

// Ec2DDR ec2 d d r
//
// swagger:model Ec2DDR
type Ec2DDR struct {

	// assignment
	Assignment []*Assignment `json:"assignment"`

	// change name
	ChangeName string `json:"changeName,omitempty"`
}

// Validate validates this ec2 d d r
func (m *Ec2DDR) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ec2DDR) validateAssignment(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignment) { // not required
		return nil
	}

	for i := 0; i < len(m.Assignment); i++ {
		if swag.IsZero(m.Assignment[i]) { // not required
			continue
		}

		if m.Assignment[i] != nil {
			if err := m.Assignment[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignment" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignment" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ec2 d d r based on the context it is used
func (m *Ec2DDR) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ec2DDR) contextValidateAssignment(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Assignment); i++ {

		if m.Assignment[i] != nil {
			if err := m.Assignment[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignment" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignment" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Ec2DDR) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ec2DDR) UnmarshalBinary(b []byte) error {
	var res Ec2DDR
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
