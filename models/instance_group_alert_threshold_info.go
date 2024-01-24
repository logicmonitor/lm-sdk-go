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

// InstanceGroupAlertThresholdInfo instance group alert threshold info
//
// swagger:model InstanceGroupAlertThresholdInfo
type InstanceGroupAlertThresholdInfo struct {

	// The count that the alert must exist for this many poll cycles before the alert will be cleared
	// Format: byte
	AlertClearTransitionInterval strfmt.Base64 `json:"alertClearTransitionInterval,omitempty"`

	// alert enabled
	// Read Only: true
	AlertEnabled *bool `json:"alertEnabled,omitempty"`

	// alert expr
	// Read Only: true
	AlertExpr string `json:"alertExpr,omitempty"`

	// The triggered alert level if we cannot collect data for this datapoint. The values can be 0-4 (0:unused alert, 1:alert ok, 2:warn alert, 2:error alert, 4:critical alert)
	// Format: byte
	AlertForNoData strfmt.Base64 `json:"alertForNoData,omitempty"`

	// The count that the alert must exist for this many poll cycles before it will be triggered
	// Format: byte
	AlertTransitionInterval strfmt.Base64 `json:"alertTransitionInterval,omitempty"`

	// enable anomaly alert generation
	// Read Only: true
	EnableAnomalyAlertGeneration string `json:"enableAnomalyAlertGeneration,omitempty"`

	// enable anomaly alert suppression
	// Read Only: true
	EnableAnomalyAlertSuppression string `json:"enableAnomalyAlertSuppression,omitempty"`

	// group Id
	// Read Only: true
	GroupID int32 `json:"groupId,omitempty"`
}

// Validate validates this instance group alert threshold info
func (m *InstanceGroupAlertThresholdInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this instance group alert threshold info based on the context it is used
func (m *InstanceGroupAlertThresholdInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertEnabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAlertExpr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnableAnomalyAlertGeneration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnableAnomalyAlertSuppression(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupAlertThresholdInfo) contextValidateAlertEnabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "alertEnabled", "body", m.AlertEnabled); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupAlertThresholdInfo) contextValidateAlertExpr(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "alertExpr", "body", string(m.AlertExpr)); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupAlertThresholdInfo) contextValidateEnableAnomalyAlertGeneration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableAnomalyAlertGeneration", "body", string(m.EnableAnomalyAlertGeneration)); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupAlertThresholdInfo) contextValidateEnableAnomalyAlertSuppression(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableAnomalyAlertSuppression", "body", string(m.EnableAnomalyAlertSuppression)); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupAlertThresholdInfo) contextValidateGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "groupId", "body", int32(m.GroupID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupAlertThresholdInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupAlertThresholdInfo) UnmarshalBinary(b []byte) error {
	var res InstanceGroupAlertThresholdInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
