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

// GetDeviceGroupPropertyListReader is a Reader for the GetDeviceGroupPropertyList structure.
type GetDeviceGroupPropertyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceGroupPropertyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceGroupPropertyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceGroupPropertyListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceGroupPropertyListOK creates a GetDeviceGroupPropertyListOK with default headers values
func NewGetDeviceGroupPropertyListOK() *GetDeviceGroupPropertyListOK {
	return &GetDeviceGroupPropertyListOK{}
}

/*
	GetDeviceGroupPropertyListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceGroupPropertyListOK struct {
	Payload *models.PropertyPaginationResponse
}

func (o *GetDeviceGroupPropertyListOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{gid}/properties][%d] getDeviceGroupPropertyListOK  %+v", 200, o.Payload)
}
func (o *GetDeviceGroupPropertyListOK) GetPayload() *models.PropertyPaginationResponse {
	return o.Payload
}

func (o *GetDeviceGroupPropertyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceGroupPropertyListDefault creates a GetDeviceGroupPropertyListDefault with default headers values
func NewGetDeviceGroupPropertyListDefault(code int) *GetDeviceGroupPropertyListDefault {
	return &GetDeviceGroupPropertyListDefault{
		_statusCode: code,
	}
}

/*
	GetDeviceGroupPropertyListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceGroupPropertyListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device group property list default response
func (o *GetDeviceGroupPropertyListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceGroupPropertyListDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{gid}/properties][%d] getDeviceGroupPropertyList default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceGroupPropertyListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceGroupPropertyListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
