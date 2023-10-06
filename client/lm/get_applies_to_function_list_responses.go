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

// GetAppliesToFunctionListReader is a Reader for the GetAppliesToFunctionList structure.
type GetAppliesToFunctionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppliesToFunctionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppliesToFunctionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAppliesToFunctionListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppliesToFunctionListOK creates a GetAppliesToFunctionListOK with default headers values
func NewGetAppliesToFunctionListOK() *GetAppliesToFunctionListOK {
	return &GetAppliesToFunctionListOK{}
}

/*
	GetAppliesToFunctionListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAppliesToFunctionListOK struct {
	Payload *models.RestAppliesToFunctionPaginationResponse
}

func (o *GetAppliesToFunctionListOK) Error() string {
	return fmt.Sprintf("[GET /setting/functions][%d] getAppliesToFunctionListOK  %+v", 200, o.Payload)
}
func (o *GetAppliesToFunctionListOK) GetPayload() *models.RestAppliesToFunctionPaginationResponse {
	return o.Payload
}

func (o *GetAppliesToFunctionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestAppliesToFunctionPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppliesToFunctionListDefault creates a GetAppliesToFunctionListDefault with default headers values
func NewGetAppliesToFunctionListDefault(code int) *GetAppliesToFunctionListDefault {
	return &GetAppliesToFunctionListDefault{
		_statusCode: code,
	}
}

/*
	GetAppliesToFunctionListDefault describes a response with status code -1, with default header values.

Error
*/
type GetAppliesToFunctionListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get applies to function list default response
func (o *GetAppliesToFunctionListDefault) Code() int {
	return o._statusCode
}

func (o *GetAppliesToFunctionListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/functions][%d] getAppliesToFunctionList default  %+v", o._statusCode, o.Payload)
}
func (o *GetAppliesToFunctionListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAppliesToFunctionListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
