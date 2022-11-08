// Code generated by go-swagger; DO NOT EDIT.

package users

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

// GetAdminByIDReader is a Reader for the GetAdminByID structure.
type GetAdminByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAdminByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAdminByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdminByIDOK creates a GetAdminByIDOK with default headers values
func NewGetAdminByIDOK() *GetAdminByIDOK {
	return &GetAdminByIDOK{}
}

/*
GetAdminByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAdminByIDOK struct {
	Payload *models.Admin
}

// IsSuccess returns true when this get admin by Id o k response has a 2xx status code
func (o *GetAdminByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get admin by Id o k response has a 3xx status code
func (o *GetAdminByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admin by Id o k response has a 4xx status code
func (o *GetAdminByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get admin by Id o k response has a 5xx status code
func (o *GetAdminByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get admin by Id o k response a status code equal to that given
func (o *GetAdminByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAdminByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/admins/{id}][%d] getAdminByIdOK  %+v", 200, o.Payload)
}

func (o *GetAdminByIDOK) String() string {
	return fmt.Sprintf("[GET /setting/admins/{id}][%d] getAdminByIdOK  %+v", 200, o.Payload)
}

func (o *GetAdminByIDOK) GetPayload() *models.Admin {
	return o.Payload
}

func (o *GetAdminByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Admin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminByIDTooManyRequests creates a GetAdminByIDTooManyRequests with default headers values
func NewGetAdminByIDTooManyRequests() *GetAdminByIDTooManyRequests {
	return &GetAdminByIDTooManyRequests{}
}

/*
GetAdminByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAdminByIDTooManyRequests struct {

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

// IsSuccess returns true when this get admin by Id too many requests response has a 2xx status code
func (o *GetAdminByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admin by Id too many requests response has a 3xx status code
func (o *GetAdminByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admin by Id too many requests response has a 4xx status code
func (o *GetAdminByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get admin by Id too many requests response has a 5xx status code
func (o *GetAdminByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get admin by Id too many requests response a status code equal to that given
func (o *GetAdminByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAdminByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/admins/{id}][%d] getAdminByIdTooManyRequests ", 429)
}

func (o *GetAdminByIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/admins/{id}][%d] getAdminByIdTooManyRequests ", 429)
}

func (o *GetAdminByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAdminByIDDefault creates a GetAdminByIDDefault with default headers values
func NewGetAdminByIDDefault(code int) *GetAdminByIDDefault {
	return &GetAdminByIDDefault{
		_statusCode: code,
	}
}

/*
GetAdminByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAdminByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get admin by Id default response
func (o *GetAdminByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get admin by Id default response has a 2xx status code
func (o *GetAdminByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get admin by Id default response has a 3xx status code
func (o *GetAdminByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get admin by Id default response has a 4xx status code
func (o *GetAdminByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get admin by Id default response has a 5xx status code
func (o *GetAdminByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get admin by Id default response a status code equal to that given
func (o *GetAdminByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAdminByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/admins/{id}][%d] getAdminById default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminByIDDefault) String() string {
	return fmt.Sprintf("[GET /setting/admins/{id}][%d] getAdminById default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAdminByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
