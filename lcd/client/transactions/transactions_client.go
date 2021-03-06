// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new transactions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transactions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetTxs(params *GetTxsParams) (*GetTxsOK, error)

	GetTxsHash(params *GetTxsHashParams) (*GetTxsHashOK, error)

	PostTxs(params *PostTxsParams) (*PostTxsOK, error)

	PostTxsEncode(params *PostTxsEncodeParams) (*PostTxsEncodeOK, error)

	PostTxsEstimateFee(params *PostTxsEstimateFeeParams) (*PostTxsEstimateFeeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTxs searches transactions

  Search transactions by events.
*/
func (a *Client) GetTxs(params *GetTxsParams) (*GetTxsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTxsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTxs",
		Method:             "GET",
		PathPattern:        "/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTxsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTxsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTxs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTxsHash gets a tx by hash

  Retrieve a transaction using its hash.
*/
func (a *Client) GetTxsHash(params *GetTxsHashParams) (*GetTxsHashOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTxsHashParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTxsHash",
		Method:             "GET",
		PathPattern:        "/txs/{hash}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTxsHashReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTxsHashOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTxsHash: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTxs broadcasts a signed tx

  Broadcast a signed tx to a full node
*/
func (a *Client) PostTxs(params *PostTxsParams) (*PostTxsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTxsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTxs",
		Method:             "POST",
		PathPattern:        "/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostTxsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTxsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTxs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTxsEncode encodes a transaction to the amino wire format

  Encode a transaction (signed or not) from JSON to base64-encoded Amino serialized bytes
*/
func (a *Client) PostTxsEncode(params *PostTxsEncodeParams) (*PostTxsEncodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTxsEncodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTxsEncode",
		Method:             "POST",
		PathPattern:        "/txs/encode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostTxsEncodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTxsEncodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTxsEncode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTxsEstimateFee estimates fees and gas of a transaction

  Estimate fees and gas of a transaction according to given parameters
*/
func (a *Client) PostTxsEstimateFee(params *PostTxsEstimateFeeParams) (*PostTxsEstimateFeeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTxsEstimateFeeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTxsEstimateFee",
		Method:             "POST",
		PathPattern:        "/txs/estimate_fee",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostTxsEstimateFeeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTxsEstimateFeeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTxsEstimateFee: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
