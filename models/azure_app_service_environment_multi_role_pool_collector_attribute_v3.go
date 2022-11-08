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

// AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3 azure app service environment multi role pool collector attribute v3
//
// swagger:model AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3
type AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3 struct {

	// period
	Period int32 `json:"period,omitempty"`
}

// Name gets the name of this subtype
func (m *AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3) Name() string {
	return "AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3"
}

// SetName sets the name of this subtype
func (m *AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3) UnmarshalJSON(raw []byte) error {
	var data struct {

		// period
		Period int32 `json:"period,omitempty"`
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

	var result AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.Period = data.Period

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// period
		Period int32 `json:"period,omitempty"`
	}{

		Period: m.Period,
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

// Validate validates this azure app service environment multi role pool collector attribute v3
func (m *AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this azure app service environment multi role pool collector attribute v3 based on the context it is used
func (m *AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3) UnmarshalBinary(b []byte) error {
	var res AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
