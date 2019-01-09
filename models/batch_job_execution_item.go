// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BatchJobExecutionItem batch job execution item
// swagger:model BatchJobExecutionItem
type BatchJobExecutionItem struct {

	// alert level
	// Read Only: true
	AlertLevel int32 `json:"alertLevel,omitempty"`

	// bj name
	// Read Only: true
	BjName string `json:"bjName,omitempty"`

	// cmdline
	// Read Only: true
	Cmdline string `json:"cmdline,omitempty"`

	// device batch job Id
	// Read Only: true
	DeviceBatchJobID int32 `json:"deviceBatchJobId,omitempty"`

	// device display name
	// Read Only: true
	DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

	// execution no
	// Read Only: true
	ExecutionNo int64 `json:"executionNo,omitempty"`

	// exit code
	// Read Only: true
	ExitCode int32 `json:"exitCode,omitempty"`

	// finished on
	// Read Only: true
	FinishedOn int64 `json:"finishedOn,omitempty"`

	// finished on local
	// Read Only: true
	FinishedOnLocal string `json:"finishedOnLocal,omitempty"`

	// id
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// started on
	// Read Only: true
	StartedOn int64 `json:"startedOn,omitempty"`

	// started on local
	// Read Only: true
	StartedOnLocal string `json:"startedOnLocal,omitempty"`

	// stderr
	// Read Only: true
	Stderr string `json:"stderr,omitempty"`

	// stdout
	// Read Only: true
	Stdout string `json:"stdout,omitempty"`

	// user data
	// Read Only: true
	UserData string `json:"userData,omitempty"`
}

// Validate validates this batch job execution item
func (m *BatchJobExecutionItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BatchJobExecutionItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchJobExecutionItem) UnmarshalBinary(b []byte) error {
	var res BatchJobExecutionItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
