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

// GetNetflowFlowListReader is a Reader for the GetNetflowFlowList structure.
type GetNetflowFlowListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetflowFlowListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetflowFlowListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetNetflowFlowListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNetflowFlowListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetflowFlowListOK creates a GetNetflowFlowListOK with default headers values
func NewGetNetflowFlowListOK() *GetNetflowFlowListOK {
	return &GetNetflowFlowListOK{}
}

/*
GetNetflowFlowListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetNetflowFlowListOK struct {
	Payload *models.FlowRecordPaginationResponse
}

// IsSuccess returns true when this get netflow flow list o k response has a 2xx status code
func (o *GetNetflowFlowListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get netflow flow list o k response has a 3xx status code
func (o *GetNetflowFlowListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get netflow flow list o k response has a 4xx status code
func (o *GetNetflowFlowListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get netflow flow list o k response has a 5xx status code
func (o *GetNetflowFlowListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get netflow flow list o k response a status code equal to that given
func (o *GetNetflowFlowListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetflowFlowListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/flows][%d] getNetflowFlowListOK  %+v", 200, o.Payload)
}

func (o *GetNetflowFlowListOK) String() string {
	return fmt.Sprintf("[GET /device/devices/{id}/flows][%d] getNetflowFlowListOK  %+v", 200, o.Payload)
}

func (o *GetNetflowFlowListOK) GetPayload() *models.FlowRecordPaginationResponse {
	return o.Payload
}

func (o *GetNetflowFlowListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowRecordPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetflowFlowListTooManyRequests creates a GetNetflowFlowListTooManyRequests with default headers values
func NewGetNetflowFlowListTooManyRequests() *GetNetflowFlowListTooManyRequests {
	return &GetNetflowFlowListTooManyRequests{}
}

/*
GetNetflowFlowListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetNetflowFlowListTooManyRequests struct {

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

// IsSuccess returns true when this get netflow flow list too many requests response has a 2xx status code
func (o *GetNetflowFlowListTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get netflow flow list too many requests response has a 3xx status code
func (o *GetNetflowFlowListTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get netflow flow list too many requests response has a 4xx status code
func (o *GetNetflowFlowListTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get netflow flow list too many requests response has a 5xx status code
func (o *GetNetflowFlowListTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get netflow flow list too many requests response a status code equal to that given
func (o *GetNetflowFlowListTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetNetflowFlowListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/flows][%d] getNetflowFlowListTooManyRequests ", 429)
}

func (o *GetNetflowFlowListTooManyRequests) String() string {
	return fmt.Sprintf("[GET /device/devices/{id}/flows][%d] getNetflowFlowListTooManyRequests ", 429)
}

func (o *GetNetflowFlowListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetNetflowFlowListDefault creates a GetNetflowFlowListDefault with default headers values
func NewGetNetflowFlowListDefault(code int) *GetNetflowFlowListDefault {
	return &GetNetflowFlowListDefault{
		_statusCode: code,
	}
}

/*
GetNetflowFlowListDefault describes a response with status code -1, with default header values.

Error
*/
type GetNetflowFlowListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get netflow flow list default response
func (o *GetNetflowFlowListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get netflow flow list default response has a 2xx status code
func (o *GetNetflowFlowListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get netflow flow list default response has a 3xx status code
func (o *GetNetflowFlowListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get netflow flow list default response has a 4xx status code
func (o *GetNetflowFlowListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get netflow flow list default response has a 5xx status code
func (o *GetNetflowFlowListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get netflow flow list default response a status code equal to that given
func (o *GetNetflowFlowListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetNetflowFlowListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/flows][%d] getNetflowFlowList default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetflowFlowListDefault) String() string {
	return fmt.Sprintf("[GET /device/devices/{id}/flows][%d] getNetflowFlowList default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetflowFlowListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetNetflowFlowListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
