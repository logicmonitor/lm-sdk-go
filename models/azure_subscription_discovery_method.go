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

// AzureSubscriptionDiscoveryMethod azure subscription discovery method
//
// swagger:model AzureSubscriptionDiscoveryMethod
type AzureSubscriptionDiscoveryMethod struct {
	AzureSubscriptionDiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *AzureSubscriptionDiscoveryMethod) Name() string {
	return "AzureSubscriptionDiscoveryMethod"
}

// SetName sets the name of this subtype
func (m *AzureSubscriptionDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AzureSubscriptionDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		AzureSubscriptionDiscoveryMethodAllOf1
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

	var result AzureSubscriptionDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.AzureSubscriptionDiscoveryMethodAllOf1 = data.AzureSubscriptionDiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AzureSubscriptionDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AzureSubscriptionDiscoveryMethodAllOf1
	}{

		AzureSubscriptionDiscoveryMethodAllOf1: m.AzureSubscriptionDiscoveryMethodAllOf1,
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

// Validate validates this azure subscription discovery method
func (m *AzureSubscriptionDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AzureSubscriptionDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this azure subscription discovery method based on the context it is used
func (m *AzureSubscriptionDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AzureSubscriptionDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AzureSubscriptionDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureSubscriptionDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res AzureSubscriptionDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AzureSubscriptionDiscoveryMethodAllOf1 azure subscription discovery method all of1
//
// swagger:model AzureSubscriptionDiscoveryMethodAllOf1
type AzureSubscriptionDiscoveryMethodAllOf1 interface{}
