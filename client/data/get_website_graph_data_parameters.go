// Code generated by go-swagger; DO NOT EDIT.

package data

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

// NewGetWebsiteGraphDataParams creates a new GetWebsiteGraphDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsiteGraphDataParams() *GetWebsiteGraphDataParams {
	return &GetWebsiteGraphDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsiteGraphDataParamsWithTimeout creates a new GetWebsiteGraphDataParams object
// with the ability to set a timeout on a request.
func NewGetWebsiteGraphDataParamsWithTimeout(timeout time.Duration) *GetWebsiteGraphDataParams {
	return &GetWebsiteGraphDataParams{
		timeout: timeout,
	}
}

// NewGetWebsiteGraphDataParamsWithContext creates a new GetWebsiteGraphDataParams object
// with the ability to set a context for a request.
func NewGetWebsiteGraphDataParamsWithContext(ctx context.Context) *GetWebsiteGraphDataParams {
	return &GetWebsiteGraphDataParams{
		Context: ctx,
	}
}

// NewGetWebsiteGraphDataParamsWithHTTPClient creates a new GetWebsiteGraphDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsiteGraphDataParamsWithHTTPClient(client *http.Client) *GetWebsiteGraphDataParams {
	return &GetWebsiteGraphDataParams{
		HTTPClient: client,
	}
}

/*
GetWebsiteGraphDataParams contains all the parameters to send to the API endpoint

	for the get website graph data operation.

	Typically these are written to a http.Request.
*/
type GetWebsiteGraphDataParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty"
	UserAgent *string

	// CheckpointID.
	//
	// Format: int32
	CheckpointID int32

	// End.
	//
	// Format: int64
	End *int64

	// Format.
	Format *string

	// GraphName.
	GraphName string

	// Start.
	//
	// Format: int64
	Start *int64

	// WebsiteID.
	//
	// Format: int32
	WebsiteID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get website graph data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteGraphDataParams) WithDefaults() *GetWebsiteGraphDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get website graph data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteGraphDataParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1-dirty")
	)

	val := GetWebsiteGraphDataParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithTimeout(timeout time.Duration) *GetWebsiteGraphDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithContext(ctx context.Context) *GetWebsiteGraphDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithHTTPClient(client *http.Client) *GetWebsiteGraphDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithUserAgent(userAgent *string) *GetWebsiteGraphDataParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithCheckpointID adds the checkpointID to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithCheckpointID(checkpointID int32) *GetWebsiteGraphDataParams {
	o.SetCheckpointID(checkpointID)
	return o
}

// SetCheckpointID adds the checkpointId to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetCheckpointID(checkpointID int32) {
	o.CheckpointID = checkpointID
}

// WithEnd adds the end to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithEnd(end *int64) *GetWebsiteGraphDataParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetEnd(end *int64) {
	o.End = end
}

// WithFormat adds the format to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithFormat(format *string) *GetWebsiteGraphDataParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetFormat(format *string) {
	o.Format = format
}

// WithGraphName adds the graphName to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithGraphName(graphName string) *GetWebsiteGraphDataParams {
	o.SetGraphName(graphName)
	return o
}

// SetGraphName adds the graphName to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetGraphName(graphName string) {
	o.GraphName = graphName
}

// WithStart adds the start to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithStart(start *int64) *GetWebsiteGraphDataParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetStart(start *int64) {
	o.Start = start
}

// WithWebsiteID adds the websiteID to the get website graph data params
func (o *GetWebsiteGraphDataParams) WithWebsiteID(websiteID int32) *GetWebsiteGraphDataParams {
	o.SetWebsiteID(websiteID)
	return o
}

// SetWebsiteID adds the websiteId to the get website graph data params
func (o *GetWebsiteGraphDataParams) SetWebsiteID(websiteID int32) {
	o.WebsiteID = websiteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteGraphDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param checkpointId
	if err := r.SetPathParam("checkpointId", swag.FormatInt32(o.CheckpointID)); err != nil {
		return err
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

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	// path param graphName
	if err := r.SetPathParam("graphName", o.GraphName); err != nil {
		return err
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

	// path param websiteId
	if err := r.SetPathParam("websiteId", swag.FormatInt32(o.WebsiteID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
