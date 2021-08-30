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
)

// NewGetDeviceDatasourceInstanceGroupByIDParams creates a new GetDeviceDatasourceInstanceGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceDatasourceInstanceGroupByIDParams() *GetDeviceDatasourceInstanceGroupByIDParams {
	return &GetDeviceDatasourceInstanceGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceInstanceGroupByIDParamsWithTimeout creates a new GetDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a timeout on a request.
func NewGetDeviceDatasourceInstanceGroupByIDParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGroupByIDParams {
	return &GetDeviceDatasourceInstanceGroupByIDParams{
		timeout: timeout,
	}
}

// NewGetDeviceDatasourceInstanceGroupByIDParamsWithContext creates a new GetDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a context for a request.
func NewGetDeviceDatasourceInstanceGroupByIDParamsWithContext(ctx context.Context) *GetDeviceDatasourceInstanceGroupByIDParams {
	return &GetDeviceDatasourceInstanceGroupByIDParams{
		Context: ctx,
	}
}

// NewGetDeviceDatasourceInstanceGroupByIDParamsWithHTTPClient creates a new GetDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceDatasourceInstanceGroupByIDParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGroupByIDParams {
	return &GetDeviceDatasourceInstanceGroupByIDParams{
		HTTPClient: client,
	}
}

/* GetDeviceDatasourceInstanceGroupByIDParams contains all the parameters to send to the API endpoint
   for the get device datasource instance group by Id operation.

   Typically these are written to a http.Request.
*/
type GetDeviceDatasourceInstanceGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	/* DeviceDsID.

	   The device-datasource ID you'd like to add an instance group for

	   Format: int32
	*/
	DeviceDsID int32

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	// Fields.
	Fields *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device datasource instance group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithDefaults() *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device datasource instance group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := GetDeviceDatasourceInstanceGroupByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithContext(ctx context.Context) *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithUserAgent(userAgent *string) *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceDsID adds the deviceDsID to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithDeviceDsID(deviceDsID int32) *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetDeviceDsID(deviceDsID)
	return o
}

// SetDeviceDsID adds the deviceDsId to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetDeviceDsID(deviceDsID int32) {
	o.DeviceDsID = deviceDsID
}

// WithDeviceID adds the deviceID to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithFields(fields *string) *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WithID(id int32) *GetDeviceDatasourceInstanceGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device datasource instance group by Id params
func (o *GetDeviceDatasourceInstanceGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceInstanceGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceDsId
	if err := r.SetPathParam("deviceDsId", swag.FormatInt32(o.DeviceDsID)); err != nil {
		return err
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
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