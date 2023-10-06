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

// UpdateDeviceGroupDatasourceAlertSettingReader is a Reader for the UpdateDeviceGroupDatasourceAlertSetting structure.
type UpdateDeviceGroupDatasourceAlertSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceGroupDatasourceAlertSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceGroupDatasourceAlertSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateDeviceGroupDatasourceAlertSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDeviceGroupDatasourceAlertSettingOK creates a UpdateDeviceGroupDatasourceAlertSettingOK with default headers values
func NewUpdateDeviceGroupDatasourceAlertSettingOK() *UpdateDeviceGroupDatasourceAlertSettingOK {
	return &UpdateDeviceGroupDatasourceAlertSettingOK{}
}

/*
	UpdateDeviceGroupDatasourceAlertSettingOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDeviceGroupDatasourceAlertSettingOK struct {
	Payload *models.DeviceGroupDataSourceAlertConfig
}

func (o *UpdateDeviceGroupDatasourceAlertSettingOK) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{deviceGroupId}/datasources/{dsId}/alertsettings][%d] updateDeviceGroupDatasourceAlertSettingOK  %+v", 200, o.Payload)
}
func (o *UpdateDeviceGroupDatasourceAlertSettingOK) GetPayload() *models.DeviceGroupDataSourceAlertConfig {
	return o.Payload
}

func (o *UpdateDeviceGroupDatasourceAlertSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroupDataSourceAlertConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceGroupDatasourceAlertSettingDefault creates a UpdateDeviceGroupDatasourceAlertSettingDefault with default headers values
func NewUpdateDeviceGroupDatasourceAlertSettingDefault(code int) *UpdateDeviceGroupDatasourceAlertSettingDefault {
	return &UpdateDeviceGroupDatasourceAlertSettingDefault{
		_statusCode: code,
	}
}

/*
	UpdateDeviceGroupDatasourceAlertSettingDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDeviceGroupDatasourceAlertSettingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device group datasource alert setting default response
func (o *UpdateDeviceGroupDatasourceAlertSettingDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDeviceGroupDatasourceAlertSettingDefault) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{deviceGroupId}/datasources/{dsId}/alertsettings][%d] updateDeviceGroupDatasourceAlertSetting default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDeviceGroupDatasourceAlertSettingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceGroupDatasourceAlertSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
