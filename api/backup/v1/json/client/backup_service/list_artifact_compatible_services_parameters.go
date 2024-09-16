// Code generated by go-swagger; DO NOT EDIT.

package backup_service

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

// NewListArtifactCompatibleServicesParams creates a new ListArtifactCompatibleServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListArtifactCompatibleServicesParams() *ListArtifactCompatibleServicesParams {
	return &ListArtifactCompatibleServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListArtifactCompatibleServicesParamsWithTimeout creates a new ListArtifactCompatibleServicesParams object
// with the ability to set a timeout on a request.
func NewListArtifactCompatibleServicesParamsWithTimeout(timeout time.Duration) *ListArtifactCompatibleServicesParams {
	return &ListArtifactCompatibleServicesParams{
		timeout: timeout,
	}
}

// NewListArtifactCompatibleServicesParamsWithContext creates a new ListArtifactCompatibleServicesParams object
// with the ability to set a context for a request.
func NewListArtifactCompatibleServicesParamsWithContext(ctx context.Context) *ListArtifactCompatibleServicesParams {
	return &ListArtifactCompatibleServicesParams{
		Context: ctx,
	}
}

// NewListArtifactCompatibleServicesParamsWithHTTPClient creates a new ListArtifactCompatibleServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListArtifactCompatibleServicesParamsWithHTTPClient(client *http.Client) *ListArtifactCompatibleServicesParams {
	return &ListArtifactCompatibleServicesParams{
		HTTPClient: client,
	}
}

/*
ListArtifactCompatibleServicesParams contains all the parameters to send to the API endpoint

	for the list artifact compatible services operation.

	Typically these are written to a http.Request.
*/
type ListArtifactCompatibleServicesParams struct {
	/* ArtifactID.

	   Artifact id used to determine restore compatibility.
	*/
	ArtifactID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list artifact compatible services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListArtifactCompatibleServicesParams) WithDefaults() *ListArtifactCompatibleServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list artifact compatible services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListArtifactCompatibleServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list artifact compatible services params
func (o *ListArtifactCompatibleServicesParams) WithTimeout(timeout time.Duration) *ListArtifactCompatibleServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list artifact compatible services params
func (o *ListArtifactCompatibleServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list artifact compatible services params
func (o *ListArtifactCompatibleServicesParams) WithContext(ctx context.Context) *ListArtifactCompatibleServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list artifact compatible services params
func (o *ListArtifactCompatibleServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list artifact compatible services params
func (o *ListArtifactCompatibleServicesParams) WithHTTPClient(client *http.Client) *ListArtifactCompatibleServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list artifact compatible services params
func (o *ListArtifactCompatibleServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactID adds the artifactID to the list artifact compatible services params
func (o *ListArtifactCompatibleServicesParams) WithArtifactID(artifactID string) *ListArtifactCompatibleServicesParams {
	o.SetArtifactID(artifactID)
	return o
}

// SetArtifactID adds the artifactId to the list artifact compatible services params
func (o *ListArtifactCompatibleServicesParams) SetArtifactID(artifactID string) {
	o.ArtifactID = artifactID
}

// WriteToRequest writes these params to a swagger request
func (o *ListArtifactCompatibleServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param artifact_id
	if err := r.SetPathParam("artifact_id", o.ArtifactID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}