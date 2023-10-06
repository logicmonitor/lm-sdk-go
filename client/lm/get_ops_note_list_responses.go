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

// GetOpsNoteListReader is a Reader for the GetOpsNoteList structure.
type GetOpsNoteListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOpsNoteListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOpsNoteListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOpsNoteListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOpsNoteListOK creates a GetOpsNoteListOK with default headers values
func NewGetOpsNoteListOK() *GetOpsNoteListOK {
	return &GetOpsNoteListOK{}
}

/*
	GetOpsNoteListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOpsNoteListOK struct {
	Payload *models.OpsNotePaginationResponse
}

func (o *GetOpsNoteListOK) Error() string {
	return fmt.Sprintf("[GET /setting/opsnotes][%d] getOpsNoteListOK  %+v", 200, o.Payload)
}
func (o *GetOpsNoteListOK) GetPayload() *models.OpsNotePaginationResponse {
	return o.Payload
}

func (o *GetOpsNoteListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpsNotePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOpsNoteListDefault creates a GetOpsNoteListDefault with default headers values
func NewGetOpsNoteListDefault(code int) *GetOpsNoteListDefault {
	return &GetOpsNoteListDefault{
		_statusCode: code,
	}
}

/*
	GetOpsNoteListDefault describes a response with status code -1, with default header values.

Error
*/
type GetOpsNoteListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get ops note list default response
func (o *GetOpsNoteListDefault) Code() int {
	return o._statusCode
}

func (o *GetOpsNoteListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/opsnotes][%d] getOpsNoteList default  %+v", o._statusCode, o.Payload)
}
func (o *GetOpsNoteListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOpsNoteListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
