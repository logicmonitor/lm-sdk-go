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

// GetWidgetListReader is a Reader for the GetWidgetList structure.
type GetWidgetListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWidgetListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWidgetListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWidgetListOK creates a GetWidgetListOK with default headers values
func NewGetWidgetListOK() *GetWidgetListOK {
	return &GetWidgetListOK{}
}

/* GetWidgetListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWidgetListOK struct {
	Payload *models.WidgetPaginationResponse
}

func (o *GetWidgetListOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/widgets][%d] getWidgetListOK  %+v", 200, o.Payload)
}
func (o *GetWidgetListOK) GetPayload() *models.WidgetPaginationResponse {
	return o.Payload
}

func (o *GetWidgetListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WidgetPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetListDefault creates a GetWidgetListDefault with default headers values
func NewGetWidgetListDefault(code int) *GetWidgetListDefault {
	return &GetWidgetListDefault{
		_statusCode: code,
	}
}

/* GetWidgetListDefault describes a response with status code -1, with default header values.

Error
*/
type GetWidgetListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get widget list default response
func (o *GetWidgetListDefault) Code() int {
	return o._statusCode
}

func (o *GetWidgetListDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/widgets][%d] getWidgetList default  %+v", o._statusCode, o.Payload)
}
func (o *GetWidgetListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWidgetListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
