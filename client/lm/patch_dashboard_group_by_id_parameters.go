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

// NewPatchDashboardGroupByIDParams creates a new PatchDashboardGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDashboardGroupByIDParams() *PatchDashboardGroupByIDParams {
	return &PatchDashboardGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDashboardGroupByIDParamsWithTimeout creates a new PatchDashboardGroupByIDParams object
// with the ability to set a timeout on a request.
func NewPatchDashboardGroupByIDParamsWithTimeout(timeout time.Duration) *PatchDashboardGroupByIDParams {
	return &PatchDashboardGroupByIDParams{
		timeout: timeout,
	}
}

// NewPatchDashboardGroupByIDParamsWithContext creates a new PatchDashboardGroupByIDParams object
// with the ability to set a context for a request.
func NewPatchDashboardGroupByIDParamsWithContext(ctx context.Context) *PatchDashboardGroupByIDParams {
	return &PatchDashboardGroupByIDParams{
		Context: ctx,
	}
}

// NewPatchDashboardGroupByIDParamsWithHTTPClient creates a new PatchDashboardGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDashboardGroupByIDParamsWithHTTPClient(client *http.Client) *PatchDashboardGroupByIDParams {
	return &PatchDashboardGroupByIDParams{
		HTTPClient: client,
	}
}

/* PatchDashboardGroupByIDParams contains all the parameters to send to the API endpoint
   for the patch dashboard group by Id operation.

   Typically these are written to a http.Request.
*/
type PatchDashboardGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.DashboardGroup

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch dashboard group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDashboardGroupByIDParams) WithDefaults() *PatchDashboardGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch dashboard group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDashboardGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := PatchDashboardGroupByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) WithTimeout(timeout time.Duration) *PatchDashboardGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) WithContext(ctx context.Context) *PatchDashboardGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) WithHTTPClient(client *http.Client) *PatchDashboardGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) WithUserAgent(userAgent *string) *PatchDashboardGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) WithBody(body *models.DashboardGroup) *PatchDashboardGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) SetBody(body *models.DashboardGroup) {
	o.Body = body
}

// WithID adds the id to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) WithID(id int32) *PatchDashboardGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch dashboard group by Id params
func (o *PatchDashboardGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDashboardGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
