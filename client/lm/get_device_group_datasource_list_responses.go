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

// GetDeviceGroupDatasourceListReader is a Reader for the GetDeviceGroupDatasourceList structure.
type GetDeviceGroupDatasourceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceGroupDatasourceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceGroupDatasourceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDeviceGroupDatasourceListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceGroupDatasourceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceGroupDatasourceListOK creates a GetDeviceGroupDatasourceListOK with default headers values
func NewGetDeviceGroupDatasourceListOK() *GetDeviceGroupDatasourceListOK {
	return &GetDeviceGroupDatasourceListOK{}
}

/* GetDeviceGroupDatasourceListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceGroupDatasourceListOK struct {
	Payload *models.DeviceGroupDatasourcePaginationResponse
}

func (o *GetDeviceGroupDatasourceListOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources][%d] getDeviceGroupDatasourceListOK  %+v", 200, o.Payload)
}
func (o *GetDeviceGroupDatasourceListOK) GetPayload() *models.DeviceGroupDatasourcePaginationResponse {
	return o.Payload
}

func (o *GetDeviceGroupDatasourceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroupDatasourcePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceGroupDatasourceListTooManyRequests creates a GetDeviceGroupDatasourceListTooManyRequests with default headers values
func NewGetDeviceGroupDatasourceListTooManyRequests() *GetDeviceGroupDatasourceListTooManyRequests {
	return &GetDeviceGroupDatasourceListTooManyRequests{}
}

/* GetDeviceGroupDatasourceListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceGroupDatasourceListTooManyRequests struct {

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

func (o *GetDeviceGroupDatasourceListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources][%d] getDeviceGroupDatasourceListTooManyRequests ", 429)
}

func (o *GetDeviceGroupDatasourceListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceGroupDatasourceListDefault creates a GetDeviceGroupDatasourceListDefault with default headers values
func NewGetDeviceGroupDatasourceListDefault(code int) *GetDeviceGroupDatasourceListDefault {
	return &GetDeviceGroupDatasourceListDefault{
		_statusCode: code,
	}
}

/* GetDeviceGroupDatasourceListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceGroupDatasourceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device group datasource list default response
func (o *GetDeviceGroupDatasourceListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceGroupDatasourceListDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources][%d] getDeviceGroupDatasourceList default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceGroupDatasourceListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceGroupDatasourceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
