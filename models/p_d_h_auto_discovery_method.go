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

// PDHAutoDiscoveryMethod p d h auto discovery method
//
// swagger:model PDHAutoDiscoveryMethod
type PDHAutoDiscoveryMethod struct {

	// category
	// Required: true
	Category *string `json:"category"`

	// obj regex
	// Required: true
	ObjRegex *string `json:"objRegex"`
}

// Name gets the name of this subtype
func (m *PDHAutoDiscoveryMethod) Name() string {
	return "PDHAutoDiscoveryMethod"
}

// SetName sets the name of this subtype
func (m *PDHAutoDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PDHAutoDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// category
		// Required: true
		Category *string `json:"category"`

		// obj regex
		// Required: true
		ObjRegex *string `json:"objRegex"`
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

	var result PDHAutoDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.Category = data.Category
	result.ObjRegex = data.ObjRegex

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PDHAutoDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// category
		// Required: true
		Category *string `json:"category"`

		// obj regex
		// Required: true
		ObjRegex *string `json:"objRegex"`
	}{

		Category: m.Category,

		ObjRegex: m.ObjRegex,
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

// Validate validates this p d h auto discovery method
func (m *PDHAutoDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjRegex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PDHAutoDiscoveryMethod) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *PDHAutoDiscoveryMethod) validateObjRegex(formats strfmt.Registry) error {

	if err := validate.Required("objRegex", "body", m.ObjRegex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p d h auto discovery method based on the context it is used
func (m *PDHAutoDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PDHAutoDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PDHAutoDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res PDHAutoDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
