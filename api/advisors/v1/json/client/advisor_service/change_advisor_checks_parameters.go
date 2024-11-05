// Code generated by go-swagger; DO NOT EDIT.

package advisor_service

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

// NewChangeAdvisorChecksParams creates a new ChangeAdvisorChecksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeAdvisorChecksParams() *ChangeAdvisorChecksParams {
	return &ChangeAdvisorChecksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeAdvisorChecksParamsWithTimeout creates a new ChangeAdvisorChecksParams object
// with the ability to set a timeout on a request.
func NewChangeAdvisorChecksParamsWithTimeout(timeout time.Duration) *ChangeAdvisorChecksParams {
	return &ChangeAdvisorChecksParams{
		timeout: timeout,
	}
}

// NewChangeAdvisorChecksParamsWithContext creates a new ChangeAdvisorChecksParams object
// with the ability to set a context for a request.
func NewChangeAdvisorChecksParamsWithContext(ctx context.Context) *ChangeAdvisorChecksParams {
	return &ChangeAdvisorChecksParams{
		Context: ctx,
	}
}

// NewChangeAdvisorChecksParamsWithHTTPClient creates a new ChangeAdvisorChecksParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeAdvisorChecksParamsWithHTTPClient(client *http.Client) *ChangeAdvisorChecksParams {
	return &ChangeAdvisorChecksParams{
		HTTPClient: client,
	}
}

/*
ChangeAdvisorChecksParams contains all the parameters to send to the API endpoint

	for the change advisor checks operation.

	Typically these are written to a http.Request.
*/
type ChangeAdvisorChecksParams struct {
	// Body.
	Body ChangeAdvisorChecksBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change advisor checks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeAdvisorChecksParams) WithDefaults() *ChangeAdvisorChecksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change advisor checks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeAdvisorChecksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change advisor checks params
func (o *ChangeAdvisorChecksParams) WithTimeout(timeout time.Duration) *ChangeAdvisorChecksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change advisor checks params
func (o *ChangeAdvisorChecksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change advisor checks params
func (o *ChangeAdvisorChecksParams) WithContext(ctx context.Context) *ChangeAdvisorChecksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change advisor checks params
func (o *ChangeAdvisorChecksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change advisor checks params
func (o *ChangeAdvisorChecksParams) WithHTTPClient(client *http.Client) *ChangeAdvisorChecksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change advisor checks params
func (o *ChangeAdvisorChecksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change advisor checks params
func (o *ChangeAdvisorChecksParams) WithBody(body ChangeAdvisorChecksBody) *ChangeAdvisorChecksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change advisor checks params
func (o *ChangeAdvisorChecksParams) SetBody(body ChangeAdvisorChecksBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeAdvisorChecksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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