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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewAddDeviceParams creates a new AddDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDeviceParams() *AddDeviceParams {
	return &AddDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDeviceParamsWithTimeout creates a new AddDeviceParams object
// with the ability to set a timeout on a request.
func NewAddDeviceParamsWithTimeout(timeout time.Duration) *AddDeviceParams {
	return &AddDeviceParams{
		timeout: timeout,
	}
}

// NewAddDeviceParamsWithContext creates a new AddDeviceParams object
// with the ability to set a context for a request.
func NewAddDeviceParamsWithContext(ctx context.Context) *AddDeviceParams {
	return &AddDeviceParams{
		Context: ctx,
	}
}

// NewAddDeviceParamsWithHTTPClient creates a new AddDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDeviceParamsWithHTTPClient(client *http.Client) *AddDeviceParams {
	return &AddDeviceParams{
		HTTPClient: client,
	}
}

/*
AddDeviceParams contains all the parameters to send to the API endpoint

	for the add device operation.

	Typically these are written to a http.Request.
*/
type AddDeviceParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// AddFromWizard.
	AddFromWizard *bool

	// Body.
	Body *models.Device

	// End.
	//
	// Format: int64
	End *int64

	// NetflowFilter.
	NetflowFilter *string

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDeviceParams) WithDefaults() *AddDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDeviceParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")

		addFromWizardDefault = bool(false)
	)

	val := AddDeviceParams{
		UserAgent:     &userAgentDefault,
		AddFromWizard: &addFromWizardDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add device params
func (o *AddDeviceParams) WithTimeout(timeout time.Duration) *AddDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device params
func (o *AddDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device params
func (o *AddDeviceParams) WithContext(ctx context.Context) *AddDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device params
func (o *AddDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device params
func (o *AddDeviceParams) WithHTTPClient(client *http.Client) *AddDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device params
func (o *AddDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add device params
func (o *AddDeviceParams) WithUserAgent(userAgent *string) *AddDeviceParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add device params
func (o *AddDeviceParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithAddFromWizard adds the addFromWizard to the add device params
func (o *AddDeviceParams) WithAddFromWizard(addFromWizard *bool) *AddDeviceParams {
	o.SetAddFromWizard(addFromWizard)
	return o
}

// SetAddFromWizard adds the addFromWizard to the add device params
func (o *AddDeviceParams) SetAddFromWizard(addFromWizard *bool) {
	o.AddFromWizard = addFromWizard
}

// WithBody adds the body to the add device params
func (o *AddDeviceParams) WithBody(body *models.Device) *AddDeviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device params
func (o *AddDeviceParams) SetBody(body *models.Device) {
	o.Body = body
}

// WithEnd adds the end to the add device params
func (o *AddDeviceParams) WithEnd(end *int64) *AddDeviceParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the add device params
func (o *AddDeviceParams) SetEnd(end *int64) {
	o.End = end
}

// WithNetflowFilter adds the netflowFilter to the add device params
func (o *AddDeviceParams) WithNetflowFilter(netflowFilter *string) *AddDeviceParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the add device params
func (o *AddDeviceParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithStart adds the start to the add device params
func (o *AddDeviceParams) WithStart(start *int64) *AddDeviceParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the add device params
func (o *AddDeviceParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *AddDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AddFromWizard != nil {

		// query param addFromWizard
		var qrAddFromWizard bool

		if o.AddFromWizard != nil {
			qrAddFromWizard = *o.AddFromWizard
		}
		qAddFromWizard := swag.FormatBool(qrAddFromWizard)
		if qAddFromWizard != "" {

			if err := r.SetQueryParam("addFromWizard", qAddFromWizard); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}
	}

	if o.NetflowFilter != nil {

		// query param netflowFilter
		var qrNetflowFilter string

		if o.NetflowFilter != nil {
			qrNetflowFilter = *o.NetflowFilter
		}
		qNetflowFilter := qrNetflowFilter
		if qNetflowFilter != "" {

			if err := r.SetQueryParam("netflowFilter", qNetflowFilter); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
