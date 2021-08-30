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
	"github.com/go-openapi/validate"
)

// ScriptAutoDiscoveryMethod script auto discovery method
//
// swagger:model ScriptAutoDiscoveryMethod
type ScriptAutoDiscoveryMethod struct {

	// groovy script
	GroovyScript string `json:"groovyScript,omitempty"`

	// linux cmdline
	LinuxCmdline string `json:"linuxCmdline,omitempty"`

	// linux script
	LinuxScript string `json:"linuxScript,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`

	// win cmdline
	WinCmdline string `json:"winCmdline,omitempty"`

	// win script
	WinScript string `json:"winScript,omitempty"`
}

// Name gets the name of this subtype
func (m *ScriptAutoDiscoveryMethod) Name() string {
	return "ad_script"
}

// SetName sets the name of this subtype
func (m *ScriptAutoDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ScriptAutoDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// groovy script
		GroovyScript string `json:"groovyScript,omitempty"`

		// linux cmdline
		LinuxCmdline string `json:"linuxCmdline,omitempty"`

		// linux script
		LinuxScript string `json:"linuxScript,omitempty"`

		// type
		// Required: true
		Type *string `json:"type"`

		// win cmdline
		WinCmdline string `json:"winCmdline,omitempty"`

		// win script
		WinScript string `json:"winScript,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ScriptAutoDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.GroovyScript = data.GroovyScript
	result.LinuxCmdline = data.LinuxCmdline
	result.LinuxScript = data.LinuxScript
	result.Type = data.Type
	result.WinCmdline = data.WinCmdline
	result.WinScript = data.WinScript

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ScriptAutoDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// groovy script
		GroovyScript string `json:"groovyScript,omitempty"`

		// linux cmdline
		LinuxCmdline string `json:"linuxCmdline,omitempty"`

		// linux script
		LinuxScript string `json:"linuxScript,omitempty"`

		// type
		// Required: true
		Type *string `json:"type"`

		// win cmdline
		WinCmdline string `json:"winCmdline,omitempty"`

		// win script
		WinScript string `json:"winScript,omitempty"`
	}{

		GroovyScript: m.GroovyScript,

		LinuxCmdline: m.LinuxCmdline,

		LinuxScript: m.LinuxScript,

		Type: m.Type,

		WinCmdline: m.WinCmdline,

		WinScript: m.WinScript,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this script auto discovery method
func (m *ScriptAutoDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScriptAutoDiscoveryMethod) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this script auto discovery method based on the context it is used
func (m *ScriptAutoDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ScriptAutoDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptAutoDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res ScriptAutoDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}