// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// GetRoleListReader is a Reader for the GetRoleList structure.
type GetRoleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetRoleListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRoleListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoleListOK creates a GetRoleListOK with default headers values
func NewGetRoleListOK() *GetRoleListOK {
	return &GetRoleListOK{}
}

/*
GetRoleListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoleListOK struct {
	Payload *models.RolePaginationResponse
}

// IsSuccess returns true when this get role list o k response has a 2xx status code
func (o *GetRoleListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get role list o k response has a 3xx status code
func (o *GetRoleListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role list o k response has a 4xx status code
func (o *GetRoleListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role list o k response has a 5xx status code
func (o *GetRoleListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get role list o k response a status code equal to that given
func (o *GetRoleListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoleListOK) Error() string {
	return fmt.Sprintf("[GET /setting/roles][%d] getRoleListOK  %+v", 200, o.Payload)
}

func (o *GetRoleListOK) String() string {
	return fmt.Sprintf("[GET /setting/roles][%d] getRoleListOK  %+v", 200, o.Payload)
}

func (o *GetRoleListOK) GetPayload() *models.RolePaginationResponse {
	return o.Payload
}

func (o *GetRoleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RolePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleListTooManyRequests creates a GetRoleListTooManyRequests with default headers values
func NewGetRoleListTooManyRequests() *GetRoleListTooManyRequests {
	return &GetRoleListTooManyRequests{}
}

/*
GetRoleListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRoleListTooManyRequests struct {

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

// IsSuccess returns true when this get role list too many requests response has a 2xx status code
func (o *GetRoleListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role list too many requests response has a 3xx status code
func (o *GetRoleListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role list too many requests response has a 4xx status code
func (o *GetRoleListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role list too many requests response has a 5xx status code
func (o *GetRoleListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get role list too many requests response a status code equal to that given
func (o *GetRoleListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoleListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/roles][%d] getRoleListTooManyRequests ", 429)
}

func (o *GetRoleListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/roles][%d] getRoleListTooManyRequests ", 429)
}

func (o *GetRoleListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRoleListDefault creates a GetRoleListDefault with default headers values
func NewGetRoleListDefault(code int) *GetRoleListDefault {
	return &GetRoleListDefault{
		_statusCode: code,
	}
}

/*
GetRoleListDefault describes a response with status code -1, with default header values.

Error
*/
type GetRoleListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get role list default response
func (o *GetRoleListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get role list default response has a 2xx status code
func (o *GetRoleListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get role list default response has a 3xx status code
func (o *GetRoleListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get role list default response has a 4xx status code
func (o *GetRoleListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get role list default response has a 5xx status code
func (o *GetRoleListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get role list default response a status code equal to that given
func (o *GetRoleListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetRoleListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/roles][%d] getRoleList default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoleListDefault) String() string {
	return fmt.Sprintf("[GET /setting/roles][%d] getRoleList default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoleListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRoleListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
