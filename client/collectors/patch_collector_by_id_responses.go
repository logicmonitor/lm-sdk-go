// Code generated by go-swagger; DO NOT EDIT.

package collectors

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

// PatchCollectorByIDReader is a Reader for the PatchCollectorByID structure.
type PatchCollectorByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCollectorByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCollectorByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchCollectorByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchCollectorByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCollectorByIDOK creates a PatchCollectorByIDOK with default headers values
func NewPatchCollectorByIDOK() *PatchCollectorByIDOK {
	return &PatchCollectorByIDOK{}
}

/*
PatchCollectorByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchCollectorByIDOK struct {
	Payload *models.Collector
}

// IsSuccess returns true when this patch collector by Id o k response has a 2xx status code
func (o *PatchCollectorByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch collector by Id o k response has a 3xx status code
func (o *PatchCollectorByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch collector by Id o k response has a 4xx status code
func (o *PatchCollectorByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch collector by Id o k response has a 5xx status code
func (o *PatchCollectorByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch collector by Id o k response a status code equal to that given
func (o *PatchCollectorByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchCollectorByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/collector/collectors/{id}][%d] patchCollectorByIdOK  %+v", 200, o.Payload)
}

func (o *PatchCollectorByIDOK) String() string {
	return fmt.Sprintf("[PATCH /setting/collector/collectors/{id}][%d] patchCollectorByIdOK  %+v", 200, o.Payload)
}

func (o *PatchCollectorByIDOK) GetPayload() *models.Collector {
	return o.Payload
}

func (o *PatchCollectorByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCollectorByIDTooManyRequests creates a PatchCollectorByIDTooManyRequests with default headers values
func NewPatchCollectorByIDTooManyRequests() *PatchCollectorByIDTooManyRequests {
	return &PatchCollectorByIDTooManyRequests{}
}

/*
PatchCollectorByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchCollectorByIDTooManyRequests struct {

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

// IsSuccess returns true when this patch collector by Id too many requests response has a 2xx status code
func (o *PatchCollectorByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch collector by Id too many requests response has a 3xx status code
func (o *PatchCollectorByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch collector by Id too many requests response has a 4xx status code
func (o *PatchCollectorByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch collector by Id too many requests response has a 5xx status code
func (o *PatchCollectorByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch collector by Id too many requests response a status code equal to that given
func (o *PatchCollectorByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchCollectorByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /setting/collector/collectors/{id}][%d] patchCollectorByIdTooManyRequests ", 429)
}

func (o *PatchCollectorByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /setting/collector/collectors/{id}][%d] patchCollectorByIdTooManyRequests ", 429)
}

func (o *PatchCollectorByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchCollectorByIDDefault creates a PatchCollectorByIDDefault with default headers values
func NewPatchCollectorByIDDefault(code int) *PatchCollectorByIDDefault {
	return &PatchCollectorByIDDefault{
		_statusCode: code,
	}
}

/*
PatchCollectorByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchCollectorByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch collector by Id default response
func (o *PatchCollectorByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch collector by Id default response has a 2xx status code
func (o *PatchCollectorByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch collector by Id default response has a 3xx status code
func (o *PatchCollectorByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch collector by Id default response has a 4xx status code
func (o *PatchCollectorByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch collector by Id default response has a 5xx status code
func (o *PatchCollectorByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch collector by Id default response a status code equal to that given
func (o *PatchCollectorByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchCollectorByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/collector/collectors/{id}][%d] patchCollectorById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCollectorByIDDefault) String() string {
	return fmt.Sprintf("[PATCH /setting/collector/collectors/{id}][%d] patchCollectorById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCollectorByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchCollectorByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
