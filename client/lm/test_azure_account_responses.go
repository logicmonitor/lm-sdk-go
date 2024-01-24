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

// TestAzureAccountReader is a Reader for the TestAzureAccount structure.
type TestAzureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestAzureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestAzureAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewTestAzureAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTestAzureAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTestAzureAccountOK creates a TestAzureAccountOK with default headers values
func NewTestAzureAccountOK() *TestAzureAccountOK {
	return &TestAzureAccountOK{}
}

/* TestAzureAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type TestAzureAccountOK struct {
	Payload models.RestCloudOkPermissionsV3
}

func (o *TestAzureAccountOK) Error() string {
	return fmt.Sprintf("[POST /azure/functions/testAccount][%d] testAzureAccountOK  %+v", 200, o.Payload)
}
func (o *TestAzureAccountOK) GetPayload() models.RestCloudOkPermissionsV3 {
	return o.Payload
}

func (o *TestAzureAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAzureAccountTooManyRequests creates a TestAzureAccountTooManyRequests with default headers values
func NewTestAzureAccountTooManyRequests() *TestAzureAccountTooManyRequests {
	return &TestAzureAccountTooManyRequests{}
}

/* TestAzureAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type TestAzureAccountTooManyRequests struct {

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

func (o *TestAzureAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /azure/functions/testAccount][%d] testAzureAccountTooManyRequests ", 429)
}

func (o *TestAzureAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTestAzureAccountDefault creates a TestAzureAccountDefault with default headers values
func NewTestAzureAccountDefault(code int) *TestAzureAccountDefault {
	return &TestAzureAccountDefault{
		_statusCode: code,
	}
}

/* TestAzureAccountDefault describes a response with status code -1, with default header values.

Error
*/
type TestAzureAccountDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the test azure account default response
func (o *TestAzureAccountDefault) Code() int {
	return o._statusCode
}

func (o *TestAzureAccountDefault) Error() string {
	return fmt.Sprintf("[POST /azure/functions/testAccount][%d] testAzureAccount default  %+v", o._statusCode, o.Payload)
}
func (o *TestAzureAccountDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TestAzureAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}