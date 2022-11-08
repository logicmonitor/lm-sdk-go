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

// AddDeviceGroupPropertyReader is a Reader for the AddDeviceGroupProperty structure.
type AddDeviceGroupPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDeviceGroupPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDeviceGroupPropertyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddDeviceGroupPropertyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddDeviceGroupPropertyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDeviceGroupPropertyOK creates a AddDeviceGroupPropertyOK with default headers values
func NewAddDeviceGroupPropertyOK() *AddDeviceGroupPropertyOK {
	return &AddDeviceGroupPropertyOK{}
}

/*
AddDeviceGroupPropertyOK describes a response with status code 200, with default header values.

successful operation
*/
type AddDeviceGroupPropertyOK struct {
	Payload *models.EntityProperty
}

// IsSuccess returns true when this add device group property o k response has a 2xx status code
func (o *AddDeviceGroupPropertyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add device group property o k response has a 3xx status code
func (o *AddDeviceGroupPropertyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add device group property o k response has a 4xx status code
func (o *AddDeviceGroupPropertyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add device group property o k response has a 5xx status code
func (o *AddDeviceGroupPropertyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add device group property o k response a status code equal to that given
func (o *AddDeviceGroupPropertyOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddDeviceGroupPropertyOK) Error() string {
	return fmt.Sprintf("[POST /device/groups/{gid}/properties][%d] addDeviceGroupPropertyOK  %+v", 200, o.Payload)
}

func (o *AddDeviceGroupPropertyOK) String() string {
	return fmt.Sprintf("[POST /device/groups/{gid}/properties][%d] addDeviceGroupPropertyOK  %+v", 200, o.Payload)
}

func (o *AddDeviceGroupPropertyOK) GetPayload() *models.EntityProperty {
	return o.Payload
}

func (o *AddDeviceGroupPropertyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityProperty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeviceGroupPropertyTooManyRequests creates a AddDeviceGroupPropertyTooManyRequests with default headers values
func NewAddDeviceGroupPropertyTooManyRequests() *AddDeviceGroupPropertyTooManyRequests {
	return &AddDeviceGroupPropertyTooManyRequests{}
}

/*
AddDeviceGroupPropertyTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddDeviceGroupPropertyTooManyRequests struct {

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

// IsSuccess returns true when this add device group property too many requests response has a 2xx status code
func (o *AddDeviceGroupPropertyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add device group property too many requests response has a 3xx status code
func (o *AddDeviceGroupPropertyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add device group property too many requests response has a 4xx status code
func (o *AddDeviceGroupPropertyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add device group property too many requests response has a 5xx status code
func (o *AddDeviceGroupPropertyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add device group property too many requests response a status code equal to that given
func (o *AddDeviceGroupPropertyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AddDeviceGroupPropertyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /device/groups/{gid}/properties][%d] addDeviceGroupPropertyTooManyRequests ", 429)
}

func (o *AddDeviceGroupPropertyTooManyRequests) String() string {
	return fmt.Sprintf("[POST /device/groups/{gid}/properties][%d] addDeviceGroupPropertyTooManyRequests ", 429)
}

func (o *AddDeviceGroupPropertyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddDeviceGroupPropertyDefault creates a AddDeviceGroupPropertyDefault with default headers values
func NewAddDeviceGroupPropertyDefault(code int) *AddDeviceGroupPropertyDefault {
	return &AddDeviceGroupPropertyDefault{
		_statusCode: code,
	}
}

/*
AddDeviceGroupPropertyDefault describes a response with status code -1, with default header values.

Error
*/
type AddDeviceGroupPropertyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add device group property default response
func (o *AddDeviceGroupPropertyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add device group property default response has a 2xx status code
func (o *AddDeviceGroupPropertyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add device group property default response has a 3xx status code
func (o *AddDeviceGroupPropertyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add device group property default response has a 4xx status code
func (o *AddDeviceGroupPropertyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add device group property default response has a 5xx status code
func (o *AddDeviceGroupPropertyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add device group property default response a status code equal to that given
func (o *AddDeviceGroupPropertyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddDeviceGroupPropertyDefault) Error() string {
	return fmt.Sprintf("[POST /device/groups/{gid}/properties][%d] addDeviceGroupProperty default  %+v", o._statusCode, o.Payload)
}

func (o *AddDeviceGroupPropertyDefault) String() string {
	return fmt.Sprintf("[POST /device/groups/{gid}/properties][%d] addDeviceGroupProperty default  %+v", o._statusCode, o.Payload)
}

func (o *AddDeviceGroupPropertyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddDeviceGroupPropertyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
