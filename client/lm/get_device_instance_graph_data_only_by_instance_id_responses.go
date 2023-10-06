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

// GetDeviceInstanceGraphDataOnlyByInstanceIDReader is a Reader for the GetDeviceInstanceGraphDataOnlyByInstanceID structure.
type GetDeviceInstanceGraphDataOnlyByInstanceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceInstanceGraphDataOnlyByInstanceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceInstanceGraphDataOnlyByInstanceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceInstanceGraphDataOnlyByInstanceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceInstanceGraphDataOnlyByInstanceIDOK creates a GetDeviceInstanceGraphDataOnlyByInstanceIDOK with default headers values
func NewGetDeviceInstanceGraphDataOnlyByInstanceIDOK() *GetDeviceInstanceGraphDataOnlyByInstanceIDOK {
	return &GetDeviceInstanceGraphDataOnlyByInstanceIDOK{}
}

/*
	GetDeviceInstanceGraphDataOnlyByInstanceIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDeviceInstanceGraphDataOnlyByInstanceIDOK struct {
	Payload *models.GraphPlot
}

func (o *GetDeviceInstanceGraphDataOnlyByInstanceIDOK) Error() string {
	return fmt.Sprintf("[GET /device/devicedatasourceinstances/{instanceId}/graphs/{graphId}/data][%d] getDeviceInstanceGraphDataOnlyByInstanceIdOK  %+v", 200, o.Payload)
}
func (o *GetDeviceInstanceGraphDataOnlyByInstanceIDOK) GetPayload() *models.GraphPlot {
	return o.Payload
}

func (o *GetDeviceInstanceGraphDataOnlyByInstanceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphPlot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceInstanceGraphDataOnlyByInstanceIDDefault creates a GetDeviceInstanceGraphDataOnlyByInstanceIDDefault with default headers values
func NewGetDeviceInstanceGraphDataOnlyByInstanceIDDefault(code int) *GetDeviceInstanceGraphDataOnlyByInstanceIDDefault {
	return &GetDeviceInstanceGraphDataOnlyByInstanceIDDefault{
		_statusCode: code,
	}
}

/*
	GetDeviceInstanceGraphDataOnlyByInstanceIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetDeviceInstanceGraphDataOnlyByInstanceIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device instance graph data only by instance Id default response
func (o *GetDeviceInstanceGraphDataOnlyByInstanceIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceInstanceGraphDataOnlyByInstanceIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/devicedatasourceinstances/{instanceId}/graphs/{graphId}/data][%d] getDeviceInstanceGraphDataOnlyByInstanceId default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceInstanceGraphDataOnlyByInstanceIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceInstanceGraphDataOnlyByInstanceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
