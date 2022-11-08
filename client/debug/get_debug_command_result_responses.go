// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// GetDebugCommandResultReader is a Reader for the GetDebugCommandResult structure.
type GetDebugCommandResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDebugCommandResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDebugCommandResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDebugCommandResultTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDebugCommandResultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDebugCommandResultOK creates a GetDebugCommandResultOK with default headers values
func NewGetDebugCommandResultOK() *GetDebugCommandResultOK {
	return &GetDebugCommandResultOK{}
}

/*
GetDebugCommandResultOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDebugCommandResultOK struct {
	Payload *models.Debug
}

// IsSuccess returns true when this get debug command result o k response has a 2xx status code
func (o *GetDebugCommandResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get debug command result o k response has a 3xx status code
func (o *GetDebugCommandResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get debug command result o k response has a 4xx status code
func (o *GetDebugCommandResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get debug command result o k response has a 5xx status code
func (o *GetDebugCommandResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get debug command result o k response a status code equal to that given
func (o *GetDebugCommandResultOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDebugCommandResultOK) Error() string {
	return fmt.Sprintf("[GET /debug/{id}][%d] getDebugCommandResultOK  %+v", 200, o.Payload)
}

func (o *GetDebugCommandResultOK) String() string {
	return fmt.Sprintf("[GET /debug/{id}][%d] getDebugCommandResultOK  %+v", 200, o.Payload)
}

func (o *GetDebugCommandResultOK) GetPayload() *models.Debug {
	return o.Payload
}

func (o *GetDebugCommandResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Debug)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDebugCommandResultTooManyRequests creates a GetDebugCommandResultTooManyRequests with default headers values
func NewGetDebugCommandResultTooManyRequests() *GetDebugCommandResultTooManyRequests {
	return &GetDebugCommandResultTooManyRequests{}
}

/*
GetDebugCommandResultTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDebugCommandResultTooManyRequests struct {

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

// IsSuccess returns true when this get debug command result too many requests response has a 2xx status code
func (o *GetDebugCommandResultTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get debug command result too many requests response has a 3xx status code
func (o *GetDebugCommandResultTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get debug command result too many requests response has a 4xx status code
func (o *GetDebugCommandResultTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get debug command result too many requests response has a 5xx status code
func (o *GetDebugCommandResultTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get debug command result too many requests response a status code equal to that given
func (o *GetDebugCommandResultTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDebugCommandResultTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /debug/{id}][%d] getDebugCommandResultTooManyRequests ", 429)
}

func (o *GetDebugCommandResultTooManyRequests) String() string {
	return fmt.Sprintf("[GET /debug/{id}][%d] getDebugCommandResultTooManyRequests ", 429)
}

func (o *GetDebugCommandResultTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDebugCommandResultDefault creates a GetDebugCommandResultDefault with default headers values
func NewGetDebugCommandResultDefault(code int) *GetDebugCommandResultDefault {
	return &GetDebugCommandResultDefault{
		_statusCode: code,
	}
}

/*
GetDebugCommandResultDefault describes a response with status code -1, with default header values.

Error
*/
type GetDebugCommandResultDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get debug command result default response
func (o *GetDebugCommandResultDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get debug command result default response has a 2xx status code
func (o *GetDebugCommandResultDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get debug command result default response has a 3xx status code
func (o *GetDebugCommandResultDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get debug command result default response has a 4xx status code
func (o *GetDebugCommandResultDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get debug command result default response has a 5xx status code
func (o *GetDebugCommandResultDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get debug command result default response a status code equal to that given
func (o *GetDebugCommandResultDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetDebugCommandResultDefault) Error() string {
	return fmt.Sprintf("[GET /debug/{id}][%d] getDebugCommandResult default  %+v", o._statusCode, o.Payload)
}

func (o *GetDebugCommandResultDefault) String() string {
	return fmt.Sprintf("[GET /debug/{id}][%d] getDebugCommandResult default  %+v", o._statusCode, o.Payload)
}

func (o *GetDebugCommandResultDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDebugCommandResultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
