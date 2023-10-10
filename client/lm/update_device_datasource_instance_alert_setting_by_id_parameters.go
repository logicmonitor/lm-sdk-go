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

// NewUpdateDeviceDatasourceInstanceAlertSettingByIDParams creates a new UpdateDeviceDatasourceInstanceAlertSettingByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceDatasourceInstanceAlertSettingByIDParams() *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	return &UpdateDeviceDatasourceInstanceAlertSettingByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceDatasourceInstanceAlertSettingByIDParamsWithTimeout creates a new UpdateDeviceDatasourceInstanceAlertSettingByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceDatasourceInstanceAlertSettingByIDParamsWithTimeout(timeout time.Duration) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	return &UpdateDeviceDatasourceInstanceAlertSettingByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceDatasourceInstanceAlertSettingByIDParamsWithContext creates a new UpdateDeviceDatasourceInstanceAlertSettingByIDParams object
// with the ability to set a context for a request.
func NewUpdateDeviceDatasourceInstanceAlertSettingByIDParamsWithContext(ctx context.Context) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	return &UpdateDeviceDatasourceInstanceAlertSettingByIDParams{
		Context: ctx,
	}
}

// NewUpdateDeviceDatasourceInstanceAlertSettingByIDParamsWithHTTPClient creates a new UpdateDeviceDatasourceInstanceAlertSettingByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceDatasourceInstanceAlertSettingByIDParamsWithHTTPClient(client *http.Client) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	return &UpdateDeviceDatasourceInstanceAlertSettingByIDParams{
		HTTPClient: client,
	}
}

/* UpdateDeviceDatasourceInstanceAlertSettingByIDParams contains all the parameters to send to the API endpoint
   for the update device datasource instance alert setting by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateDeviceDatasourceInstanceAlertSettingByIDParams struct {

	// Body.
	Body *models.DeviceDataSourceInstanceAlertSetting

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	/* HdsID.

	   Device-DataSource ID

	   Format: int32
	*/
	HdsID int32

	// ID.
	//
	// Format: int32
	ID int32

	// InstanceID.
	//
	// Format: int32
	InstanceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device datasource instance alert setting by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithDefaults() *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device datasource instance alert setting by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithTimeout(timeout time.Duration) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithContext(ctx context.Context) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithHTTPClient(client *http.Client) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithBody(body *models.DeviceDataSourceInstanceAlertSetting) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetBody(body *models.DeviceDataSourceInstanceAlertSetting) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithDeviceID(deviceID int32) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithHdsID adds the hdsID to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithHdsID(hdsID int32) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithID adds the id to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithID(id int32) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetID(id int32) {
	o.ID = id
}

// WithInstanceID adds the instanceID to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WithInstanceID(instanceID int32) *UpdateDeviceDatasourceInstanceAlertSettingByIDParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the update device datasource instance alert setting by Id params
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) SetInstanceID(instanceID int32) {
	o.InstanceID = instanceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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

	// path param instanceId
	if err := r.SetPathParam("instanceId", swag.FormatInt32(o.InstanceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
