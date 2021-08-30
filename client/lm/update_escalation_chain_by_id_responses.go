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

// UpdateEscalationChainByIDReader is a Reader for the UpdateEscalationChainByID structure.
type UpdateEscalationChainByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEscalationChainByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEscalationChainByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateEscalationChainByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateEscalationChainByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateEscalationChainByIDOK creates a UpdateEscalationChainByIDOK with default headers values
func NewUpdateEscalationChainByIDOK() *UpdateEscalationChainByIDOK {
	return &UpdateEscalationChainByIDOK{}
}

/* UpdateEscalationChainByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateEscalationChainByIDOK struct {
	Payload *models.EscalatingChain
}

func (o *UpdateEscalationChainByIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/alert/chains/{id}][%d] updateEscalationChainByIdOK  %+v", 200, o.Payload)
}
func (o *UpdateEscalationChainByIDOK) GetPayload() *models.EscalatingChain {
	return o.Payload
}

func (o *UpdateEscalationChainByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EscalatingChain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEscalationChainByIDTooManyRequests creates a UpdateEscalationChainByIDTooManyRequests with default headers values
func NewUpdateEscalationChainByIDTooManyRequests() *UpdateEscalationChainByIDTooManyRequests {
	return &UpdateEscalationChainByIDTooManyRequests{}
}

/* UpdateEscalationChainByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateEscalationChainByIDTooManyRequests struct {

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

func (o *UpdateEscalationChainByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /setting/alert/chains/{id}][%d] updateEscalationChainByIdTooManyRequests ", 429)
}

func (o *UpdateEscalationChainByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateEscalationChainByIDDefault creates a UpdateEscalationChainByIDDefault with default headers values
func NewUpdateEscalationChainByIDDefault(code int) *UpdateEscalationChainByIDDefault {
	return &UpdateEscalationChainByIDDefault{
		_statusCode: code,
	}
}

/* UpdateEscalationChainByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateEscalationChainByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update escalation chain by Id default response
func (o *UpdateEscalationChainByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateEscalationChainByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/alert/chains/{id}][%d] updateEscalationChainById default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateEscalationChainByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateEscalationChainByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}