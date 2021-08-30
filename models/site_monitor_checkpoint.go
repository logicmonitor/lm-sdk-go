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

// SiteMonitorCheckpoint site monitor checkpoint
//
// swagger:model SiteMonitorCheckpoint
type SiteMonitorCheckpoint struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// display prio
	DisplayPrio int32 `json:"displayPrio,omitempty"`

	// geo info
	// Read Only: true
	GeoInfo string `json:"geoInfo,omitempty"`

	// id
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// is enabled in root
	// Read Only: true
	IsEnabledInRoot *bool `json:"isEnabledInRoot,omitempty"`

	// name
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this site monitor checkpoint
func (m *SiteMonitorCheckpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteMonitorCheckpoint) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this site monitor checkpoint based on the context it is used
func (m *SiteMonitorCheckpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGeoInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsEnabledInRoot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteMonitorCheckpoint) contextValidateGeoInfo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "geoInfo", "body", string(m.GeoInfo)); err != nil {
		return err
	}

	return nil
}

func (m *SiteMonitorCheckpoint) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *SiteMonitorCheckpoint) contextValidateIsEnabledInRoot(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isEnabledInRoot", "body", m.IsEnabledInRoot); err != nil {
		return err
	}

	return nil
}

func (m *SiteMonitorCheckpoint) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteMonitorCheckpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteMonitorCheckpoint) UnmarshalBinary(b []byte) error {
	var res SiteMonitorCheckpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}