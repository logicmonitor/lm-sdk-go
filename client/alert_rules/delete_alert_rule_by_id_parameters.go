// Code generated by go-swagger; DO NOT EDIT.

package alert_rules

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

// NewDeleteAlertRuleByIDParams creates a new DeleteAlertRuleByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAlertRuleByIDParams() *DeleteAlertRuleByIDParams {
	return &DeleteAlertRuleByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAlertRuleByIDParamsWithTimeout creates a new DeleteAlertRuleByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAlertRuleByIDParamsWithTimeout(timeout time.Duration) *DeleteAlertRuleByIDParams {
	return &DeleteAlertRuleByIDParams{
		timeout: timeout,
	}
}

// NewDeleteAlertRuleByIDParamsWithContext creates a new DeleteAlertRuleByIDParams object
// with the ability to set a context for a request.
func NewDeleteAlertRuleByIDParamsWithContext(ctx context.Context) *DeleteAlertRuleByIDParams {
	return &DeleteAlertRuleByIDParams{
		Context: ctx,
	}
}

// NewDeleteAlertRuleByIDParamsWithHTTPClient creates a new DeleteAlertRuleByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAlertRuleByIDParamsWithHTTPClient(client *http.Client) *DeleteAlertRuleByIDParams {
	return &DeleteAlertRuleByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAlertRuleByIDParams contains all the parameters to send to the API endpoint

	for the delete alert rule by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteAlertRuleByIDParams struct {

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

// WithDefaults hydrates default values in the delete alert rule by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAlertRuleByIDParams) WithDefaults() *DeleteAlertRuleByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete alert rule by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAlertRuleByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")
	)

	val := DeleteAlertRuleByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithTimeout(timeout time.Duration) *DeleteAlertRuleByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithContext(ctx context.Context) *DeleteAlertRuleByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithHTTPClient(client *http.Client) *DeleteAlertRuleByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithUserAgent(userAgent *string) *DeleteAlertRuleByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithID(id int32) *DeleteAlertRuleByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAlertRuleByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
