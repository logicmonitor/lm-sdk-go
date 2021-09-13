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

// PatchDeviceGroupByIDReader is a Reader for the PatchDeviceGroupByID structure.
type PatchDeviceGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDeviceGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDeviceGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchDeviceGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDeviceGroupByIDOK creates a PatchDeviceGroupByIDOK with default headers values
func NewPatchDeviceGroupByIDOK() *PatchDeviceGroupByIDOK {
	return &PatchDeviceGroupByIDOK{}
}

/* PatchDeviceGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchDeviceGroupByIDOK struct {
	Payload *models.DeviceGroup
}

func (o *PatchDeviceGroupByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /device/groups/{id}][%d] patchDeviceGroupByIdOK  %+v", 200, o.Payload)
}
func (o *PatchDeviceGroupByIDOK) GetPayload() *models.DeviceGroup {
	return o.Payload
}

func (o *PatchDeviceGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceGroupByIDDefault creates a PatchDeviceGroupByIDDefault with default headers values
func NewPatchDeviceGroupByIDDefault(code int) *PatchDeviceGroupByIDDefault {
	return &PatchDeviceGroupByIDDefault{
		_statusCode: code,
	}
}

/* PatchDeviceGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchDeviceGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch device group by Id default response
func (o *PatchDeviceGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchDeviceGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /device/groups/{id}][%d] patchDeviceGroupById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchDeviceGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchDeviceGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
