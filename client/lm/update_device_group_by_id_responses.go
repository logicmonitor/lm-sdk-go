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

// UpdateDeviceGroupByIDReader is a Reader for the UpdateDeviceGroupByID structure.
type UpdateDeviceGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateDeviceGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDeviceGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDeviceGroupByIDOK creates a UpdateDeviceGroupByIDOK with default headers values
func NewUpdateDeviceGroupByIDOK() *UpdateDeviceGroupByIDOK {
	return &UpdateDeviceGroupByIDOK{}
}

/* UpdateDeviceGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDeviceGroupByIDOK struct {
	Payload *models.DeviceGroup
}

func (o *UpdateDeviceGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{id}][%d] updateDeviceGroupByIdOK  %+v", 200, o.Payload)
}
func (o *UpdateDeviceGroupByIDOK) GetPayload() *models.DeviceGroup {
	return o.Payload
}

func (o *UpdateDeviceGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceGroupByIDTooManyRequests creates a UpdateDeviceGroupByIDTooManyRequests with default headers values
func NewUpdateDeviceGroupByIDTooManyRequests() *UpdateDeviceGroupByIDTooManyRequests {
	return &UpdateDeviceGroupByIDTooManyRequests{}
}

/* UpdateDeviceGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDeviceGroupByIDTooManyRequests struct {

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

func (o *UpdateDeviceGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{id}][%d] updateDeviceGroupByIdTooManyRequests ", 429)
}

func (o *UpdateDeviceGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDeviceGroupByIDDefault creates a UpdateDeviceGroupByIDDefault with default headers values
func NewUpdateDeviceGroupByIDDefault(code int) *UpdateDeviceGroupByIDDefault {
	return &UpdateDeviceGroupByIDDefault{
		_statusCode: code,
	}
}

/* UpdateDeviceGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDeviceGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device group by Id default response
func (o *UpdateDeviceGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDeviceGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{id}][%d] updateDeviceGroupById default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDeviceGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}