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

// GetUnmonitoredDeviceListReader is a Reader for the GetUnmonitoredDeviceList structure.
type GetUnmonitoredDeviceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUnmonitoredDeviceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUnmonitoredDeviceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUnmonitoredDeviceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUnmonitoredDeviceListOK creates a GetUnmonitoredDeviceListOK with default headers values
func NewGetUnmonitoredDeviceListOK() *GetUnmonitoredDeviceListOK {
	return &GetUnmonitoredDeviceListOK{}
}

/*
	GetUnmonitoredDeviceListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUnmonitoredDeviceListOK struct {
	Payload *models.UnmonitoredDevicePaginationResponse
}

func (o *GetUnmonitoredDeviceListOK) Error() string {
	return fmt.Sprintf("[GET /device/unmonitoreddevices][%d] getUnmonitoredDeviceListOK  %+v", 200, o.Payload)
}
func (o *GetUnmonitoredDeviceListOK) GetPayload() *models.UnmonitoredDevicePaginationResponse {
	return o.Payload
}

func (o *GetUnmonitoredDeviceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnmonitoredDevicePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnmonitoredDeviceListDefault creates a GetUnmonitoredDeviceListDefault with default headers values
func NewGetUnmonitoredDeviceListDefault(code int) *GetUnmonitoredDeviceListDefault {
	return &GetUnmonitoredDeviceListDefault{
		_statusCode: code,
	}
}

/*
	GetUnmonitoredDeviceListDefault describes a response with status code -1, with default header values.

Error
*/
type GetUnmonitoredDeviceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get unmonitored device list default response
func (o *GetUnmonitoredDeviceListDefault) Code() int {
	return o._statusCode
}

func (o *GetUnmonitoredDeviceListDefault) Error() string {
	return fmt.Sprintf("[GET /device/unmonitoreddevices][%d] getUnmonitoredDeviceList default  %+v", o._statusCode, o.Payload)
}
func (o *GetUnmonitoredDeviceListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUnmonitoredDeviceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
