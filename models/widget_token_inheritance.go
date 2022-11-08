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

// WidgetTokenInheritance widget token inheritance
//
// swagger:model WidgetTokenInheritance
type WidgetTokenInheritance struct {

	// The fullpath for the widget token
	// Read Only: true
	Fullpath string `json:"fullpath,omitempty"`

	// The property value for the group
	// Read Only: true
	Value string `json:"value,omitempty"`
}

// Validate validates this widget token inheritance
func (m *WidgetTokenInheritance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this widget token inheritance based on the context it is used
func (m *WidgetTokenInheritance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullpath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WidgetTokenInheritance) contextValidateFullpath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fullpath", "body", string(m.Fullpath)); err != nil {
		return err
	}

	return nil
}

func (m *WidgetTokenInheritance) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", string(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WidgetTokenInheritance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WidgetTokenInheritance) UnmarshalBinary(b []byte) error {
	var res WidgetTokenInheritance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
