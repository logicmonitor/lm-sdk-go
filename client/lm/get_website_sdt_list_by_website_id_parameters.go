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

// NewGetWebsiteSDTListByWebsiteIDParams creates a new GetWebsiteSDTListByWebsiteIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsiteSDTListByWebsiteIDParams() *GetWebsiteSDTListByWebsiteIDParams {
	return &GetWebsiteSDTListByWebsiteIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsiteSDTListByWebsiteIDParamsWithTimeout creates a new GetWebsiteSDTListByWebsiteIDParams object
// with the ability to set a timeout on a request.
func NewGetWebsiteSDTListByWebsiteIDParamsWithTimeout(timeout time.Duration) *GetWebsiteSDTListByWebsiteIDParams {
	return &GetWebsiteSDTListByWebsiteIDParams{
		timeout: timeout,
	}
}

// NewGetWebsiteSDTListByWebsiteIDParamsWithContext creates a new GetWebsiteSDTListByWebsiteIDParams object
// with the ability to set a context for a request.
func NewGetWebsiteSDTListByWebsiteIDParamsWithContext(ctx context.Context) *GetWebsiteSDTListByWebsiteIDParams {
	return &GetWebsiteSDTListByWebsiteIDParams{
		Context: ctx,
	}
}

// NewGetWebsiteSDTListByWebsiteIDParamsWithHTTPClient creates a new GetWebsiteSDTListByWebsiteIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsiteSDTListByWebsiteIDParamsWithHTTPClient(client *http.Client) *GetWebsiteSDTListByWebsiteIDParams {
	return &GetWebsiteSDTListByWebsiteIDParams{
		HTTPClient: client,
	}
}

/*
GetWebsiteSDTListByWebsiteIDParams contains all the parameters to send to the API endpoint

	for the get website SDT list by website Id operation.

	Typically these are written to a http.Request.
*/
type GetWebsiteSDTListByWebsiteIDParams struct {

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

// WithDefaults hydrates default values in the get website SDT list by website Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteSDTListByWebsiteIDParams) WithDefaults() *GetWebsiteSDTListByWebsiteIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get website SDT list by website Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteSDTListByWebsiteIDParams) SetDefaults() {
	var (
		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetWebsiteSDTListByWebsiteIDParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) WithTimeout(timeout time.Duration) *GetWebsiteSDTListByWebsiteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) WithContext(ctx context.Context) *GetWebsiteSDTListByWebsiteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) WithHTTPClient(client *http.Client) *GetWebsiteSDTListByWebsiteIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) WithFields(fields *string) *GetWebsiteSDTListByWebsiteIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) WithFilter(filter *string) *GetWebsiteSDTListByWebsiteIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) WithID(id int32) *GetWebsiteSDTListByWebsiteIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) SetID(id int32) {
	o.ID = id
}

// WithOffset adds the offset to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) WithOffset(offset *int32) *GetWebsiteSDTListByWebsiteIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) WithSize(size *int32) *GetWebsiteSDTListByWebsiteIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get website SDT list by website Id params
func (o *GetWebsiteSDTListByWebsiteIDParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteSDTListByWebsiteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
