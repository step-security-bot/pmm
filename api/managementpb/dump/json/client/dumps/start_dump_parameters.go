// Code generated by go-swagger; DO NOT EDIT.

package dumps

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

// NewStartDumpParams creates a new StartDumpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartDumpParams() *StartDumpParams {
	return &StartDumpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartDumpParamsWithTimeout creates a new StartDumpParams object
// with the ability to set a timeout on a request.
func NewStartDumpParamsWithTimeout(timeout time.Duration) *StartDumpParams {
	return &StartDumpParams{
		timeout: timeout,
	}
}

// NewStartDumpParamsWithContext creates a new StartDumpParams object
// with the ability to set a context for a request.
func NewStartDumpParamsWithContext(ctx context.Context) *StartDumpParams {
	return &StartDumpParams{
		Context: ctx,
	}
}

// NewStartDumpParamsWithHTTPClient creates a new StartDumpParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartDumpParamsWithHTTPClient(client *http.Client) *StartDumpParams {
	return &StartDumpParams{
		HTTPClient: client,
	}
}

/*
StartDumpParams contains all the parameters to send to the API endpoint

	for the start dump operation.

	Typically these are written to a http.Request.
*/
type StartDumpParams struct {
	// Body.
	Body StartDumpBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start dump params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartDumpParams) WithDefaults() *StartDumpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start dump params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartDumpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start dump params
func (o *StartDumpParams) WithTimeout(timeout time.Duration) *StartDumpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start dump params
func (o *StartDumpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start dump params
func (o *StartDumpParams) WithContext(ctx context.Context) *StartDumpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start dump params
func (o *StartDumpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start dump params
func (o *StartDumpParams) WithHTTPClient(client *http.Client) *StartDumpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start dump params
func (o *StartDumpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start dump params
func (o *StartDumpParams) WithBody(body StartDumpBody) *StartDumpParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start dump params
func (o *StartDumpParams) SetBody(body StartDumpBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StartDumpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}