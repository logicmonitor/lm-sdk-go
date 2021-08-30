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

// GetAlertRuleByIDReader is a Reader for the GetAlertRuleByID structure.
type GetAlertRuleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRuleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertRuleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAlertRuleByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAlertRuleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertRuleByIDOK creates a GetAlertRuleByIDOK with default headers values
func NewGetAlertRuleByIDOK() *GetAlertRuleByIDOK {
	return &GetAlertRuleByIDOK{}
}

/* GetAlertRuleByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertRuleByIDOK struct {
	Payload *models.AlertRule
}

func (o *GetAlertRuleByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules/{id}][%d] getAlertRuleByIdOK  %+v", 200, o.Payload)
}
func (o *GetAlertRuleByIDOK) GetPayload() *models.AlertRule {
	return o.Payload
}

func (o *GetAlertRuleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRuleByIDTooManyRequests creates a GetAlertRuleByIDTooManyRequests with default headers values
func NewGetAlertRuleByIDTooManyRequests() *GetAlertRuleByIDTooManyRequests {
	return &GetAlertRuleByIDTooManyRequests{}
}

/* GetAlertRuleByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAlertRuleByIDTooManyRequests struct {

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

func (o *GetAlertRuleByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules/{id}][%d] getAlertRuleByIdTooManyRequests ", 429)
}

func (o *GetAlertRuleByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertRuleByIDDefault creates a GetAlertRuleByIDDefault with default headers values
func NewGetAlertRuleByIDDefault(code int) *GetAlertRuleByIDDefault {
	return &GetAlertRuleByIDDefault{
		_statusCode: code,
	}
}

/* GetAlertRuleByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertRuleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get alert rule by Id default response
func (o *GetAlertRuleByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertRuleByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules/{id}][%d] getAlertRuleById default  %+v", o._statusCode, o.Payload)
}
func (o *GetAlertRuleByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertRuleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}