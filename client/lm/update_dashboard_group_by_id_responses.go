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

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// UpdateDashboardGroupByIDReader is a Reader for the UpdateDashboardGroupByID structure.
type UpdateDashboardGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDashboardGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateDashboardGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDashboardGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDashboardGroupByIDOK creates a UpdateDashboardGroupByIDOK with default headers values
func NewUpdateDashboardGroupByIDOK() *UpdateDashboardGroupByIDOK {
	return &UpdateDashboardGroupByIDOK{}
}

/* UpdateDashboardGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDashboardGroupByIDOK struct {
	Payload *models.DashboardGroup
}

func (o *UpdateDashboardGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /dashboard/groups/{id}][%d] updateDashboardGroupByIdOK  %+v", 200, o.Payload)
}
func (o *UpdateDashboardGroupByIDOK) GetPayload() *models.DashboardGroup {
	return o.Payload
}

func (o *UpdateDashboardGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardGroupByIDTooManyRequests creates a UpdateDashboardGroupByIDTooManyRequests with default headers values
func NewUpdateDashboardGroupByIDTooManyRequests() *UpdateDashboardGroupByIDTooManyRequests {
	return &UpdateDashboardGroupByIDTooManyRequests{}
}

/* UpdateDashboardGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDashboardGroupByIDTooManyRequests struct {

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

func (o *UpdateDashboardGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /dashboard/groups/{id}][%d] updateDashboardGroupByIdTooManyRequests ", 429)
}

func (o *UpdateDashboardGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDashboardGroupByIDDefault creates a UpdateDashboardGroupByIDDefault with default headers values
func NewUpdateDashboardGroupByIDDefault(code int) *UpdateDashboardGroupByIDDefault {
	return &UpdateDashboardGroupByIDDefault{
		_statusCode: code,
	}
}

/* UpdateDashboardGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDashboardGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update dashboard group by Id default response
func (o *UpdateDashboardGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDashboardGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /dashboard/groups/{id}][%d] updateDashboardGroupById default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDashboardGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDashboardGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
