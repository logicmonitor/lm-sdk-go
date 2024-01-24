// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestGcpAccountTestV3 rest gcp account test v3
//
// swagger:model RestGcpAccountTestV3
type RestGcpAccountTestV3 struct {

	// checked services
	CheckedServices string `json:"checkedServices,omitempty"`

	// group Id
	GroupID int32 `json:"groupId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// service account key
	ServiceAccountKey string `json:"serviceAccountKey,omitempty"`
}

// Validate validates this rest gcp account test v3
func (m *RestGcpAccountTestV3) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest gcp account test v3 based on context it is used
func (m *RestGcpAccountTestV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestGcpAccountTestV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestGcpAccountTestV3) UnmarshalBinary(b []byte) error {
	var res RestGcpAccountTestV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}