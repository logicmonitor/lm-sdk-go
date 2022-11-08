// Code generated by go-swagger; DO NOT EDIT.

package collector_groups

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

// UpdateCollectorGroupByIDReader is a Reader for the UpdateCollectorGroupByID structure.
type UpdateCollectorGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCollectorGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCollectorGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateCollectorGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateCollectorGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCollectorGroupByIDOK creates a UpdateCollectorGroupByIDOK with default headers values
func NewUpdateCollectorGroupByIDOK() *UpdateCollectorGroupByIDOK {
	return &UpdateCollectorGroupByIDOK{}
}

/*
UpdateCollectorGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateCollectorGroupByIDOK struct {
	Payload *models.CollectorGroup
}

// IsSuccess returns true when this update collector group by Id o k response has a 2xx status code
func (o *UpdateCollectorGroupByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update collector group by Id o k response has a 3xx status code
func (o *UpdateCollectorGroupByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update collector group by Id o k response has a 4xx status code
func (o *UpdateCollectorGroupByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update collector group by Id o k response has a 5xx status code
func (o *UpdateCollectorGroupByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update collector group by Id o k response a status code equal to that given
func (o *UpdateCollectorGroupByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateCollectorGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/collector/groups/{id}][%d] updateCollectorGroupByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateCollectorGroupByIDOK) String() string {
	return fmt.Sprintf("[PUT /setting/collector/groups/{id}][%d] updateCollectorGroupByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateCollectorGroupByIDOK) GetPayload() *models.CollectorGroup {
	return o.Payload
}

func (o *UpdateCollectorGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectorGroupByIDTooManyRequests creates a UpdateCollectorGroupByIDTooManyRequests with default headers values
func NewUpdateCollectorGroupByIDTooManyRequests() *UpdateCollectorGroupByIDTooManyRequests {
	return &UpdateCollectorGroupByIDTooManyRequests{}
}

/*
UpdateCollectorGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateCollectorGroupByIDTooManyRequests struct {

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

// IsSuccess returns true when this update collector group by Id too many requests response has a 2xx status code
func (o *UpdateCollectorGroupByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update collector group by Id too many requests response has a 3xx status code
func (o *UpdateCollectorGroupByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update collector group by Id too many requests response has a 4xx status code
func (o *UpdateCollectorGroupByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update collector group by Id too many requests response has a 5xx status code
func (o *UpdateCollectorGroupByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update collector group by Id too many requests response a status code equal to that given
func (o *UpdateCollectorGroupByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateCollectorGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /setting/collector/groups/{id}][%d] updateCollectorGroupByIdTooManyRequests ", 429)
}

func (o *UpdateCollectorGroupByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /setting/collector/groups/{id}][%d] updateCollectorGroupByIdTooManyRequests ", 429)
}

func (o *UpdateCollectorGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCollectorGroupByIDDefault creates a UpdateCollectorGroupByIDDefault with default headers values
func NewUpdateCollectorGroupByIDDefault(code int) *UpdateCollectorGroupByIDDefault {
	return &UpdateCollectorGroupByIDDefault{
		_statusCode: code,
	}
}

/*
UpdateCollectorGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateCollectorGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update collector group by Id default response
func (o *UpdateCollectorGroupByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update collector group by Id default response has a 2xx status code
func (o *UpdateCollectorGroupByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update collector group by Id default response has a 3xx status code
func (o *UpdateCollectorGroupByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update collector group by Id default response has a 4xx status code
func (o *UpdateCollectorGroupByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update collector group by Id default response has a 5xx status code
func (o *UpdateCollectorGroupByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update collector group by Id default response a status code equal to that given
func (o *UpdateCollectorGroupByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateCollectorGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/collector/groups/{id}][%d] updateCollectorGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCollectorGroupByIDDefault) String() string {
	return fmt.Sprintf("[PUT /setting/collector/groups/{id}][%d] updateCollectorGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCollectorGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateCollectorGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
