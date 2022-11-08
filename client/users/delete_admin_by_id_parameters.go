// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewDeleteAdminByIDParams creates a new DeleteAdminByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAdminByIDParams() *DeleteAdminByIDParams {
	return &DeleteAdminByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAdminByIDParamsWithTimeout creates a new DeleteAdminByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAdminByIDParamsWithTimeout(timeout time.Duration) *DeleteAdminByIDParams {
	return &DeleteAdminByIDParams{
		timeout: timeout,
	}
}

// NewDeleteAdminByIDParamsWithContext creates a new DeleteAdminByIDParams object
// with the ability to set a context for a request.
func NewDeleteAdminByIDParamsWithContext(ctx context.Context) *DeleteAdminByIDParams {
	return &DeleteAdminByIDParams{
		Context: ctx,
	}
}

// NewDeleteAdminByIDParamsWithHTTPClient creates a new DeleteAdminByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAdminByIDParamsWithHTTPClient(client *http.Client) *DeleteAdminByIDParams {
	return &DeleteAdminByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAdminByIDParams contains all the parameters to send to the API endpoint

	for the delete admin by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteAdminByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete admin by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdminByIDParams) WithDefaults() *DeleteAdminByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete admin by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdminByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")
	)

	val := DeleteAdminByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete admin by Id params
func (o *DeleteAdminByIDParams) WithTimeout(timeout time.Duration) *DeleteAdminByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete admin by Id params
func (o *DeleteAdminByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete admin by Id params
func (o *DeleteAdminByIDParams) WithContext(ctx context.Context) *DeleteAdminByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete admin by Id params
func (o *DeleteAdminByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete admin by Id params
func (o *DeleteAdminByIDParams) WithHTTPClient(client *http.Client) *DeleteAdminByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete admin by Id params
func (o *DeleteAdminByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete admin by Id params
func (o *DeleteAdminByIDParams) WithUserAgent(userAgent *string) *DeleteAdminByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete admin by Id params
func (o *DeleteAdminByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the delete admin by Id params
func (o *DeleteAdminByIDParams) WithID(id int32) *DeleteAdminByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete admin by Id params
func (o *DeleteAdminByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAdminByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserAgent != nil {

		// header param User-Agent
		if err := r.SetHeaderParam("User-Agent", *o.UserAgent); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
