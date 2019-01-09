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

// NewGetDeviceGroupDatasourceAlertSettingParams creates a new GetDeviceGroupDatasourceAlertSettingParams object
// with the default values initialized.
func NewGetDeviceGroupDatasourceAlertSettingParams() *GetDeviceGroupDatasourceAlertSettingParams {
	var ()
	return &GetDeviceGroupDatasourceAlertSettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupDatasourceAlertSettingParamsWithTimeout creates a new GetDeviceGroupDatasourceAlertSettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceGroupDatasourceAlertSettingParamsWithTimeout(timeout time.Duration) *GetDeviceGroupDatasourceAlertSettingParams {
	var ()
	return &GetDeviceGroupDatasourceAlertSettingParams{

		timeout: timeout,
	}
}

// NewGetDeviceGroupDatasourceAlertSettingParamsWithContext creates a new GetDeviceGroupDatasourceAlertSettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceGroupDatasourceAlertSettingParamsWithContext(ctx context.Context) *GetDeviceGroupDatasourceAlertSettingParams {
	var ()
	return &GetDeviceGroupDatasourceAlertSettingParams{

		Context: ctx,
	}
}

// NewGetDeviceGroupDatasourceAlertSettingParamsWithHTTPClient creates a new GetDeviceGroupDatasourceAlertSettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceGroupDatasourceAlertSettingParamsWithHTTPClient(client *http.Client) *GetDeviceGroupDatasourceAlertSettingParams {
	var ()
	return &GetDeviceGroupDatasourceAlertSettingParams{
		HTTPClient: client,
	}
}

/*GetDeviceGroupDatasourceAlertSettingParams contains all the parameters to send to the API endpoint
for the get device group datasource alert setting operation typically these are written to a http.Request
*/
type GetDeviceGroupDatasourceAlertSettingParams struct {

	/*DeviceGroupID*/
	DeviceGroupID int32
	/*DsID*/
	DsID int32
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) WithTimeout(timeout time.Duration) *GetDeviceGroupDatasourceAlertSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) WithContext(ctx context.Context) *GetDeviceGroupDatasourceAlertSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) WithHTTPClient(client *http.Client) *GetDeviceGroupDatasourceAlertSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceGroupID adds the deviceGroupID to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) WithDeviceGroupID(deviceGroupID int32) *GetDeviceGroupDatasourceAlertSettingParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithDsID adds the dsID to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) WithDsID(dsID int32) *GetDeviceGroupDatasourceAlertSettingParams {
	o.SetDsID(dsID)
	return o
}

// SetDsID adds the dsId to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) SetDsID(dsID int32) {
	o.DsID = dsID
}

// WithFields adds the fields to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) WithFields(fields *string) *GetDeviceGroupDatasourceAlertSettingParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device group datasource alert setting params
func (o *GetDeviceGroupDatasourceAlertSettingParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupDatasourceAlertSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
	}

	// path param dsId
	if err := r.SetPathParam("dsId", swag.FormatInt32(o.DsID)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}