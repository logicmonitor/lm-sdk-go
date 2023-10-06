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

// PatchAdminByIDReader is a Reader for the PatchAdminByID structure.
type PatchAdminByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAdminByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAdminByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchAdminByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAdminByIDOK creates a PatchAdminByIDOK with default headers values
func NewPatchAdminByIDOK() *PatchAdminByIDOK {
	return &PatchAdminByIDOK{}
}

/*
	PatchAdminByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchAdminByIDOK struct {
	Payload *models.Admin
}

func (o *PatchAdminByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/admins/{id}][%d] patchAdminByIdOK  %+v", 200, o.Payload)
}
func (o *PatchAdminByIDOK) GetPayload() *models.Admin {
	return o.Payload
}

func (o *PatchAdminByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Admin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAdminByIDDefault creates a PatchAdminByIDDefault with default headers values
func NewPatchAdminByIDDefault(code int) *PatchAdminByIDDefault {
	return &PatchAdminByIDDefault{
		_statusCode: code,
	}
}

/*
	PatchAdminByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchAdminByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch admin by Id default response
func (o *PatchAdminByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchAdminByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/admins/{id}][%d] patchAdminById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchAdminByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchAdminByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
