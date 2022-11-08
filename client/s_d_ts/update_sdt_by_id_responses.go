// Code generated by go-swagger; DO NOT EDIT.

package s_d_ts

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

// UpdateSDTByIDReader is a Reader for the UpdateSDTByID structure.
type UpdateSDTByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSDTByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSDTByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateSDTByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateSDTByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSDTByIDOK creates a UpdateSDTByIDOK with default headers values
func NewUpdateSDTByIDOK() *UpdateSDTByIDOK {
	return &UpdateSDTByIDOK{}
}

/*
UpdateSDTByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateSDTByIDOK struct {
	Payload models.SDT
}

// IsSuccess returns true when this update Sdt by Id o k response has a 2xx status code
func (o *UpdateSDTByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update Sdt by Id o k response has a 3xx status code
func (o *UpdateSDTByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Sdt by Id o k response has a 4xx status code
func (o *UpdateSDTByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Sdt by Id o k response has a 5xx status code
func (o *UpdateSDTByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update Sdt by Id o k response a status code equal to that given
func (o *UpdateSDTByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateSDTByIDOK) Error() string {
	return fmt.Sprintf("[PUT /sdt/sdts/{id}][%d] updateSdtByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateSDTByIDOK) String() string {
	return fmt.Sprintf("[PUT /sdt/sdts/{id}][%d] updateSdtByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateSDTByIDOK) GetPayload() models.SDT {
	return o.Payload
}

func (o *UpdateSDTByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalSDT(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateSDTByIDTooManyRequests creates a UpdateSDTByIDTooManyRequests with default headers values
func NewUpdateSDTByIDTooManyRequests() *UpdateSDTByIDTooManyRequests {
	return &UpdateSDTByIDTooManyRequests{}
}

/*
UpdateSDTByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateSDTByIDTooManyRequests struct {

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

// IsSuccess returns true when this update Sdt by Id too many requests response has a 2xx status code
func (o *UpdateSDTByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Sdt by Id too many requests response has a 3xx status code
func (o *UpdateSDTByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Sdt by Id too many requests response has a 4xx status code
func (o *UpdateSDTByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Sdt by Id too many requests response has a 5xx status code
func (o *UpdateSDTByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update Sdt by Id too many requests response a status code equal to that given
func (o *UpdateSDTByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateSDTByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /sdt/sdts/{id}][%d] updateSdtByIdTooManyRequests ", 429)
}

func (o *UpdateSDTByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /sdt/sdts/{id}][%d] updateSdtByIdTooManyRequests ", 429)
}

func (o *UpdateSDTByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateSDTByIDDefault creates a UpdateSDTByIDDefault with default headers values
func NewUpdateSDTByIDDefault(code int) *UpdateSDTByIDDefault {
	return &UpdateSDTByIDDefault{
		_statusCode: code,
	}
}

/*
UpdateSDTByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateSDTByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update SDT by Id default response
func (o *UpdateSDTByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update SDT by Id default response has a 2xx status code
func (o *UpdateSDTByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update SDT by Id default response has a 3xx status code
func (o *UpdateSDTByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update SDT by Id default response has a 4xx status code
func (o *UpdateSDTByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update SDT by Id default response has a 5xx status code
func (o *UpdateSDTByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update SDT by Id default response a status code equal to that given
func (o *UpdateSDTByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateSDTByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /sdt/sdts/{id}][%d] updateSDTById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSDTByIDDefault) String() string {
	return fmt.Sprintf("[PUT /sdt/sdts/{id}][%d] updateSDTById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSDTByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateSDTByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
