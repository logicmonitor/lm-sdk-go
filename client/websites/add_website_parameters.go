// Code generated by go-swagger; DO NOT EDIT.

package websites

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

// NewAddWebsiteParams creates a new AddWebsiteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddWebsiteParams() *AddWebsiteParams {
	return &AddWebsiteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddWebsiteParamsWithTimeout creates a new AddWebsiteParams object
// with the ability to set a timeout on a request.
func NewAddWebsiteParamsWithTimeout(timeout time.Duration) *AddWebsiteParams {
	return &AddWebsiteParams{
		timeout: timeout,
	}
}

// NewAddWebsiteParamsWithContext creates a new AddWebsiteParams object
// with the ability to set a context for a request.
func NewAddWebsiteParamsWithContext(ctx context.Context) *AddWebsiteParams {
	return &AddWebsiteParams{
		Context: ctx,
	}
}

// NewAddWebsiteParamsWithHTTPClient creates a new AddWebsiteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddWebsiteParamsWithHTTPClient(client *http.Client) *AddWebsiteParams {
	return &AddWebsiteParams{
		HTTPClient: client,
	}
}

/*
AddWebsiteParams contains all the parameters to send to the API endpoint

	for the add website operation.

	Typically these are written to a http.Request.
*/
type AddWebsiteParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// Body.
	Body models.Website

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add website params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddWebsiteParams) WithDefaults() *AddWebsiteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add website params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddWebsiteParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")
	)

	val := AddWebsiteParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add website params
func (o *AddWebsiteParams) WithTimeout(timeout time.Duration) *AddWebsiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add website params
func (o *AddWebsiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add website params
func (o *AddWebsiteParams) WithContext(ctx context.Context) *AddWebsiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add website params
func (o *AddWebsiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add website params
func (o *AddWebsiteParams) WithHTTPClient(client *http.Client) *AddWebsiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add website params
func (o *AddWebsiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add website params
func (o *AddWebsiteParams) WithUserAgent(userAgent *string) *AddWebsiteParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add website params
func (o *AddWebsiteParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add website params
func (o *AddWebsiteParams) WithBody(body models.Website) *AddWebsiteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add website params
func (o *AddWebsiteParams) SetBody(body models.Website) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddWebsiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
