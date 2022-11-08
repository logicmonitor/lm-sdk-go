// Code generated by go-swagger; DO NOT EDIT.

package widgets

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

// UpdateWidgetByIDReader is a Reader for the UpdateWidgetByID structure.
type UpdateWidgetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWidgetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateWidgetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateWidgetByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateWidgetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateWidgetByIDOK creates a UpdateWidgetByIDOK with default headers values
func NewUpdateWidgetByIDOK() *UpdateWidgetByIDOK {
	return &UpdateWidgetByIDOK{}
}

/*
UpdateWidgetByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateWidgetByIDOK struct {
	Payload models.Widget
}

// IsSuccess returns true when this update widget by Id o k response has a 2xx status code
func (o *UpdateWidgetByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update widget by Id o k response has a 3xx status code
func (o *UpdateWidgetByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update widget by Id o k response has a 4xx status code
func (o *UpdateWidgetByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update widget by Id o k response has a 5xx status code
func (o *UpdateWidgetByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update widget by Id o k response a status code equal to that given
func (o *UpdateWidgetByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateWidgetByIDOK) Error() string {
	return fmt.Sprintf("[PUT /dashboard/widgets/{id}][%d] updateWidgetByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateWidgetByIDOK) String() string {
	return fmt.Sprintf("[PUT /dashboard/widgets/{id}][%d] updateWidgetByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateWidgetByIDOK) GetPayload() models.Widget {
	return o.Payload
}

func (o *UpdateWidgetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalWidget(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateWidgetByIDTooManyRequests creates a UpdateWidgetByIDTooManyRequests with default headers values
func NewUpdateWidgetByIDTooManyRequests() *UpdateWidgetByIDTooManyRequests {
	return &UpdateWidgetByIDTooManyRequests{}
}

/*
UpdateWidgetByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateWidgetByIDTooManyRequests struct {

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

// IsSuccess returns true when this update widget by Id too many requests response has a 2xx status code
func (o *UpdateWidgetByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update widget by Id too many requests response has a 3xx status code
func (o *UpdateWidgetByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update widget by Id too many requests response has a 4xx status code
func (o *UpdateWidgetByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update widget by Id too many requests response has a 5xx status code
func (o *UpdateWidgetByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update widget by Id too many requests response a status code equal to that given
func (o *UpdateWidgetByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateWidgetByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /dashboard/widgets/{id}][%d] updateWidgetByIdTooManyRequests ", 429)
}

func (o *UpdateWidgetByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /dashboard/widgets/{id}][%d] updateWidgetByIdTooManyRequests ", 429)
}

func (o *UpdateWidgetByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateWidgetByIDDefault creates a UpdateWidgetByIDDefault with default headers values
func NewUpdateWidgetByIDDefault(code int) *UpdateWidgetByIDDefault {
	return &UpdateWidgetByIDDefault{
		_statusCode: code,
	}
}

/*
UpdateWidgetByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateWidgetByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update widget by Id default response
func (o *UpdateWidgetByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update widget by Id default response has a 2xx status code
func (o *UpdateWidgetByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update widget by Id default response has a 3xx status code
func (o *UpdateWidgetByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update widget by Id default response has a 4xx status code
func (o *UpdateWidgetByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update widget by Id default response has a 5xx status code
func (o *UpdateWidgetByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update widget by Id default response a status code equal to that given
func (o *UpdateWidgetByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateWidgetByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /dashboard/widgets/{id}][%d] updateWidgetById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateWidgetByIDDefault) String() string {
	return fmt.Sprintf("[PUT /dashboard/widgets/{id}][%d] updateWidgetById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateWidgetByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateWidgetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
