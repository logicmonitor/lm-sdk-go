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

// DeleteAPITokenByIDReader is a Reader for the DeleteAPITokenByID structure.
type DeleteAPITokenByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPITokenByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPITokenByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAPITokenByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAPITokenByIDOK creates a DeleteAPITokenByIDOK with default headers values
func NewDeleteAPITokenByIDOK() *DeleteAPITokenByIDOK {
	return &DeleteAPITokenByIDOK{}
}

/* DeleteAPITokenByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteAPITokenByIDOK struct {
	Payload interface{}
}

func (o *DeleteAPITokenByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/admins/{adminId}/apitokens/{apitokenId}][%d] deleteApiTokenByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteAPITokenByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAPITokenByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPITokenByIDDefault creates a DeleteAPITokenByIDDefault with default headers values
func NewDeleteAPITokenByIDDefault(code int) *DeleteAPITokenByIDDefault {
	return &DeleteAPITokenByIDDefault{
		_statusCode: code,
	}
}

/* DeleteAPITokenByIDDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteAPITokenByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete Api token by Id default response
func (o *DeleteAPITokenByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAPITokenByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/admins/{adminId}/apitokens/{apitokenId}][%d] deleteApiTokenById default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteAPITokenByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteAPITokenByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
