// Code generated by go-swagger; DO NOT EDIT.

package collector_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"reflect"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new collector versions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
		userAgent: nil,
	}
}

/*
Client for collector versions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	userAgent *string
}

/*
GetCollectorVersionList gets collector version list
*/
func (a *Client) GetCollectorVersionList(params *GetCollectorVersionListParams) (*GetCollectorVersionListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectorVersionListParams()
	}
	reflect.ValueOf(params).MethodByName("SetUserAgent").Call([]reflect.Value{reflect.ValueOf(a.userAgent)})

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCollectorVersionList",
		Method:             "GET",
		PathPattern:        "/setting/collector/collectors/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCollectorVersionListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCollectorVersionListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// SetUserAgent changes the userAgent on the client
func (a *Client) SetUserAgent(userAgent *string) {
	a.userAgent = userAgent
}
