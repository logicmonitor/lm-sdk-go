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

// GetSDTHistoryByWebsiteIDReader is a Reader for the GetSDTHistoryByWebsiteID structure.
type GetSDTHistoryByWebsiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSDTHistoryByWebsiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSDTHistoryByWebsiteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetSDTHistoryByWebsiteIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSDTHistoryByWebsiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSDTHistoryByWebsiteIDOK creates a GetSDTHistoryByWebsiteIDOK with default headers values
func NewGetSDTHistoryByWebsiteIDOK() *GetSDTHistoryByWebsiteIDOK {
	return &GetSDTHistoryByWebsiteIDOK{}
}

/* GetSDTHistoryByWebsiteIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSDTHistoryByWebsiteIDOK struct {
	Payload *models.WebsiteSDTHistoryPaginationResponse
}

func (o *GetSDTHistoryByWebsiteIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/historysdts][%d] getSdtHistoryByWebsiteIdOK  %+v", 200, o.Payload)
}
func (o *GetSDTHistoryByWebsiteIDOK) GetPayload() *models.WebsiteSDTHistoryPaginationResponse {
	return o.Payload
}

func (o *GetSDTHistoryByWebsiteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteSDTHistoryPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDTHistoryByWebsiteIDTooManyRequests creates a GetSDTHistoryByWebsiteIDTooManyRequests with default headers values
func NewGetSDTHistoryByWebsiteIDTooManyRequests() *GetSDTHistoryByWebsiteIDTooManyRequests {
	return &GetSDTHistoryByWebsiteIDTooManyRequests{}
}

/* GetSDTHistoryByWebsiteIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSDTHistoryByWebsiteIDTooManyRequests struct {

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

func (o *GetSDTHistoryByWebsiteIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/historysdts][%d] getSdtHistoryByWebsiteIdTooManyRequests ", 429)
}

func (o *GetSDTHistoryByWebsiteIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSDTHistoryByWebsiteIDDefault creates a GetSDTHistoryByWebsiteIDDefault with default headers values
func NewGetSDTHistoryByWebsiteIDDefault(code int) *GetSDTHistoryByWebsiteIDDefault {
	return &GetSDTHistoryByWebsiteIDDefault{
		_statusCode: code,
	}
}

/* GetSDTHistoryByWebsiteIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetSDTHistoryByWebsiteIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get SDT history by website Id default response
func (o *GetSDTHistoryByWebsiteIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSDTHistoryByWebsiteIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/historysdts][%d] getSDTHistoryByWebsiteId default  %+v", o._statusCode, o.Payload)
}
func (o *GetSDTHistoryByWebsiteIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSDTHistoryByWebsiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
