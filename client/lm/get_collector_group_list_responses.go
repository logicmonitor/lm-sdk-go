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

// GetCollectorGroupListReader is a Reader for the GetCollectorGroupList structure.
type GetCollectorGroupListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectorGroupListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectorGroupListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCollectorGroupListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCollectorGroupListOK creates a GetCollectorGroupListOK with default headers values
func NewGetCollectorGroupListOK() *GetCollectorGroupListOK {
	return &GetCollectorGroupListOK{}
}

/*
	GetCollectorGroupListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCollectorGroupListOK struct {
	Payload *models.CollectorGroupPaginationResponse
}

func (o *GetCollectorGroupListOK) Error() string {
	return fmt.Sprintf("[GET /setting/collector/groups][%d] getCollectorGroupListOK  %+v", 200, o.Payload)
}
func (o *GetCollectorGroupListOK) GetPayload() *models.CollectorGroupPaginationResponse {
	return o.Payload
}

func (o *GetCollectorGroupListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectorGroupPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectorGroupListDefault creates a GetCollectorGroupListDefault with default headers values
func NewGetCollectorGroupListDefault(code int) *GetCollectorGroupListDefault {
	return &GetCollectorGroupListDefault{
		_statusCode: code,
	}
}

/*
	GetCollectorGroupListDefault describes a response with status code -1, with default header values.

Error
*/
type GetCollectorGroupListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get collector group list default response
func (o *GetCollectorGroupListDefault) Code() int {
	return o._statusCode
}

func (o *GetCollectorGroupListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/collector/groups][%d] getCollectorGroupList default  %+v", o._statusCode, o.Payload)
}
func (o *GetCollectorGroupListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCollectorGroupListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
