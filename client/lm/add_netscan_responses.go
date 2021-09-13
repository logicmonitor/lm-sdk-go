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

// AddNetscanReader is a Reader for the AddNetscan structure.
type AddNetscanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNetscanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddNetscanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddNetscanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddNetscanOK creates a AddNetscanOK with default headers values
func NewAddNetscanOK() *AddNetscanOK {
	return &AddNetscanOK{}
}

/* AddNetscanOK describes a response with status code 200, with default header values.

successful operation
*/
type AddNetscanOK struct {
	Payload models.Netscan
}

func (o *AddNetscanOK) Error() string {
	return fmt.Sprintf("[POST /setting/netscans][%d] addNetscanOK  %+v", 200, o.Payload)
}
func (o *AddNetscanOK) GetPayload() models.Netscan {
	return o.Payload
}

func (o *AddNetscanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalNetscan(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewAddNetscanDefault creates a AddNetscanDefault with default headers values
func NewAddNetscanDefault(code int) *AddNetscanDefault {
	return &AddNetscanDefault{
		_statusCode: code,
	}
}

/* AddNetscanDefault describes a response with status code -1, with default header values.

Error
*/
type AddNetscanDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add netscan default response
func (o *AddNetscanDefault) Code() int {
	return o._statusCode
}

func (o *AddNetscanDefault) Error() string {
	return fmt.Sprintf("[POST /setting/netscans][%d] addNetscan default  %+v", o._statusCode, o.Payload)
}
func (o *AddNetscanDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddNetscanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
