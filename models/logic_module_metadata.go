// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LogicModuleMetadata logic module metadata
// swagger:model LogicModuleMetadata
type LogicModuleMetadata struct {

	// The publication status of this module, along with its "locator" token
	// Read Only: true
	LMLocator string `json:"lmLocator,omitempty"`

	// The author of this particular module, which is based on the account name of the form [accountname].logicmonitor.com. As such, the modules are exclusively linked to the individual publisher, not their company's account.
	// Read Only: true
	Namespace string `json:"namespace,omitempty"`

	// The quality specification of this module, as determined by our Monitoring Engineers
	// Read Only: true
	Quality string `json:"quality,omitempty"`

	// Indicates the specific version of this module
	// Read Only: true
	RegistryVersion string `json:"registryVersion,omitempty"`
}

// Validate validates this logic module metadata
func (m *LogicModuleMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogicModuleMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogicModuleMetadata) UnmarshalBinary(b []byte) error {
	var res LogicModuleMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}