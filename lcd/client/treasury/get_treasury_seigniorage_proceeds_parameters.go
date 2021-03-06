// Code generated by go-swagger; DO NOT EDIT.

package treasury

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetTreasurySeigniorageProceedsParams creates a new GetTreasurySeigniorageProceedsParams object
// with the default values initialized.
func NewGetTreasurySeigniorageProceedsParams() *GetTreasurySeigniorageProceedsParams {

	return &GetTreasurySeigniorageProceedsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTreasurySeigniorageProceedsParamsWithTimeout creates a new GetTreasurySeigniorageProceedsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTreasurySeigniorageProceedsParamsWithTimeout(timeout time.Duration) *GetTreasurySeigniorageProceedsParams {

	return &GetTreasurySeigniorageProceedsParams{

		timeout: timeout,
	}
}

// NewGetTreasurySeigniorageProceedsParamsWithContext creates a new GetTreasurySeigniorageProceedsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTreasurySeigniorageProceedsParamsWithContext(ctx context.Context) *GetTreasurySeigniorageProceedsParams {

	return &GetTreasurySeigniorageProceedsParams{

		Context: ctx,
	}
}

// NewGetTreasurySeigniorageProceedsParamsWithHTTPClient creates a new GetTreasurySeigniorageProceedsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTreasurySeigniorageProceedsParamsWithHTTPClient(client *http.Client) *GetTreasurySeigniorageProceedsParams {

	return &GetTreasurySeigniorageProceedsParams{
		HTTPClient: client,
	}
}

/*GetTreasurySeigniorageProceedsParams contains all the parameters to send to the API endpoint
for the get treasury seigniorage proceeds operation typically these are written to a http.Request
*/
type GetTreasurySeigniorageProceedsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get treasury seigniorage proceeds params
func (o *GetTreasurySeigniorageProceedsParams) WithTimeout(timeout time.Duration) *GetTreasurySeigniorageProceedsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get treasury seigniorage proceeds params
func (o *GetTreasurySeigniorageProceedsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get treasury seigniorage proceeds params
func (o *GetTreasurySeigniorageProceedsParams) WithContext(ctx context.Context) *GetTreasurySeigniorageProceedsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get treasury seigniorage proceeds params
func (o *GetTreasurySeigniorageProceedsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get treasury seigniorage proceeds params
func (o *GetTreasurySeigniorageProceedsParams) WithHTTPClient(client *http.Client) *GetTreasurySeigniorageProceedsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get treasury seigniorage proceeds params
func (o *GetTreasurySeigniorageProceedsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTreasurySeigniorageProceedsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
