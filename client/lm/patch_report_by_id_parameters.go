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

// NewPatchReportByIDParams creates a new PatchReportByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchReportByIDParams() *PatchReportByIDParams {
	return &PatchReportByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchReportByIDParamsWithTimeout creates a new PatchReportByIDParams object
// with the ability to set a timeout on a request.
func NewPatchReportByIDParamsWithTimeout(timeout time.Duration) *PatchReportByIDParams {
	return &PatchReportByIDParams{
		timeout: timeout,
	}
}

// NewPatchReportByIDParamsWithContext creates a new PatchReportByIDParams object
// with the ability to set a context for a request.
func NewPatchReportByIDParamsWithContext(ctx context.Context) *PatchReportByIDParams {
	return &PatchReportByIDParams{
		Context: ctx,
	}
}

// NewPatchReportByIDParamsWithHTTPClient creates a new PatchReportByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchReportByIDParamsWithHTTPClient(client *http.Client) *PatchReportByIDParams {
	return &PatchReportByIDParams{
		HTTPClient: client,
	}
}

/*
PatchReportByIDParams contains all the parameters to send to the API endpoint

	for the patch report by Id operation.

	Typically these are written to a http.Request.
*/
type PatchReportByIDParams struct {

	// Body.
	Body models.ReportBase

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch report by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchReportByIDParams) WithDefaults() *PatchReportByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch report by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchReportByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch report by Id params
func (o *PatchReportByIDParams) WithTimeout(timeout time.Duration) *PatchReportByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch report by Id params
func (o *PatchReportByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch report by Id params
func (o *PatchReportByIDParams) WithContext(ctx context.Context) *PatchReportByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch report by Id params
func (o *PatchReportByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch report by Id params
func (o *PatchReportByIDParams) WithHTTPClient(client *http.Client) *PatchReportByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch report by Id params
func (o *PatchReportByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch report by Id params
func (o *PatchReportByIDParams) WithBody(body models.ReportBase) *PatchReportByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch report by Id params
func (o *PatchReportByIDParams) SetBody(body models.ReportBase) {
	o.Body = body
}

// WithID adds the id to the patch report by Id params
func (o *PatchReportByIDParams) WithID(id int32) *PatchReportByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch report by Id params
func (o *PatchReportByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchReportByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
