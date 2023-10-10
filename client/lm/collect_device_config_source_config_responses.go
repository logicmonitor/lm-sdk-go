// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// CollectDeviceConfigSourceConfigReader is a Reader for the CollectDeviceConfigSourceConfig structure.
type CollectDeviceConfigSourceConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectDeviceConfigSourceConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectDeviceConfigSourceConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCollectDeviceConfigSourceConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCollectDeviceConfigSourceConfigOK creates a CollectDeviceConfigSourceConfigOK with default headers values
func NewCollectDeviceConfigSourceConfigOK() *CollectDeviceConfigSourceConfigOK {
	return &CollectDeviceConfigSourceConfigOK{}
}

/* CollectDeviceConfigSourceConfigOK describes a response with status code 200, with default header values.

successful operation
*/
type CollectDeviceConfigSourceConfigOK struct {
	Payload interface{}
}

func (o *CollectDeviceConfigSourceConfigOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config/collectNow][%d] collectDeviceConfigSourceConfigOK  %+v", 200, o.Payload)
}
func (o *CollectDeviceConfigSourceConfigOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CollectDeviceConfigSourceConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectDeviceConfigSourceConfigDefault creates a CollectDeviceConfigSourceConfigDefault with default headers values
func NewCollectDeviceConfigSourceConfigDefault(code int) *CollectDeviceConfigSourceConfigDefault {
	return &CollectDeviceConfigSourceConfigDefault{
		_statusCode: code,
	}
}

/* CollectDeviceConfigSourceConfigDefault describes a response with status code -1, with default header values.

Error
*/
type CollectDeviceConfigSourceConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the collect device config source config default response
func (o *CollectDeviceConfigSourceConfigDefault) Code() int {
	return o._statusCode
}

func (o *CollectDeviceConfigSourceConfigDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/config/collectNow][%d] collectDeviceConfigSourceConfig default  %+v", o._statusCode, o.Payload)
}
func (o *CollectDeviceConfigSourceConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CollectDeviceConfigSourceConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
