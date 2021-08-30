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

// NewGetWidgetDataByIDParams creates a new GetWidgetDataByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWidgetDataByIDParams() *GetWidgetDataByIDParams {
	return &GetWidgetDataByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWidgetDataByIDParamsWithTimeout creates a new GetWidgetDataByIDParams object
// with the ability to set a timeout on a request.
func NewGetWidgetDataByIDParamsWithTimeout(timeout time.Duration) *GetWidgetDataByIDParams {
	return &GetWidgetDataByIDParams{
		timeout: timeout,
	}
}

// NewGetWidgetDataByIDParamsWithContext creates a new GetWidgetDataByIDParams object
// with the ability to set a context for a request.
func NewGetWidgetDataByIDParamsWithContext(ctx context.Context) *GetWidgetDataByIDParams {
	return &GetWidgetDataByIDParams{
		Context: ctx,
	}
}

// NewGetWidgetDataByIDParamsWithHTTPClient creates a new GetWidgetDataByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWidgetDataByIDParamsWithHTTPClient(client *http.Client) *GetWidgetDataByIDParams {
	return &GetWidgetDataByIDParams{
		HTTPClient: client,
	}
}

/* GetWidgetDataByIDParams contains all the parameters to send to the API endpoint
   for the get widget data by Id operation.

   Typically these are written to a http.Request.
*/
type GetWidgetDataByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// End.
	//
	// Format: int64
	End *int64

	// Format.
	Format *string

	// ID.
	//
	// Format: int32
	ID int32

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get widget data by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWidgetDataByIDParams) WithDefaults() *GetWidgetDataByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get widget data by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWidgetDataByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := GetWidgetDataByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get widget data by Id params
func (o *GetWidgetDataByIDParams) WithTimeout(timeout time.Duration) *GetWidgetDataByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get widget data by Id params
func (o *GetWidgetDataByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get widget data by Id params
func (o *GetWidgetDataByIDParams) WithContext(ctx context.Context) *GetWidgetDataByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get widget data by Id params
func (o *GetWidgetDataByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get widget data by Id params
func (o *GetWidgetDataByIDParams) WithHTTPClient(client *http.Client) *GetWidgetDataByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get widget data by Id params
func (o *GetWidgetDataByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get widget data by Id params
func (o *GetWidgetDataByIDParams) WithUserAgent(userAgent *string) *GetWidgetDataByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get widget data by Id params
func (o *GetWidgetDataByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithEnd adds the end to the get widget data by Id params
func (o *GetWidgetDataByIDParams) WithEnd(end *int64) *GetWidgetDataByIDParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get widget data by Id params
func (o *GetWidgetDataByIDParams) SetEnd(end *int64) {
	o.End = end
}

// WithFormat adds the format to the get widget data by Id params
func (o *GetWidgetDataByIDParams) WithFormat(format *string) *GetWidgetDataByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get widget data by Id params
func (o *GetWidgetDataByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get widget data by Id params
func (o *GetWidgetDataByIDParams) WithID(id int32) *GetWidgetDataByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get widget data by Id params
func (o *GetWidgetDataByIDParams) SetID(id int32) {
	o.ID = id
}

// WithStart adds the start to the get widget data by Id params
func (o *GetWidgetDataByIDParams) WithStart(start *int64) *GetWidgetDataByIDParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get widget data by Id params
func (o *GetWidgetDataByIDParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetWidgetDataByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}