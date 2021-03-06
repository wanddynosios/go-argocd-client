// Code generated by go-swagger; DO NOT EDIT.

package repository_service

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
	"github.com/go-openapi/swag"
)

// NewRepositoryServiceGetHelmChartsParams creates a new RepositoryServiceGetHelmChartsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepositoryServiceGetHelmChartsParams() *RepositoryServiceGetHelmChartsParams {
	return &RepositoryServiceGetHelmChartsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryServiceGetHelmChartsParamsWithTimeout creates a new RepositoryServiceGetHelmChartsParams object
// with the ability to set a timeout on a request.
func NewRepositoryServiceGetHelmChartsParamsWithTimeout(timeout time.Duration) *RepositoryServiceGetHelmChartsParams {
	return &RepositoryServiceGetHelmChartsParams{
		timeout: timeout,
	}
}

// NewRepositoryServiceGetHelmChartsParamsWithContext creates a new RepositoryServiceGetHelmChartsParams object
// with the ability to set a context for a request.
func NewRepositoryServiceGetHelmChartsParamsWithContext(ctx context.Context) *RepositoryServiceGetHelmChartsParams {
	return &RepositoryServiceGetHelmChartsParams{
		Context: ctx,
	}
}

// NewRepositoryServiceGetHelmChartsParamsWithHTTPClient creates a new RepositoryServiceGetHelmChartsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepositoryServiceGetHelmChartsParamsWithHTTPClient(client *http.Client) *RepositoryServiceGetHelmChartsParams {
	return &RepositoryServiceGetHelmChartsParams{
		HTTPClient: client,
	}
}

/* RepositoryServiceGetHelmChartsParams contains all the parameters to send to the API endpoint
   for the repository service get helm charts operation.

   Typically these are written to a http.Request.
*/
type RepositoryServiceGetHelmChartsParams struct {

	/* ForceRefresh.

	   Whether to force a cache refresh on repo's connection state.
	*/
	ForceRefresh *bool

	/* Repo.

	   Repo URL for query
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repository service get helm charts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceGetHelmChartsParams) WithDefaults() *RepositoryServiceGetHelmChartsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repository service get helm charts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceGetHelmChartsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) WithTimeout(timeout time.Duration) *RepositoryServiceGetHelmChartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) WithContext(ctx context.Context) *RepositoryServiceGetHelmChartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) WithHTTPClient(client *http.Client) *RepositoryServiceGetHelmChartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceRefresh adds the forceRefresh to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) WithForceRefresh(forceRefresh *bool) *RepositoryServiceGetHelmChartsParams {
	o.SetForceRefresh(forceRefresh)
	return o
}

// SetForceRefresh adds the forceRefresh to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) SetForceRefresh(forceRefresh *bool) {
	o.ForceRefresh = forceRefresh
}

// WithRepo adds the repo to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) WithRepo(repo string) *RepositoryServiceGetHelmChartsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repository service get helm charts params
func (o *RepositoryServiceGetHelmChartsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryServiceGetHelmChartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceRefresh != nil {

		// query param forceRefresh
		var qrForceRefresh bool

		if o.ForceRefresh != nil {
			qrForceRefresh = *o.ForceRefresh
		}
		qForceRefresh := swag.FormatBool(qrForceRefresh)
		if qForceRefresh != "" {

			if err := r.SetQueryParam("forceRefresh", qForceRefresh); err != nil {
				return err
			}
		}
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
