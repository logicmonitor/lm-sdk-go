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

// NewGetReportGroupByIDParams creates a new GetReportGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReportGroupByIDParams() *GetReportGroupByIDParams {
	return &GetReportGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportGroupByIDParamsWithTimeout creates a new GetReportGroupByIDParams object
// with the ability to set a timeout on a request.
func NewGetReportGroupByIDParamsWithTimeout(timeout time.Duration) *GetReportGroupByIDParams {
	return &GetReportGroupByIDParams{
		timeout: timeout,
	}
}

// NewGetReportGroupByIDParamsWithContext creates a new GetReportGroupByIDParams object
// with the ability to set a context for a request.
func NewGetReportGroupByIDParamsWithContext(ctx context.Context) *GetReportGroupByIDParams {
	return &GetReportGroupByIDParams{
		Context: ctx,
	}
}

// NewGetReportGroupByIDParamsWithHTTPClient creates a new GetReportGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReportGroupByIDParamsWithHTTPClient(client *http.Client) *GetReportGroupByIDParams {
	return &GetReportGroupByIDParams{
		HTTPClient: client,
	}
}

/* GetReportGroupByIDParams contains all the parameters to send to the API endpoint
   for the get report group by Id operation.

   Typically these are written to a http.Request.
*/
type GetReportGroupByIDParams struct {

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

// WithDefaults hydrates default values in the get report group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportGroupByIDParams) WithDefaults() *GetReportGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get report group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := GetReportGroupByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get report group by Id params
func (o *GetReportGroupByIDParams) WithTimeout(timeout time.Duration) *GetReportGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get report group by Id params
func (o *GetReportGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get report group by Id params
func (o *GetReportGroupByIDParams) WithContext(ctx context.Context) *GetReportGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get report group by Id params
func (o *GetReportGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get report group by Id params
func (o *GetReportGroupByIDParams) WithHTTPClient(client *http.Client) *GetReportGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get report group by Id params
func (o *GetReportGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get report group by Id params
func (o *GetReportGroupByIDParams) WithUserAgent(userAgent *string) *GetReportGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get report group by Id params
func (o *GetReportGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the get report group by Id params
func (o *GetReportGroupByIDParams) WithID(id int32) *GetReportGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get report group by Id params
func (o *GetReportGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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