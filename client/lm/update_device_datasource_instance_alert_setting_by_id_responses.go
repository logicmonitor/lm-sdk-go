// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// UpdateDeviceDatasourceInstanceAlertSettingByIDReader is a Reader for the UpdateDeviceDatasourceInstanceAlertSettingByID structure.
type UpdateDeviceDatasourceInstanceAlertSettingByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceDatasourceInstanceAlertSettingByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateDeviceDatasourceInstanceAlertSettingByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDeviceDatasourceInstanceAlertSettingByIDOK creates a UpdateDeviceDatasourceInstanceAlertSettingByIDOK with default headers values
func NewUpdateDeviceDatasourceInstanceAlertSettingByIDOK() *UpdateDeviceDatasourceInstanceAlertSettingByIDOK {
	return &UpdateDeviceDatasourceInstanceAlertSettingByIDOK{}
}

/*
	UpdateDeviceDatasourceInstanceAlertSettingByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDeviceDatasourceInstanceAlertSettingByIDOK struct {
	Payload *models.DeviceDataSourceInstanceAlertSetting
}

func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDOK) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] updateDeviceDatasourceInstanceAlertSettingByIdOK  %+v", 200, o.Payload)
}
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDOK) GetPayload() *models.DeviceDataSourceInstanceAlertSetting {
	return o.Payload
}

func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstanceAlertSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceDatasourceInstanceAlertSettingByIDDefault creates a UpdateDeviceDatasourceInstanceAlertSettingByIDDefault with default headers values
func NewUpdateDeviceDatasourceInstanceAlertSettingByIDDefault(code int) *UpdateDeviceDatasourceInstanceAlertSettingByIDDefault {
	return &UpdateDeviceDatasourceInstanceAlertSettingByIDDefault{
		_statusCode: code,
	}
}

/*
	UpdateDeviceDatasourceInstanceAlertSettingByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDeviceDatasourceInstanceAlertSettingByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device datasource instance alert setting by Id default response
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] updateDeviceDatasourceInstanceAlertSettingById default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceDatasourceInstanceAlertSettingByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
