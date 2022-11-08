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

// SaasOffice365HealthCollectorAttributeV3 saas office365 health collector attribute v3
//
// swagger:model SaasOffice365HealthCollectorAttributeV3
type SaasOffice365HealthCollectorAttributeV3 struct {
	SaasOffice365HealthCollectorAttributeV3AllOf1
}

// Name gets the name of this subtype
func (m *SaasOffice365HealthCollectorAttributeV3) Name() string {
	return "SaasOffice365HealthCollectorAttributeV3"
}

// SetName sets the name of this subtype
func (m *SaasOffice365HealthCollectorAttributeV3) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SaasOffice365HealthCollectorAttributeV3) UnmarshalJSON(raw []byte) error {
	var data struct {
		SaasOffice365HealthCollectorAttributeV3AllOf1
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

	var result SaasOffice365HealthCollectorAttributeV3

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.SaasOffice365HealthCollectorAttributeV3AllOf1 = data.SaasOffice365HealthCollectorAttributeV3AllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SaasOffice365HealthCollectorAttributeV3) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		SaasOffice365HealthCollectorAttributeV3AllOf1
	}{

		SaasOffice365HealthCollectorAttributeV3AllOf1: m.SaasOffice365HealthCollectorAttributeV3AllOf1,
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

// Validate validates this saas office365 health collector attribute v3
func (m *SaasOffice365HealthCollectorAttributeV3) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SaasOffice365HealthCollectorAttributeV3AllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this saas office365 health collector attribute v3 based on the context it is used
func (m *SaasOffice365HealthCollectorAttributeV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SaasOffice365HealthCollectorAttributeV3AllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SaasOffice365HealthCollectorAttributeV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SaasOffice365HealthCollectorAttributeV3) UnmarshalBinary(b []byte) error {
	var res SaasOffice365HealthCollectorAttributeV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SaasOffice365HealthCollectorAttributeV3AllOf1 saas office365 health collector attribute v3 all of1
//
// swagger:model SaasOffice365HealthCollectorAttributeV3AllOf1
type SaasOffice365HealthCollectorAttributeV3AllOf1 interface{}
