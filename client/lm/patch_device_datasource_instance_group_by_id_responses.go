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

// PatchDeviceDatasourceInstanceGroupByIDReader is a Reader for the PatchDeviceDatasourceInstanceGroupByID structure.
type PatchDeviceDatasourceInstanceGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDeviceDatasourceInstanceGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDeviceDatasourceInstanceGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchDeviceDatasourceInstanceGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDeviceDatasourceInstanceGroupByIDOK creates a PatchDeviceDatasourceInstanceGroupByIDOK with default headers values
func NewPatchDeviceDatasourceInstanceGroupByIDOK() *PatchDeviceDatasourceInstanceGroupByIDOK {
	return &PatchDeviceDatasourceInstanceGroupByIDOK{}
}

/*
	PatchDeviceDatasourceInstanceGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchDeviceDatasourceInstanceGroupByIDOK struct {
	Payload *models.DeviceDataSourceInstanceGroup
}

func (o *PatchDeviceDatasourceInstanceGroupByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{id}][%d] patchDeviceDatasourceInstanceGroupByIdOK  %+v", 200, o.Payload)
}
func (o *PatchDeviceDatasourceInstanceGroupByIDOK) GetPayload() *models.DeviceDataSourceInstanceGroup {
	return o.Payload
}

func (o *PatchDeviceDatasourceInstanceGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstanceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceDatasourceInstanceGroupByIDDefault creates a PatchDeviceDatasourceInstanceGroupByIDDefault with default headers values
func NewPatchDeviceDatasourceInstanceGroupByIDDefault(code int) *PatchDeviceDatasourceInstanceGroupByIDDefault {
	return &PatchDeviceDatasourceInstanceGroupByIDDefault{
		_statusCode: code,
	}
}

/*
	PatchDeviceDatasourceInstanceGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchDeviceDatasourceInstanceGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch device datasource instance group by Id default response
func (o *PatchDeviceDatasourceInstanceGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchDeviceDatasourceInstanceGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{id}][%d] patchDeviceDatasourceInstanceGroupById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchDeviceDatasourceInstanceGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchDeviceDatasourceInstanceGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
