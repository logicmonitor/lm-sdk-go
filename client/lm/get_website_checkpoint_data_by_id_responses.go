// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// GetWebsiteCheckpointDataByIDReader is a Reader for the GetWebsiteCheckpointDataByID structure.
type GetWebsiteCheckpointDataByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteCheckpointDataByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsiteCheckpointDataByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetWebsiteCheckpointDataByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWebsiteCheckpointDataByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteCheckpointDataByIDOK creates a GetWebsiteCheckpointDataByIDOK with default headers values
func NewGetWebsiteCheckpointDataByIDOK() *GetWebsiteCheckpointDataByIDOK {
	return &GetWebsiteCheckpointDataByIDOK{}
}

/* GetWebsiteCheckpointDataByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebsiteCheckpointDataByIDOK struct {
	Payload *models.WebsiteCheckpointRawData
}

func (o *GetWebsiteCheckpointDataByIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{srvId}/checkpoints/{checkId}/data][%d] getWebsiteCheckpointDataByIdOK  %+v", 200, o.Payload)
}
func (o *GetWebsiteCheckpointDataByIDOK) GetPayload() *models.WebsiteCheckpointRawData {
	return o.Payload
}

func (o *GetWebsiteCheckpointDataByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteCheckpointRawData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteCheckpointDataByIDTooManyRequests creates a GetWebsiteCheckpointDataByIDTooManyRequests with default headers values
func NewGetWebsiteCheckpointDataByIDTooManyRequests() *GetWebsiteCheckpointDataByIDTooManyRequests {
	return &GetWebsiteCheckpointDataByIDTooManyRequests{}
}

/* GetWebsiteCheckpointDataByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetWebsiteCheckpointDataByIDTooManyRequests struct {

	/* Request limit per X-Rate-Limit-Window
	 */
	XRateLimitLimit int64

	/* The number of requests left for the time window
	 */
	XRateLimitRemaining int64

	/* The rolling time window length with the unit of second
	 */
	XRateLimitWindow int64
}

func (o *GetWebsiteCheckpointDataByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /website/websites/{srvId}/checkpoints/{checkId}/data][%d] getWebsiteCheckpointDataByIdTooManyRequests ", 429)
}

func (o *GetWebsiteCheckpointDataByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-rate-limit-limit
	hdrXRateLimitLimit := response.GetHeader("x-rate-limit-limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("x-rate-limit-limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header x-rate-limit-remaining
	hdrXRateLimitRemaining := response.GetHeader("x-rate-limit-remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("x-rate-limit-remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header x-rate-limit-window
	hdrXRateLimitWindow := response.GetHeader("x-rate-limit-window")

	if hdrXRateLimitWindow != "" {
		valxRateLimitWindow, err := swag.ConvertInt64(hdrXRateLimitWindow)
		if err != nil {
			return errors.InvalidType("x-rate-limit-window", "header", "int64", hdrXRateLimitWindow)
		}
		o.XRateLimitWindow = valxRateLimitWindow
	}

	return nil
}

// NewGetWebsiteCheckpointDataByIDDefault creates a GetWebsiteCheckpointDataByIDDefault with default headers values
func NewGetWebsiteCheckpointDataByIDDefault(code int) *GetWebsiteCheckpointDataByIDDefault {
	return &GetWebsiteCheckpointDataByIDDefault{
		_statusCode: code,
	}
}

/* GetWebsiteCheckpointDataByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetWebsiteCheckpointDataByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website checkpoint data by Id default response
func (o *GetWebsiteCheckpointDataByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteCheckpointDataByIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{srvId}/checkpoints/{checkId}/data][%d] getWebsiteCheckpointDataById default  %+v", o._statusCode, o.Payload)
}
func (o *GetWebsiteCheckpointDataByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWebsiteCheckpointDataByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
