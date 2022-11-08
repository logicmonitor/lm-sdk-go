// Code generated by go-swagger; DO NOT EDIT.

package ops_notes

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

// AddOpsNoteReader is a Reader for the AddOpsNote structure.
type AddOpsNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOpsNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddOpsNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddOpsNoteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddOpsNoteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOpsNoteOK creates a AddOpsNoteOK with default headers values
func NewAddOpsNoteOK() *AddOpsNoteOK {
	return &AddOpsNoteOK{}
}

/*
AddOpsNoteOK describes a response with status code 200, with default header values.

successful operation
*/
type AddOpsNoteOK struct {
	Payload *models.OpsNote
}

// IsSuccess returns true when this add ops note o k response has a 2xx status code
func (o *AddOpsNoteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add ops note o k response has a 3xx status code
func (o *AddOpsNoteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add ops note o k response has a 4xx status code
func (o *AddOpsNoteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add ops note o k response has a 5xx status code
func (o *AddOpsNoteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add ops note o k response a status code equal to that given
func (o *AddOpsNoteOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddOpsNoteOK) Error() string {
	return fmt.Sprintf("[POST /setting/opsnotes][%d] addOpsNoteOK  %+v", 200, o.Payload)
}

func (o *AddOpsNoteOK) String() string {
	return fmt.Sprintf("[POST /setting/opsnotes][%d] addOpsNoteOK  %+v", 200, o.Payload)
}

func (o *AddOpsNoteOK) GetPayload() *models.OpsNote {
	return o.Payload
}

func (o *AddOpsNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpsNote)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOpsNoteTooManyRequests creates a AddOpsNoteTooManyRequests with default headers values
func NewAddOpsNoteTooManyRequests() *AddOpsNoteTooManyRequests {
	return &AddOpsNoteTooManyRequests{}
}

/*
AddOpsNoteTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddOpsNoteTooManyRequests struct {

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

// IsSuccess returns true when this add ops note too many requests response has a 2xx status code
func (o *AddOpsNoteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add ops note too many requests response has a 3xx status code
func (o *AddOpsNoteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add ops note too many requests response has a 4xx status code
func (o *AddOpsNoteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add ops note too many requests response has a 5xx status code
func (o *AddOpsNoteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add ops note too many requests response a status code equal to that given
func (o *AddOpsNoteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AddOpsNoteTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /setting/opsnotes][%d] addOpsNoteTooManyRequests ", 429)
}

func (o *AddOpsNoteTooManyRequests) String() string {
	return fmt.Sprintf("[POST /setting/opsnotes][%d] addOpsNoteTooManyRequests ", 429)
}

func (o *AddOpsNoteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddOpsNoteDefault creates a AddOpsNoteDefault with default headers values
func NewAddOpsNoteDefault(code int) *AddOpsNoteDefault {
	return &AddOpsNoteDefault{
		_statusCode: code,
	}
}

/*
AddOpsNoteDefault describes a response with status code -1, with default header values.

Error
*/
type AddOpsNoteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add ops note default response
func (o *AddOpsNoteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add ops note default response has a 2xx status code
func (o *AddOpsNoteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add ops note default response has a 3xx status code
func (o *AddOpsNoteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add ops note default response has a 4xx status code
func (o *AddOpsNoteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add ops note default response has a 5xx status code
func (o *AddOpsNoteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add ops note default response a status code equal to that given
func (o *AddOpsNoteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddOpsNoteDefault) Error() string {
	return fmt.Sprintf("[POST /setting/opsnotes][%d] addOpsNote default  %+v", o._statusCode, o.Payload)
}

func (o *AddOpsNoteDefault) String() string {
	return fmt.Sprintf("[POST /setting/opsnotes][%d] addOpsNote default  %+v", o._statusCode, o.Payload)
}

func (o *AddOpsNoteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddOpsNoteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
