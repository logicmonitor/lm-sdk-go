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

// AddAPITokenByAdminIDReader is a Reader for the AddAPITokenByAdminID structure.
type AddAPITokenByAdminIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAPITokenByAdminIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAPITokenByAdminIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddAPITokenByAdminIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddAPITokenByAdminIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddAPITokenByAdminIDOK creates a AddAPITokenByAdminIDOK with default headers values
func NewAddAPITokenByAdminIDOK() *AddAPITokenByAdminIDOK {
	return &AddAPITokenByAdminIDOK{}
}

/* AddAPITokenByAdminIDOK describes a response with status code 200, with default header values.

successful operation
*/
type AddAPITokenByAdminIDOK struct {
	Payload *models.APIToken
}

func (o *AddAPITokenByAdminIDOK) Error() string {
	return fmt.Sprintf("[POST /setting/admins/{adminId}/apitokens][%d] addApiTokenByAdminIdOK  %+v", 200, o.Payload)
}
func (o *AddAPITokenByAdminIDOK) GetPayload() *models.APIToken {
	return o.Payload
}

func (o *AddAPITokenByAdminIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAPITokenByAdminIDTooManyRequests creates a AddAPITokenByAdminIDTooManyRequests with default headers values
func NewAddAPITokenByAdminIDTooManyRequests() *AddAPITokenByAdminIDTooManyRequests {
	return &AddAPITokenByAdminIDTooManyRequests{}
}

/* AddAPITokenByAdminIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddAPITokenByAdminIDTooManyRequests struct {

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

func (o *AddAPITokenByAdminIDTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /setting/admins/{adminId}/apitokens][%d] addApiTokenByAdminIdTooManyRequests ", 429)
}

func (o *AddAPITokenByAdminIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddAPITokenByAdminIDDefault creates a AddAPITokenByAdminIDDefault with default headers values
func NewAddAPITokenByAdminIDDefault(code int) *AddAPITokenByAdminIDDefault {
	return &AddAPITokenByAdminIDDefault{
		_statusCode: code,
	}
}

/* AddAPITokenByAdminIDDefault describes a response with status code -1, with default header values.

Error
*/
type AddAPITokenByAdminIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add Api token by admin Id default response
func (o *AddAPITokenByAdminIDDefault) Code() int {
	return o._statusCode
}

func (o *AddAPITokenByAdminIDDefault) Error() string {
	return fmt.Sprintf("[POST /setting/admins/{adminId}/apitokens][%d] addApiTokenByAdminId default  %+v", o._statusCode, o.Payload)
}
func (o *AddAPITokenByAdminIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddAPITokenByAdminIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}