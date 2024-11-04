// Code generated by go-swagger; DO NOT EDIT.

package ha_proxy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ha proxy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ha proxy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddHAProxy(params *AddHAProxyParams, opts ...ClientOption) (*AddHAProxyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddHAProxy adds HA proxy

Adds HAProxy service and external exporter. It automatically adds a service to inventory, which is running on the provided "node_id", then adds an "external exporter" agent to the inventory.
*/
func (a *Client) AddHAProxy(params *AddHAProxyParams, opts ...ClientOption) (*AddHAProxyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddHAProxyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddHAProxy",
		Method:             "POST",
		PathPattern:        "/v1/management/HAProxy/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddHAProxyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddHAProxyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddHAProxyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}