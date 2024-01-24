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

// NewAddAlertRuleParams creates a new AddAlertRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAlertRuleParams() *AddAlertRuleParams {
	return &AddAlertRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAlertRuleParamsWithTimeout creates a new AddAlertRuleParams object
// with the ability to set a timeout on a request.
func NewAddAlertRuleParamsWithTimeout(timeout time.Duration) *AddAlertRuleParams {
	return &AddAlertRuleParams{
		timeout: timeout,
	}
}

// NewAddAlertRuleParamsWithContext creates a new AddAlertRuleParams object
// with the ability to set a context for a request.
func NewAddAlertRuleParamsWithContext(ctx context.Context) *AddAlertRuleParams {
	return &AddAlertRuleParams{
		Context: ctx,
	}
}

// NewAddAlertRuleParamsWithHTTPClient creates a new AddAlertRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAlertRuleParamsWithHTTPClient(client *http.Client) *AddAlertRuleParams {
	return &AddAlertRuleParams{
		HTTPClient: client,
	}
}

/* AddAlertRuleParams contains all the parameters to send to the API endpoint
   for the add alert rule operation.

   Typically these are written to a http.Request.
*/
type AddAlertRuleParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.AlertRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAlertRuleParams) WithDefaults() *AddAlertRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAlertRuleParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := AddAlertRuleParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add alert rule params
func (o *AddAlertRuleParams) WithTimeout(timeout time.Duration) *AddAlertRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add alert rule params
func (o *AddAlertRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add alert rule params
func (o *AddAlertRuleParams) WithContext(ctx context.Context) *AddAlertRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add alert rule params
func (o *AddAlertRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add alert rule params
func (o *AddAlertRuleParams) WithHTTPClient(client *http.Client) *AddAlertRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add alert rule params
func (o *AddAlertRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add alert rule params
func (o *AddAlertRuleParams) WithUserAgent(userAgent *string) *AddAlertRuleParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add alert rule params
func (o *AddAlertRuleParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add alert rule params
func (o *AddAlertRuleParams) WithBody(body *models.AlertRule) *AddAlertRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add alert rule params
func (o *AddAlertRuleParams) SetBody(body *models.AlertRule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAlertRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
