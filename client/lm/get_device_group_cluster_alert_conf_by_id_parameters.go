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

// NewGetDeviceGroupClusterAlertConfByIDParams creates a new GetDeviceGroupClusterAlertConfByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceGroupClusterAlertConfByIDParams() *GetDeviceGroupClusterAlertConfByIDParams {
	return &GetDeviceGroupClusterAlertConfByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupClusterAlertConfByIDParamsWithTimeout creates a new GetDeviceGroupClusterAlertConfByIDParams object
// with the ability to set a timeout on a request.
func NewGetDeviceGroupClusterAlertConfByIDParamsWithTimeout(timeout time.Duration) *GetDeviceGroupClusterAlertConfByIDParams {
	return &GetDeviceGroupClusterAlertConfByIDParams{
		timeout: timeout,
	}
}

// NewGetDeviceGroupClusterAlertConfByIDParamsWithContext creates a new GetDeviceGroupClusterAlertConfByIDParams object
// with the ability to set a context for a request.
func NewGetDeviceGroupClusterAlertConfByIDParamsWithContext(ctx context.Context) *GetDeviceGroupClusterAlertConfByIDParams {
	return &GetDeviceGroupClusterAlertConfByIDParams{
		Context: ctx,
	}
}

// NewGetDeviceGroupClusterAlertConfByIDParamsWithHTTPClient creates a new GetDeviceGroupClusterAlertConfByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceGroupClusterAlertConfByIDParamsWithHTTPClient(client *http.Client) *GetDeviceGroupClusterAlertConfByIDParams {
	return &GetDeviceGroupClusterAlertConfByIDParams{
		HTTPClient: client,
	}
}

/* GetDeviceGroupClusterAlertConfByIDParams contains all the parameters to send to the API endpoint
   for the get device group cluster alert conf by Id operation.

   Typically these are written to a http.Request.
*/
type GetDeviceGroupClusterAlertConfByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// DeviceGroupID.
	//
	// Format: int32
	DeviceGroupID int32

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device group cluster alert conf by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupClusterAlertConfByIDParams) WithDefaults() *GetDeviceGroupClusterAlertConfByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device group cluster alert conf by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupClusterAlertConfByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := GetDeviceGroupClusterAlertConfByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) WithTimeout(timeout time.Duration) *GetDeviceGroupClusterAlertConfByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) WithContext(ctx context.Context) *GetDeviceGroupClusterAlertConfByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) WithHTTPClient(client *http.Client) *GetDeviceGroupClusterAlertConfByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) WithUserAgent(userAgent *string) *GetDeviceGroupClusterAlertConfByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceGroupID adds the deviceGroupID to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) WithDeviceGroupID(deviceGroupID int32) *GetDeviceGroupClusterAlertConfByIDParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithID adds the id to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) WithID(id int32) *GetDeviceGroupClusterAlertConfByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device group cluster alert conf by Id params
func (o *GetDeviceGroupClusterAlertConfByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupClusterAlertConfByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
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