// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// NewUpdateDashboardByIDParams creates a new UpdateDashboardByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDashboardByIDParams() *UpdateDashboardByIDParams {
	return &UpdateDashboardByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDashboardByIDParamsWithTimeout creates a new UpdateDashboardByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDashboardByIDParamsWithTimeout(timeout time.Duration) *UpdateDashboardByIDParams {
	return &UpdateDashboardByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDashboardByIDParamsWithContext creates a new UpdateDashboardByIDParams object
// with the ability to set a context for a request.
func NewUpdateDashboardByIDParamsWithContext(ctx context.Context) *UpdateDashboardByIDParams {
	return &UpdateDashboardByIDParams{
		Context: ctx,
	}
}

// NewUpdateDashboardByIDParamsWithHTTPClient creates a new UpdateDashboardByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDashboardByIDParamsWithHTTPClient(client *http.Client) *UpdateDashboardByIDParams {
	return &UpdateDashboardByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateDashboardByIDParams contains all the parameters to send to the API endpoint

	for the update dashboard by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateDashboardByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// Body.
	Body *models.Dashboard

	// ID.
	//
	// Format: int32
	ID int32

	// OverwriteGroupFields.
	OverwriteGroupFields *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update dashboard by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardByIDParams) WithDefaults() *UpdateDashboardByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update dashboard by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")

		overwriteGroupFieldsDefault = bool(false)
	)

	val := UpdateDashboardByIDParams{
		UserAgent:            &userAgentDefault,
		OverwriteGroupFields: &overwriteGroupFieldsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) WithTimeout(timeout time.Duration) *UpdateDashboardByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) WithContext(ctx context.Context) *UpdateDashboardByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) WithHTTPClient(client *http.Client) *UpdateDashboardByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) WithUserAgent(userAgent *string) *UpdateDashboardByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) WithBody(body *models.Dashboard) *UpdateDashboardByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) SetBody(body *models.Dashboard) {
	o.Body = body
}

// WithID adds the id to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) WithID(id int32) *UpdateDashboardByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOverwriteGroupFields adds the overwriteGroupFields to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) WithOverwriteGroupFields(overwriteGroupFields *bool) *UpdateDashboardByIDParams {
	o.SetOverwriteGroupFields(overwriteGroupFields)
	return o
}

// SetOverwriteGroupFields adds the overwriteGroupFields to the update dashboard by Id params
func (o *UpdateDashboardByIDParams) SetOverwriteGroupFields(overwriteGroupFields *bool) {
	o.OverwriteGroupFields = overwriteGroupFields
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDashboardByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.OverwriteGroupFields != nil {

		// query param overwriteGroupFields
		var qrOverwriteGroupFields bool

		if o.OverwriteGroupFields != nil {
			qrOverwriteGroupFields = *o.OverwriteGroupFields
		}
		qOverwriteGroupFields := swag.FormatBool(qrOverwriteGroupFields)
		if qOverwriteGroupFields != "" {

			if err := r.SetQueryParam("overwriteGroupFields", qOverwriteGroupFields); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
