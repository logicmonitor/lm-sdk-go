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

// GetDeviceDatasourceInstanceDataReader is a Reader for the GetDeviceDatasourceInstanceData structure.
type GetDeviceDatasourceInstanceDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceDatasourceInstanceDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceDatasourceInstanceDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceDataOK creates a GetDeviceDatasourceInstanceDataOK with default headers values
func NewGetDeviceDatasourceInstanceDataOK() *GetDeviceDatasourceInstanceDataOK {
	return &GetDeviceDatasourceInstanceDataOK{}
}

/*
	GetDeviceDatasourceInstanceDataOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceDataOK struct {
	Payload *models.DeviceDataSourceInstanceData
}

func (o *GetDeviceDatasourceInstanceDataOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data][%d] getDeviceDatasourceInstanceDataOK  %+v", 200, o.Payload)
}
func (o *GetDeviceDatasourceInstanceDataOK) GetPayload() *models.DeviceDataSourceInstanceData {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstanceData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceDataDefault creates a GetDeviceDatasourceInstanceDataDefault with default headers values
func NewGetDeviceDatasourceInstanceDataDefault(code int) *GetDeviceDatasourceInstanceDataDefault {
	return &GetDeviceDatasourceInstanceDataDefault{
		_statusCode: code,
	}
}

/*
	GetDeviceDatasourceInstanceDataDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceDatasourceInstanceDataDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance data default response
func (o *GetDeviceDatasourceInstanceDataDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceDataDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/data][%d] getDeviceDatasourceInstanceData default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceDatasourceInstanceDataDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDatasourceInstanceDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
