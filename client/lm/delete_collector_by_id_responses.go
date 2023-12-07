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

// DeleteCollectorByIDReader is a Reader for the DeleteCollectorByID structure.
type DeleteCollectorByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCollectorByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCollectorByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewDeleteCollectorByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCollectorByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCollectorByIDOK creates a DeleteCollectorByIDOK with default headers values
func NewDeleteCollectorByIDOK() *DeleteCollectorByIDOK {
	return &DeleteCollectorByIDOK{}
}

/* DeleteCollectorByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteCollectorByIDOK struct {
	Payload interface{}
}

func (o *DeleteCollectorByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/collector/collectors/{id}][%d] deleteCollectorByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteCollectorByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteCollectorByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCollectorByIDTooManyRequests creates a DeleteCollectorByIDTooManyRequests with default headers values
func NewDeleteCollectorByIDTooManyRequests() *DeleteCollectorByIDTooManyRequests {
	return &DeleteCollectorByIDTooManyRequests{}
}

/* DeleteCollectorByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteCollectorByIDTooManyRequests struct {

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

func (o *DeleteCollectorByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /setting/collector/collectors/{id}][%d] deleteCollectorByIdTooManyRequests ", 429)
}

func (o *DeleteCollectorByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCollectorByIDDefault creates a DeleteCollectorByIDDefault with default headers values
func NewDeleteCollectorByIDDefault(code int) *DeleteCollectorByIDDefault {
	return &DeleteCollectorByIDDefault{
		_statusCode: code,
	}
}

/* DeleteCollectorByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteCollectorByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete collector by Id default response
func (o *DeleteCollectorByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCollectorByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/collector/collectors/{id}][%d] deleteCollectorById default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteCollectorByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteCollectorByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
