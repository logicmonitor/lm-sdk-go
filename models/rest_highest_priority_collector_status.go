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

// RestHighestPriorityCollectorStatus rest highest priority collector status
//
// swagger:model RestHighestPriorityCollectorStatus
type RestHighestPriorityCollectorStatus struct {

	// The acked status of the highest priority sub collector
	// Read Only: true
	Acked *bool `json:"acked,omitempty"`

	// The SDT status of the highest priority sub collector
	// Read Only: true
	InSDT *bool `json:"inSDT,omitempty"`

	// The down status of the highest priority sub collector
	// Read Only: true
	IsDown *bool `json:"isDown,omitempty"`

	// The status of the highest priority sub collector
	// Read Only: true
	Status int32 `json:"status,omitempty"`
}

// Validate validates this rest highest priority collector status
func (m *RestHighestPriorityCollectorStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this rest highest priority collector status based on the context it is used
func (m *RestHighestPriorityCollectorStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcked(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInSDT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsDown(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestHighestPriorityCollectorStatus) contextValidateAcked(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "acked", "body", m.Acked); err != nil {
		return err
	}

	return nil
}

func (m *RestHighestPriorityCollectorStatus) contextValidateInSDT(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "inSDT", "body", m.InSDT); err != nil {
		return err
	}

	return nil
}

func (m *RestHighestPriorityCollectorStatus) contextValidateIsDown(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isDown", "body", m.IsDown); err != nil {
		return err
	}

	return nil
}

func (m *RestHighestPriorityCollectorStatus) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", int32(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestHighestPriorityCollectorStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestHighestPriorityCollectorStatus) UnmarshalBinary(b []byte) error {
	var res RestHighestPriorityCollectorStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
