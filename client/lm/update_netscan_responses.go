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

// UpdateNetscanReader is a Reader for the UpdateNetscan structure.
type UpdateNetscanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetscanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetscanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateNetscanTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateNetscanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNetscanOK creates a UpdateNetscanOK with default headers values
func NewUpdateNetscanOK() *UpdateNetscanOK {
	return &UpdateNetscanOK{}
}

/* UpdateNetscanOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateNetscanOK struct {
	Payload models.Netscan
}

func (o *UpdateNetscanOK) Error() string {
	return fmt.Sprintf("[PUT /setting/netscans/{id}][%d] updateNetscanOK  %+v", 200, o.Payload)
}
func (o *UpdateNetscanOK) GetPayload() models.Netscan {
	return o.Payload
}

func (o *UpdateNetscanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalNetscan(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateNetscanTooManyRequests creates a UpdateNetscanTooManyRequests with default headers values
func NewUpdateNetscanTooManyRequests() *UpdateNetscanTooManyRequests {
	return &UpdateNetscanTooManyRequests{}
}

/* UpdateNetscanTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateNetscanTooManyRequests struct {

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

func (o *UpdateNetscanTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /setting/netscans/{id}][%d] updateNetscanTooManyRequests ", 429)
}

func (o *UpdateNetscanTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateNetscanDefault creates a UpdateNetscanDefault with default headers values
func NewUpdateNetscanDefault(code int) *UpdateNetscanDefault {
	return &UpdateNetscanDefault{
		_statusCode: code,
	}
}

/* UpdateNetscanDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateNetscanDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update netscan default response
func (o *UpdateNetscanDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNetscanDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/netscans/{id}][%d] updateNetscan default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateNetscanDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateNetscanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}