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

// PatchWebsiteGroupByIDReader is a Reader for the PatchWebsiteGroupByID structure.
type PatchWebsiteGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWebsiteGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWebsiteGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchWebsiteGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchWebsiteGroupByIDOK creates a PatchWebsiteGroupByIDOK with default headers values
func NewPatchWebsiteGroupByIDOK() *PatchWebsiteGroupByIDOK {
	return &PatchWebsiteGroupByIDOK{}
}

/*
	PatchWebsiteGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchWebsiteGroupByIDOK struct {
	Payload *models.WebsiteGroup
}

func (o *PatchWebsiteGroupByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /website/groups/{id}][%d] patchWebsiteGroupByIdOK  %+v", 200, o.Payload)
}
func (o *PatchWebsiteGroupByIDOK) GetPayload() *models.WebsiteGroup {
	return o.Payload
}

func (o *PatchWebsiteGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWebsiteGroupByIDDefault creates a PatchWebsiteGroupByIDDefault with default headers values
func NewPatchWebsiteGroupByIDDefault(code int) *PatchWebsiteGroupByIDDefault {
	return &PatchWebsiteGroupByIDDefault{
		_statusCode: code,
	}
}

/*
	PatchWebsiteGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type PatchWebsiteGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch website group by Id default response
func (o *PatchWebsiteGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchWebsiteGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /website/groups/{id}][%d] patchWebsiteGroupById default  %+v", o._statusCode, o.Payload)
}
func (o *PatchWebsiteGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchWebsiteGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
