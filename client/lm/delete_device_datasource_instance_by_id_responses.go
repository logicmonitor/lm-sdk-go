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

// DeleteDeviceDatasourceInstanceByIDReader is a Reader for the DeleteDeviceDatasourceInstanceByID structure.
type DeleteDeviceDatasourceInstanceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceDatasourceInstanceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeviceDatasourceInstanceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewDeleteDeviceDatasourceInstanceByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDeviceDatasourceInstanceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDeviceDatasourceInstanceByIDOK creates a DeleteDeviceDatasourceInstanceByIDOK with default headers values
func NewDeleteDeviceDatasourceInstanceByIDOK() *DeleteDeviceDatasourceInstanceByIDOK {
	return &DeleteDeviceDatasourceInstanceByIDOK{}
}

/* DeleteDeviceDatasourceInstanceByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteDeviceDatasourceInstanceByIDOK struct {
	Payload interface{}
}

func (o *DeleteDeviceDatasourceInstanceByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}][%d] deleteDeviceDatasourceInstanceByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteDeviceDatasourceInstanceByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDeviceDatasourceInstanceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceDatasourceInstanceByIDTooManyRequests creates a DeleteDeviceDatasourceInstanceByIDTooManyRequests with default headers values
func NewDeleteDeviceDatasourceInstanceByIDTooManyRequests() *DeleteDeviceDatasourceInstanceByIDTooManyRequests {
	return &DeleteDeviceDatasourceInstanceByIDTooManyRequests{}
}

/* DeleteDeviceDatasourceInstanceByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteDeviceDatasourceInstanceByIDTooManyRequests struct {

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

func (o *DeleteDeviceDatasourceInstanceByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}][%d] deleteDeviceDatasourceInstanceByIdTooManyRequests ", 429)
}

func (o *DeleteDeviceDatasourceInstanceByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteDeviceDatasourceInstanceByIDDefault creates a DeleteDeviceDatasourceInstanceByIDDefault with default headers values
func NewDeleteDeviceDatasourceInstanceByIDDefault(code int) *DeleteDeviceDatasourceInstanceByIDDefault {
	return &DeleteDeviceDatasourceInstanceByIDDefault{
		_statusCode: code,
	}
}

/* DeleteDeviceDatasourceInstanceByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteDeviceDatasourceInstanceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete device datasource instance by Id default response
func (o *DeleteDeviceDatasourceInstanceByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDeviceDatasourceInstanceByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}][%d] deleteDeviceDatasourceInstanceById default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteDeviceDatasourceInstanceByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDeviceDatasourceInstanceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}