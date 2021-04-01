// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteDeviceDatasourceInstanceByIDParams creates a new DeleteDeviceDatasourceInstanceByIDParams object
// with the default values initialized.
func NewDeleteDeviceDatasourceInstanceByIDParams() *DeleteDeviceDatasourceInstanceByIDParams {
	var ()
	return &DeleteDeviceDatasourceInstanceByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceDatasourceInstanceByIDParamsWithTimeout creates a new DeleteDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeviceDatasourceInstanceByIDParamsWithTimeout(timeout time.Duration) *DeleteDeviceDatasourceInstanceByIDParams {
	var ()
	return &DeleteDeviceDatasourceInstanceByIDParams{

		timeout: timeout,
	}
}

// NewDeleteDeviceDatasourceInstanceByIDParamsWithContext creates a new DeleteDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeviceDatasourceInstanceByIDParamsWithContext(ctx context.Context) *DeleteDeviceDatasourceInstanceByIDParams {
	var ()
	return &DeleteDeviceDatasourceInstanceByIDParams{

		Context: ctx,
	}
}

// NewDeleteDeviceDatasourceInstanceByIDParamsWithHTTPClient creates a new DeleteDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeviceDatasourceInstanceByIDParamsWithHTTPClient(client *http.Client) *DeleteDeviceDatasourceInstanceByIDParams {
	var ()
	return &DeleteDeviceDatasourceInstanceByIDParams{
		HTTPClient: client,
	}
}

/*DeleteDeviceDatasourceInstanceByIDParams contains all the parameters to send to the API endpoint
for the delete device datasource instance by Id operation typically these are written to a http.Request
*/
type DeleteDeviceDatasourceInstanceByIDParams struct {

	/*DeviceID*/
	DeviceID int32
	/*HdsID
	  The device-datasource ID

	*/
	HdsID int32
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) WithTimeout(timeout time.Duration) *DeleteDeviceDatasourceInstanceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) WithContext(ctx context.Context) *DeleteDeviceDatasourceInstanceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) WithHTTPClient(client *http.Client) *DeleteDeviceDatasourceInstanceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) WithDeviceID(deviceID int32) *DeleteDeviceDatasourceInstanceByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithHdsID adds the hdsID to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) WithHdsID(hdsID int32) *DeleteDeviceDatasourceInstanceByIDParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithID adds the id to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) WithID(id int32) *DeleteDeviceDatasourceInstanceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete device datasource instance by Id params
func (o *DeleteDeviceDatasourceInstanceByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceDatasourceInstanceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}