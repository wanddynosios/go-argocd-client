// Code generated by go-swagger; DO NOT EDIT.

package application_service

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

	"github.com/wanddynosios/go-argocd-client/models"
)

// NewApplicationServiceSyncParams creates a new ApplicationServiceSyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationServiceSyncParams() *ApplicationServiceSyncParams {
	return &ApplicationServiceSyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationServiceSyncParamsWithTimeout creates a new ApplicationServiceSyncParams object
// with the ability to set a timeout on a request.
func NewApplicationServiceSyncParamsWithTimeout(timeout time.Duration) *ApplicationServiceSyncParams {
	return &ApplicationServiceSyncParams{
		timeout: timeout,
	}
}

// NewApplicationServiceSyncParamsWithContext creates a new ApplicationServiceSyncParams object
// with the ability to set a context for a request.
func NewApplicationServiceSyncParamsWithContext(ctx context.Context) *ApplicationServiceSyncParams {
	return &ApplicationServiceSyncParams{
		Context: ctx,
	}
}

// NewApplicationServiceSyncParamsWithHTTPClient creates a new ApplicationServiceSyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationServiceSyncParamsWithHTTPClient(client *http.Client) *ApplicationServiceSyncParams {
	return &ApplicationServiceSyncParams{
		HTTPClient: client,
	}
}

/* ApplicationServiceSyncParams contains all the parameters to send to the API endpoint
   for the application service sync operation.

   Typically these are written to a http.Request.
*/
type ApplicationServiceSyncParams struct {

	// Body.
	Body *models.ApplicationApplicationSyncRequest

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application service sync params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceSyncParams) WithDefaults() *ApplicationServiceSyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application service sync params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceSyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application service sync params
func (o *ApplicationServiceSyncParams) WithTimeout(timeout time.Duration) *ApplicationServiceSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application service sync params
func (o *ApplicationServiceSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application service sync params
func (o *ApplicationServiceSyncParams) WithContext(ctx context.Context) *ApplicationServiceSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application service sync params
func (o *ApplicationServiceSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application service sync params
func (o *ApplicationServiceSyncParams) WithHTTPClient(client *http.Client) *ApplicationServiceSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application service sync params
func (o *ApplicationServiceSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the application service sync params
func (o *ApplicationServiceSyncParams) WithBody(body *models.ApplicationApplicationSyncRequest) *ApplicationServiceSyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the application service sync params
func (o *ApplicationServiceSyncParams) SetBody(body *models.ApplicationApplicationSyncRequest) {
	o.Body = body
}

// WithName adds the name to the application service sync params
func (o *ApplicationServiceSyncParams) WithName(name string) *ApplicationServiceSyncParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application service sync params
func (o *ApplicationServiceSyncParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationServiceSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
