// Code generated by go-swagger; DO NOT EDIT.

package device_groups

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

// NewPatchDeviceGroupPropertyByNameParams creates a new PatchDeviceGroupPropertyByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDeviceGroupPropertyByNameParams() *PatchDeviceGroupPropertyByNameParams {
	return &PatchDeviceGroupPropertyByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDeviceGroupPropertyByNameParamsWithTimeout creates a new PatchDeviceGroupPropertyByNameParams object
// with the ability to set a timeout on a request.
func NewPatchDeviceGroupPropertyByNameParamsWithTimeout(timeout time.Duration) *PatchDeviceGroupPropertyByNameParams {
	return &PatchDeviceGroupPropertyByNameParams{
		timeout: timeout,
	}
}

// NewPatchDeviceGroupPropertyByNameParamsWithContext creates a new PatchDeviceGroupPropertyByNameParams object
// with the ability to set a context for a request.
func NewPatchDeviceGroupPropertyByNameParamsWithContext(ctx context.Context) *PatchDeviceGroupPropertyByNameParams {
	return &PatchDeviceGroupPropertyByNameParams{
		Context: ctx,
	}
}

// NewPatchDeviceGroupPropertyByNameParamsWithHTTPClient creates a new PatchDeviceGroupPropertyByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDeviceGroupPropertyByNameParamsWithHTTPClient(client *http.Client) *PatchDeviceGroupPropertyByNameParams {
	return &PatchDeviceGroupPropertyByNameParams{
		HTTPClient: client,
	}
}

/*
PatchDeviceGroupPropertyByNameParams contains all the parameters to send to the API endpoint

	for the patch device group property by name operation.

	Typically these are written to a http.Request.
*/
type PatchDeviceGroupPropertyByNameParams struct {

	// PatchFields.
	PatchFields *string

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// Body.
	Body *models.EntityProperty

	/* Gid.

	   group ID

	   Format: int32
	*/
	Gid int32

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch device group property by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceGroupPropertyByNameParams) WithDefaults() *PatchDeviceGroupPropertyByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch device group property by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceGroupPropertyByNameParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")
	)

	val := PatchDeviceGroupPropertyByNameParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) WithTimeout(timeout time.Duration) *PatchDeviceGroupPropertyByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) WithContext(ctx context.Context) *PatchDeviceGroupPropertyByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) WithHTTPClient(client *http.Client) *PatchDeviceGroupPropertyByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchFields adds the patchFields to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) WithPatchFields(patchFields *string) *PatchDeviceGroupPropertyByNameParams {
	o.SetPatchFields(patchFields)
	return o
}

// SetPatchFields adds the patchFields to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) SetPatchFields(patchFields *string) {
	o.PatchFields = patchFields
}

// WithUserAgent adds the userAgent to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) WithUserAgent(userAgent *string) *PatchDeviceGroupPropertyByNameParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) WithBody(body *models.EntityProperty) *PatchDeviceGroupPropertyByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) SetBody(body *models.EntityProperty) {
	o.Body = body
}

// WithGid adds the gid to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) WithGid(gid int32) *PatchDeviceGroupPropertyByNameParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) SetGid(gid int32) {
	o.Gid = gid
}

// WithName adds the name to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) WithName(name string) *PatchDeviceGroupPropertyByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch device group property by name params
func (o *PatchDeviceGroupPropertyByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDeviceGroupPropertyByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PatchFields != nil {

		// query param PatchFields
		var qrPatchFields string

		if o.PatchFields != nil {
			qrPatchFields = *o.PatchFields
		}
		qPatchFields := qrPatchFields
		if qPatchFields != "" {

			if err := r.SetQueryParam("PatchFields", qPatchFields); err != nil {
				return err
			}
		}
	}

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

	// path param gid
	if err := r.SetPathParam("gid", swag.FormatInt32(o.Gid)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
