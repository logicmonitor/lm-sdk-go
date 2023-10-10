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

// PatchOpsNoteByIDReader is a Reader for the PatchOpsNoteByID structure.
type PatchOpsNoteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOpsNoteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchOpsNoteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchOpsNoteByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchOpsNoteByIDOK creates a PatchOpsNoteByIDOK with default headers values
func NewPatchOpsNoteByIDOK() *PatchOpsNoteByIDOK {
	return &PatchOpsNoteByIDOK{}
}

/* PatchOpsNoteByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchOpsNoteByIDOK struct {
	Payload *models.OpsNote
}

func (o *PatchOpsNoteByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/opsnotes/{id}][%d] patchOpsNoteByIdOK  %+v", 200, o.Payload)
}
func (o *PatchOpsNoteByIDOK) GetPayload() *models.OpsNote {
	return o.Payload
}

func (o *PatchOpsNoteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpsNote)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOpsNoteByIDDefault creates a PatchOpsNoteByIDDefault with default headers values
func NewPatchOpsNoteByIDDefault(code int) *PatchOpsNoteByIDDefault {
	return &PatchOpsNoteByIDDefault{
		_statusCode: code,
	}
}

/* PatchOpsNoteByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchOpsNoteByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch ops note by Id default response
func (o *PatchOpsNoteByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchOpsNoteByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/opsnotes/{id}][%d] patchOpsNoteById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchOpsNoteByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchOpsNoteByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
