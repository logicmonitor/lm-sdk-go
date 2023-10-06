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

// GetDatasourceListReader is a Reader for the GetDatasourceList structure.
type GetDatasourceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatasourceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatasourceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatasourceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatasourceListOK creates a GetDatasourceListOK with default headers values
func NewGetDatasourceListOK() *GetDatasourceListOK {
	return &GetDatasourceListOK{}
}

/*
	GetDatasourceListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDatasourceListOK struct {
	Payload *models.DatasourcePaginationResponse
}

func (o *GetDatasourceListOK) Error() string {
	return fmt.Sprintf("[GET /setting/datasources][%d] getDatasourceListOK  %+v", 200, o.Payload)
}
func (o *GetDatasourceListOK) GetPayload() *models.DatasourcePaginationResponse {
	return o.Payload
}

func (o *GetDatasourceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatasourcePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasourceListDefault creates a GetDatasourceListDefault with default headers values
func NewGetDatasourceListDefault(code int) *GetDatasourceListDefault {
	return &GetDatasourceListDefault{
		_statusCode: code,
	}
}

/*
	GetDatasourceListDefault describes a response with status code -1, with default header values.

Error
*/
type GetDatasourceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get datasource list default response
func (o *GetDatasourceListDefault) Code() int {
	return o._statusCode
}

func (o *GetDatasourceListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/datasources][%d] getDatasourceList default  %+v", o._statusCode, o.Payload)
}
func (o *GetDatasourceListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDatasourceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
