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

// AddRecipientGroupReader is a Reader for the AddRecipientGroup structure.
type AddRecipientGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRecipientGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRecipientGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddRecipientGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddRecipientGroupOK creates a AddRecipientGroupOK with default headers values
func NewAddRecipientGroupOK() *AddRecipientGroupOK {
	return &AddRecipientGroupOK{}
}

/* AddRecipientGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type AddRecipientGroupOK struct {
	Payload *models.RecipientGroup
}

func (o *AddRecipientGroupOK) Error() string {
	return fmt.Sprintf("[POST /setting/recipientgroups][%d] addRecipientGroupOK  %+v", 200, o.Payload)
}
func (o *AddRecipientGroupOK) GetPayload() *models.RecipientGroup {
	return o.Payload
}

func (o *AddRecipientGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecipientGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRecipientGroupDefault creates a AddRecipientGroupDefault with default headers values
func NewAddRecipientGroupDefault(code int) *AddRecipientGroupDefault {
	return &AddRecipientGroupDefault{
		_statusCode: code,
	}
}

/* AddRecipientGroupDefault describes a response with status code -1, with default header values.

Error
*/
type AddRecipientGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add recipient group default response
func (o *AddRecipientGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddRecipientGroupDefault) Error() string {
	return fmt.Sprintf("[POST /setting/recipientgroups][%d] addRecipientGroup default  %+v", o._statusCode, o.Payload)
}
func (o *AddRecipientGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddRecipientGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
