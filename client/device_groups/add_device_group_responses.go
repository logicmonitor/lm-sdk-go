// Code generated by go-swagger; DO NOT EDIT.

package device_groups

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

// AddDeviceGroupReader is a Reader for the AddDeviceGroup structure.
type AddDeviceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDeviceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDeviceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddDeviceGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddDeviceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDeviceGroupOK creates a AddDeviceGroupOK with default headers values
func NewAddDeviceGroupOK() *AddDeviceGroupOK {
	return &AddDeviceGroupOK{}
}

/*
AddDeviceGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type AddDeviceGroupOK struct {
	Payload *models.DeviceGroup
}

// IsSuccess returns true when this add device group o k response has a 2xx status code
func (o *AddDeviceGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add device group o k response has a 3xx status code
func (o *AddDeviceGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add device group o k response has a 4xx status code
func (o *AddDeviceGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add device group o k response has a 5xx status code
func (o *AddDeviceGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add device group o k response a status code equal to that given
func (o *AddDeviceGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddDeviceGroupOK) Error() string {
	return fmt.Sprintf("[POST /device/groups][%d] addDeviceGroupOK  %+v", 200, o.Payload)
}

func (o *AddDeviceGroupOK) String() string {
	return fmt.Sprintf("[POST /device/groups][%d] addDeviceGroupOK  %+v", 200, o.Payload)
}

func (o *AddDeviceGroupOK) GetPayload() *models.DeviceGroup {
	return o.Payload
}

func (o *AddDeviceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeviceGroupTooManyRequests creates a AddDeviceGroupTooManyRequests with default headers values
func NewAddDeviceGroupTooManyRequests() *AddDeviceGroupTooManyRequests {
	return &AddDeviceGroupTooManyRequests{}
}

/*
AddDeviceGroupTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddDeviceGroupTooManyRequests struct {

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

// IsSuccess returns true when this add device group too many requests response has a 2xx status code
func (o *AddDeviceGroupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add device group too many requests response has a 3xx status code
func (o *AddDeviceGroupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add device group too many requests response has a 4xx status code
func (o *AddDeviceGroupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add device group too many requests response has a 5xx status code
func (o *AddDeviceGroupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add device group too many requests response a status code equal to that given
func (o *AddDeviceGroupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AddDeviceGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /device/groups][%d] addDeviceGroupTooManyRequests ", 429)
}

func (o *AddDeviceGroupTooManyRequests) String() string {
	return fmt.Sprintf("[POST /device/groups][%d] addDeviceGroupTooManyRequests ", 429)
}

func (o *AddDeviceGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddDeviceGroupDefault creates a AddDeviceGroupDefault with default headers values
func NewAddDeviceGroupDefault(code int) *AddDeviceGroupDefault {
	return &AddDeviceGroupDefault{
		_statusCode: code,
	}
}

/*
AddDeviceGroupDefault describes a response with status code -1, with default header values.

Error
*/
type AddDeviceGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add device group default response
func (o *AddDeviceGroupDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add device group default response has a 2xx status code
func (o *AddDeviceGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add device group default response has a 3xx status code
func (o *AddDeviceGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add device group default response has a 4xx status code
func (o *AddDeviceGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add device group default response has a 5xx status code
func (o *AddDeviceGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add device group default response a status code equal to that given
func (o *AddDeviceGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddDeviceGroupDefault) Error() string {
	return fmt.Sprintf("[POST /device/groups][%d] addDeviceGroup default  %+v", o._statusCode, o.Payload)
}

func (o *AddDeviceGroupDefault) String() string {
	return fmt.Sprintf("[POST /device/groups][%d] addDeviceGroup default  %+v", o._statusCode, o.Payload)
}

func (o *AddDeviceGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddDeviceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
