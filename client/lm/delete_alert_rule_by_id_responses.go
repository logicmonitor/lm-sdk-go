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

// DeleteAlertRuleByIDReader is a Reader for the DeleteAlertRuleByID structure.
type DeleteAlertRuleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertRuleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAlertRuleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewDeleteAlertRuleByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAlertRuleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAlertRuleByIDOK creates a DeleteAlertRuleByIDOK with default headers values
func NewDeleteAlertRuleByIDOK() *DeleteAlertRuleByIDOK {
	return &DeleteAlertRuleByIDOK{}
}

/* DeleteAlertRuleByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteAlertRuleByIDOK struct {
	Payload interface{}
}

func (o *DeleteAlertRuleByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/alert/rules/{id}][%d] deleteAlertRuleByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteAlertRuleByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAlertRuleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertRuleByIDTooManyRequests creates a DeleteAlertRuleByIDTooManyRequests with default headers values
func NewDeleteAlertRuleByIDTooManyRequests() *DeleteAlertRuleByIDTooManyRequests {
	return &DeleteAlertRuleByIDTooManyRequests{}
}

/* DeleteAlertRuleByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteAlertRuleByIDTooManyRequests struct {

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

func (o *DeleteAlertRuleByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /setting/alert/rules/{id}][%d] deleteAlertRuleByIdTooManyRequests ", 429)
}

func (o *DeleteAlertRuleByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAlertRuleByIDDefault creates a DeleteAlertRuleByIDDefault with default headers values
func NewDeleteAlertRuleByIDDefault(code int) *DeleteAlertRuleByIDDefault {
	return &DeleteAlertRuleByIDDefault{
		_statusCode: code,
	}
}

/* DeleteAlertRuleByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteAlertRuleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete alert rule by Id default response
func (o *DeleteAlertRuleByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAlertRuleByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/alert/rules/{id}][%d] deleteAlertRuleById default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteAlertRuleByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteAlertRuleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}