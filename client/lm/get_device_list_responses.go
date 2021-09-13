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

// GetDeviceListReader is a Reader for the GetDeviceList structure.
type GetDeviceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceListOK creates a GetDeviceListOK with default headers values
func NewGetDeviceListOK() *GetDeviceListOK {
	return &GetDeviceListOK{}
}

/* GetDeviceListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceListOK struct {
	Payload *models.DevicePaginationResponse
}

func (o *GetDeviceListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices][%d] getDeviceListOK  %+v", 200, o.Payload)
}
func (o *GetDeviceListOK) GetPayload() *models.DevicePaginationResponse {
	return o.Payload
}

func (o *GetDeviceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevicePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceListDefault creates a GetDeviceListDefault with default headers values
func NewGetDeviceListDefault(code int) *GetDeviceListDefault {
	return &GetDeviceListDefault{
		_statusCode: code,
	}
}

/* GetDeviceListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device list default response
func (o *GetDeviceListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices][%d] getDeviceList default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
