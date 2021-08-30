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

// NewDeleteReportByIDParams creates a new DeleteReportByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteReportByIDParams() *DeleteReportByIDParams {
	return &DeleteReportByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReportByIDParamsWithTimeout creates a new DeleteReportByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteReportByIDParamsWithTimeout(timeout time.Duration) *DeleteReportByIDParams {
	return &DeleteReportByIDParams{
		timeout: timeout,
	}
}

// NewDeleteReportByIDParamsWithContext creates a new DeleteReportByIDParams object
// with the ability to set a context for a request.
func NewDeleteReportByIDParamsWithContext(ctx context.Context) *DeleteReportByIDParams {
	return &DeleteReportByIDParams{
		Context: ctx,
	}
}

// NewDeleteReportByIDParamsWithHTTPClient creates a new DeleteReportByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteReportByIDParamsWithHTTPClient(client *http.Client) *DeleteReportByIDParams {
	return &DeleteReportByIDParams{
		HTTPClient: client,
	}
}

/* DeleteReportByIDParams contains all the parameters to send to the API endpoint
   for the delete report by Id operation.

   Typically these are written to a http.Request.
*/
type DeleteReportByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete report by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteReportByIDParams) WithDefaults() *DeleteReportByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete report by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteReportByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := DeleteReportByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete report by Id params
func (o *DeleteReportByIDParams) WithTimeout(timeout time.Duration) *DeleteReportByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete report by Id params
func (o *DeleteReportByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete report by Id params
func (o *DeleteReportByIDParams) WithContext(ctx context.Context) *DeleteReportByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete report by Id params
func (o *DeleteReportByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete report by Id params
func (o *DeleteReportByIDParams) WithHTTPClient(client *http.Client) *DeleteReportByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete report by Id params
func (o *DeleteReportByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete report by Id params
func (o *DeleteReportByIDParams) WithUserAgent(userAgent *string) *DeleteReportByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete report by Id params
func (o *DeleteReportByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the delete report by Id params
func (o *DeleteReportByIDParams) WithID(id int32) *DeleteReportByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete report by Id params
func (o *DeleteReportByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReportByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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