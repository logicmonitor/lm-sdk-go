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

// GetWebsitePropertyListByWebsiteIDReader is a Reader for the GetWebsitePropertyListByWebsiteID structure.
type GetWebsitePropertyListByWebsiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsitePropertyListByWebsiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsitePropertyListByWebsiteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWebsitePropertyListByWebsiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsitePropertyListByWebsiteIDOK creates a GetWebsitePropertyListByWebsiteIDOK with default headers values
func NewGetWebsitePropertyListByWebsiteIDOK() *GetWebsitePropertyListByWebsiteIDOK {
	return &GetWebsitePropertyListByWebsiteIDOK{}
}

/*
	GetWebsitePropertyListByWebsiteIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsitePropertyListByWebsiteIDOK struct {
	Payload *models.PropertyPaginationResponse
}

func (o *GetWebsitePropertyListByWebsiteIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/properties][%d] getWebsitePropertyListByWebsiteIdOK  %+v", 200, o.Payload)
}
func (o *GetWebsitePropertyListByWebsiteIDOK) GetPayload() *models.PropertyPaginationResponse {
	return o.Payload
}

func (o *GetWebsitePropertyListByWebsiteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsitePropertyListByWebsiteIDDefault creates a GetWebsitePropertyListByWebsiteIDDefault with default headers values
func NewGetWebsitePropertyListByWebsiteIDDefault(code int) *GetWebsitePropertyListByWebsiteIDDefault {
	return &GetWebsitePropertyListByWebsiteIDDefault{
		_statusCode: code,
	}
}

/*
	GetWebsitePropertyListByWebsiteIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsitePropertyListByWebsiteIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website property list by website Id default response
func (o *GetWebsitePropertyListByWebsiteIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsitePropertyListByWebsiteIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/properties][%d] getWebsitePropertyListByWebsiteId default  %+v", o._statusCode, o.Payload)
}
func (o *GetWebsitePropertyListByWebsiteIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsitePropertyListByWebsiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
