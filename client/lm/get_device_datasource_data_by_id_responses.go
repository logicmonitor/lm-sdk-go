// Code generated by go-swagger; DO NOT EDIT.

package lm

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

// GetDeviceDatasourceDataByIDReader is a Reader for the GetDeviceDatasourceDataByID structure.
type GetDeviceDatasourceDataByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceDataByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceDatasourceDataByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceDatasourceDataByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceDatasourceDataByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceDataByIDOK creates a GetDeviceDatasourceDataByIDOK with default headers values
func NewGetDeviceDatasourceDataByIDOK() *GetDeviceDatasourceDataByIDOK {
	return &GetDeviceDatasourceDataByIDOK{}
}

/* GetDeviceDatasourceDataByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceDatasourceDataByIDOK struct {
	Payload *models.DeviceDataSourceData
}

func (o *GetDeviceDatasourceDataByIDOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{id}/data][%d] getDeviceDatasourceDataByIdOK  %+v", 200, o.Payload)
}
func (o *GetDeviceDatasourceDataByIDOK) GetPayload() *models.DeviceDataSourceData {
	return o.Payload
}

func (o *GetDeviceDatasourceDataByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceDataByIDTooManyRequests creates a GetDeviceDatasourceDataByIDTooManyRequests with default headers values
func NewGetDeviceDatasourceDataByIDTooManyRequests() *GetDeviceDatasourceDataByIDTooManyRequests {
	return &GetDeviceDatasourceDataByIDTooManyRequests{}
}

/* GetDeviceDatasourceDataByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceDatasourceDataByIDTooManyRequests struct {

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

func (o *GetDeviceDatasourceDataByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{id}/data][%d] getDeviceDatasourceDataByIdTooManyRequests ", 429)
}

func (o *GetDeviceDatasourceDataByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceDatasourceDataByIDDefault creates a GetDeviceDatasourceDataByIDDefault with default headers values
func NewGetDeviceDatasourceDataByIDDefault(code int) *GetDeviceDatasourceDataByIDDefault {
	return &GetDeviceDatasourceDataByIDDefault{
		_statusCode: code,
	}
}

/* GetDeviceDatasourceDataByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceDatasourceDataByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource data by Id default response
func (o *GetDeviceDatasourceDataByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceDataByIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{id}/data][%d] getDeviceDatasourceDataById default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceDatasourceDataByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceDataByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}