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

// AddWebsiteGroupReader is a Reader for the AddWebsiteGroup structure.
type AddWebsiteGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWebsiteGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddWebsiteGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddWebsiteGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddWebsiteGroupOK creates a AddWebsiteGroupOK with default headers values
func NewAddWebsiteGroupOK() *AddWebsiteGroupOK {
	return &AddWebsiteGroupOK{}
}

/* AddWebsiteGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type AddWebsiteGroupOK struct {
	Payload *models.WebsiteGroup
}

func (o *AddWebsiteGroupOK) Error() string {
	return fmt.Sprintf("[POST /website/groups][%d] addWebsiteGroupOK  %+v", 200, o.Payload)
}
func (o *AddWebsiteGroupOK) GetPayload() *models.WebsiteGroup {
	return o.Payload
}

func (o *AddWebsiteGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddWebsiteGroupDefault creates a AddWebsiteGroupDefault with default headers values
func NewAddWebsiteGroupDefault(code int) *AddWebsiteGroupDefault {
	return &AddWebsiteGroupDefault{
		_statusCode: code,
	}
}

/* AddWebsiteGroupDefault describes a response with status code -1, with default header values.

Error
*/
type AddWebsiteGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add website group default response
func (o *AddWebsiteGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddWebsiteGroupDefault) Error() string {
	return fmt.Sprintf("[POST /website/groups][%d] addWebsiteGroup default  %+v", o._statusCode, o.Payload)
}
func (o *AddWebsiteGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddWebsiteGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
