// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SDTPaginationResponse SDT pagination response
//
// swagger:model SDTPaginationResponse
type SDTPaginationResponse struct {
	itemsField []SDT

	// search Id
	// Read Only: true
	SearchID string `json:"searchId,omitempty"`

	// total
	// Read Only: true
	Total int32 `json:"total,omitempty"`
}

// Items gets the items of this base type
func (m *SDTPaginationResponse) Items() []SDT {
	return m.itemsField
}

// SetItems sets the items of this base type
func (m *SDTPaginationResponse) SetItems(val []SDT) {
	m.itemsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SDTPaginationResponse) UnmarshalJSON(raw []byte) error {
	var data struct {
		Items json.RawMessage `json:"items"`

		SearchID string `json:"searchId,omitempty"`

		Total int32 `json:"total,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propItems []SDT
	if string(data.Items) != "null" {
		items, err := UnmarshalSDTSlice(bytes.NewBuffer(data.Items), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propItems = items
	}

	var result SDTPaginationResponse

	// items
	result.itemsField = propItems

	// searchId
	result.SearchID = data.SearchID

	// total
	result.Total = data.Total

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SDTPaginationResponse) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		SearchID string `json:"searchId,omitempty"`

		Total int32 `json:"total,omitempty"`
	}{

		SearchID: m.SearchID,

		Total: m.Total,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Items []SDT `json:"items"`
	}{

		Items: m.itemsField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this SDT pagination response
func (m *SDTPaginationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDTPaginationResponse) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items()) { // not required
		return nil
	}

	for i := 0; i < len(m.Items()); i++ {

		if err := m.itemsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("items" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this SDT pagination response based on the context it is used
func (m *SDTPaginationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearchID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDTPaginationResponse) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items()); i++ {

		if err := m.itemsField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("items" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SDTPaginationResponse) contextValidateSearchID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "searchId", "body", string(m.SearchID)); err != nil {
		return err
	}

	return nil
}

func (m *SDTPaginationResponse) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total", "body", int32(m.Total)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SDTPaginationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDTPaginationResponse) UnmarshalBinary(b []byte) error {
	var res SDTPaginationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
