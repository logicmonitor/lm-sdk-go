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

// DeleteOpsNoteByIDReader is a Reader for the DeleteOpsNoteByID structure.
type DeleteOpsNoteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOpsNoteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOpsNoteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteOpsNoteByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOpsNoteByIDOK creates a DeleteOpsNoteByIDOK with default headers values
func NewDeleteOpsNoteByIDOK() *DeleteOpsNoteByIDOK {
	return &DeleteOpsNoteByIDOK{}
}

/* DeleteOpsNoteByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteOpsNoteByIDOK struct {
	Payload interface{}
}

func (o *DeleteOpsNoteByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/opsnotes/{id}][%d] deleteOpsNoteByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteOpsNoteByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOpsNoteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOpsNoteByIDDefault creates a DeleteOpsNoteByIDDefault with default headers values
func NewDeleteOpsNoteByIDDefault(code int) *DeleteOpsNoteByIDDefault {
	return &DeleteOpsNoteByIDDefault{
		_statusCode: code,
	}
}

/* DeleteOpsNoteByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteOpsNoteByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete ops note by Id default response
func (o *DeleteOpsNoteByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOpsNoteByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/opsnotes/{id}][%d] deleteOpsNoteById default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteOpsNoteByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteOpsNoteByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
