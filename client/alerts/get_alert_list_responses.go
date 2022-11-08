// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// GetAlertListReader is a Reader for the GetAlertList structure.
type GetAlertListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAlertListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAlertListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertListOK creates a GetAlertListOK with default headers values
func NewGetAlertListOK() *GetAlertListOK {
	return &GetAlertListOK{}
}

/*
GetAlertListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertListOK struct {
	Payload *models.AlertPaginationResponse
}

// IsSuccess returns true when this get alert list o k response has a 2xx status code
func (o *GetAlertListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert list o k response has a 3xx status code
func (o *GetAlertListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert list o k response has a 4xx status code
func (o *GetAlertListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert list o k response has a 5xx status code
func (o *GetAlertListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert list o k response a status code equal to that given
func (o *GetAlertListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAlertListOK) Error() string {
	return fmt.Sprintf("[GET /alert/alerts][%d] getAlertListOK  %+v", 200, o.Payload)
}

func (o *GetAlertListOK) String() string {
	return fmt.Sprintf("[GET /alert/alerts][%d] getAlertListOK  %+v", 200, o.Payload)
}

func (o *GetAlertListOK) GetPayload() *models.AlertPaginationResponse {
	return o.Payload
}

func (o *GetAlertListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertListTooManyRequests creates a GetAlertListTooManyRequests with default headers values
func NewGetAlertListTooManyRequests() *GetAlertListTooManyRequests {
	return &GetAlertListTooManyRequests{}
}

/*
GetAlertListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAlertListTooManyRequests struct {

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

// IsSuccess returns true when this get alert list too many requests response has a 2xx status code
func (o *GetAlertListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert list too many requests response has a 3xx status code
func (o *GetAlertListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert list too many requests response has a 4xx status code
func (o *GetAlertListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert list too many requests response has a 5xx status code
func (o *GetAlertListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert list too many requests response a status code equal to that given
func (o *GetAlertListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAlertListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /alert/alerts][%d] getAlertListTooManyRequests ", 429)
}

func (o *GetAlertListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /alert/alerts][%d] getAlertListTooManyRequests ", 429)
}

func (o *GetAlertListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertListDefault creates a GetAlertListDefault with default headers values
func NewGetAlertListDefault(code int) *GetAlertListDefault {
	return &GetAlertListDefault{
		_statusCode: code,
	}
}

/*
GetAlertListDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get alert list default response
func (o *GetAlertListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get alert list default response has a 2xx status code
func (o *GetAlertListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get alert list default response has a 3xx status code
func (o *GetAlertListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get alert list default response has a 4xx status code
func (o *GetAlertListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get alert list default response has a 5xx status code
func (o *GetAlertListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get alert list default response a status code equal to that given
func (o *GetAlertListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAlertListDefault) Error() string {
	return fmt.Sprintf("[GET /alert/alerts][%d] getAlertList default  %+v", o._statusCode, o.Payload)
}

func (o *GetAlertListDefault) String() string {
	return fmt.Sprintf("[GET /alert/alerts][%d] getAlertList default  %+v", o._statusCode, o.Payload)
}

func (o *GetAlertListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
