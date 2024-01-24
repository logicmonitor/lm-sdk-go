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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewTestGCPAccountParams creates a new TestGCPAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestGCPAccountParams() *TestGCPAccountParams {
	return &TestGCPAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestGCPAccountParamsWithTimeout creates a new TestGCPAccountParams object
// with the ability to set a timeout on a request.
func NewTestGCPAccountParamsWithTimeout(timeout time.Duration) *TestGCPAccountParams {
	return &TestGCPAccountParams{
		timeout: timeout,
	}
}

// NewTestGCPAccountParamsWithContext creates a new TestGCPAccountParams object
// with the ability to set a context for a request.
func NewTestGCPAccountParamsWithContext(ctx context.Context) *TestGCPAccountParams {
	return &TestGCPAccountParams{
		Context: ctx,
	}
}

// NewTestGCPAccountParamsWithHTTPClient creates a new TestGCPAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestGCPAccountParamsWithHTTPClient(client *http.Client) *TestGCPAccountParams {
	return &TestGCPAccountParams{
		HTTPClient: client,
	}
}

/* TestGCPAccountParams contains all the parameters to send to the API endpoint
   for the test g c p account operation.

   Typically these are written to a http.Request.
*/
type TestGCPAccountParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.RestGcpAccountTestV3

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test g c p account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestGCPAccountParams) WithDefaults() *TestGCPAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test g c p account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestGCPAccountParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := TestGCPAccountParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the test g c p account params
func (o *TestGCPAccountParams) WithTimeout(timeout time.Duration) *TestGCPAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test g c p account params
func (o *TestGCPAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test g c p account params
func (o *TestGCPAccountParams) WithContext(ctx context.Context) *TestGCPAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test g c p account params
func (o *TestGCPAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test g c p account params
func (o *TestGCPAccountParams) WithHTTPClient(client *http.Client) *TestGCPAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test g c p account params
func (o *TestGCPAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the test g c p account params
func (o *TestGCPAccountParams) WithUserAgent(userAgent *string) *TestGCPAccountParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the test g c p account params
func (o *TestGCPAccountParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the test g c p account params
func (o *TestGCPAccountParams) WithBody(body *models.RestGcpAccountTestV3) *TestGCPAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test g c p account params
func (o *TestGCPAccountParams) SetBody(body *models.RestGcpAccountTestV3) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestGCPAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}