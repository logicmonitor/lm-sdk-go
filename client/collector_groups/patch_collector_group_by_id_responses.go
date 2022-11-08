// Code generated by go-swagger; DO NOT EDIT.

package collector_groups

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

// PatchCollectorGroupByIDReader is a Reader for the PatchCollectorGroupByID structure.
type PatchCollectorGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCollectorGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCollectorGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchCollectorGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchCollectorGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCollectorGroupByIDOK creates a PatchCollectorGroupByIDOK with default headers values
func NewPatchCollectorGroupByIDOK() *PatchCollectorGroupByIDOK {
	return &PatchCollectorGroupByIDOK{}
}

/*
PatchCollectorGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchCollectorGroupByIDOK struct {
	Payload *models.CollectorGroup
}

// IsSuccess returns true when this patch collector group by Id o k response has a 2xx status code
func (o *PatchCollectorGroupByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch collector group by Id o k response has a 3xx status code
func (o *PatchCollectorGroupByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch collector group by Id o k response has a 4xx status code
func (o *PatchCollectorGroupByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch collector group by Id o k response has a 5xx status code
func (o *PatchCollectorGroupByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch collector group by Id o k response a status code equal to that given
func (o *PatchCollectorGroupByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchCollectorGroupByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/collector/groups/{id}][%d] patchCollectorGroupByIdOK  %+v", 200, o.Payload)
}

func (o *PatchCollectorGroupByIDOK) String() string {
	return fmt.Sprintf("[PATCH /setting/collector/groups/{id}][%d] patchCollectorGroupByIdOK  %+v", 200, o.Payload)
}

func (o *PatchCollectorGroupByIDOK) GetPayload() *models.CollectorGroup {
	return o.Payload
}

func (o *PatchCollectorGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCollectorGroupByIDTooManyRequests creates a PatchCollectorGroupByIDTooManyRequests with default headers values
func NewPatchCollectorGroupByIDTooManyRequests() *PatchCollectorGroupByIDTooManyRequests {
	return &PatchCollectorGroupByIDTooManyRequests{}
}

/*
PatchCollectorGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchCollectorGroupByIDTooManyRequests struct {

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

// IsSuccess returns true when this patch collector group by Id too many requests response has a 2xx status code
func (o *PatchCollectorGroupByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch collector group by Id too many requests response has a 3xx status code
func (o *PatchCollectorGroupByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch collector group by Id too many requests response has a 4xx status code
func (o *PatchCollectorGroupByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch collector group by Id too many requests response has a 5xx status code
func (o *PatchCollectorGroupByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch collector group by Id too many requests response a status code equal to that given
func (o *PatchCollectorGroupByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchCollectorGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /setting/collector/groups/{id}][%d] patchCollectorGroupByIdTooManyRequests ", 429)
}

func (o *PatchCollectorGroupByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /setting/collector/groups/{id}][%d] patchCollectorGroupByIdTooManyRequests ", 429)
}

func (o *PatchCollectorGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchCollectorGroupByIDDefault creates a PatchCollectorGroupByIDDefault with default headers values
func NewPatchCollectorGroupByIDDefault(code int) *PatchCollectorGroupByIDDefault {
	return &PatchCollectorGroupByIDDefault{
		_statusCode: code,
	}
}

/*
PatchCollectorGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchCollectorGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch collector group by Id default response
func (o *PatchCollectorGroupByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch collector group by Id default response has a 2xx status code
func (o *PatchCollectorGroupByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch collector group by Id default response has a 3xx status code
func (o *PatchCollectorGroupByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch collector group by Id default response has a 4xx status code
func (o *PatchCollectorGroupByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch collector group by Id default response has a 5xx status code
func (o *PatchCollectorGroupByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch collector group by Id default response a status code equal to that given
func (o *PatchCollectorGroupByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchCollectorGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/collector/groups/{id}][%d] patchCollectorGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCollectorGroupByIDDefault) String() string {
	return fmt.Sprintf("[PATCH /setting/collector/groups/{id}][%d] patchCollectorGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCollectorGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchCollectorGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
