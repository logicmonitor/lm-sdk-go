// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// GetDeviceListReader is a Reader for the GetDeviceList structure.
type GetDeviceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceListOK creates a GetDeviceListOK with default headers values
func NewGetDeviceListOK() *GetDeviceListOK {
	return &GetDeviceListOK{}
}

/*
GetDeviceListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceListOK struct {
	Payload *models.DevicePaginationResponse
}

// IsSuccess returns true when this get device list o k response has a 2xx status code
func (o *GetDeviceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device list o k response has a 3xx status code
func (o *GetDeviceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device list o k response has a 4xx status code
func (o *GetDeviceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device list o k response has a 5xx status code
func (o *GetDeviceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device list o k response a status code equal to that given
func (o *GetDeviceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices][%d] getDeviceListOK  %+v", 200, o.Payload)
}

func (o *GetDeviceListOK) String() string {
	return fmt.Sprintf("[GET /device/devices][%d] getDeviceListOK  %+v", 200, o.Payload)
}

func (o *GetDeviceListOK) GetPayload() *models.DevicePaginationResponse {
	return o.Payload
}

func (o *GetDeviceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevicePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceListTooManyRequests creates a GetDeviceListTooManyRequests with default headers values
func NewGetDeviceListTooManyRequests() *GetDeviceListTooManyRequests {
	return &GetDeviceListTooManyRequests{}
}

/*
GetDeviceListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceListTooManyRequests struct {

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

// IsSuccess returns true when this get device list too many requests response has a 2xx status code
func (o *GetDeviceListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device list too many requests response has a 3xx status code
func (o *GetDeviceListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device list too many requests response has a 4xx status code
func (o *GetDeviceListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device list too many requests response has a 5xx status code
func (o *GetDeviceListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get device list too many requests response a status code equal to that given
func (o *GetDeviceListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDeviceListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices][%d] getDeviceListTooManyRequests ", 429)
}

func (o *GetDeviceListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /device/devices][%d] getDeviceListTooManyRequests ", 429)
}

func (o *GetDeviceListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceListDefault creates a GetDeviceListDefault with default headers values
func NewGetDeviceListDefault(code int) *GetDeviceListDefault {
	return &GetDeviceListDefault{
		_statusCode: code,
	}
}

/*
GetDeviceListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device list default response
func (o *GetDeviceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get device list default response has a 2xx status code
func (o *GetDeviceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get device list default response has a 3xx status code
func (o *GetDeviceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get device list default response has a 4xx status code
func (o *GetDeviceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get device list default response has a 5xx status code
func (o *GetDeviceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get device list default response a status code equal to that given
func (o *GetDeviceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetDeviceListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices][%d] getDeviceList default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceListDefault) String() string {
	return fmt.Sprintf("[GET /device/devices][%d] getDeviceList default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
