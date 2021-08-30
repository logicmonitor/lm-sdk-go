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

// AutoDiscoveryConfiguration auto discovery configuration
//
// swagger:model AutoDiscoveryConfiguration
type AutoDiscoveryConfiguration struct {

	// data source name
	// Read Only: true
	DataSourceName string `json:"dataSourceName,omitempty"`

	// delete inactive instance
	DeleteInactiveInstance bool `json:"deleteInactiveInstance,omitempty"`

	// disable instance
	DisableInstance bool `json:"disableInstance,omitempty"`

	// filters
	Filters []*AutoDiscoveryFilter `json:"filters,omitempty"`

	// instance auto group method
	InstanceAutoGroupMethod string `json:"instanceAutoGroupMethod,omitempty"`

	// instance auto group method params
	InstanceAutoGroupMethodParams string `json:"instanceAutoGroupMethodParams,omitempty"`

	methodField AutoDiscoveryMethod

	// persistent instance
	PersistentInstance bool `json:"persistentInstance,omitempty"`

	// schedule interval
	ScheduleInterval int32 `json:"scheduleInterval,omitempty"`
}

// Method gets the method of this base type
func (m *AutoDiscoveryConfiguration) Method() AutoDiscoveryMethod {
	return m.methodField
}

// SetMethod sets the method of this base type
func (m *AutoDiscoveryConfiguration) SetMethod(val AutoDiscoveryMethod) {
	m.methodField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AutoDiscoveryConfiguration) UnmarshalJSON(raw []byte) error {
	var data struct {
		DataSourceName string `json:"dataSourceName,omitempty"`

		DeleteInactiveInstance bool `json:"deleteInactiveInstance,omitempty"`

		DisableInstance bool `json:"disableInstance,omitempty"`

		Filters []*AutoDiscoveryFilter `json:"filters,omitempty"`

		InstanceAutoGroupMethod string `json:"instanceAutoGroupMethod,omitempty"`

		InstanceAutoGroupMethodParams string `json:"instanceAutoGroupMethodParams,omitempty"`

		Method json.RawMessage `json:"method"`

		PersistentInstance bool `json:"persistentInstance,omitempty"`

		ScheduleInterval int32 `json:"scheduleInterval,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	propMethod, err := UnmarshalAutoDiscoveryMethod(bytes.NewBuffer(data.Method), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result AutoDiscoveryConfiguration

	// dataSourceName
	result.DataSourceName = data.DataSourceName

	// deleteInactiveInstance
	result.DeleteInactiveInstance = data.DeleteInactiveInstance

	// disableInstance
	result.DisableInstance = data.DisableInstance

	// filters
	result.Filters = data.Filters

	// instanceAutoGroupMethod
	result.InstanceAutoGroupMethod = data.InstanceAutoGroupMethod

	// instanceAutoGroupMethodParams
	result.InstanceAutoGroupMethodParams = data.InstanceAutoGroupMethodParams

	// method
	result.methodField = propMethod

	// persistentInstance
	result.PersistentInstance = data.PersistentInstance

	// scheduleInterval
	result.ScheduleInterval = data.ScheduleInterval

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AutoDiscoveryConfiguration) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		DataSourceName string `json:"dataSourceName,omitempty"`

		DeleteInactiveInstance bool `json:"deleteInactiveInstance,omitempty"`

		DisableInstance bool `json:"disableInstance,omitempty"`

		Filters []*AutoDiscoveryFilter `json:"filters,omitempty"`

		InstanceAutoGroupMethod string `json:"instanceAutoGroupMethod,omitempty"`

		InstanceAutoGroupMethodParams string `json:"instanceAutoGroupMethodParams,omitempty"`

		PersistentInstance bool `json:"persistentInstance,omitempty"`

		ScheduleInterval int32 `json:"scheduleInterval,omitempty"`
	}{

		DataSourceName: m.DataSourceName,

		DeleteInactiveInstance: m.DeleteInactiveInstance,

		DisableInstance: m.DisableInstance,

		Filters: m.Filters,

		InstanceAutoGroupMethod: m.InstanceAutoGroupMethod,

		InstanceAutoGroupMethodParams: m.InstanceAutoGroupMethodParams,

		PersistentInstance: m.PersistentInstance,

		ScheduleInterval: m.ScheduleInterval,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Method AutoDiscoveryMethod `json:"method"`
	}{

		Method: m.methodField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this auto discovery configuration
func (m *AutoDiscoveryConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoDiscoveryConfiguration) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AutoDiscoveryConfiguration) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method()); err != nil {
		return err
	}

	if err := m.Method().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("method")
		}
		return err
	}

	return nil
}

// ContextValidate validate this auto discovery configuration based on the context it is used
func (m *AutoDiscoveryConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataSourceName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoDiscoveryConfiguration) contextValidateDataSourceName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceName", "body", string(m.DataSourceName)); err != nil {
		return err
	}

	return nil
}

func (m *AutoDiscoveryConfiguration) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AutoDiscoveryConfiguration) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Method().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("method")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoDiscoveryConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoDiscoveryConfiguration) UnmarshalBinary(b []byte) error {
	var res AutoDiscoveryConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}