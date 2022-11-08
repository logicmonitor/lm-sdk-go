// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// GetDeviceDatasourceInstanceGraphDataReader is a Reader for the GetDeviceDatasourceInstanceGraphData structure.
type GetDeviceDatasourceInstanceGraphDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceGraphDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceDatasourceInstanceGraphDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceDatasourceInstanceGraphDataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceDatasourceInstanceGraphDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceGraphDataOK creates a GetDeviceDatasourceInstanceGraphDataOK with default headers values
func NewGetDeviceDatasourceInstanceGraphDataOK() *GetDeviceDatasourceInstanceGraphDataOK {
	return &GetDeviceDatasourceInstanceGraphDataOK{}
}

/*
GetDeviceDatasourceInstanceGraphDataOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceGraphDataOK struct {
	Payload *models.GraphPlot
}

// IsSuccess returns true when this get device datasource instance graph data o k response has a 2xx status code
func (o *GetDeviceDatasourceInstanceGraphDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device datasource instance graph data o k response has a 3xx status code
func (o *GetDeviceDatasourceInstanceGraphDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device datasource instance graph data o k response has a 4xx status code
func (o *GetDeviceDatasourceInstanceGraphDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device datasource instance graph data o k response has a 5xx status code
func (o *GetDeviceDatasourceInstanceGraphDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device datasource instance graph data o k response a status code equal to that given
func (o *GetDeviceDatasourceInstanceGraphDataOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceDatasourceInstanceGraphDataOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/graphs/{graphId}/data][%d] getDeviceDatasourceInstanceGraphDataOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceInstanceGraphDataOK) String() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/graphs/{graphId}/data][%d] getDeviceDatasourceInstanceGraphDataOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceInstanceGraphDataOK) GetPayload() *models.GraphPlot {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceGraphDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphPlot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceGraphDataTooManyRequests creates a GetDeviceDatasourceInstanceGraphDataTooManyRequests with default headers values
func NewGetDeviceDatasourceInstanceGraphDataTooManyRequests() *GetDeviceDatasourceInstanceGraphDataTooManyRequests {
	return &GetDeviceDatasourceInstanceGraphDataTooManyRequests{}
}

/*
GetDeviceDatasourceInstanceGraphDataTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceDatasourceInstanceGraphDataTooManyRequests struct {

	/* Request limit per X-Rate-Limit-Window
	 */
	XRateLimitLimit int64

	/* The number of requests left for the time window
	 */
	XRateLimitRemaining int64

	/* The rolling time window length with the unit of second
	 */
	XRateLimitWindow int64
}

// IsSuccess returns true when this get device datasource instance graph data too many requests response has a 2xx status code
func (o *GetDeviceDatasourceInstanceGraphDataTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device datasource instance graph data too many requests response has a 3xx status code
func (o *GetDeviceDatasourceInstanceGraphDataTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device datasource instance graph data too many requests response has a 4xx status code
func (o *GetDeviceDatasourceInstanceGraphDataTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device datasource instance graph data too many requests response has a 5xx status code
func (o *GetDeviceDatasourceInstanceGraphDataTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get device datasource instance graph data too many requests response a status code equal to that given
func (o *GetDeviceDatasourceInstanceGraphDataTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDeviceDatasourceInstanceGraphDataTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/graphs/{graphId}/data][%d] getDeviceDatasourceInstanceGraphDataTooManyRequests ", 429)
}

func (o *GetDeviceDatasourceInstanceGraphDataTooManyRequests) String() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/graphs/{graphId}/data][%d] getDeviceDatasourceInstanceGraphDataTooManyRequests ", 429)
}

func (o *GetDeviceDatasourceInstanceGraphDataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-rate-limit-limit
	hdrXRateLimitLimit := response.GetHeader("x-rate-limit-limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("x-rate-limit-limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header x-rate-limit-remaining
	hdrXRateLimitRemaining := response.GetHeader("x-rate-limit-remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("x-rate-limit-remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header x-rate-limit-window
	hdrXRateLimitWindow := response.GetHeader("x-rate-limit-window")

	if hdrXRateLimitWindow != "" {
		valxRateLimitWindow, err := swag.ConvertInt64(hdrXRateLimitWindow)
		if err != nil {
			return errors.InvalidType("x-rate-limit-window", "header", "int64", hdrXRateLimitWindow)
		}
		o.XRateLimitWindow = valxRateLimitWindow
	}

	return nil
}

// NewGetDeviceDatasourceInstanceGraphDataDefault creates a GetDeviceDatasourceInstanceGraphDataDefault with default headers values
func NewGetDeviceDatasourceInstanceGraphDataDefault(code int) *GetDeviceDatasourceInstanceGraphDataDefault {
	return &GetDeviceDatasourceInstanceGraphDataDefault{
		_statusCode: code,
	}
}

/*
GetDeviceDatasourceInstanceGraphDataDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceDatasourceInstanceGraphDataDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance graph data default response
func (o *GetDeviceDatasourceInstanceGraphDataDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get device datasource instance graph data default response has a 2xx status code
func (o *GetDeviceDatasourceInstanceGraphDataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get device datasource instance graph data default response has a 3xx status code
func (o *GetDeviceDatasourceInstanceGraphDataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get device datasource instance graph data default response has a 4xx status code
func (o *GetDeviceDatasourceInstanceGraphDataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get device datasource instance graph data default response has a 5xx status code
func (o *GetDeviceDatasourceInstanceGraphDataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get device datasource instance graph data default response a status code equal to that given
func (o *GetDeviceDatasourceInstanceGraphDataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetDeviceDatasourceInstanceGraphDataDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/graphs/{graphId}/data][%d] getDeviceDatasourceInstanceGraphData default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceInstanceGraphDataDefault) String() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/graphs/{graphId}/data][%d] getDeviceDatasourceInstanceGraphData default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceInstanceGraphDataDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceGraphDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
