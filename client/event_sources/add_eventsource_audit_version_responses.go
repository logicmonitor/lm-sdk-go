// Code generated by go-swagger; DO NOT EDIT.

package event_sources

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

// AddEventsourceAuditVersionReader is a Reader for the AddEventsourceAuditVersion structure.
type AddEventsourceAuditVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddEventsourceAuditVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddEventsourceAuditVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddEventsourceAuditVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddEventsourceAuditVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddEventsourceAuditVersionOK creates a AddEventsourceAuditVersionOK with default headers values
func NewAddEventsourceAuditVersionOK() *AddEventsourceAuditVersionOK {
	return &AddEventsourceAuditVersionOK{}
}

/*
AddEventsourceAuditVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type AddEventsourceAuditVersionOK struct {
	Payload models.EventSource
}

// IsSuccess returns true when this add eventsource audit version o k response has a 2xx status code
func (o *AddEventsourceAuditVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add eventsource audit version o k response has a 3xx status code
func (o *AddEventsourceAuditVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add eventsource audit version o k response has a 4xx status code
func (o *AddEventsourceAuditVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add eventsource audit version o k response has a 5xx status code
func (o *AddEventsourceAuditVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add eventsource audit version o k response a status code equal to that given
func (o *AddEventsourceAuditVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddEventsourceAuditVersionOK) Error() string {
	return fmt.Sprintf("[POST /setting/eventsources/{id}/audit][%d] addEventsourceAuditVersionOK  %+v", 200, o.Payload)
}

func (o *AddEventsourceAuditVersionOK) String() string {
	return fmt.Sprintf("[POST /setting/eventsources/{id}/audit][%d] addEventsourceAuditVersionOK  %+v", 200, o.Payload)
}

func (o *AddEventsourceAuditVersionOK) GetPayload() models.EventSource {
	return o.Payload
}

func (o *AddEventsourceAuditVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalEventSource(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewAddEventsourceAuditVersionTooManyRequests creates a AddEventsourceAuditVersionTooManyRequests with default headers values
func NewAddEventsourceAuditVersionTooManyRequests() *AddEventsourceAuditVersionTooManyRequests {
	return &AddEventsourceAuditVersionTooManyRequests{}
}

/*
AddEventsourceAuditVersionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddEventsourceAuditVersionTooManyRequests struct {

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

// IsSuccess returns true when this add eventsource audit version too many requests response has a 2xx status code
func (o *AddEventsourceAuditVersionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add eventsource audit version too many requests response has a 3xx status code
func (o *AddEventsourceAuditVersionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add eventsource audit version too many requests response has a 4xx status code
func (o *AddEventsourceAuditVersionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add eventsource audit version too many requests response has a 5xx status code
func (o *AddEventsourceAuditVersionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add eventsource audit version too many requests response a status code equal to that given
func (o *AddEventsourceAuditVersionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AddEventsourceAuditVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /setting/eventsources/{id}/audit][%d] addEventsourceAuditVersionTooManyRequests ", 429)
}

func (o *AddEventsourceAuditVersionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /setting/eventsources/{id}/audit][%d] addEventsourceAuditVersionTooManyRequests ", 429)
}

func (o *AddEventsourceAuditVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddEventsourceAuditVersionDefault creates a AddEventsourceAuditVersionDefault with default headers values
func NewAddEventsourceAuditVersionDefault(code int) *AddEventsourceAuditVersionDefault {
	return &AddEventsourceAuditVersionDefault{
		_statusCode: code,
	}
}

/*
AddEventsourceAuditVersionDefault describes a response with status code -1, with default header values.

Error
*/
type AddEventsourceAuditVersionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add eventsource audit version default response
func (o *AddEventsourceAuditVersionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add eventsource audit version default response has a 2xx status code
func (o *AddEventsourceAuditVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add eventsource audit version default response has a 3xx status code
func (o *AddEventsourceAuditVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add eventsource audit version default response has a 4xx status code
func (o *AddEventsourceAuditVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add eventsource audit version default response has a 5xx status code
func (o *AddEventsourceAuditVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add eventsource audit version default response a status code equal to that given
func (o *AddEventsourceAuditVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddEventsourceAuditVersionDefault) Error() string {
	return fmt.Sprintf("[POST /setting/eventsources/{id}/audit][%d] addEventsourceAuditVersion default  %+v", o._statusCode, o.Payload)
}

func (o *AddEventsourceAuditVersionDefault) String() string {
	return fmt.Sprintf("[POST /setting/eventsources/{id}/audit][%d] addEventsourceAuditVersion default  %+v", o._statusCode, o.Payload)
}

func (o *AddEventsourceAuditVersionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddEventsourceAuditVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
