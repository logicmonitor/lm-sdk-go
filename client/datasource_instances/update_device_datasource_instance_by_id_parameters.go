// Code generated by go-swagger; DO NOT EDIT.

package datasource_instances

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

// NewUpdateDeviceDatasourceInstanceByIDParams creates a new UpdateDeviceDatasourceInstanceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceDatasourceInstanceByIDParams() *UpdateDeviceDatasourceInstanceByIDParams {
	return &UpdateDeviceDatasourceInstanceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceDatasourceInstanceByIDParamsWithTimeout creates a new UpdateDeviceDatasourceInstanceByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceDatasourceInstanceByIDParamsWithTimeout(timeout time.Duration) *UpdateDeviceDatasourceInstanceByIDParams {
	return &UpdateDeviceDatasourceInstanceByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceDatasourceInstanceByIDParamsWithContext creates a new UpdateDeviceDatasourceInstanceByIDParams object
// with the ability to set a context for a request.
func NewUpdateDeviceDatasourceInstanceByIDParamsWithContext(ctx context.Context) *UpdateDeviceDatasourceInstanceByIDParams {
	return &UpdateDeviceDatasourceInstanceByIDParams{
		Context: ctx,
	}
}

// NewUpdateDeviceDatasourceInstanceByIDParamsWithHTTPClient creates a new UpdateDeviceDatasourceInstanceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceDatasourceInstanceByIDParamsWithHTTPClient(client *http.Client) *UpdateDeviceDatasourceInstanceByIDParams {
	return &UpdateDeviceDatasourceInstanceByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateDeviceDatasourceInstanceByIDParams contains all the parameters to send to the API endpoint

	for the update device datasource instance by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateDeviceDatasourceInstanceByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// Body.
	Body *models.DeviceDataSourceInstance

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	/* HdsID.

	   The device-datasource ID

	   Format: int32
	*/
	HdsID int32

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

// WithDefaults hydrates default values in the update device datasource instance by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithDefaults() *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device datasource instance by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")

		opTypeDefault = string("refresh")
	)

	val := UpdateDeviceDatasourceInstanceByIDParams{
		UserAgent: &userAgentDefault,
		OpType:    &opTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithTimeout(timeout time.Duration) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithContext(ctx context.Context) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithHTTPClient(client *http.Client) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithUserAgent(userAgent *string) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithBody(body *models.DeviceDataSourceInstance) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetBody(body *models.DeviceDataSourceInstance) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithDeviceID(deviceID int32) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithHdsID adds the hdsID to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithHdsID(hdsID int32) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithID adds the id to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithID(id int32) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOpType adds the opType to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithOpType(opType *string) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetOpType(opType)
	return o
}

// SetOpType adds the opType to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetOpType(opType *string) {
	o.OpType = opType
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceDatasourceInstanceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
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
