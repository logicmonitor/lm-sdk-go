// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// AckAlertByIDReader is a Reader for the AckAlertByID structure.
type AckAlertByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AckAlertByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAckAlertByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAckAlertByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAckAlertByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAckAlertByIDOK creates a AckAlertByIDOK with default headers values
func NewAckAlertByIDOK() *AckAlertByIDOK {
	return &AckAlertByIDOK{}
}

/*
AckAlertByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type AckAlertByIDOK struct {
	Payload interface{}
}

// IsSuccess returns true when this ack alert by Id o k response has a 2xx status code
func (o *AckAlertByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ack alert by Id o k response has a 3xx status code
func (o *AckAlertByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ack alert by Id o k response has a 4xx status code
func (o *AckAlertByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ack alert by Id o k response has a 5xx status code
func (o *AckAlertByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ack alert by Id o k response a status code equal to that given
func (o *AckAlertByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *AckAlertByIDOK) Error() string {
	return fmt.Sprintf("[POST /alert/alerts/{id}/ack][%d] ackAlertByIdOK  %+v", 200, o.Payload)
}

func (o *AckAlertByIDOK) String() string {
	return fmt.Sprintf("[POST /alert/alerts/{id}/ack][%d] ackAlertByIdOK  %+v", 200, o.Payload)
}

func (o *AckAlertByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AckAlertByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAckAlertByIDTooManyRequests creates a AckAlertByIDTooManyRequests with default headers values
func NewAckAlertByIDTooManyRequests() *AckAlertByIDTooManyRequests {
	return &AckAlertByIDTooManyRequests{}
}

/*
AckAlertByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AckAlertByIDTooManyRequests struct {

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

// IsSuccess returns true when this ack alert by Id too many requests response has a 2xx status code
func (o *AckAlertByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ack alert by Id too many requests response has a 3xx status code
func (o *AckAlertByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ack alert by Id too many requests response has a 4xx status code
func (o *AckAlertByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this ack alert by Id too many requests response has a 5xx status code
func (o *AckAlertByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this ack alert by Id too many requests response a status code equal to that given
func (o *AckAlertByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AckAlertByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /alert/alerts/{id}/ack][%d] ackAlertByIdTooManyRequests ", 429)
}

func (o *AckAlertByIDTooManyRequests) String() string {
	return fmt.Sprintf("[POST /alert/alerts/{id}/ack][%d] ackAlertByIdTooManyRequests ", 429)
}

func (o *AckAlertByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAckAlertByIDDefault creates a AckAlertByIDDefault with default headers values
func NewAckAlertByIDDefault(code int) *AckAlertByIDDefault {
	return &AckAlertByIDDefault{
		_statusCode: code,
	}
}

/*
AckAlertByIDDefault describes a response with status code -1, with default header values.

Error
*/
type AckAlertByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ack alert by Id default response
func (o *AckAlertByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ack alert by Id default response has a 2xx status code
func (o *AckAlertByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ack alert by Id default response has a 3xx status code
func (o *AckAlertByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ack alert by Id default response has a 4xx status code
func (o *AckAlertByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ack alert by Id default response has a 5xx status code
func (o *AckAlertByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ack alert by Id default response a status code equal to that given
func (o *AckAlertByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AckAlertByIDDefault) Error() string {
	return fmt.Sprintf("[POST /alert/alerts/{id}/ack][%d] ackAlertById default  %+v", o._statusCode, o.Payload)
}

func (o *AckAlertByIDDefault) String() string {
	return fmt.Sprintf("[POST /alert/alerts/{id}/ack][%d] ackAlertById default  %+v", o._statusCode, o.Payload)
}

func (o *AckAlertByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AckAlertByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
