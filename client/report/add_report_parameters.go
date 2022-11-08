// Code generated by go-swagger; DO NOT EDIT.

package report

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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewAddReportParams creates a new AddReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddReportParams() *AddReportParams {
	return &AddReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddReportParamsWithTimeout creates a new AddReportParams object
// with the ability to set a timeout on a request.
func NewAddReportParamsWithTimeout(timeout time.Duration) *AddReportParams {
	return &AddReportParams{
		timeout: timeout,
	}
}

// NewAddReportParamsWithContext creates a new AddReportParams object
// with the ability to set a context for a request.
func NewAddReportParamsWithContext(ctx context.Context) *AddReportParams {
	return &AddReportParams{
		Context: ctx,
	}
}

// NewAddReportParamsWithHTTPClient creates a new AddReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddReportParamsWithHTTPClient(client *http.Client) *AddReportParams {
	return &AddReportParams{
		HTTPClient: client,
	}
}

/*
AddReportParams contains all the parameters to send to the API endpoint

	for the add report operation.

	Typically these are written to a http.Request.
*/
type AddReportParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// Body.
	Body models.ReportBase

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddReportParams) WithDefaults() *AddReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddReportParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")
	)

	val := AddReportParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add report params
func (o *AddReportParams) WithTimeout(timeout time.Duration) *AddReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add report params
func (o *AddReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add report params
func (o *AddReportParams) WithContext(ctx context.Context) *AddReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add report params
func (o *AddReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add report params
func (o *AddReportParams) WithHTTPClient(client *http.Client) *AddReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add report params
func (o *AddReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add report params
func (o *AddReportParams) WithUserAgent(userAgent *string) *AddReportParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add report params
func (o *AddReportParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add report params
func (o *AddReportParams) WithBody(body models.ReportBase) *AddReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add report params
func (o *AddReportParams) SetBody(body models.ReportBase) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
