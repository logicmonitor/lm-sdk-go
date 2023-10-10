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

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// NewAddRoleParams creates a new AddRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRoleParams() *AddRoleParams {
	return &AddRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRoleParamsWithTimeout creates a new AddRoleParams object
// with the ability to set a timeout on a request.
func NewAddRoleParamsWithTimeout(timeout time.Duration) *AddRoleParams {
	return &AddRoleParams{
		timeout: timeout,
	}
}

// NewAddRoleParamsWithContext creates a new AddRoleParams object
// with the ability to set a context for a request.
func NewAddRoleParamsWithContext(ctx context.Context) *AddRoleParams {
	return &AddRoleParams{
		Context: ctx,
	}
}

// NewAddRoleParamsWithHTTPClient creates a new AddRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRoleParamsWithHTTPClient(client *http.Client) *AddRoleParams {
	return &AddRoleParams{
		HTTPClient: client,
	}
}

/* AddRoleParams contains all the parameters to send to the API endpoint
   for the add role operation.

   Typically these are written to a http.Request.
*/
type AddRoleParams struct {

	// Body.
	Body *models.Role

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRoleParams) WithDefaults() *AddRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add role params
func (o *AddRoleParams) WithTimeout(timeout time.Duration) *AddRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add role params
func (o *AddRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add role params
func (o *AddRoleParams) WithContext(ctx context.Context) *AddRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add role params
func (o *AddRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add role params
func (o *AddRoleParams) WithHTTPClient(client *http.Client) *AddRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add role params
func (o *AddRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add role params
func (o *AddRoleParams) WithBody(body *models.Role) *AddRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add role params
func (o *AddRoleParams) SetBody(body *models.Role) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
