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

// AuditLog audit log
//
// swagger:model AuditLog
type AuditLog struct {

	// The description of the action recorded in the access log entry
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The time, in epoch seconds, that the action recorded in the access log entry occurred
	// Read Only: true
	HappenedOn int64 `json:"happenedOn,omitempty"`

	// The date and time that the action recorded in the access log entry occured
	// Read Only: true
	HappenedOnLocal string `json:"happenedOnLocal,omitempty"`

	// The Id of the access log entry
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The IP address that the action was performed from
	// Read Only: true
	IP string `json:"ip,omitempty"`

	// The username associated with the user that performed the action recorded in the access log entry
	// Read Only: true
	Username string `json:"username,omitempty"`
}

// Validate validates this audit log
func (m *AuditLog) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this audit log based on the context it is used
func (m *AuditLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHappenedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHappenedOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsername(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) contextValidateHappenedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "happenedOn", "body", int64(m.HappenedOn)); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) contextValidateHappenedOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "happenedOnLocal", "body", string(m.HappenedOnLocal)); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ip", "body", string(m.IP)); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) contextValidateUsername(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "username", "body", string(m.Username)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLog) UnmarshalBinary(b []byte) error {
	var res AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}