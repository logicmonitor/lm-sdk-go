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

// AzureBillingCollectorAttribute azure billing collector attribute
//
// swagger:model AzureBillingCollectorAttribute
type AzureBillingCollectorAttribute struct {
	AzureBillingCollectorAttributeAllOf1
}

// Name gets the name of this subtype
func (m *AzureBillingCollectorAttribute) Name() string {
	return "AzureBillingCollectorAttribute"
}

// SetName sets the name of this subtype
func (m *AzureBillingCollectorAttribute) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AzureBillingCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {
		AzureBillingCollectorAttributeAllOf1
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

	var result AzureBillingCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.AzureBillingCollectorAttributeAllOf1 = data.AzureBillingCollectorAttributeAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AzureBillingCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AzureBillingCollectorAttributeAllOf1
	}{

		AzureBillingCollectorAttributeAllOf1: m.AzureBillingCollectorAttributeAllOf1,
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

// Validate validates this azure billing collector attribute
func (m *AzureBillingCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AzureBillingCollectorAttributeAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this azure billing collector attribute based on the context it is used
func (m *AzureBillingCollectorAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AzureBillingCollectorAttributeAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AzureBillingCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureBillingCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res AzureBillingCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AzureBillingCollectorAttributeAllOf1 azure billing collector attribute all of1
//
// swagger:model AzureBillingCollectorAttributeAllOf1
type AzureBillingCollectorAttributeAllOf1 interface{}
