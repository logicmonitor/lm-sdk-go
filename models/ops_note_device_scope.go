// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpsNoteDeviceScope ops note device scope
//
// swagger:model OpsNoteDeviceScope
type OpsNoteDeviceScope struct {

	// device Id
	// Example: 56
	DeviceID int32 `json:"deviceId,omitempty"`

	// device name
	DeviceName string `json:"deviceName,omitempty"`

	// full path
	FullPath string `json:"fullPath,omitempty"`

	// group Id
	GroupID int32 `json:"groupId,omitempty"`
}

// Type gets the type of this subtype
func (m *OpsNoteDeviceScope) Type() string {
	return "OpsNoteDeviceScope"
}

// SetType sets the type of this subtype
func (m *OpsNoteDeviceScope) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OpsNoteDeviceScope) UnmarshalJSON(raw []byte) error {
	var data struct {

		// device Id
		// Example: 56
		DeviceID int32 `json:"deviceId,omitempty"`

		// device name
		DeviceName string `json:"deviceName,omitempty"`

		// full path
		FullPath string `json:"fullPath,omitempty"`

		// group Id
		GroupID int32 `json:"groupId,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result OpsNoteDeviceScope

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.DeviceID = data.DeviceID
	result.DeviceName = data.DeviceName
	result.FullPath = data.FullPath
	result.GroupID = data.GroupID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OpsNoteDeviceScope) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// device Id
		// Example: 56
		DeviceID int32 `json:"deviceId,omitempty"`

		// device name
		DeviceName string `json:"deviceName,omitempty"`

		// full path
		FullPath string `json:"fullPath,omitempty"`

		// group Id
		GroupID int32 `json:"groupId,omitempty"`
	}{

		DeviceID: m.DeviceID,

		DeviceName: m.DeviceName,

		FullPath: m.FullPath,

		GroupID: m.GroupID,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this ops note device scope
func (m *OpsNoteDeviceScope) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this ops note device scope based on the context it is used
func (m *OpsNoteDeviceScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpsNoteDeviceScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpsNoteDeviceScope) UnmarshalBinary(b []byte) error {
	var res OpsNoteDeviceScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
