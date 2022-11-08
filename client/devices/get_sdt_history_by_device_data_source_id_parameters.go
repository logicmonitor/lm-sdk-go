// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewGetSDTHistoryByDeviceDataSourceIDParams creates a new GetSDTHistoryByDeviceDataSourceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSDTHistoryByDeviceDataSourceIDParams() *GetSDTHistoryByDeviceDataSourceIDParams {
	return &GetSDTHistoryByDeviceDataSourceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSDTHistoryByDeviceDataSourceIDParamsWithTimeout creates a new GetSDTHistoryByDeviceDataSourceIDParams object
// with the ability to set a timeout on a request.
func NewGetSDTHistoryByDeviceDataSourceIDParamsWithTimeout(timeout time.Duration) *GetSDTHistoryByDeviceDataSourceIDParams {
	return &GetSDTHistoryByDeviceDataSourceIDParams{
		timeout: timeout,
	}
}

// NewGetSDTHistoryByDeviceDataSourceIDParamsWithContext creates a new GetSDTHistoryByDeviceDataSourceIDParams object
// with the ability to set a context for a request.
func NewGetSDTHistoryByDeviceDataSourceIDParamsWithContext(ctx context.Context) *GetSDTHistoryByDeviceDataSourceIDParams {
	return &GetSDTHistoryByDeviceDataSourceIDParams{
		Context: ctx,
	}
}

// NewGetSDTHistoryByDeviceDataSourceIDParamsWithHTTPClient creates a new GetSDTHistoryByDeviceDataSourceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSDTHistoryByDeviceDataSourceIDParamsWithHTTPClient(client *http.Client) *GetSDTHistoryByDeviceDataSourceIDParams {
	return &GetSDTHistoryByDeviceDataSourceIDParams{
		HTTPClient: client,
	}
}

/*
GetSDTHistoryByDeviceDataSourceIDParams contains all the parameters to send to the API endpoint

	for the get SDT history by device data source Id operation.

	Typically these are written to a http.Request.
*/
type GetSDTHistoryByDeviceDataSourceIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get SDT history by device data source Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithDefaults() *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get SDT history by device data source Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetSDTHistoryByDeviceDataSourceIDParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithTimeout(timeout time.Duration) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithContext(ctx context.Context) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithHTTPClient(client *http.Client) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithUserAgent(userAgent *string) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceID adds the deviceID to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithDeviceID(deviceID int32) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithFields(fields *string) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithFilter(filter *string) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithID(id int32) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetID(id int32) {
	o.ID = id
}

// WithOffset adds the offset to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithOffset(offset *int32) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WithSize(size *int32) *GetSDTHistoryByDeviceDataSourceIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get SDT history by device data source Id params
func (o *GetSDTHistoryByDeviceDataSourceIDParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetSDTHistoryByDeviceDataSourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
