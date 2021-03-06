// Code generated by go-swagger; DO NOT EDIT.

package cluster_service

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

// NewClusterServiceInvalidateCacheParams creates a new ClusterServiceInvalidateCacheParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterServiceInvalidateCacheParams() *ClusterServiceInvalidateCacheParams {
	return &ClusterServiceInvalidateCacheParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterServiceInvalidateCacheParamsWithTimeout creates a new ClusterServiceInvalidateCacheParams object
// with the ability to set a timeout on a request.
func NewClusterServiceInvalidateCacheParamsWithTimeout(timeout time.Duration) *ClusterServiceInvalidateCacheParams {
	return &ClusterServiceInvalidateCacheParams{
		timeout: timeout,
	}
}

// NewClusterServiceInvalidateCacheParamsWithContext creates a new ClusterServiceInvalidateCacheParams object
// with the ability to set a context for a request.
func NewClusterServiceInvalidateCacheParamsWithContext(ctx context.Context) *ClusterServiceInvalidateCacheParams {
	return &ClusterServiceInvalidateCacheParams{
		Context: ctx,
	}
}

// NewClusterServiceInvalidateCacheParamsWithHTTPClient creates a new ClusterServiceInvalidateCacheParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterServiceInvalidateCacheParamsWithHTTPClient(client *http.Client) *ClusterServiceInvalidateCacheParams {
	return &ClusterServiceInvalidateCacheParams{
		HTTPClient: client,
	}
}

/* ClusterServiceInvalidateCacheParams contains all the parameters to send to the API endpoint
   for the cluster service invalidate cache operation.

   Typically these are written to a http.Request.
*/
type ClusterServiceInvalidateCacheParams struct {

	/* IDValue.

	   value holds the cluster server URL or cluster name
	*/
	IDValue string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster service invalidate cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterServiceInvalidateCacheParams) WithDefaults() *ClusterServiceInvalidateCacheParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster service invalidate cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterServiceInvalidateCacheParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster service invalidate cache params
func (o *ClusterServiceInvalidateCacheParams) WithTimeout(timeout time.Duration) *ClusterServiceInvalidateCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster service invalidate cache params
func (o *ClusterServiceInvalidateCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster service invalidate cache params
func (o *ClusterServiceInvalidateCacheParams) WithContext(ctx context.Context) *ClusterServiceInvalidateCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster service invalidate cache params
func (o *ClusterServiceInvalidateCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster service invalidate cache params
func (o *ClusterServiceInvalidateCacheParams) WithHTTPClient(client *http.Client) *ClusterServiceInvalidateCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster service invalidate cache params
func (o *ClusterServiceInvalidateCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDValue adds the iDValue to the cluster service invalidate cache params
func (o *ClusterServiceInvalidateCacheParams) WithIDValue(iDValue string) *ClusterServiceInvalidateCacheParams {
	o.SetIDValue(iDValue)
	return o
}

// SetIDValue adds the idValue to the cluster service invalidate cache params
func (o *ClusterServiceInvalidateCacheParams) SetIDValue(iDValue string) {
	o.IDValue = iDValue
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterServiceInvalidateCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id.value
	if err := r.SetPathParam("id.value", o.IDValue); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
