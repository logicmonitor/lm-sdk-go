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

// AddDashboardGroupReader is a Reader for the AddDashboardGroup structure.
type AddDashboardGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDashboardGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDashboardGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddDashboardGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddDashboardGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDashboardGroupOK creates a AddDashboardGroupOK with default headers values
func NewAddDashboardGroupOK() *AddDashboardGroupOK {
	return &AddDashboardGroupOK{}
}

/* AddDashboardGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type AddDashboardGroupOK struct {
	Payload *models.DashboardGroup
}

func (o *AddDashboardGroupOK) Error() string {
	return fmt.Sprintf("[POST /dashboard/groups][%d] addDashboardGroupOK  %+v", 200, o.Payload)
}
func (o *AddDashboardGroupOK) GetPayload() *models.DashboardGroup {
	return o.Payload
}

func (o *AddDashboardGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDashboardGroupTooManyRequests creates a AddDashboardGroupTooManyRequests with default headers values
func NewAddDashboardGroupTooManyRequests() *AddDashboardGroupTooManyRequests {
	return &AddDashboardGroupTooManyRequests{}
}

/* AddDashboardGroupTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddDashboardGroupTooManyRequests struct {

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

func (o *AddDashboardGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /dashboard/groups][%d] addDashboardGroupTooManyRequests ", 429)
}

func (o *AddDashboardGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddDashboardGroupDefault creates a AddDashboardGroupDefault with default headers values
func NewAddDashboardGroupDefault(code int) *AddDashboardGroupDefault {
	return &AddDashboardGroupDefault{
		_statusCode: code,
	}
}

/* AddDashboardGroupDefault describes a response with status code -1, with default header values.

Error
*/
type AddDashboardGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add dashboard group default response
func (o *AddDashboardGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddDashboardGroupDefault) Error() string {
	return fmt.Sprintf("[POST /dashboard/groups][%d] addDashboardGroup default  %+v", o._statusCode, o.Payload)
}
func (o *AddDashboardGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddDashboardGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}