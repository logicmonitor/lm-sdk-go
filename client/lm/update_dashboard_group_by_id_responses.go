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

// UpdateDashboardGroupByIDReader is a Reader for the UpdateDashboardGroupByID structure.
type UpdateDashboardGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDashboardGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateDashboardGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDashboardGroupByIDOK creates a UpdateDashboardGroupByIDOK with default headers values
func NewUpdateDashboardGroupByIDOK() *UpdateDashboardGroupByIDOK {
	return &UpdateDashboardGroupByIDOK{}
}

/* UpdateDashboardGroupByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDashboardGroupByIDOK struct {
	Payload *models.DashboardGroup
}

func (o *UpdateDashboardGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /dashboard/groups/{id}][%d] updateDashboardGroupByIdOK  %+v", 200, o.Payload)
}
func (o *UpdateDashboardGroupByIDOK) GetPayload() *models.DashboardGroup {
	return o.Payload
}

func (o *UpdateDashboardGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardGroupByIDDefault creates a UpdateDashboardGroupByIDDefault with default headers values
func NewUpdateDashboardGroupByIDDefault(code int) *UpdateDashboardGroupByIDDefault {
	return &UpdateDashboardGroupByIDDefault{
		_statusCode: code,
	}
}

/* UpdateDashboardGroupByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDashboardGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update dashboard group by Id default response
func (o *UpdateDashboardGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDashboardGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /dashboard/groups/{id}][%d] updateDashboardGroupById default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDashboardGroupByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDashboardGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
