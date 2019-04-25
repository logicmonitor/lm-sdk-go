// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BigNumberInfo big number info
// swagger:model BigNumberInfo
type BigNumberInfo struct {

	// The datapoints and virtual datapoints whose values should be displayed in the big number widget
	// Required: true
	BigNumberItems []*BigNumberItem `json:"bigNumberItems"`

	// The counter is used for saving applyTo expression, it's mainly used for count device
	Counters []*Counter `json:"counters,omitempty"`

	// The datapoints included in the widget. Note that a datapoint must be referenced in the bigNumberItems object in order to be displayed
	// Required: true
	DataPoints []*BigNumberDataPoint `json:"dataPoints"`

	// The virtual datapoints included in the widget. Note that a virtual datapoint must be referenced in the bigNumberItems object in order to be displayed
	VirtualDataPoints []*VirtualDataPoint `json:"virtualDataPoints,omitempty"`
}

// Validate validates this big number info
func (m *BigNumberInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBigNumberItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCounters(formats); err != nil {
		res = append(res, err)
	}

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

func (m *BigNumberInfo) validateBigNumberItems(formats strfmt.Registry) error {

	if err := validate.Required("bigNumberItems", "body", m.BigNumberItems); err != nil {
		return err
	}

	for i := 0; i < len(m.BigNumberItems); i++ {
		if swag.IsZero(m.BigNumberItems[i]) { // not required
			continue
		}

		if m.BigNumberItems[i] != nil {
			if err := m.BigNumberItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bigNumberItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BigNumberInfo) validateCounters(formats strfmt.Registry) error {

	if swag.IsZero(m.Counters) { // not required
		return nil
	}

	for i := 0; i < len(m.Counters); i++ {
		if swag.IsZero(m.Counters[i]) { // not required
			continue
		}

		if m.Counters[i] != nil {
			if err := m.Counters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BigNumberInfo) validateDataPoints(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *BigNumberInfo) validateVirtualDataPoints(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BigNumberInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BigNumberInfo) UnmarshalBinary(b []byte) error {
	var res BigNumberInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
