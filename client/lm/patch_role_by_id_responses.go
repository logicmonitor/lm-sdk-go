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

// PatchRoleByIDReader is a Reader for the PatchRoleByID structure.
type PatchRoleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRoleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRoleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchRoleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchRoleByIDOK creates a PatchRoleByIDOK with default headers values
func NewPatchRoleByIDOK() *PatchRoleByIDOK {
	return &PatchRoleByIDOK{}
}

/* PatchRoleByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchRoleByIDOK struct {
	Payload *models.Role
}

func (o *PatchRoleByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/roles/{id}][%d] patchRoleByIdOK  %+v", 200, o.Payload)
}
func (o *PatchRoleByIDOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *PatchRoleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoleByIDDefault creates a PatchRoleByIDDefault with default headers values
func NewPatchRoleByIDDefault(code int) *PatchRoleByIDDefault {
	return &PatchRoleByIDDefault{
		_statusCode: code,
	}
}

/* PatchRoleByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchRoleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch role by Id default response
func (o *PatchRoleByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchRoleByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/roles/{id}][%d] patchRoleById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchRoleByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchRoleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
