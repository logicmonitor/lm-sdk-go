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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewUpdateAlertRuleByIDParams creates a new UpdateAlertRuleByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAlertRuleByIDParams() *UpdateAlertRuleByIDParams {
	return &UpdateAlertRuleByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAlertRuleByIDParamsWithTimeout creates a new UpdateAlertRuleByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateAlertRuleByIDParamsWithTimeout(timeout time.Duration) *UpdateAlertRuleByIDParams {
	return &UpdateAlertRuleByIDParams{
		timeout: timeout,
	}
}

// NewUpdateAlertRuleByIDParamsWithContext creates a new UpdateAlertRuleByIDParams object
// with the ability to set a context for a request.
func NewUpdateAlertRuleByIDParamsWithContext(ctx context.Context) *UpdateAlertRuleByIDParams {
	return &UpdateAlertRuleByIDParams{
		Context: ctx,
	}
}

// NewUpdateAlertRuleByIDParamsWithHTTPClient creates a new UpdateAlertRuleByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAlertRuleByIDParamsWithHTTPClient(client *http.Client) *UpdateAlertRuleByIDParams {
	return &UpdateAlertRuleByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateAlertRuleByIDParams contains all the parameters to send to the API endpoint

	for the update alert rule by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateAlertRuleByIDParams struct {

	// Body.
	Body *models.AlertRule

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update alert rule by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAlertRuleByIDParams) WithDefaults() *UpdateAlertRuleByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update alert rule by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAlertRuleByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) WithTimeout(timeout time.Duration) *UpdateAlertRuleByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) WithContext(ctx context.Context) *UpdateAlertRuleByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) WithHTTPClient(client *http.Client) *UpdateAlertRuleByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) WithBody(body *models.AlertRule) *UpdateAlertRuleByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) SetBody(body *models.AlertRule) {
	o.Body = body
}

// WithID adds the id to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) WithID(id int32) *UpdateAlertRuleByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update alert rule by Id params
func (o *UpdateAlertRuleByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAlertRuleByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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
