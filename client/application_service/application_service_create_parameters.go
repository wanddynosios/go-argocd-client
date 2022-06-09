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
	"github.com/go-openapi/swag"

	"github.com/wanddynosios/go-argocd-client/models"
)

// NewApplicationServiceCreateParams creates a new ApplicationServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationServiceCreateParams() *ApplicationServiceCreateParams {
	return &ApplicationServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationServiceCreateParamsWithTimeout creates a new ApplicationServiceCreateParams object
// with the ability to set a timeout on a request.
func NewApplicationServiceCreateParamsWithTimeout(timeout time.Duration) *ApplicationServiceCreateParams {
	return &ApplicationServiceCreateParams{
		timeout: timeout,
	}
}

// NewApplicationServiceCreateParamsWithContext creates a new ApplicationServiceCreateParams object
// with the ability to set a context for a request.
func NewApplicationServiceCreateParamsWithContext(ctx context.Context) *ApplicationServiceCreateParams {
	return &ApplicationServiceCreateParams{
		Context: ctx,
	}
}

// NewApplicationServiceCreateParamsWithHTTPClient creates a new ApplicationServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationServiceCreateParamsWithHTTPClient(client *http.Client) *ApplicationServiceCreateParams {
	return &ApplicationServiceCreateParams{
		HTTPClient: client,
	}
}

/* ApplicationServiceCreateParams contains all the parameters to send to the API endpoint
   for the application service create operation.

   Typically these are written to a http.Request.
*/
type ApplicationServiceCreateParams struct {

	// Body.
	Body *models.V1alpha1Application

	// Upsert.
	Upsert *bool

	// Validate.
	Validate *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceCreateParams) WithDefaults() *ApplicationServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application service create params
func (o *ApplicationServiceCreateParams) WithTimeout(timeout time.Duration) *ApplicationServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application service create params
func (o *ApplicationServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application service create params
func (o *ApplicationServiceCreateParams) WithContext(ctx context.Context) *ApplicationServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application service create params
func (o *ApplicationServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application service create params
func (o *ApplicationServiceCreateParams) WithHTTPClient(client *http.Client) *ApplicationServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application service create params
func (o *ApplicationServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the application service create params
func (o *ApplicationServiceCreateParams) WithBody(body *models.V1alpha1Application) *ApplicationServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the application service create params
func (o *ApplicationServiceCreateParams) SetBody(body *models.V1alpha1Application) {
	o.Body = body
}

// WithUpsert adds the upsert to the application service create params
func (o *ApplicationServiceCreateParams) WithUpsert(upsert *bool) *ApplicationServiceCreateParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the application service create params
func (o *ApplicationServiceCreateParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WithValidate adds the validate to the application service create params
func (o *ApplicationServiceCreateParams) WithValidate(validate *bool) *ApplicationServiceCreateParams {
	o.SetValidate(validate)
	return o
}

// SetValidate adds the validate to the application service create params
func (o *ApplicationServiceCreateParams) SetValidate(validate *bool) {
	o.Validate = validate
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Validate != nil {

		// query param validate
		var qrValidate bool

		if o.Validate != nil {
			qrValidate = *o.Validate
		}
		qValidate := swag.FormatBool(qrValidate)
		if qValidate != "" {

			if err := r.SetQueryParam("validate", qValidate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
