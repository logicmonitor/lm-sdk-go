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

// AzureVNGConnectionDiscoveryMethod azure v n g connection discovery method
//
// swagger:model AzureVNGConnectionDiscoveryMethod
type AzureVNGConnectionDiscoveryMethod struct {
	AzureVNGConnectionDiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *AzureVNGConnectionDiscoveryMethod) Name() string {
	return "AzureVNGConnectionDiscoveryMethod"
}

// SetName sets the name of this subtype
func (m *AzureVNGConnectionDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AzureVNGConnectionDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		AzureVNGConnectionDiscoveryMethodAllOf1
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

	var result AzureVNGConnectionDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.AzureVNGConnectionDiscoveryMethodAllOf1 = data.AzureVNGConnectionDiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AzureVNGConnectionDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AzureVNGConnectionDiscoveryMethodAllOf1
	}{

		AzureVNGConnectionDiscoveryMethodAllOf1: m.AzureVNGConnectionDiscoveryMethodAllOf1,
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

// Validate validates this azure v n g connection discovery method
func (m *AzureVNGConnectionDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AzureVNGConnectionDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this azure v n g connection discovery method based on the context it is used
func (m *AzureVNGConnectionDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AzureVNGConnectionDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AzureVNGConnectionDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureVNGConnectionDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res AzureVNGConnectionDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AzureVNGConnectionDiscoveryMethodAllOf1 azure v n g connection discovery method all of1
//
// swagger:model AzureVNGConnectionDiscoveryMethodAllOf1
type AzureVNGConnectionDiscoveryMethodAllOf1 interface{}
