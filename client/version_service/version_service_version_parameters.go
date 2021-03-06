// Code generated by go-swagger; DO NOT EDIT.

package version_service

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

// NewVersionServiceVersionParams creates a new VersionServiceVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVersionServiceVersionParams() *VersionServiceVersionParams {
	return &VersionServiceVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVersionServiceVersionParamsWithTimeout creates a new VersionServiceVersionParams object
// with the ability to set a timeout on a request.
func NewVersionServiceVersionParamsWithTimeout(timeout time.Duration) *VersionServiceVersionParams {
	return &VersionServiceVersionParams{
		timeout: timeout,
	}
}

// NewVersionServiceVersionParamsWithContext creates a new VersionServiceVersionParams object
// with the ability to set a context for a request.
func NewVersionServiceVersionParamsWithContext(ctx context.Context) *VersionServiceVersionParams {
	return &VersionServiceVersionParams{
		Context: ctx,
	}
}

// NewVersionServiceVersionParamsWithHTTPClient creates a new VersionServiceVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewVersionServiceVersionParamsWithHTTPClient(client *http.Client) *VersionServiceVersionParams {
	return &VersionServiceVersionParams{
		HTTPClient: client,
	}
}

/* VersionServiceVersionParams contains all the parameters to send to the API endpoint
   for the version service version operation.

   Typically these are written to a http.Request.
*/
type VersionServiceVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the version service version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionServiceVersionParams) WithDefaults() *VersionServiceVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the version service version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionServiceVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the version service version params
func (o *VersionServiceVersionParams) WithTimeout(timeout time.Duration) *VersionServiceVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the version service version params
func (o *VersionServiceVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the version service version params
func (o *VersionServiceVersionParams) WithContext(ctx context.Context) *VersionServiceVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the version service version params
func (o *VersionServiceVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the version service version params
func (o *VersionServiceVersionParams) WithHTTPClient(client *http.Client) *VersionServiceVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the version service version params
func (o *VersionServiceVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VersionServiceVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
