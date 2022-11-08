// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomGraph custom graph
//
// swagger:model CustomGraph
type CustomGraph struct {

	// true: You can set this field to true to aggregate results into one line.
	// false: Results will not be aggregated
	// the default value is true
	Aggregate bool `json:"aggregate,omitempty"`

	// The datapoints added to the widget (note that a datapoint must be referenced in a graph line to be displayed)
	// Required: true
	DataPoints []*CustomFlexibleVirtualDataSourceEx `json:"dataPoints"`

	// Whether the top X are displayed (false) or the bottom X are displayed (true), the default value is true
	Desc bool `json:"desc,omitempty"`

	// The function for global consolidate
	GlobalConsolidateFunction string `json:"globalConsolidateFunction,omitempty"`

	// The unique id of the custom graph displayed by this widget (not to be confused with the widget id)
	ID int32 `json:"id,omitempty"`

	// The maximum value that should be displayed on the y-axis
	MaxValue float64 `json:"maxValue,omitempty"`

	// The minimum value that should be displayed on the y-axis
	MinValue float64 `json:"minValue,omitempty"`

	// The base scale unit (1000 or 1024)
	ScaleUnit int32 `json:"scaleUnit,omitempty"`

	// The number of lines to display for each configured datapoint
	TopX int32 `json:"topX,omitempty"`

	// The label that will be display along the y axis
	VerticalLabel string `json:"verticalLabel,omitempty"`

	// The virtual datapoints added to the widget (note that a virtual datapoint must be referenced in a graph line to be displayed)
	VirtualDataPoints []*CustomVirtualDataPoint `json:"virtualDataPoints"`
}

// Validate validates this custom graph
func (m *CustomGraph) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualDataPoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomGraph) validateDataPoints(formats strfmt.Registry) error {

	if err := validate.Required("dataPoints", "body", m.DataPoints); err != nil {
		return err
	}

	for i := 0; i < len(m.DataPoints); i++ {
		if swag.IsZero(m.DataPoints[i]) { // not required
			continue
		}

		if m.DataPoints[i] != nil {
			if err := m.DataPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomGraph) validateVirtualDataPoints(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualDataPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualDataPoints); i++ {
		if swag.IsZero(m.VirtualDataPoints[i]) { // not required
			continue
		}

		if m.VirtualDataPoints[i] != nil {
			if err := m.VirtualDataPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualDataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualDataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this custom graph based on the context it is used
func (m *CustomGraph) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataPoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualDataPoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomGraph) contextValidateDataPoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataPoints); i++ {

		if m.DataPoints[i] != nil {
			if err := m.DataPoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomGraph) contextValidateVirtualDataPoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VirtualDataPoints); i++ {

		if m.VirtualDataPoints[i] != nil {
			if err := m.VirtualDataPoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualDataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualDataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomGraph) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomGraph) UnmarshalBinary(b []byte) error {
	var res CustomGraph
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
