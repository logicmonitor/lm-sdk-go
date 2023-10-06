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

// NewAddDashboardGroupParams creates a new AddDashboardGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDashboardGroupParams() *AddDashboardGroupParams {
	return &AddDashboardGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDashboardGroupParamsWithTimeout creates a new AddDashboardGroupParams object
// with the ability to set a timeout on a request.
func NewAddDashboardGroupParamsWithTimeout(timeout time.Duration) *AddDashboardGroupParams {
	return &AddDashboardGroupParams{
		timeout: timeout,
	}
}

// NewAddDashboardGroupParamsWithContext creates a new AddDashboardGroupParams object
// with the ability to set a context for a request.
func NewAddDashboardGroupParamsWithContext(ctx context.Context) *AddDashboardGroupParams {
	return &AddDashboardGroupParams{
		Context: ctx,
	}
}

// NewAddDashboardGroupParamsWithHTTPClient creates a new AddDashboardGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDashboardGroupParamsWithHTTPClient(client *http.Client) *AddDashboardGroupParams {
	return &AddDashboardGroupParams{
		HTTPClient: client,
	}
}

/*
AddDashboardGroupParams contains all the parameters to send to the API endpoint

	for the add dashboard group operation.

	Typically these are written to a http.Request.
*/
type AddDashboardGroupParams struct {

	// Body.
	Body *models.DashboardGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add dashboard group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDashboardGroupParams) WithDefaults() *AddDashboardGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add dashboard group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDashboardGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add dashboard group params
func (o *AddDashboardGroupParams) WithTimeout(timeout time.Duration) *AddDashboardGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add dashboard group params
func (o *AddDashboardGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add dashboard group params
func (o *AddDashboardGroupParams) WithContext(ctx context.Context) *AddDashboardGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add dashboard group params
func (o *AddDashboardGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add dashboard group params
func (o *AddDashboardGroupParams) WithHTTPClient(client *http.Client) *AddDashboardGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add dashboard group params
func (o *AddDashboardGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add dashboard group params
func (o *AddDashboardGroupParams) WithBody(body *models.DashboardGroup) *AddDashboardGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add dashboard group params
func (o *AddDashboardGroupParams) SetBody(body *models.DashboardGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDashboardGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
