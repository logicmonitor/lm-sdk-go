// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GraphLine graph line
//
// swagger:model GraphLine
type GraphLine struct {

	// color name
	ColorName string `json:"colorName,omitempty"`

	// data point Id
	DataPointID int32 `json:"dataPointId,omitempty"`

	// data point name
	DataPointName string `json:"dataPointName,omitempty"`

	// is virtual data point
	IsVirtualDataPoint bool `json:"isVirtualDataPoint,omitempty"`

	// legend
	Legend string `json:"legend,omitempty"`

	// type
	Type int32 `json:"type,omitempty"`
}

// Validate validates this graph line
func (m *GraphLine) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this graph line based on context it is used
func (m *GraphLine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GraphLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphLine) UnmarshalBinary(b []byte) error {
	var res GraphLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
