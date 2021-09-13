// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ColorThreshold color threshold
//
// swagger:model ColorThreshold
type ColorThreshold struct {

	// level
	// Required: true
	Level *int32 `json:"level"`

	// relation
	// Required: true
	Relation *string `json:"relation"`

	// threshold
	// Required: true
	Threshold *float64 `json:"threshold"`
}

// Validate validates this color threshold
func (m *ColorThreshold) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ColorThreshold) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *ColorThreshold) validateRelation(formats strfmt.Registry) error {

	if err := validate.Required("relation", "body", m.Relation); err != nil {
		return err
	}

	return nil
}

func (m *ColorThreshold) validateThreshold(formats strfmt.Registry) error {

	if err := validate.Required("threshold", "body", m.Threshold); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this color threshold based on context it is used
func (m *ColorThreshold) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ColorThreshold) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ColorThreshold) UnmarshalBinary(b []byte) error {
	var res ColorThreshold
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
