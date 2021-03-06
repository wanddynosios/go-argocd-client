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

// NewClusterServiceGetParams creates a new ClusterServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterServiceGetParams() *ClusterServiceGetParams {
	return &ClusterServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterServiceGetParamsWithTimeout creates a new ClusterServiceGetParams object
// with the ability to set a timeout on a request.
func NewClusterServiceGetParamsWithTimeout(timeout time.Duration) *ClusterServiceGetParams {
	return &ClusterServiceGetParams{
		timeout: timeout,
	}
}

// NewClusterServiceGetParamsWithContext creates a new ClusterServiceGetParams object
// with the ability to set a context for a request.
func NewClusterServiceGetParamsWithContext(ctx context.Context) *ClusterServiceGetParams {
	return &ClusterServiceGetParams{
		Context: ctx,
	}
}

// NewClusterServiceGetParamsWithHTTPClient creates a new ClusterServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterServiceGetParamsWithHTTPClient(client *http.Client) *ClusterServiceGetParams {
	return &ClusterServiceGetParams{
		HTTPClient: client,
	}
}

/* ClusterServiceGetParams contains all the parameters to send to the API endpoint
   for the cluster service get operation.

   Typically these are written to a http.Request.
*/
type ClusterServiceGetParams struct {

	/* IDType.

	   type is the type of the specified cluster identifier ( "server" - default, "name" ).
	*/
	IDType *string

	/* IDValue.

	   value holds the cluster server URL or cluster name
	*/
	IDValue string

	// Name.
	Name *string

	// Server.
	Server *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterServiceGetParams) WithDefaults() *ClusterServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster service get params
func (o *ClusterServiceGetParams) WithTimeout(timeout time.Duration) *ClusterServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster service get params
func (o *ClusterServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster service get params
func (o *ClusterServiceGetParams) WithContext(ctx context.Context) *ClusterServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster service get params
func (o *ClusterServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster service get params
func (o *ClusterServiceGetParams) WithHTTPClient(client *http.Client) *ClusterServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster service get params
func (o *ClusterServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDType adds the iDType to the cluster service get params
func (o *ClusterServiceGetParams) WithIDType(iDType *string) *ClusterServiceGetParams {
	o.SetIDType(iDType)
	return o
}

// SetIDType adds the idType to the cluster service get params
func (o *ClusterServiceGetParams) SetIDType(iDType *string) {
	o.IDType = iDType
}

// WithIDValue adds the iDValue to the cluster service get params
func (o *ClusterServiceGetParams) WithIDValue(iDValue string) *ClusterServiceGetParams {
	o.SetIDValue(iDValue)
	return o
}

// SetIDValue adds the idValue to the cluster service get params
func (o *ClusterServiceGetParams) SetIDValue(iDValue string) {
	o.IDValue = iDValue
}

// WithName adds the name to the cluster service get params
func (o *ClusterServiceGetParams) WithName(name *string) *ClusterServiceGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the cluster service get params
func (o *ClusterServiceGetParams) SetName(name *string) {
	o.Name = name
}

// WithServer adds the server to the cluster service get params
func (o *ClusterServiceGetParams) WithServer(server *string) *ClusterServiceGetParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the cluster service get params
func (o *ClusterServiceGetParams) SetServer(server *string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IDType != nil {

		// query param id.type
		var qrIDType string

		if o.IDType != nil {
			qrIDType = *o.IDType
		}
		qIDType := qrIDType
		if qIDType != "" {

			if err := r.SetQueryParam("id.type", qIDType); err != nil {
				return err
			}
		}
	}

	// path param id.value
	if err := r.SetPathParam("id.value", o.IDValue); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Server != nil {

		// query param server
		var qrServer string

		if o.Server != nil {
			qrServer = *o.Server
		}
		qServer := qrServer
		if qServer != "" {

			if err := r.SetQueryParam("server", qServer); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
