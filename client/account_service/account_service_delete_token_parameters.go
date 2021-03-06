// Code generated by go-swagger; DO NOT EDIT.

package account_service

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

// NewAccountServiceDeleteTokenParams creates a new AccountServiceDeleteTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountServiceDeleteTokenParams() *AccountServiceDeleteTokenParams {
	return &AccountServiceDeleteTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountServiceDeleteTokenParamsWithTimeout creates a new AccountServiceDeleteTokenParams object
// with the ability to set a timeout on a request.
func NewAccountServiceDeleteTokenParamsWithTimeout(timeout time.Duration) *AccountServiceDeleteTokenParams {
	return &AccountServiceDeleteTokenParams{
		timeout: timeout,
	}
}

// NewAccountServiceDeleteTokenParamsWithContext creates a new AccountServiceDeleteTokenParams object
// with the ability to set a context for a request.
func NewAccountServiceDeleteTokenParamsWithContext(ctx context.Context) *AccountServiceDeleteTokenParams {
	return &AccountServiceDeleteTokenParams{
		Context: ctx,
	}
}

// NewAccountServiceDeleteTokenParamsWithHTTPClient creates a new AccountServiceDeleteTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountServiceDeleteTokenParamsWithHTTPClient(client *http.Client) *AccountServiceDeleteTokenParams {
	return &AccountServiceDeleteTokenParams{
		HTTPClient: client,
	}
}

/* AccountServiceDeleteTokenParams contains all the parameters to send to the API endpoint
   for the account service delete token operation.

   Typically these are written to a http.Request.
*/
type AccountServiceDeleteTokenParams struct {

	// ID.
	ID string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account service delete token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountServiceDeleteTokenParams) WithDefaults() *AccountServiceDeleteTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account service delete token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountServiceDeleteTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account service delete token params
func (o *AccountServiceDeleteTokenParams) WithTimeout(timeout time.Duration) *AccountServiceDeleteTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account service delete token params
func (o *AccountServiceDeleteTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account service delete token params
func (o *AccountServiceDeleteTokenParams) WithContext(ctx context.Context) *AccountServiceDeleteTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account service delete token params
func (o *AccountServiceDeleteTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account service delete token params
func (o *AccountServiceDeleteTokenParams) WithHTTPClient(client *http.Client) *AccountServiceDeleteTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account service delete token params
func (o *AccountServiceDeleteTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the account service delete token params
func (o *AccountServiceDeleteTokenParams) WithID(id string) *AccountServiceDeleteTokenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the account service delete token params
func (o *AccountServiceDeleteTokenParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the account service delete token params
func (o *AccountServiceDeleteTokenParams) WithName(name string) *AccountServiceDeleteTokenParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the account service delete token params
func (o *AccountServiceDeleteTokenParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *AccountServiceDeleteTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
