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

// EC2ScheduledEventAutoDiscoveryMethod e c2 scheduled event auto discovery method
//
// swagger:model EC2ScheduledEventAutoDiscoveryMethod
type EC2ScheduledEventAutoDiscoveryMethod struct {
	EC2ScheduledEventAutoDiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *EC2ScheduledEventAutoDiscoveryMethod) Name() string {
	return "EC2ScheduledEventAutoDiscoveryMethod"
}

// SetName sets the name of this subtype
func (m *EC2ScheduledEventAutoDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *EC2ScheduledEventAutoDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		EC2ScheduledEventAutoDiscoveryMethodAllOf1
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

	var result EC2ScheduledEventAutoDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.EC2ScheduledEventAutoDiscoveryMethodAllOf1 = data.EC2ScheduledEventAutoDiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m EC2ScheduledEventAutoDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		EC2ScheduledEventAutoDiscoveryMethodAllOf1
	}{

		EC2ScheduledEventAutoDiscoveryMethodAllOf1: m.EC2ScheduledEventAutoDiscoveryMethodAllOf1,
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

// Validate validates this e c2 scheduled event auto discovery method
func (m *EC2ScheduledEventAutoDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EC2ScheduledEventAutoDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this e c2 scheduled event auto discovery method based on the context it is used
func (m *EC2ScheduledEventAutoDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EC2ScheduledEventAutoDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *EC2ScheduledEventAutoDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EC2ScheduledEventAutoDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res EC2ScheduledEventAutoDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EC2ScheduledEventAutoDiscoveryMethodAllOf1 e c2 scheduled event auto discovery method all of1
//
// swagger:model EC2ScheduledEventAutoDiscoveryMethodAllOf1
type EC2ScheduledEventAutoDiscoveryMethodAllOf1 interface{}
