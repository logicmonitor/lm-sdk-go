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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewAddAPITokenByAdminIDParams creates a new AddAPITokenByAdminIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAPITokenByAdminIDParams() *AddAPITokenByAdminIDParams {
	return &AddAPITokenByAdminIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAPITokenByAdminIDParamsWithTimeout creates a new AddAPITokenByAdminIDParams object
// with the ability to set a timeout on a request.
func NewAddAPITokenByAdminIDParamsWithTimeout(timeout time.Duration) *AddAPITokenByAdminIDParams {
	return &AddAPITokenByAdminIDParams{
		timeout: timeout,
	}
}

// NewAddAPITokenByAdminIDParamsWithContext creates a new AddAPITokenByAdminIDParams object
// with the ability to set a context for a request.
func NewAddAPITokenByAdminIDParamsWithContext(ctx context.Context) *AddAPITokenByAdminIDParams {
	return &AddAPITokenByAdminIDParams{
		Context: ctx,
	}
}

// NewAddAPITokenByAdminIDParamsWithHTTPClient creates a new AddAPITokenByAdminIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAPITokenByAdminIDParamsWithHTTPClient(client *http.Client) *AddAPITokenByAdminIDParams {
	return &AddAPITokenByAdminIDParams{
		HTTPClient: client,
	}
}

/*
AddAPITokenByAdminIDParams contains all the parameters to send to the API endpoint

	for the add Api token by admin Id operation.

	Typically these are written to a http.Request.
*/
type AddAPITokenByAdminIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// AdminID.
	//
	// Format: int32
	AdminID int32

	// Body.
	Body *models.APIToken

	// Type.
	//
	// Default: "API Token"
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add Api token by admin Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAPITokenByAdminIDParams) WithDefaults() *AddAPITokenByAdminIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add Api token by admin Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAPITokenByAdminIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")

		typeVarDefault = string("API Token")
	)

	val := AddAPITokenByAdminIDParams{
		UserAgent: &userAgentDefault,
		Type:      &typeVarDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithTimeout(timeout time.Duration) *AddAPITokenByAdminIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithContext(ctx context.Context) *AddAPITokenByAdminIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithHTTPClient(client *http.Client) *AddAPITokenByAdminIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithUserAgent(userAgent *string) *AddAPITokenByAdminIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithAdminID adds the adminID to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithAdminID(adminID int32) *AddAPITokenByAdminIDParams {
	o.SetAdminID(adminID)
	return o
}

// SetAdminID adds the adminId to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetAdminID(adminID int32) {
	o.AdminID = adminID
}

// WithBody adds the body to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithBody(body *models.APIToken) *AddAPITokenByAdminIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetBody(body *models.APIToken) {
	o.Body = body
}

// WithType adds the typeVar to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithType(typeVar *string) *AddAPITokenByAdminIDParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *AddAPITokenByAdminIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param adminId
	if err := r.SetPathParam("adminId", swag.FormatInt32(o.AdminID)); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
