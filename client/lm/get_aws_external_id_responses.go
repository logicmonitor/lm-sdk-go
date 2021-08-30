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

// GetAwsExternalIDReader is a Reader for the GetAwsExternalID structure.
type GetAwsExternalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsExternalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsExternalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAwsExternalIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAwsExternalIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAwsExternalIDOK creates a GetAwsExternalIDOK with default headers values
func NewGetAwsExternalIDOK() *GetAwsExternalIDOK {
	return &GetAwsExternalIDOK{}
}

/* GetAwsExternalIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAwsExternalIDOK struct {
	Payload *models.AwsExternalID
}

func (o *GetAwsExternalIDOK) Error() string {
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalIdOK  %+v", 200, o.Payload)
}
func (o *GetAwsExternalIDOK) GetPayload() *models.AwsExternalID {
	return o.Payload
}

func (o *GetAwsExternalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsExternalID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsExternalIDTooManyRequests creates a GetAwsExternalIDTooManyRequests with default headers values
func NewGetAwsExternalIDTooManyRequests() *GetAwsExternalIDTooManyRequests {
	return &GetAwsExternalIDTooManyRequests{}
}

/* GetAwsExternalIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAwsExternalIDTooManyRequests struct {

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

func (o *GetAwsExternalIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalIdTooManyRequests ", 429)
}

func (o *GetAwsExternalIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAwsExternalIDDefault creates a GetAwsExternalIDDefault with default headers values
func NewGetAwsExternalIDDefault(code int) *GetAwsExternalIDDefault {
	return &GetAwsExternalIDDefault{
		_statusCode: code,
	}
}

/* GetAwsExternalIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAwsExternalIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get aws external Id default response
func (o *GetAwsExternalIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAwsExternalIDDefault) Error() string {
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalId default  %+v", o._statusCode, o.Payload)
}
func (o *GetAwsExternalIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAwsExternalIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}