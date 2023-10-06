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

// GetAllSDTListByDeviceIDReader is a Reader for the GetAllSDTListByDeviceID structure.
type GetAllSDTListByDeviceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSDTListByDeviceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllSDTListByDeviceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAllSDTListByDeviceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllSDTListByDeviceIDOK creates a GetAllSDTListByDeviceIDOK with default headers values
func NewGetAllSDTListByDeviceIDOK() *GetAllSDTListByDeviceIDOK {
	return &GetAllSDTListByDeviceIDOK{}
}

/*
	GetAllSDTListByDeviceIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllSDTListByDeviceIDOK struct {
	Payload *models.SDTPaginationResponse
}

func (o *GetAllSDTListByDeviceIDOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/sdts][%d] getAllSdtListByDeviceIdOK  %+v", 200, o.Payload)
}
func (o *GetAllSDTListByDeviceIDOK) GetPayload() *models.SDTPaginationResponse {
	return o.Payload
}

func (o *GetAllSDTListByDeviceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDTPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSDTListByDeviceIDDefault creates a GetAllSDTListByDeviceIDDefault with default headers values
func NewGetAllSDTListByDeviceIDDefault(code int) *GetAllSDTListByDeviceIDDefault {
	return &GetAllSDTListByDeviceIDDefault{
		_statusCode: code,
	}
}

/*
	GetAllSDTListByDeviceIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAllSDTListByDeviceIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get all SDT list by device Id default response
func (o *GetAllSDTListByDeviceIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAllSDTListByDeviceIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/sdts][%d] getAllSDTListByDeviceId default  %+v", o._statusCode, o.Payload)
}
func (o *GetAllSDTListByDeviceIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllSDTListByDeviceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
