// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new application service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ApplicationServiceCreate(params *ApplicationServiceCreateParams, opts ...ClientOption) (*ApplicationServiceCreateOK, error)

	ApplicationServiceDelete(params *ApplicationServiceDeleteParams, opts ...ClientOption) (*ApplicationServiceDeleteOK, error)

	ApplicationServiceDeleteResource(params *ApplicationServiceDeleteResourceParams, opts ...ClientOption) (*ApplicationServiceDeleteResourceOK, error)

	ApplicationServiceGet(params *ApplicationServiceGetParams, opts ...ClientOption) (*ApplicationServiceGetOK, error)

	ApplicationServiceGetApplicationSyncWindows(params *ApplicationServiceGetApplicationSyncWindowsParams, opts ...ClientOption) (*ApplicationServiceGetApplicationSyncWindowsOK, error)

	ApplicationServiceGetManifests(params *ApplicationServiceGetManifestsParams, opts ...ClientOption) (*ApplicationServiceGetManifestsOK, error)

	ApplicationServiceGetResource(params *ApplicationServiceGetResourceParams, opts ...ClientOption) (*ApplicationServiceGetResourceOK, error)

	ApplicationServiceList(params *ApplicationServiceListParams, opts ...ClientOption) (*ApplicationServiceListOK, error)

	ApplicationServiceListResourceActions(params *ApplicationServiceListResourceActionsParams, opts ...ClientOption) (*ApplicationServiceListResourceActionsOK, error)

	ApplicationServiceListResourceEvents(params *ApplicationServiceListResourceEventsParams, opts ...ClientOption) (*ApplicationServiceListResourceEventsOK, error)

	ApplicationServiceManagedResources(params *ApplicationServiceManagedResourcesParams, opts ...ClientOption) (*ApplicationServiceManagedResourcesOK, error)

	ApplicationServicePatch(params *ApplicationServicePatchParams, opts ...ClientOption) (*ApplicationServicePatchOK, error)

	ApplicationServicePatchResource(params *ApplicationServicePatchResourceParams, opts ...ClientOption) (*ApplicationServicePatchResourceOK, error)

	ApplicationServicePodLogs(params *ApplicationServicePodLogsParams, opts ...ClientOption) (*ApplicationServicePodLogsOK, error)

	ApplicationServicePodLogs2(params *ApplicationServicePodLogs2Params, opts ...ClientOption) (*ApplicationServicePodLogs2OK, error)

	ApplicationServiceResourceTree(params *ApplicationServiceResourceTreeParams, opts ...ClientOption) (*ApplicationServiceResourceTreeOK, error)

	ApplicationServiceRevisionMetadata(params *ApplicationServiceRevisionMetadataParams, opts ...ClientOption) (*ApplicationServiceRevisionMetadataOK, error)

	ApplicationServiceRollback(params *ApplicationServiceRollbackParams, opts ...ClientOption) (*ApplicationServiceRollbackOK, error)

	ApplicationServiceRunResourceAction(params *ApplicationServiceRunResourceActionParams, opts ...ClientOption) (*ApplicationServiceRunResourceActionOK, error)

	ApplicationServiceSync(params *ApplicationServiceSyncParams, opts ...ClientOption) (*ApplicationServiceSyncOK, error)

	ApplicationServiceTerminateOperation(params *ApplicationServiceTerminateOperationParams, opts ...ClientOption) (*ApplicationServiceTerminateOperationOK, error)

	ApplicationServiceUpdate(params *ApplicationServiceUpdateParams, opts ...ClientOption) (*ApplicationServiceUpdateOK, error)

	ApplicationServiceUpdateSpec(params *ApplicationServiceUpdateSpecParams, opts ...ClientOption) (*ApplicationServiceUpdateSpecOK, error)

	ApplicationServiceWatch(params *ApplicationServiceWatchParams, opts ...ClientOption) (*ApplicationServiceWatchOK, error)

	ApplicationServiceWatchResourceTree(params *ApplicationServiceWatchResourceTreeParams, opts ...ClientOption) (*ApplicationServiceWatchResourceTreeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ApplicationServiceCreate creates creates an application
*/
func (a *Client) ApplicationServiceCreate(params *ApplicationServiceCreateParams, opts ...ClientOption) (*ApplicationServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_Create",
		Method:             "POST",
		PathPattern:        "/api/v1/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceDelete deletes deletes an application
*/
func (a *Client) ApplicationServiceDelete(params *ApplicationServiceDeleteParams, opts ...ClientOption) (*ApplicationServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/applications/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceDeleteResource deletes resource deletes a single application resource
*/
func (a *Client) ApplicationServiceDeleteResource(params *ApplicationServiceDeleteResourceParams, opts ...ClientOption) (*ApplicationServiceDeleteResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceDeleteResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_DeleteResource",
		Method:             "DELETE",
		PathPattern:        "/api/v1/applications/{name}/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceDeleteResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceDeleteResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceDeleteResourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceGet gets returns an application by name
*/
func (a *Client) ApplicationServiceGet(params *ApplicationServiceGetParams, opts ...ClientOption) (*ApplicationServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_Get",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceGetApplicationSyncWindows gets returns sync windows of the application
*/
func (a *Client) ApplicationServiceGetApplicationSyncWindows(params *ApplicationServiceGetApplicationSyncWindowsParams, opts ...ClientOption) (*ApplicationServiceGetApplicationSyncWindowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceGetApplicationSyncWindowsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_GetApplicationSyncWindows",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/syncwindows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceGetApplicationSyncWindowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceGetApplicationSyncWindowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceGetApplicationSyncWindowsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceGetManifests gets manifests returns application manifests
*/
func (a *Client) ApplicationServiceGetManifests(params *ApplicationServiceGetManifestsParams, opts ...ClientOption) (*ApplicationServiceGetManifestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceGetManifestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_GetManifests",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/manifests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceGetManifestsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceGetManifestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceGetManifestsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceGetResource gets resource returns single application resource
*/
func (a *Client) ApplicationServiceGetResource(params *ApplicationServiceGetResourceParams, opts ...ClientOption) (*ApplicationServiceGetResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceGetResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_GetResource",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceGetResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceGetResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceGetResourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceList lists returns list of applications
*/
func (a *Client) ApplicationServiceList(params *ApplicationServiceListParams, opts ...ClientOption) (*ApplicationServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_List",
		Method:             "GET",
		PathPattern:        "/api/v1/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceListResourceActions lists resource actions returns list of resource actions
*/
func (a *Client) ApplicationServiceListResourceActions(params *ApplicationServiceListResourceActionsParams, opts ...ClientOption) (*ApplicationServiceListResourceActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceListResourceActionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_ListResourceActions",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/resource/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceListResourceActionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceListResourceActionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceListResourceActionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceListResourceEvents lists resource events returns a list of event resources
*/
func (a *Client) ApplicationServiceListResourceEvents(params *ApplicationServiceListResourceEventsParams, opts ...ClientOption) (*ApplicationServiceListResourceEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceListResourceEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_ListResourceEvents",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceListResourceEventsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceListResourceEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceListResourceEventsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceManagedResources manageds resources returns list of managed resources
*/
func (a *Client) ApplicationServiceManagedResources(params *ApplicationServiceManagedResourcesParams, opts ...ClientOption) (*ApplicationServiceManagedResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceManagedResourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_ManagedResources",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{applicationName}/managed-resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceManagedResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceManagedResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceManagedResourcesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServicePatch patches patch an application
*/
func (a *Client) ApplicationServicePatch(params *ApplicationServicePatchParams, opts ...ClientOption) (*ApplicationServicePatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServicePatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_Patch",
		Method:             "PATCH",
		PathPattern:        "/api/v1/applications/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServicePatchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServicePatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServicePatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServicePatchResource patches resource patch single application resource
*/
func (a *Client) ApplicationServicePatchResource(params *ApplicationServicePatchResourceParams, opts ...ClientOption) (*ApplicationServicePatchResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServicePatchResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_PatchResource",
		Method:             "POST",
		PathPattern:        "/api/v1/applications/{name}/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServicePatchResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServicePatchResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServicePatchResourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServicePodLogs pods logs returns stream of log entries for the specified pod pod
*/
func (a *Client) ApplicationServicePodLogs(params *ApplicationServicePodLogsParams, opts ...ClientOption) (*ApplicationServicePodLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServicePodLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_PodLogs",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/pods/{podName}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServicePodLogsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServicePodLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServicePodLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServicePodLogs2 pods logs returns stream of log entries for the specified pod pod
*/
func (a *Client) ApplicationServicePodLogs2(params *ApplicationServicePodLogs2Params, opts ...ClientOption) (*ApplicationServicePodLogs2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServicePodLogs2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_PodLogs2",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServicePodLogs2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServicePodLogs2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServicePodLogs2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceResourceTree resources tree returns resource tree
*/
func (a *Client) ApplicationServiceResourceTree(params *ApplicationServiceResourceTreeParams, opts ...ClientOption) (*ApplicationServiceResourceTreeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceResourceTreeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_ResourceTree",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{applicationName}/resource-tree",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceResourceTreeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceResourceTreeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceResourceTreeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceRevisionMetadata gets the meta data author date tags message for a specific revision of the application
*/
func (a *Client) ApplicationServiceRevisionMetadata(params *ApplicationServiceRevisionMetadataParams, opts ...ClientOption) (*ApplicationServiceRevisionMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceRevisionMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_RevisionMetadata",
		Method:             "GET",
		PathPattern:        "/api/v1/applications/{name}/revisions/{revision}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceRevisionMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceRevisionMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceRevisionMetadataDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceRollback rollbacks syncs an application to its target state
*/
func (a *Client) ApplicationServiceRollback(params *ApplicationServiceRollbackParams, opts ...ClientOption) (*ApplicationServiceRollbackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceRollbackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_Rollback",
		Method:             "POST",
		PathPattern:        "/api/v1/applications/{name}/rollback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceRollbackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceRollbackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceRollbackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceRunResourceAction runs resource action run resource action
*/
func (a *Client) ApplicationServiceRunResourceAction(params *ApplicationServiceRunResourceActionParams, opts ...ClientOption) (*ApplicationServiceRunResourceActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceRunResourceActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_RunResourceAction",
		Method:             "POST",
		PathPattern:        "/api/v1/applications/{name}/resource/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceRunResourceActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceRunResourceActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceRunResourceActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceSync syncs syncs an application to its target state
*/
func (a *Client) ApplicationServiceSync(params *ApplicationServiceSyncParams, opts ...ClientOption) (*ApplicationServiceSyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceSyncParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_Sync",
		Method:             "POST",
		PathPattern:        "/api/v1/applications/{name}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceSyncReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceSyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceSyncDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceTerminateOperation terminates operation terminates the currently running operation
*/
func (a *Client) ApplicationServiceTerminateOperation(params *ApplicationServiceTerminateOperationParams, opts ...ClientOption) (*ApplicationServiceTerminateOperationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceTerminateOperationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_TerminateOperation",
		Method:             "DELETE",
		PathPattern:        "/api/v1/applications/{name}/operation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceTerminateOperationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceTerminateOperationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceTerminateOperationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceUpdate updates updates an application
*/
func (a *Client) ApplicationServiceUpdate(params *ApplicationServiceUpdateParams, opts ...ClientOption) (*ApplicationServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_Update",
		Method:             "PUT",
		PathPattern:        "/api/v1/applications/{application.metadata.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceUpdateSpec updates spec updates an application spec
*/
func (a *Client) ApplicationServiceUpdateSpec(params *ApplicationServiceUpdateSpecParams, opts ...ClientOption) (*ApplicationServiceUpdateSpecOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceUpdateSpecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_UpdateSpec",
		Method:             "PUT",
		PathPattern:        "/api/v1/applications/{name}/spec",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceUpdateSpecReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceUpdateSpecOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceUpdateSpecDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceWatch watches returns stream of application change events
*/
func (a *Client) ApplicationServiceWatch(params *ApplicationServiceWatchParams, opts ...ClientOption) (*ApplicationServiceWatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceWatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_Watch",
		Method:             "GET",
		PathPattern:        "/api/v1/stream/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceWatchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceWatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceWatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ApplicationServiceWatchResourceTree watches returns stream of application resource tree
*/
func (a *Client) ApplicationServiceWatchResourceTree(params *ApplicationServiceWatchResourceTreeParams, opts ...ClientOption) (*ApplicationServiceWatchResourceTreeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationServiceWatchResourceTreeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplicationService_WatchResourceTree",
		Method:             "GET",
		PathPattern:        "/api/v1/stream/applications/{applicationName}/resource-tree",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationServiceWatchResourceTreeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationServiceWatchResourceTreeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationServiceWatchResourceTreeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
