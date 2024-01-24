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

// NewTestSaaSAccountParams creates a new TestSaaSAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestSaaSAccountParams() *TestSaaSAccountParams {
	return &TestSaaSAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestSaaSAccountParamsWithTimeout creates a new TestSaaSAccountParams object
// with the ability to set a timeout on a request.
func NewTestSaaSAccountParamsWithTimeout(timeout time.Duration) *TestSaaSAccountParams {
	return &TestSaaSAccountParams{
		timeout: timeout,
	}
}

// NewTestSaaSAccountParamsWithContext creates a new TestSaaSAccountParams object
// with the ability to set a context for a request.
func NewTestSaaSAccountParamsWithContext(ctx context.Context) *TestSaaSAccountParams {
	return &TestSaaSAccountParams{
		Context: ctx,
	}
}

// NewTestSaaSAccountParamsWithHTTPClient creates a new TestSaaSAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestSaaSAccountParamsWithHTTPClient(client *http.Client) *TestSaaSAccountParams {
	return &TestSaaSAccountParams{
		HTTPClient: client,
	}
}

/* TestSaaSAccountParams contains all the parameters to send to the API endpoint
   for the test saa s account operation.

   Typically these are written to a http.Request.
*/
type TestSaaSAccountParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.RestSaaSAccountTestV3

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test saa s account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestSaaSAccountParams) WithDefaults() *TestSaaSAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test saa s account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestSaaSAccountParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := TestSaaSAccountParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the test saa s account params
func (o *TestSaaSAccountParams) WithTimeout(timeout time.Duration) *TestSaaSAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test saa s account params
func (o *TestSaaSAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test saa s account params
func (o *TestSaaSAccountParams) WithContext(ctx context.Context) *TestSaaSAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test saa s account params
func (o *TestSaaSAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test saa s account params
func (o *TestSaaSAccountParams) WithHTTPClient(client *http.Client) *TestSaaSAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test saa s account params
func (o *TestSaaSAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the test saa s account params
func (o *TestSaaSAccountParams) WithUserAgent(userAgent *string) *TestSaaSAccountParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the test saa s account params
func (o *TestSaaSAccountParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the test saa s account params
func (o *TestSaaSAccountParams) WithBody(body *models.RestSaaSAccountTestV3) *TestSaaSAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test saa s account params
func (o *TestSaaSAccountParams) SetBody(body *models.RestSaaSAccountTestV3) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestSaaSAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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