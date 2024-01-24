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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewFetchDeviceInstancesDataParams creates a new FetchDeviceInstancesDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFetchDeviceInstancesDataParams() *FetchDeviceInstancesDataParams {
	return &FetchDeviceInstancesDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFetchDeviceInstancesDataParamsWithTimeout creates a new FetchDeviceInstancesDataParams object
// with the ability to set a timeout on a request.
func NewFetchDeviceInstancesDataParamsWithTimeout(timeout time.Duration) *FetchDeviceInstancesDataParams {
	return &FetchDeviceInstancesDataParams{
		timeout: timeout,
	}
}

// NewFetchDeviceInstancesDataParamsWithContext creates a new FetchDeviceInstancesDataParams object
// with the ability to set a context for a request.
func NewFetchDeviceInstancesDataParamsWithContext(ctx context.Context) *FetchDeviceInstancesDataParams {
	return &FetchDeviceInstancesDataParams{
		Context: ctx,
	}
}

// NewFetchDeviceInstancesDataParamsWithHTTPClient creates a new FetchDeviceInstancesDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewFetchDeviceInstancesDataParamsWithHTTPClient(client *http.Client) *FetchDeviceInstancesDataParams {
	return &FetchDeviceInstancesDataParams{
		HTTPClient: client,
	}
}

/* FetchDeviceInstancesDataParams contains all the parameters to send to the API endpoint
   for the fetch device instances data operation.

   Typically these are written to a http.Request.
*/
type FetchDeviceInstancesDataParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	/* Aggregate.

	   the aggregate option

	   Default: "none"
	*/
	Aggregate *string

	// Body.
	Body *models.DeviceInstances

	// End.
	//
	// Format: int64
	End *int64

	// Period.
	//
	// Format: double
	// Default: 1
	Period *float64

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fetch device instances data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchDeviceInstancesDataParams) WithDefaults() *FetchDeviceInstancesDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fetch device instances data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchDeviceInstancesDataParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		aggregateDefault = string("none")

		endDefault = int64(0)

		periodDefault = float64(1)

		startDefault = int64(0)
	)

	val := FetchDeviceInstancesDataParams{
		UserAgent: &userAgentDefault,
		Aggregate: &aggregateDefault,
		End:       &endDefault,
		Period:    &periodDefault,
		Start:     &startDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithTimeout(timeout time.Duration) *FetchDeviceInstancesDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithContext(ctx context.Context) *FetchDeviceInstancesDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithHTTPClient(client *http.Client) *FetchDeviceInstancesDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithUserAgent(userAgent *string) *FetchDeviceInstancesDataParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithAggregate adds the aggregate to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithAggregate(aggregate *string) *FetchDeviceInstancesDataParams {
	o.SetAggregate(aggregate)
	return o
}

// SetAggregate adds the aggregate to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetAggregate(aggregate *string) {
	o.Aggregate = aggregate
}

// WithBody adds the body to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithBody(body *models.DeviceInstances) *FetchDeviceInstancesDataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetBody(body *models.DeviceInstances) {
	o.Body = body
}

// WithEnd adds the end to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithEnd(end *int64) *FetchDeviceInstancesDataParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetEnd(end *int64) {
	o.End = end
}

// WithPeriod adds the period to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithPeriod(period *float64) *FetchDeviceInstancesDataParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetPeriod(period *float64) {
	o.Period = period
}

// WithStart adds the start to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) WithStart(start *int64) *FetchDeviceInstancesDataParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the fetch device instances data params
func (o *FetchDeviceInstancesDataParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *FetchDeviceInstancesDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Aggregate != nil {

		// query param aggregate
		var qrAggregate string

		if o.Aggregate != nil {
			qrAggregate = *o.Aggregate
		}
		qAggregate := qrAggregate
		if qAggregate != "" {

			if err := r.SetQueryParam("aggregate", qAggregate); err != nil {
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

	if o.Period != nil {

		// query param period
		var qrPeriod float64

		if o.Period != nil {
			qrPeriod = *o.Period
		}
		qPeriod := swag.FormatFloat64(qrPeriod)
		if qPeriod != "" {

			if err := r.SetQueryParam("period", qPeriod); err != nil {
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
