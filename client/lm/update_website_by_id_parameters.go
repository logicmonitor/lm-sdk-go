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

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// NewUpdateWebsiteByIDParams creates a new UpdateWebsiteByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateWebsiteByIDParams() *UpdateWebsiteByIDParams {
	return &UpdateWebsiteByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWebsiteByIDParamsWithTimeout creates a new UpdateWebsiteByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateWebsiteByIDParamsWithTimeout(timeout time.Duration) *UpdateWebsiteByIDParams {
	return &UpdateWebsiteByIDParams{
		timeout: timeout,
	}
}

// NewUpdateWebsiteByIDParamsWithContext creates a new UpdateWebsiteByIDParams object
// with the ability to set a context for a request.
func NewUpdateWebsiteByIDParamsWithContext(ctx context.Context) *UpdateWebsiteByIDParams {
	return &UpdateWebsiteByIDParams{
		Context: ctx,
	}
}

// NewUpdateWebsiteByIDParamsWithHTTPClient creates a new UpdateWebsiteByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateWebsiteByIDParamsWithHTTPClient(client *http.Client) *UpdateWebsiteByIDParams {
	return &UpdateWebsiteByIDParams{
		HTTPClient: client,
	}
}

/* UpdateWebsiteByIDParams contains all the parameters to send to the API endpoint
   for the update website by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateWebsiteByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body models.Website

	// ID.
	//
	// Format: int32
	ID int32

	// OpType.
	//
	// Default: "refresh"
	OpType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update website by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWebsiteByIDParams) WithDefaults() *UpdateWebsiteByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update website by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWebsiteByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		opTypeDefault = string("refresh")
	)

	val := UpdateWebsiteByIDParams{
		UserAgent: &userAgentDefault,
		OpType:    &opTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithTimeout(timeout time.Duration) *UpdateWebsiteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithContext(ctx context.Context) *UpdateWebsiteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithHTTPClient(client *http.Client) *UpdateWebsiteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithUserAgent(userAgent *string) *UpdateWebsiteByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithBody(body models.Website) *UpdateWebsiteByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetBody(body models.Website) {
	o.Body = body
}

// WithID adds the id to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithID(id int32) *UpdateWebsiteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOpType adds the opType to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithOpType(opType *string) *UpdateWebsiteByIDParams {
	o.SetOpType(opType)
	return o
}

// SetOpType adds the opType to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetOpType(opType *string) {
	o.OpType = opType
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWebsiteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.OpType != nil {

		// query param opType
		var qrOpType string

		if o.OpType != nil {
			qrOpType = *o.OpType
		}
		qOpType := qrOpType
		if qOpType != "" {

			if err := r.SetQueryParam("opType", qOpType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
