// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteEscalationChainByIDParams creates a new DeleteEscalationChainByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEscalationChainByIDParams() *DeleteEscalationChainByIDParams {
	return &DeleteEscalationChainByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEscalationChainByIDParamsWithTimeout creates a new DeleteEscalationChainByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteEscalationChainByIDParamsWithTimeout(timeout time.Duration) *DeleteEscalationChainByIDParams {
	return &DeleteEscalationChainByIDParams{
		timeout: timeout,
	}
}

// NewDeleteEscalationChainByIDParamsWithContext creates a new DeleteEscalationChainByIDParams object
// with the ability to set a context for a request.
func NewDeleteEscalationChainByIDParamsWithContext(ctx context.Context) *DeleteEscalationChainByIDParams {
	return &DeleteEscalationChainByIDParams{
		Context: ctx,
	}
}

// NewDeleteEscalationChainByIDParamsWithHTTPClient creates a new DeleteEscalationChainByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEscalationChainByIDParamsWithHTTPClient(client *http.Client) *DeleteEscalationChainByIDParams {
	return &DeleteEscalationChainByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteEscalationChainByIDParams contains all the parameters to send to the API endpoint

	for the delete escalation chain by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteEscalationChainByIDParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete escalation chain by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEscalationChainByIDParams) WithDefaults() *DeleteEscalationChainByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete escalation chain by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEscalationChainByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete escalation chain by Id params
func (o *DeleteEscalationChainByIDParams) WithTimeout(timeout time.Duration) *DeleteEscalationChainByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete escalation chain by Id params
func (o *DeleteEscalationChainByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete escalation chain by Id params
func (o *DeleteEscalationChainByIDParams) WithContext(ctx context.Context) *DeleteEscalationChainByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete escalation chain by Id params
func (o *DeleteEscalationChainByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete escalation chain by Id params
func (o *DeleteEscalationChainByIDParams) WithHTTPClient(client *http.Client) *DeleteEscalationChainByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete escalation chain by Id params
func (o *DeleteEscalationChainByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete escalation chain by Id params
func (o *DeleteEscalationChainByIDParams) WithID(id int32) *DeleteEscalationChainByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete escalation chain by Id params
func (o *DeleteEscalationChainByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEscalationChainByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
