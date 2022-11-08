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

// SaaSO365SharePointSiteDiscoveryMethod saa s o365 share point site discovery method
//
// swagger:model SaaSO365SharePointSiteDiscoveryMethod
type SaaSO365SharePointSiteDiscoveryMethod struct {
	SaaSO365SharePointSiteDiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *SaaSO365SharePointSiteDiscoveryMethod) Name() string {
	return "SaaSO365SharePointSiteDiscoveryMethod"
}

// SetName sets the name of this subtype
func (m *SaaSO365SharePointSiteDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SaaSO365SharePointSiteDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		SaaSO365SharePointSiteDiscoveryMethodAllOf1
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

	var result SaaSO365SharePointSiteDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.SaaSO365SharePointSiteDiscoveryMethodAllOf1 = data.SaaSO365SharePointSiteDiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SaaSO365SharePointSiteDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		SaaSO365SharePointSiteDiscoveryMethodAllOf1
	}{

		SaaSO365SharePointSiteDiscoveryMethodAllOf1: m.SaaSO365SharePointSiteDiscoveryMethodAllOf1,
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

// Validate validates this saa s o365 share point site discovery method
func (m *SaaSO365SharePointSiteDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SaaSO365SharePointSiteDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this saa s o365 share point site discovery method based on the context it is used
func (m *SaaSO365SharePointSiteDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SaaSO365SharePointSiteDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SaaSO365SharePointSiteDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SaaSO365SharePointSiteDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res SaaSO365SharePointSiteDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SaaSO365SharePointSiteDiscoveryMethodAllOf1 saa s o365 share point site discovery method all of1
//
// swagger:model SaaSO365SharePointSiteDiscoveryMethodAllOf1
type SaaSO365SharePointSiteDiscoveryMethodAllOf1 interface{}
