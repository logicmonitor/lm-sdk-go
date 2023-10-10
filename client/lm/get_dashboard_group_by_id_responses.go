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

// GetDashboardGroupByIDReader is a Reader for the GetDashboardGroupByID structure.
type GetDashboardGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDashboardGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDashboardGroupByIDOK creates a GetDashboardGroupByIDOK with default headers values
func NewGetDashboardGroupByIDOK() *GetDashboardGroupByIDOK {
	return &GetDashboardGroupByIDOK{}
}

/* GetDashboardGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDashboardGroupByIDOK struct {
	Payload *models.DashboardGroup
}

func (o *GetDashboardGroupByIDOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/groups/{id}][%d] getDashboardGroupByIdOK  %+v", 200, o.Payload)
}
func (o *GetDashboardGroupByIDOK) GetPayload() *models.DashboardGroup {
	return o.Payload
}

func (o *GetDashboardGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardGroupByIDDefault creates a GetDashboardGroupByIDDefault with default headers values
func NewGetDashboardGroupByIDDefault(code int) *GetDashboardGroupByIDDefault {
	return &GetDashboardGroupByIDDefault{
		_statusCode: code,
	}
}

/* GetDashboardGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetDashboardGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get dashboard group by Id default response
func (o *GetDashboardGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDashboardGroupByIDDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/groups/{id}][%d] getDashboardGroupById default  %+v", o._statusCode, o.Payload)
}
func (o *GetDashboardGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDashboardGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
