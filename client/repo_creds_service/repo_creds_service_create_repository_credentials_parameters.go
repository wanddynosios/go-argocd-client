// Code generated by go-swagger; DO NOT EDIT.

package repo_creds_service

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

	"github.com/wanddynosios/go-argocd-client/models"
)

// NewRepoCredsServiceCreateRepositoryCredentialsParams creates a new RepoCredsServiceCreateRepositoryCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoCredsServiceCreateRepositoryCredentialsParams() *RepoCredsServiceCreateRepositoryCredentialsParams {
	return &RepoCredsServiceCreateRepositoryCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoCredsServiceCreateRepositoryCredentialsParamsWithTimeout creates a new RepoCredsServiceCreateRepositoryCredentialsParams object
// with the ability to set a timeout on a request.
func NewRepoCredsServiceCreateRepositoryCredentialsParamsWithTimeout(timeout time.Duration) *RepoCredsServiceCreateRepositoryCredentialsParams {
	return &RepoCredsServiceCreateRepositoryCredentialsParams{
		timeout: timeout,
	}
}

// NewRepoCredsServiceCreateRepositoryCredentialsParamsWithContext creates a new RepoCredsServiceCreateRepositoryCredentialsParams object
// with the ability to set a context for a request.
func NewRepoCredsServiceCreateRepositoryCredentialsParamsWithContext(ctx context.Context) *RepoCredsServiceCreateRepositoryCredentialsParams {
	return &RepoCredsServiceCreateRepositoryCredentialsParams{
		Context: ctx,
	}
}

// NewRepoCredsServiceCreateRepositoryCredentialsParamsWithHTTPClient creates a new RepoCredsServiceCreateRepositoryCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoCredsServiceCreateRepositoryCredentialsParamsWithHTTPClient(client *http.Client) *RepoCredsServiceCreateRepositoryCredentialsParams {
	return &RepoCredsServiceCreateRepositoryCredentialsParams{
		HTTPClient: client,
	}
}

/* RepoCredsServiceCreateRepositoryCredentialsParams contains all the parameters to send to the API endpoint
   for the repo creds service create repository credentials operation.

   Typically these are written to a http.Request.
*/
type RepoCredsServiceCreateRepositoryCredentialsParams struct {

	/* Body.

	   Repository definition
	*/
	Body *models.V1alpha1RepoCreds

	/* Upsert.

	   Whether to create in upsert mode.
	*/
	Upsert *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo creds service create repository credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) WithDefaults() *RepoCredsServiceCreateRepositoryCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo creds service create repository credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) WithTimeout(timeout time.Duration) *RepoCredsServiceCreateRepositoryCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) WithContext(ctx context.Context) *RepoCredsServiceCreateRepositoryCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) WithHTTPClient(client *http.Client) *RepoCredsServiceCreateRepositoryCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) WithBody(body *models.V1alpha1RepoCreds) *RepoCredsServiceCreateRepositoryCredentialsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) SetBody(body *models.V1alpha1RepoCreds) {
	o.Body = body
}

// WithUpsert adds the upsert to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) WithUpsert(upsert *bool) *RepoCredsServiceCreateRepositoryCredentialsParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the repo creds service create repository credentials params
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *RepoCredsServiceCreateRepositoryCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Upsert != nil {

		// query param upsert
		var qrUpsert bool

		if o.Upsert != nil {
			qrUpsert = *o.Upsert
		}
		qUpsert := swag.FormatBool(qrUpsert)
		if qUpsert != "" {

			if err := r.SetQueryParam("upsert", qUpsert); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}