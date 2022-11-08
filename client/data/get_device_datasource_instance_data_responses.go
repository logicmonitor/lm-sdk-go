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

// GetDeviceDatasourceInstanceDataReader is a Reader for the GetDeviceDatasourceInstanceData structure.
type GetDeviceDatasourceInstanceDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceDatasourceInstanceDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceDatasourceInstanceDataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceDatasourceInstanceDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceDataOK creates a GetDeviceDatasourceInstanceDataOK with default headers values
func NewGetDeviceDatasourceInstanceDataOK() *GetDeviceDatasourceInstanceDataOK {
	return &GetDeviceDatasourceInstanceDataOK{}
}

/*
GetDeviceDatasourceInstanceDataOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceDataOK struct {
	Payload *models.DeviceDataSourceInstanceData
}

// IsSuccess returns true when this get device datasource instance data o k response has a 2xx status code
func (o *GetDeviceDatasourceInstanceDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device datasource instance data o k response has a 3xx status code
func (o *GetDeviceDatasourceInstanceDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device datasource instance data o k response has a 4xx status code
func (o *GetDeviceDatasourceInstanceDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device datasource instance data o k response has a 5xx status code
func (o *GetDeviceDatasourceInstanceDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device datasource instance data o k response a status code equal to that given
func (o *GetDeviceDatasourceInstanceDataOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceDatasourceInstanceDataOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data][%d] getDeviceDatasourceInstanceDataOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceInstanceDataOK) String() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data][%d] getDeviceDatasourceInstanceDataOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceInstanceDataOK) GetPayload() *models.DeviceDataSourceInstanceData {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstanceData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceDataTooManyRequests creates a GetDeviceDatasourceInstanceDataTooManyRequests with default headers values
func NewGetDeviceDatasourceInstanceDataTooManyRequests() *GetDeviceDatasourceInstanceDataTooManyRequests {
	return &GetDeviceDatasourceInstanceDataTooManyRequests{}
}

/*
GetDeviceDatasourceInstanceDataTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceDatasourceInstanceDataTooManyRequests struct {

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

// IsSuccess returns true when this get device datasource instance data too many requests response has a 2xx status code
func (o *GetDeviceDatasourceInstanceDataTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device datasource instance data too many requests response has a 3xx status code
func (o *GetDeviceDatasourceInstanceDataTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device datasource instance data too many requests response has a 4xx status code
func (o *GetDeviceDatasourceInstanceDataTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device datasource instance data too many requests response has a 5xx status code
func (o *GetDeviceDatasourceInstanceDataTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get device datasource instance data too many requests response a status code equal to that given
func (o *GetDeviceDatasourceInstanceDataTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDeviceDatasourceInstanceDataTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data][%d] getDeviceDatasourceInstanceDataTooManyRequests ", 429)
}

func (o *GetDeviceDatasourceInstanceDataTooManyRequests) String() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data][%d] getDeviceDatasourceInstanceDataTooManyRequests ", 429)
}

func (o *GetDeviceDatasourceInstanceDataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceDatasourceInstanceDataDefault creates a GetDeviceDatasourceInstanceDataDefault with default headers values
func NewGetDeviceDatasourceInstanceDataDefault(code int) *GetDeviceDatasourceInstanceDataDefault {
	return &GetDeviceDatasourceInstanceDataDefault{
		_statusCode: code,
	}
}

/*
GetDeviceDatasourceInstanceDataDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceDatasourceInstanceDataDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance data default response
func (o *GetDeviceDatasourceInstanceDataDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get device datasource instance data default response has a 2xx status code
func (o *GetDeviceDatasourceInstanceDataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get device datasource instance data default response has a 3xx status code
func (o *GetDeviceDatasourceInstanceDataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get device datasource instance data default response has a 4xx status code
func (o *GetDeviceDatasourceInstanceDataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get device datasource instance data default response has a 5xx status code
func (o *GetDeviceDatasourceInstanceDataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get device datasource instance data default response a status code equal to that given
func (o *GetDeviceDatasourceInstanceDataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetDeviceDatasourceInstanceDataDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data][%d] getDeviceDatasourceInstanceData default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceInstanceDataDefault) String() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data][%d] getDeviceDatasourceInstanceData default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceInstanceDataDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
