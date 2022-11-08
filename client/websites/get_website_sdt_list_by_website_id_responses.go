// Code generated by go-swagger; DO NOT EDIT.

package websites

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

// GetWebsiteSDTListByWebsiteIDReader is a Reader for the GetWebsiteSDTListByWebsiteID structure.
type GetWebsiteSDTListByWebsiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteSDTListByWebsiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteSDTListByWebsiteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWebsiteSDTListByWebsiteIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWebsiteSDTListByWebsiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteSDTListByWebsiteIDOK creates a GetWebsiteSDTListByWebsiteIDOK with default headers values
func NewGetWebsiteSDTListByWebsiteIDOK() *GetWebsiteSDTListByWebsiteIDOK {
	return &GetWebsiteSDTListByWebsiteIDOK{}
}

/*
GetWebsiteSDTListByWebsiteIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteSDTListByWebsiteIDOK struct {
	Payload *models.SDTPaginationResponse
}

// IsSuccess returns true when this get website Sdt list by website Id o k response has a 2xx status code
func (o *GetWebsiteSDTListByWebsiteIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get website Sdt list by website Id o k response has a 3xx status code
func (o *GetWebsiteSDTListByWebsiteIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get website Sdt list by website Id o k response has a 4xx status code
func (o *GetWebsiteSDTListByWebsiteIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get website Sdt list by website Id o k response has a 5xx status code
func (o *GetWebsiteSDTListByWebsiteIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get website Sdt list by website Id o k response a status code equal to that given
func (o *GetWebsiteSDTListByWebsiteIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWebsiteSDTListByWebsiteIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/sdts][%d] getWebsiteSdtListByWebsiteIdOK  %+v", 200, o.Payload)
}

func (o *GetWebsiteSDTListByWebsiteIDOK) String() string {
	return fmt.Sprintf("[GET /website/websites/{id}/sdts][%d] getWebsiteSdtListByWebsiteIdOK  %+v", 200, o.Payload)
}

func (o *GetWebsiteSDTListByWebsiteIDOK) GetPayload() *models.SDTPaginationResponse {
	return o.Payload
}

func (o *GetWebsiteSDTListByWebsiteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDTPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteSDTListByWebsiteIDTooManyRequests creates a GetWebsiteSDTListByWebsiteIDTooManyRequests with default headers values
func NewGetWebsiteSDTListByWebsiteIDTooManyRequests() *GetWebsiteSDTListByWebsiteIDTooManyRequests {
	return &GetWebsiteSDTListByWebsiteIDTooManyRequests{}
}

/*
GetWebsiteSDTListByWebsiteIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWebsiteSDTListByWebsiteIDTooManyRequests struct {

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

// IsSuccess returns true when this get website Sdt list by website Id too many requests response has a 2xx status code
func (o *GetWebsiteSDTListByWebsiteIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get website Sdt list by website Id too many requests response has a 3xx status code
func (o *GetWebsiteSDTListByWebsiteIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get website Sdt list by website Id too many requests response has a 4xx status code
func (o *GetWebsiteSDTListByWebsiteIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get website Sdt list by website Id too many requests response has a 5xx status code
func (o *GetWebsiteSDTListByWebsiteIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get website Sdt list by website Id too many requests response a status code equal to that given
func (o *GetWebsiteSDTListByWebsiteIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWebsiteSDTListByWebsiteIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/sdts][%d] getWebsiteSdtListByWebsiteIdTooManyRequests ", 429)
}

func (o *GetWebsiteSDTListByWebsiteIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /website/websites/{id}/sdts][%d] getWebsiteSdtListByWebsiteIdTooManyRequests ", 429)
}

func (o *GetWebsiteSDTListByWebsiteIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebsiteSDTListByWebsiteIDDefault creates a GetWebsiteSDTListByWebsiteIDDefault with default headers values
func NewGetWebsiteSDTListByWebsiteIDDefault(code int) *GetWebsiteSDTListByWebsiteIDDefault {
	return &GetWebsiteSDTListByWebsiteIDDefault{
		_statusCode: code,
	}
}

/*
GetWebsiteSDTListByWebsiteIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteSDTListByWebsiteIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website SDT list by website Id default response
func (o *GetWebsiteSDTListByWebsiteIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get website SDT list by website Id default response has a 2xx status code
func (o *GetWebsiteSDTListByWebsiteIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get website SDT list by website Id default response has a 3xx status code
func (o *GetWebsiteSDTListByWebsiteIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get website SDT list by website Id default response has a 4xx status code
func (o *GetWebsiteSDTListByWebsiteIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get website SDT list by website Id default response has a 5xx status code
func (o *GetWebsiteSDTListByWebsiteIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get website SDT list by website Id default response a status code equal to that given
func (o *GetWebsiteSDTListByWebsiteIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetWebsiteSDTListByWebsiteIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/sdts][%d] getWebsiteSDTListByWebsiteId default  %+v", o._statusCode, o.Payload)
}

func (o *GetWebsiteSDTListByWebsiteIDDefault) String() string {
	return fmt.Sprintf("[GET /website/websites/{id}/sdts][%d] getWebsiteSDTListByWebsiteId default  %+v", o._statusCode, o.Payload)
}

func (o *GetWebsiteSDTListByWebsiteIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteSDTListByWebsiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
