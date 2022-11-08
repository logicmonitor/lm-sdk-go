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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewAddAdminParams creates a new AddAdminParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAdminParams() *AddAdminParams {
	return &AddAdminParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAdminParamsWithTimeout creates a new AddAdminParams object
// with the ability to set a timeout on a request.
func NewAddAdminParamsWithTimeout(timeout time.Duration) *AddAdminParams {
	return &AddAdminParams{
		timeout: timeout,
	}
}

// NewAddAdminParamsWithContext creates a new AddAdminParams object
// with the ability to set a context for a request.
func NewAddAdminParamsWithContext(ctx context.Context) *AddAdminParams {
	return &AddAdminParams{
		Context: ctx,
	}
}

// NewAddAdminParamsWithHTTPClient creates a new AddAdminParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAdminParamsWithHTTPClient(client *http.Client) *AddAdminParams {
	return &AddAdminParams{
		HTTPClient: client,
	}
}

/*
AddAdminParams contains all the parameters to send to the API endpoint

	for the add admin operation.

	Typically these are written to a http.Request.
*/
type AddAdminParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// Body.
	Body *models.Admin

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAdminParams) WithDefaults() *AddAdminParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAdminParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")
	)

	val := AddAdminParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add admin params
func (o *AddAdminParams) WithTimeout(timeout time.Duration) *AddAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add admin params
func (o *AddAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add admin params
func (o *AddAdminParams) WithContext(ctx context.Context) *AddAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add admin params
func (o *AddAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add admin params
func (o *AddAdminParams) WithHTTPClient(client *http.Client) *AddAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add admin params
func (o *AddAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add admin params
func (o *AddAdminParams) WithUserAgent(userAgent *string) *AddAdminParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add admin params
func (o *AddAdminParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add admin params
func (o *AddAdminParams) WithBody(body *models.Admin) *AddAdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add admin params
func (o *AddAdminParams) SetBody(body *models.Admin) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
