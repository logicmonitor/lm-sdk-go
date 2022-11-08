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

// TableWidgetColumn table widget column
//
// swagger:model TableWidgetColumn
type TableWidgetColumn struct {

	// alternate data points
	AlternateDataPoints []*TableWidgetDataPoint `json:"alternateDataPoints"`

	// column name
	// Required: true
	ColumnName *string `json:"columnName"`

	// data point
	// Required: true
	DataPoint *TableWidgetDataPoint `json:"dataPoint"`

	// enable forecast
	EnableForecast bool `json:"enableForecast,omitempty"`

	// rounding decimal
	RoundingDecimal int32 `json:"roundingDecimal,omitempty"`

	// rpn
	Rpn string `json:"rpn,omitempty"`

	// The unit label
	UnitLabel string `json:"unitLabel,omitempty"`
}

// Validate validates this table widget column
func (m *TableWidgetColumn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternateDataPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColumnName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableWidgetColumn) validateAlternateDataPoints(formats strfmt.Registry) error {
	if swag.IsZero(m.AlternateDataPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.AlternateDataPoints); i++ {
		if swag.IsZero(m.AlternateDataPoints[i]) { // not required
			continue
		}

		if m.AlternateDataPoints[i] != nil {
			if err := m.AlternateDataPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alternateDataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alternateDataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TableWidgetColumn) validateColumnName(formats strfmt.Registry) error {

	if err := validate.Required("columnName", "body", m.ColumnName); err != nil {
		return err
	}

	return nil
}

func (m *TableWidgetColumn) validateDataPoint(formats strfmt.Registry) error {

	if err := validate.Required("dataPoint", "body", m.DataPoint); err != nil {
		return err
	}

	if m.DataPoint != nil {
		if err := m.DataPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataPoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataPoint")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this table widget column based on the context it is used
func (m *TableWidgetColumn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlternateDataPoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataPoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableWidgetColumn) contextValidateAlternateDataPoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AlternateDataPoints); i++ {

		if m.AlternateDataPoints[i] != nil {
			if err := m.AlternateDataPoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alternateDataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alternateDataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TableWidgetColumn) contextValidateDataPoint(ctx context.Context, formats strfmt.Registry) error {

	if m.DataPoint != nil {
		if err := m.DataPoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataPoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataPoint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TableWidgetColumn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableWidgetColumn) UnmarshalBinary(b []byte) error {
	var res TableWidgetColumn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
