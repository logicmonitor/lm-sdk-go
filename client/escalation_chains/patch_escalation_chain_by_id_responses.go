// Code generated by go-swagger; DO NOT EDIT.

package escalation_chains

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

// PatchEscalationChainByIDReader is a Reader for the PatchEscalationChainByID structure.
type PatchEscalationChainByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEscalationChainByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEscalationChainByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchEscalationChainByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEscalationChainByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEscalationChainByIDOK creates a PatchEscalationChainByIDOK with default headers values
func NewPatchEscalationChainByIDOK() *PatchEscalationChainByIDOK {
	return &PatchEscalationChainByIDOK{}
}

/*
PatchEscalationChainByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchEscalationChainByIDOK struct {
	Payload *models.EscalatingChain
}

// IsSuccess returns true when this patch escalation chain by Id o k response has a 2xx status code
func (o *PatchEscalationChainByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch escalation chain by Id o k response has a 3xx status code
func (o *PatchEscalationChainByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch escalation chain by Id o k response has a 4xx status code
func (o *PatchEscalationChainByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch escalation chain by Id o k response has a 5xx status code
func (o *PatchEscalationChainByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch escalation chain by Id o k response a status code equal to that given
func (o *PatchEscalationChainByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchEscalationChainByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/alert/chains/{id}][%d] patchEscalationChainByIdOK  %+v", 200, o.Payload)
}

func (o *PatchEscalationChainByIDOK) String() string {
	return fmt.Sprintf("[PATCH /setting/alert/chains/{id}][%d] patchEscalationChainByIdOK  %+v", 200, o.Payload)
}

func (o *PatchEscalationChainByIDOK) GetPayload() *models.EscalatingChain {
	return o.Payload
}

func (o *PatchEscalationChainByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EscalatingChain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEscalationChainByIDTooManyRequests creates a PatchEscalationChainByIDTooManyRequests with default headers values
func NewPatchEscalationChainByIDTooManyRequests() *PatchEscalationChainByIDTooManyRequests {
	return &PatchEscalationChainByIDTooManyRequests{}
}

/*
PatchEscalationChainByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchEscalationChainByIDTooManyRequests struct {

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

// IsSuccess returns true when this patch escalation chain by Id too many requests response has a 2xx status code
func (o *PatchEscalationChainByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch escalation chain by Id too many requests response has a 3xx status code
func (o *PatchEscalationChainByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch escalation chain by Id too many requests response has a 4xx status code
func (o *PatchEscalationChainByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch escalation chain by Id too many requests response has a 5xx status code
func (o *PatchEscalationChainByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch escalation chain by Id too many requests response a status code equal to that given
func (o *PatchEscalationChainByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchEscalationChainByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /setting/alert/chains/{id}][%d] patchEscalationChainByIdTooManyRequests ", 429)
}

func (o *PatchEscalationChainByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /setting/alert/chains/{id}][%d] patchEscalationChainByIdTooManyRequests ", 429)
}

func (o *PatchEscalationChainByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEscalationChainByIDDefault creates a PatchEscalationChainByIDDefault with default headers values
func NewPatchEscalationChainByIDDefault(code int) *PatchEscalationChainByIDDefault {
	return &PatchEscalationChainByIDDefault{
		_statusCode: code,
	}
}

/*
PatchEscalationChainByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchEscalationChainByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch escalation chain by Id default response
func (o *PatchEscalationChainByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch escalation chain by Id default response has a 2xx status code
func (o *PatchEscalationChainByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch escalation chain by Id default response has a 3xx status code
func (o *PatchEscalationChainByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch escalation chain by Id default response has a 4xx status code
func (o *PatchEscalationChainByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch escalation chain by Id default response has a 5xx status code
func (o *PatchEscalationChainByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch escalation chain by Id default response a status code equal to that given
func (o *PatchEscalationChainByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchEscalationChainByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/alert/chains/{id}][%d] patchEscalationChainById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEscalationChainByIDDefault) String() string {
	return fmt.Sprintf("[PATCH /setting/alert/chains/{id}][%d] patchEscalationChainById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEscalationChainByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchEscalationChainByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
