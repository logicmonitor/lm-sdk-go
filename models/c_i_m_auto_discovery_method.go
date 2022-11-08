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

// CIMAutoDiscoveryMethod c i m auto discovery method
//
// swagger:model CIMAutoDiscoveryMethod
type CIMAutoDiscoveryMethod struct {

	// cim class
	// Required: true
	CimClass *string `json:"cimClass"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// property
	// Required: true
	Property *string `json:"property"`
}

// Name gets the name of this subtype
func (m *CIMAutoDiscoveryMethod) Name() string {
	return "CIMAutoDiscoveryMethod"
}

// SetName sets the name of this subtype
func (m *CIMAutoDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CIMAutoDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// cim class
		// Required: true
		CimClass *string `json:"cimClass"`

		// namespace
		// Required: true
		Namespace *string `json:"namespace"`

		// property
		// Required: true
		Property *string `json:"property"`
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

	var result CIMAutoDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.CimClass = data.CimClass
	result.Namespace = data.Namespace
	result.Property = data.Property

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CIMAutoDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// cim class
		// Required: true
		CimClass *string `json:"cimClass"`

		// namespace
		// Required: true
		Namespace *string `json:"namespace"`

		// property
		// Required: true
		Property *string `json:"property"`
	}{

		CimClass: m.CimClass,

		Namespace: m.Namespace,

		Property: m.Property,
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

// Validate validates this c i m auto discovery method
func (m *CIMAutoDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCimClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CIMAutoDiscoveryMethod) validateCimClass(formats strfmt.Registry) error {

	if err := validate.Required("cimClass", "body", m.CimClass); err != nil {
		return err
	}

	return nil
}

func (m *CIMAutoDiscoveryMethod) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *CIMAutoDiscoveryMethod) validateProperty(formats strfmt.Registry) error {

	if err := validate.Required("property", "body", m.Property); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this c i m auto discovery method based on the context it is used
func (m *CIMAutoDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CIMAutoDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CIMAutoDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res CIMAutoDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
