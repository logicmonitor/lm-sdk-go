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

// GetCollectorGroupByIDReader is a Reader for the GetCollectorGroupByID structure.
type GetCollectorGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectorGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectorGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetCollectorGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCollectorGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCollectorGroupByIDOK creates a GetCollectorGroupByIDOK with default headers values
func NewGetCollectorGroupByIDOK() *GetCollectorGroupByIDOK {
	return &GetCollectorGroupByIDOK{}
}

/*
GetCollectorGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCollectorGroupByIDOK struct {
	Payload *models.CollectorGroup
}

// IsSuccess returns true when this get collector group by Id o k response has a 2xx status code
func (o *GetCollectorGroupByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get collector group by Id o k response has a 3xx status code
func (o *GetCollectorGroupByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collector group by Id o k response has a 4xx status code
func (o *GetCollectorGroupByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get collector group by Id o k response has a 5xx status code
func (o *GetCollectorGroupByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get collector group by Id o k response a status code equal to that given
func (o *GetCollectorGroupByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCollectorGroupByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/collector/groups/{id}][%d] getCollectorGroupByIdOK  %+v", 200, o.Payload)
}

func (o *GetCollectorGroupByIDOK) String() string {
	return fmt.Sprintf("[GET /setting/collector/groups/{id}][%d] getCollectorGroupByIdOK  %+v", 200, o.Payload)
}

func (o *GetCollectorGroupByIDOK) GetPayload() *models.CollectorGroup {
	return o.Payload
}

func (o *GetCollectorGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectorGroupByIDTooManyRequests creates a GetCollectorGroupByIDTooManyRequests with default headers values
func NewGetCollectorGroupByIDTooManyRequests() *GetCollectorGroupByIDTooManyRequests {
	return &GetCollectorGroupByIDTooManyRequests{}
}

/*
GetCollectorGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCollectorGroupByIDTooManyRequests struct {

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

// IsSuccess returns true when this get collector group by Id too many requests response has a 2xx status code
func (o *GetCollectorGroupByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get collector group by Id too many requests response has a 3xx status code
func (o *GetCollectorGroupByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collector group by Id too many requests response has a 4xx status code
func (o *GetCollectorGroupByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get collector group by Id too many requests response has a 5xx status code
func (o *GetCollectorGroupByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get collector group by Id too many requests response a status code equal to that given
func (o *GetCollectorGroupByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCollectorGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/collector/groups/{id}][%d] getCollectorGroupByIdTooManyRequests ", 429)
}

func (o *GetCollectorGroupByIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/collector/groups/{id}][%d] getCollectorGroupByIdTooManyRequests ", 429)
}

func (o *GetCollectorGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCollectorGroupByIDDefault creates a GetCollectorGroupByIDDefault with default headers values
func NewGetCollectorGroupByIDDefault(code int) *GetCollectorGroupByIDDefault {
	return &GetCollectorGroupByIDDefault{
		_statusCode: code,
	}
}

/*
GetCollectorGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetCollectorGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get collector group by Id default response
func (o *GetCollectorGroupByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get collector group by Id default response has a 2xx status code
func (o *GetCollectorGroupByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get collector group by Id default response has a 3xx status code
func (o *GetCollectorGroupByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get collector group by Id default response has a 4xx status code
func (o *GetCollectorGroupByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get collector group by Id default response has a 5xx status code
func (o *GetCollectorGroupByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get collector group by Id default response a status code equal to that given
func (o *GetCollectorGroupByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCollectorGroupByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/collector/groups/{id}][%d] getCollectorGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *GetCollectorGroupByIDDefault) String() string {
	return fmt.Sprintf("[GET /setting/collector/groups/{id}][%d] getCollectorGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *GetCollectorGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCollectorGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
