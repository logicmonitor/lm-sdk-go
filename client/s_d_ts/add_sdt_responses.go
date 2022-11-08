// Code generated by go-swagger; DO NOT EDIT.

package s_d_ts

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

// AddSDTReader is a Reader for the AddSDT structure.
type AddSDTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSDTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddSDTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddSDTTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddSDTDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddSDTOK creates a AddSDTOK with default headers values
func NewAddSDTOK() *AddSDTOK {
	return &AddSDTOK{}
}

/*
AddSDTOK describes a response with status code 200, with default header values.

successful operation
*/
type AddSDTOK struct {
	Payload models.SDT
}

// IsSuccess returns true when this add Sdt o k response has a 2xx status code
func (o *AddSDTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add Sdt o k response has a 3xx status code
func (o *AddSDTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add Sdt o k response has a 4xx status code
func (o *AddSDTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add Sdt o k response has a 5xx status code
func (o *AddSDTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add Sdt o k response a status code equal to that given
func (o *AddSDTOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddSDTOK) Error() string {
	return fmt.Sprintf("[POST /sdt/sdts][%d] addSdtOK  %+v", 200, o.Payload)
}

func (o *AddSDTOK) String() string {
	return fmt.Sprintf("[POST /sdt/sdts][%d] addSdtOK  %+v", 200, o.Payload)
}

func (o *AddSDTOK) GetPayload() models.SDT {
	return o.Payload
}

func (o *AddSDTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalSDT(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewAddSDTTooManyRequests creates a AddSDTTooManyRequests with default headers values
func NewAddSDTTooManyRequests() *AddSDTTooManyRequests {
	return &AddSDTTooManyRequests{}
}

/*
AddSDTTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddSDTTooManyRequests struct {

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

// IsSuccess returns true when this add Sdt too many requests response has a 2xx status code
func (o *AddSDTTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add Sdt too many requests response has a 3xx status code
func (o *AddSDTTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add Sdt too many requests response has a 4xx status code
func (o *AddSDTTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add Sdt too many requests response has a 5xx status code
func (o *AddSDTTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add Sdt too many requests response a status code equal to that given
func (o *AddSDTTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AddSDTTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /sdt/sdts][%d] addSdtTooManyRequests ", 429)
}

func (o *AddSDTTooManyRequests) String() string {
	return fmt.Sprintf("[POST /sdt/sdts][%d] addSdtTooManyRequests ", 429)
}

func (o *AddSDTTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddSDTDefault creates a AddSDTDefault with default headers values
func NewAddSDTDefault(code int) *AddSDTDefault {
	return &AddSDTDefault{
		_statusCode: code,
	}
}

/*
AddSDTDefault describes a response with status code -1, with default header values.

Error
*/
type AddSDTDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add SDT default response
func (o *AddSDTDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add SDT default response has a 2xx status code
func (o *AddSDTDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add SDT default response has a 3xx status code
func (o *AddSDTDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add SDT default response has a 4xx status code
func (o *AddSDTDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add SDT default response has a 5xx status code
func (o *AddSDTDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add SDT default response a status code equal to that given
func (o *AddSDTDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddSDTDefault) Error() string {
	return fmt.Sprintf("[POST /sdt/sdts][%d] addSDT default  %+v", o._statusCode, o.Payload)
}

func (o *AddSDTDefault) String() string {
	return fmt.Sprintf("[POST /sdt/sdts][%d] addSDT default  %+v", o._statusCode, o.Payload)
}

func (o *AddSDTDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddSDTDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
