// Code generated by go-swagger; DO NOT EDIT.

package api_tokens

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

// NewGetAPITokenListParams creates a new GetAPITokenListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPITokenListParams() *GetAPITokenListParams {
	return &GetAPITokenListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPITokenListParamsWithTimeout creates a new GetAPITokenListParams object
// with the ability to set a timeout on a request.
func NewGetAPITokenListParamsWithTimeout(timeout time.Duration) *GetAPITokenListParams {
	return &GetAPITokenListParams{
		timeout: timeout,
	}
}

// NewGetAPITokenListParamsWithContext creates a new GetAPITokenListParams object
// with the ability to set a context for a request.
func NewGetAPITokenListParamsWithContext(ctx context.Context) *GetAPITokenListParams {
	return &GetAPITokenListParams{
		Context: ctx,
	}
}

// NewGetAPITokenListParamsWithHTTPClient creates a new GetAPITokenListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPITokenListParamsWithHTTPClient(client *http.Client) *GetAPITokenListParams {
	return &GetAPITokenListParams{
		HTTPClient: client,
	}
}

/*
GetAPITokenListParams contains all the parameters to send to the API endpoint

	for the get Api token list operation.

	Typically these are written to a http.Request.
*/
type GetAPITokenListParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Permission.
	Permission *string

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	// Type.
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Api token list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITokenListParams) WithDefaults() *GetAPITokenListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Api token list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITokenListParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetAPITokenListParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Api token list params
func (o *GetAPITokenListParams) WithTimeout(timeout time.Duration) *GetAPITokenListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api token list params
func (o *GetAPITokenListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api token list params
func (o *GetAPITokenListParams) WithContext(ctx context.Context) *GetAPITokenListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api token list params
func (o *GetAPITokenListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api token list params
func (o *GetAPITokenListParams) WithHTTPClient(client *http.Client) *GetAPITokenListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api token list params
func (o *GetAPITokenListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get Api token list params
func (o *GetAPITokenListParams) WithUserAgent(userAgent *string) *GetAPITokenListParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get Api token list params
func (o *GetAPITokenListParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get Api token list params
func (o *GetAPITokenListParams) WithFields(fields *string) *GetAPITokenListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get Api token list params
func (o *GetAPITokenListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get Api token list params
func (o *GetAPITokenListParams) WithFilter(filter *string) *GetAPITokenListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get Api token list params
func (o *GetAPITokenListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get Api token list params
func (o *GetAPITokenListParams) WithOffset(offset *int32) *GetAPITokenListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get Api token list params
func (o *GetAPITokenListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithPermission adds the permission to the get Api token list params
func (o *GetAPITokenListParams) WithPermission(permission *string) *GetAPITokenListParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the get Api token list params
func (o *GetAPITokenListParams) SetPermission(permission *string) {
	o.Permission = permission
}

// WithSize adds the size to the get Api token list params
func (o *GetAPITokenListParams) WithSize(size *int32) *GetAPITokenListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get Api token list params
func (o *GetAPITokenListParams) SetSize(size *int32) {
	o.Size = size
}

// WithType adds the typeVar to the get Api token list params
func (o *GetAPITokenListParams) WithType(typeVar *string) *GetAPITokenListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get Api token list params
func (o *GetAPITokenListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPITokenListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Permission != nil {

		// query param permission
		var qrPermission string

		if o.Permission != nil {
			qrPermission = *o.Permission
		}
		qPermission := qrPermission
		if qPermission != "" {

			if err := r.SetQueryParam("permission", qPermission); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
