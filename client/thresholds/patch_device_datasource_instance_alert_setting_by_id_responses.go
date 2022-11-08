// Code generated by go-swagger; DO NOT EDIT.

package thresholds

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

// PatchDeviceDatasourceInstanceAlertSettingByIDReader is a Reader for the PatchDeviceDatasourceInstanceAlertSettingByID structure.
type PatchDeviceDatasourceInstanceAlertSettingByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDeviceDatasourceInstanceAlertSettingByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchDeviceDatasourceInstanceAlertSettingByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDeviceDatasourceInstanceAlertSettingByIDOK creates a PatchDeviceDatasourceInstanceAlertSettingByIDOK with default headers values
func NewPatchDeviceDatasourceInstanceAlertSettingByIDOK() *PatchDeviceDatasourceInstanceAlertSettingByIDOK {
	return &PatchDeviceDatasourceInstanceAlertSettingByIDOK{}
}

/*
PatchDeviceDatasourceInstanceAlertSettingByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchDeviceDatasourceInstanceAlertSettingByIDOK struct {
	Payload *models.DeviceDataSourceInstanceAlertSetting
}

// IsSuccess returns true when this patch device datasource instance alert setting by Id o k response has a 2xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch device datasource instance alert setting by Id o k response has a 3xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch device datasource instance alert setting by Id o k response has a 4xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch device datasource instance alert setting by Id o k response has a 5xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch device datasource instance alert setting by Id o k response a status code equal to that given
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] patchDeviceDatasourceInstanceAlertSettingByIdOK  %+v", 200, o.Payload)
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) String() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] patchDeviceDatasourceInstanceAlertSettingByIdOK  %+v", 200, o.Payload)
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) GetPayload() *models.DeviceDataSourceInstanceAlertSetting {
	return o.Payload
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstanceAlertSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests creates a PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests with default headers values
func NewPatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests() *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests {
	return &PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests{}
}

/*
PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests struct {

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

// IsSuccess returns true when this patch device datasource instance alert setting by Id too many requests response has a 2xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch device datasource instance alert setting by Id too many requests response has a 3xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch device datasource instance alert setting by Id too many requests response has a 4xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch device datasource instance alert setting by Id too many requests response has a 5xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch device datasource instance alert setting by Id too many requests response a status code equal to that given
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] patchDeviceDatasourceInstanceAlertSettingByIdTooManyRequests ", 429)
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] patchDeviceDatasourceInstanceAlertSettingByIdTooManyRequests ", 429)
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchDeviceDatasourceInstanceAlertSettingByIDDefault creates a PatchDeviceDatasourceInstanceAlertSettingByIDDefault with default headers values
func NewPatchDeviceDatasourceInstanceAlertSettingByIDDefault(code int) *PatchDeviceDatasourceInstanceAlertSettingByIDDefault {
	return &PatchDeviceDatasourceInstanceAlertSettingByIDDefault{
		_statusCode: code,
	}
}

/*
PatchDeviceDatasourceInstanceAlertSettingByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchDeviceDatasourceInstanceAlertSettingByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch device datasource instance alert setting by Id default response
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch device datasource instance alert setting by Id default response has a 2xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch device datasource instance alert setting by Id default response has a 3xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch device datasource instance alert setting by Id default response has a 4xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch device datasource instance alert setting by Id default response has a 5xx status code
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch device datasource instance alert setting by Id default response a status code equal to that given
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] patchDeviceDatasourceInstanceAlertSettingById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) String() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] patchDeviceDatasourceInstanceAlertSettingById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
