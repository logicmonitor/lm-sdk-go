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

// PatchWidgetByIDReader is a Reader for the PatchWidgetByID structure.
type PatchWidgetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWidgetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWidgetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchWidgetByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchWidgetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchWidgetByIDOK creates a PatchWidgetByIDOK with default headers values
func NewPatchWidgetByIDOK() *PatchWidgetByIDOK {
	return &PatchWidgetByIDOK{}
}

/* PatchWidgetByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchWidgetByIDOK struct {
	Payload models.Widget
}

func (o *PatchWidgetByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetByIdOK  %+v", 200, o.Payload)
}
func (o *PatchWidgetByIDOK) GetPayload() models.Widget {
	return o.Payload
}

func (o *PatchWidgetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalWidget(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchWidgetByIDTooManyRequests creates a PatchWidgetByIDTooManyRequests with default headers values
func NewPatchWidgetByIDTooManyRequests() *PatchWidgetByIDTooManyRequests {
	return &PatchWidgetByIDTooManyRequests{}
}

/* PatchWidgetByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchWidgetByIDTooManyRequests struct {

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

func (o *PatchWidgetByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetByIdTooManyRequests ", 429)
}

func (o *PatchWidgetByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchWidgetByIDDefault creates a PatchWidgetByIDDefault with default headers values
func NewPatchWidgetByIDDefault(code int) *PatchWidgetByIDDefault {
	return &PatchWidgetByIDDefault{
		_statusCode: code,
	}
}

/* PatchWidgetByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchWidgetByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch widget by Id default response
func (o *PatchWidgetByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchWidgetByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /dashboard/widgets/{id}][%d] patchWidgetById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchWidgetByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchWidgetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
