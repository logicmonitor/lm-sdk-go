// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestDashboardGroupAsyncCloneResponse rest dashboard group async clone response
//
// swagger:model RestDashboardGroupAsyncCloneResponse
type RestDashboardGroupAsyncCloneResponse struct {

	// The jobId of the Clone Dashboard Group
	// Read Only: true
	JobID int32 `json:"jobId,omitempty"`
}

// Validate validates this rest dashboard group async clone response
func (m *RestDashboardGroupAsyncCloneResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this rest dashboard group async clone response based on the context it is used
func (m *RestDashboardGroupAsyncCloneResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestDashboardGroupAsyncCloneResponse) contextValidateJobID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "jobId", "body", int32(m.JobID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestDashboardGroupAsyncCloneResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestDashboardGroupAsyncCloneResponse) UnmarshalBinary(b []byte) error {
	var res RestDashboardGroupAsyncCloneResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
