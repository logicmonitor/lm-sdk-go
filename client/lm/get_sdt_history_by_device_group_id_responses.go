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

// GetSDTHistoryByDeviceGroupIDReader is a Reader for the GetSDTHistoryByDeviceGroupID structure.
type GetSDTHistoryByDeviceGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSDTHistoryByDeviceGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSDTHistoryByDeviceGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSDTHistoryByDeviceGroupIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSDTHistoryByDeviceGroupIDOK creates a GetSDTHistoryByDeviceGroupIDOK with default headers values
func NewGetSDTHistoryByDeviceGroupIDOK() *GetSDTHistoryByDeviceGroupIDOK {
	return &GetSDTHistoryByDeviceGroupIDOK{}
}

/*
	GetSDTHistoryByDeviceGroupIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSDTHistoryByDeviceGroupIDOK struct {
	Payload *models.DeviceGroupSDTHistoryPaginationResponse
}

func (o *GetSDTHistoryByDeviceGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{id}/historysdts][%d] getSdtHistoryByDeviceGroupIdOK  %+v", 200, o.Payload)
}
func (o *GetSDTHistoryByDeviceGroupIDOK) GetPayload() *models.DeviceGroupSDTHistoryPaginationResponse {
	return o.Payload
}

func (o *GetSDTHistoryByDeviceGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroupSDTHistoryPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDTHistoryByDeviceGroupIDDefault creates a GetSDTHistoryByDeviceGroupIDDefault with default headers values
func NewGetSDTHistoryByDeviceGroupIDDefault(code int) *GetSDTHistoryByDeviceGroupIDDefault {
	return &GetSDTHistoryByDeviceGroupIDDefault{
		_statusCode: code,
	}
}

/*
	GetSDTHistoryByDeviceGroupIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetSDTHistoryByDeviceGroupIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get SDT history by device group Id default response
func (o *GetSDTHistoryByDeviceGroupIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSDTHistoryByDeviceGroupIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{id}/historysdts][%d] getSDTHistoryByDeviceGroupId default  %+v", o._statusCode, o.Payload)
}
func (o *GetSDTHistoryByDeviceGroupIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSDTHistoryByDeviceGroupIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
