// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePool(params *DeletePoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	GetPool(params *GetPoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPoolOK, error)

	ListPools(params *ListPoolsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPoolsOK, error)

	UpdatePool(params *UpdatePoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePoolOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeletePool deletes pool by ID
*/
func (a *Client) DeletePool(params *DeletePoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePool",
		Method:             "DELETE",
		PathPattern:        "/pools/{poolID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetPool gets pool by ID
*/
func (a *Client) GetPool(params *GetPoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPool",
		Method:             "GET",
		PathPattern:        "/pools/{poolID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPoolDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListPools lists all pools
*/
func (a *Client) ListPools(params *ListPoolsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPoolsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListPools",
		Method:             "GET",
		PathPattern:        "/pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*ListPoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListPoolsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdatePool updates pool by ID
*/
func (a *Client) UpdatePool(params *UpdatePoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdatePool",
		Method:             "PUT",
		PathPattern:        "/pools/{poolID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdatePoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePoolDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}