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

// GetWebsiteAlertListByWebsiteIDReader is a Reader for the GetWebsiteAlertListByWebsiteID structure.
type GetWebsiteAlertListByWebsiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteAlertListByWebsiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteAlertListByWebsiteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWebsiteAlertListByWebsiteIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWebsiteAlertListByWebsiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteAlertListByWebsiteIDOK creates a GetWebsiteAlertListByWebsiteIDOK with default headers values
func NewGetWebsiteAlertListByWebsiteIDOK() *GetWebsiteAlertListByWebsiteIDOK {
	return &GetWebsiteAlertListByWebsiteIDOK{}
}

/* GetWebsiteAlertListByWebsiteIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteAlertListByWebsiteIDOK struct {
	Payload *models.AlertPaginationResponse
}

func (o *GetWebsiteAlertListByWebsiteIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/alerts][%d] getWebsiteAlertListByWebsiteIdOK  %+v", 200, o.Payload)
}
func (o *GetWebsiteAlertListByWebsiteIDOK) GetPayload() *models.AlertPaginationResponse {
	return o.Payload
}

func (o *GetWebsiteAlertListByWebsiteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteAlertListByWebsiteIDTooManyRequests creates a GetWebsiteAlertListByWebsiteIDTooManyRequests with default headers values
func NewGetWebsiteAlertListByWebsiteIDTooManyRequests() *GetWebsiteAlertListByWebsiteIDTooManyRequests {
	return &GetWebsiteAlertListByWebsiteIDTooManyRequests{}
}

/* GetWebsiteAlertListByWebsiteIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWebsiteAlertListByWebsiteIDTooManyRequests struct {

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

func (o *GetWebsiteAlertListByWebsiteIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/alerts][%d] getWebsiteAlertListByWebsiteIdTooManyRequests ", 429)
}

func (o *GetWebsiteAlertListByWebsiteIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebsiteAlertListByWebsiteIDDefault creates a GetWebsiteAlertListByWebsiteIDDefault with default headers values
func NewGetWebsiteAlertListByWebsiteIDDefault(code int) *GetWebsiteAlertListByWebsiteIDDefault {
	return &GetWebsiteAlertListByWebsiteIDDefault{
		_statusCode: code,
	}
}

/* GetWebsiteAlertListByWebsiteIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteAlertListByWebsiteIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website alert list by website Id default response
func (o *GetWebsiteAlertListByWebsiteIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteAlertListByWebsiteIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/alerts][%d] getWebsiteAlertListByWebsiteId default  %+v", o._statusCode, o.Payload)
}
func (o *GetWebsiteAlertListByWebsiteIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteAlertListByWebsiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}