// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// UpdateAPITokenByAdminIDReader is a Reader for the UpdateAPITokenByAdminID structure.
type UpdateAPITokenByAdminIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAPITokenByAdminIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAPITokenByAdminIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateAPITokenByAdminIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAPITokenByAdminIDOK creates a UpdateAPITokenByAdminIDOK with default headers values
func NewUpdateAPITokenByAdminIDOK() *UpdateAPITokenByAdminIDOK {
	return &UpdateAPITokenByAdminIDOK{}
}

/*UpdateAPITokenByAdminIDOK handles this case with default header values.

successful operation
*/
type UpdateAPITokenByAdminIDOK struct {
	Payload *models.APIToken
}

func (o *UpdateAPITokenByAdminIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/admins/{adminId}/apitokens/{apitokenId}][%d] updateApiTokenByAdminIdOK  %+v", 200, o.Payload)
}

func (o *UpdateAPITokenByAdminIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPITokenByAdminIDDefault creates a UpdateAPITokenByAdminIDDefault with default headers values
func NewUpdateAPITokenByAdminIDDefault(code int) *UpdateAPITokenByAdminIDDefault {
	return &UpdateAPITokenByAdminIDDefault{
		_statusCode: code,
	}
}

/*UpdateAPITokenByAdminIDDefault handles this case with default header values.

Error
*/
type UpdateAPITokenByAdminIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update Api token by admin Id default response
func (o *UpdateAPITokenByAdminIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAPITokenByAdminIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/admins/{adminId}/apitokens/{apitokenId}][%d] updateApiTokenByAdminId default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAPITokenByAdminIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
