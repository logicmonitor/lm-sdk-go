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

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// AddPropertysourceAuditVersionReader is a Reader for the AddPropertysourceAuditVersion structure.
type AddPropertysourceAuditVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPropertysourceAuditVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPropertysourceAuditVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddPropertysourceAuditVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddPropertysourceAuditVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPropertysourceAuditVersionOK creates a AddPropertysourceAuditVersionOK with default headers values
func NewAddPropertysourceAuditVersionOK() *AddPropertysourceAuditVersionOK {
	return &AddPropertysourceAuditVersionOK{}
}

/* AddPropertysourceAuditVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type AddPropertysourceAuditVersionOK struct {
	Payload *models.PropertyRule
}

func (o *AddPropertysourceAuditVersionOK) Error() string {
	return fmt.Sprintf("[POST /setting/propertyrules/{id}/audit][%d] addPropertysourceAuditVersionOK  %+v", 200, o.Payload)
}
func (o *AddPropertysourceAuditVersionOK) GetPayload() *models.PropertyRule {
	return o.Payload
}

func (o *AddPropertysourceAuditVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPropertysourceAuditVersionTooManyRequests creates a AddPropertysourceAuditVersionTooManyRequests with default headers values
func NewAddPropertysourceAuditVersionTooManyRequests() *AddPropertysourceAuditVersionTooManyRequests {
	return &AddPropertysourceAuditVersionTooManyRequests{}
}

/* AddPropertysourceAuditVersionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddPropertysourceAuditVersionTooManyRequests struct {

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

func (o *AddPropertysourceAuditVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /setting/propertyrules/{id}/audit][%d] addPropertysourceAuditVersionTooManyRequests ", 429)
}

func (o *AddPropertysourceAuditVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddPropertysourceAuditVersionDefault creates a AddPropertysourceAuditVersionDefault with default headers values
func NewAddPropertysourceAuditVersionDefault(code int) *AddPropertysourceAuditVersionDefault {
	return &AddPropertysourceAuditVersionDefault{
		_statusCode: code,
	}
}

/* AddPropertysourceAuditVersionDefault describes a response with status code -1, with default header values.

Error
*/
type AddPropertysourceAuditVersionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add propertysource audit version default response
func (o *AddPropertysourceAuditVersionDefault) Code() int {
	return o._statusCode
}

func (o *AddPropertysourceAuditVersionDefault) Error() string {
	return fmt.Sprintf("[POST /setting/propertyrules/{id}/audit][%d] addPropertysourceAuditVersion default  %+v", o._statusCode, o.Payload)
}
func (o *AddPropertysourceAuditVersionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddPropertysourceAuditVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
