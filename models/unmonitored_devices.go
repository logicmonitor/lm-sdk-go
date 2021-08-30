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

// UnmonitoredDevices unmonitored devices
//
// swagger:model UnmonitoredDevices
type UnmonitoredDevices struct {

	// collector description
	// Read Only: true
	CollectorDescription string `json:"collectorDescription,omitempty"`

	// collector Id
	// Read Only: true
	CollectorID int32 `json:"collectorId,omitempty"`

	// device status
	// Read Only: true
	DeviceStatus string `json:"deviceStatus,omitempty"`

	// device type
	// Read Only: true
	DeviceType string `json:"deviceType,omitempty"`

	// display as
	// Read Only: true
	DisplayAs string `json:"displayAs,omitempty"`

	// dns
	// Read Only: true
	DNS string `json:"dns,omitempty"`

	// end date
	// Read Only: true
	EndDate string `json:"endDate,omitempty"`

	// end timestamp
	// Read Only: true
	EndTimestamp int64 `json:"endTimestamp,omitempty"`

	// forward Ip
	// Read Only: true
	ForwardIP string `json:"forwardIp,omitempty"`

	// id
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// ip
	// Read Only: true
	IP string `json:"ip,omitempty"`

	// manufacturer
	// Read Only: true
	Manufacturer string `json:"manufacturer,omitempty"`

	// nse Id
	// Read Only: true
	NseID int32 `json:"nseId,omitempty"`

	// nse scan Id
	// Read Only: true
	NseScanID string `json:"nseScanId,omitempty"`

	// nsp Id
	// Read Only: true
	NspID int32 `json:"nspId,omitempty"`

	// nsp name
	// Read Only: true
	NspName string `json:"nspName,omitempty"`

	// ports
	// Read Only: true
	Ports string `json:"ports,omitempty"`

	// status
	// Read Only: true
	Status string `json:"status,omitempty"`

	// sys name
	// Read Only: true
	SysName string `json:"sysName,omitempty"`
}

// Validate validates this unmonitored devices
func (m *UnmonitoredDevices) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this unmonitored devices based on the context it is used
func (m *UnmonitoredDevices) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectorDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplayAs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForwardIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManufacturer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNseID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNseScanID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNspID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNspName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSysName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnmonitoredDevices) contextValidateCollectorDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorDescription", "body", string(m.CollectorDescription)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateCollectorID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorId", "body", int32(m.CollectorID)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateDeviceStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceStatus", "body", string(m.DeviceStatus)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceType", "body", string(m.DeviceType)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateDisplayAs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "displayAs", "body", string(m.DisplayAs)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateDNS(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dns", "body", string(m.DNS)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateEndDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endDate", "body", string(m.EndDate)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateEndTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endTimestamp", "body", int64(m.EndTimestamp)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateForwardIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "forwardIp", "body", string(m.ForwardIP)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ip", "body", string(m.IP)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateManufacturer(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "manufacturer", "body", string(m.Manufacturer)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateNseID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nseId", "body", int32(m.NseID)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateNseScanID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nseScanId", "body", string(m.NseScanID)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateNspID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nspId", "body", int32(m.NspID)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateNspName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nspName", "body", string(m.NspName)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ports", "body", string(m.Ports)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *UnmonitoredDevices) contextValidateSysName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sysName", "body", string(m.SysName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnmonitoredDevices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnmonitoredDevices) UnmarshalBinary(b []byte) error {
	var res UnmonitoredDevices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}