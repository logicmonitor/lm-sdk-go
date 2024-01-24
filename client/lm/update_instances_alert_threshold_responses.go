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

// UpdateInstancesAlertThresholdReader is a Reader for the UpdateInstancesAlertThreshold structure.
type UpdateInstancesAlertThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstancesAlertThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInstancesAlertThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateInstancesAlertThresholdTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateInstancesAlertThresholdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateInstancesAlertThresholdOK creates a UpdateInstancesAlertThresholdOK with default headers values
func NewUpdateInstancesAlertThresholdOK() *UpdateInstancesAlertThresholdOK {
	return &UpdateInstancesAlertThresholdOK{}
}

/* UpdateInstancesAlertThresholdOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateInstancesAlertThresholdOK struct {
	Payload *models.DeviceDataSourceInstanceGroup
}

func (o *UpdateInstancesAlertThresholdOK) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{dsigId}/datapoints/{dpId}/alertconfig][%d] updateInstancesAlertThresholdOK  %+v", 200, o.Payload)
}
func (o *UpdateInstancesAlertThresholdOK) GetPayload() *models.DeviceDataSourceInstanceGroup {
	return o.Payload
}

func (o *UpdateInstancesAlertThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstanceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstancesAlertThresholdTooManyRequests creates a UpdateInstancesAlertThresholdTooManyRequests with default headers values
func NewUpdateInstancesAlertThresholdTooManyRequests() *UpdateInstancesAlertThresholdTooManyRequests {
	return &UpdateInstancesAlertThresholdTooManyRequests{}
}

/* UpdateInstancesAlertThresholdTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateInstancesAlertThresholdTooManyRequests struct {

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

func (o *UpdateInstancesAlertThresholdTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{dsigId}/datapoints/{dpId}/alertconfig][%d] updateInstancesAlertThresholdTooManyRequests ", 429)
}

func (o *UpdateInstancesAlertThresholdTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateInstancesAlertThresholdDefault creates a UpdateInstancesAlertThresholdDefault with default headers values
func NewUpdateInstancesAlertThresholdDefault(code int) *UpdateInstancesAlertThresholdDefault {
	return &UpdateInstancesAlertThresholdDefault{
		_statusCode: code,
	}
}

/* UpdateInstancesAlertThresholdDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateInstancesAlertThresholdDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update instances alert threshold default response
func (o *UpdateInstancesAlertThresholdDefault) Code() int {
	return o._statusCode
}

func (o *UpdateInstancesAlertThresholdDefault) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{dsigId}/datapoints/{dpId}/alertconfig][%d] updateInstancesAlertThreshold default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateInstancesAlertThresholdDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateInstancesAlertThresholdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}