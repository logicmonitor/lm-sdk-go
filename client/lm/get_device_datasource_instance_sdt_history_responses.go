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

// GetDeviceDatasourceInstanceSDTHistoryReader is a Reader for the GetDeviceDatasourceInstanceSDTHistory structure.
type GetDeviceDatasourceInstanceSDTHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceSDTHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceDatasourceInstanceSDTHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceDatasourceInstanceSDTHistoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceDatasourceInstanceSDTHistoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceSDTHistoryOK creates a GetDeviceDatasourceInstanceSDTHistoryOK with default headers values
func NewGetDeviceDatasourceInstanceSDTHistoryOK() *GetDeviceDatasourceInstanceSDTHistoryOK {
	return &GetDeviceDatasourceInstanceSDTHistoryOK{}
}

/* GetDeviceDatasourceInstanceSDTHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceSDTHistoryOK struct {
	Payload *models.DeviceGroupSDTHistoryPaginationResponse
}

func (o *GetDeviceDatasourceInstanceSDTHistoryOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/historysdts][%d] getDeviceDatasourceInstanceSdtHistoryOK  %+v", 200, o.Payload)
}
func (o *GetDeviceDatasourceInstanceSDTHistoryOK) GetPayload() *models.DeviceGroupSDTHistoryPaginationResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceSDTHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroupSDTHistoryPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceSDTHistoryTooManyRequests creates a GetDeviceDatasourceInstanceSDTHistoryTooManyRequests with default headers values
func NewGetDeviceDatasourceInstanceSDTHistoryTooManyRequests() *GetDeviceDatasourceInstanceSDTHistoryTooManyRequests {
	return &GetDeviceDatasourceInstanceSDTHistoryTooManyRequests{}
}

/* GetDeviceDatasourceInstanceSDTHistoryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceDatasourceInstanceSDTHistoryTooManyRequests struct {

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

func (o *GetDeviceDatasourceInstanceSDTHistoryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/historysdts][%d] getDeviceDatasourceInstanceSdtHistoryTooManyRequests ", 429)
}

func (o *GetDeviceDatasourceInstanceSDTHistoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceDatasourceInstanceSDTHistoryDefault creates a GetDeviceDatasourceInstanceSDTHistoryDefault with default headers values
func NewGetDeviceDatasourceInstanceSDTHistoryDefault(code int) *GetDeviceDatasourceInstanceSDTHistoryDefault {
	return &GetDeviceDatasourceInstanceSDTHistoryDefault{
		_statusCode: code,
	}
}

/* GetDeviceDatasourceInstanceSDTHistoryDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceDatasourceInstanceSDTHistoryDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance SDT history default response
func (o *GetDeviceDatasourceInstanceSDTHistoryDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceSDTHistoryDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/historysdts][%d] getDeviceDatasourceInstanceSDTHistory default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceDatasourceInstanceSDTHistoryDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceSDTHistoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}