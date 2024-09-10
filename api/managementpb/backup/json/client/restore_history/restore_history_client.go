// Code generated by go-swagger; DO NOT EDIT.

package restore_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new restore history API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for restore history API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListRestoreHistory(params *ListRestoreHistoryParams, opts ...ClientOption) (*ListRestoreHistoryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListRestoreHistory lists restore history returns a list of all backup restore history items
*/
func (a *Client) ListRestoreHistory(params *ListRestoreHistoryParams, opts ...ClientOption) (*ListRestoreHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRestoreHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListRestoreHistory",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/RestoreHistory/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRestoreHistoryReader{formats: a.formats},
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
	success, ok := result.(*ListRestoreHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRestoreHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}