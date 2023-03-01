// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CoreBwlimit(params *CoreBwlimitParams) (*CoreBwlimitOK, error)

	CoreStatsDelete(params *CoreStatsDeleteParams) (*CoreStatsDeleteOK, error)

	CoreStatsReset(params *CoreStatsResetParams) (*CoreStatsResetOK, error)

	FreeOSMemory(params *FreeOSMemoryParams) (*FreeOSMemoryOK, error)

	JobInfo(params *JobInfoParams) (*JobInfoOK, error)

	JobProgress(params *JobProgressParams) (*JobProgressOK, error)

	JobStop(params *JobStopParams) (*JobStopOK, error)

	NodeInfo(params *NodeInfoParams) (*NodeInfoOK, error)

	OperationsAbout(params *OperationsAboutParams) (*OperationsAboutOK, error)

	OperationsCheckPermissions(params *OperationsCheckPermissionsParams) (*OperationsCheckPermissionsOK, error)

	OperationsDeletefile(params *OperationsDeletefileParams) (*OperationsDeletefileOK, error)

	OperationsFileInfo(params *OperationsFileInfoParams) (*OperationsFileInfoOK, error)

	OperationsList(params *OperationsListParams) (*OperationsListOK, error)

	OperationsMovefile(params *OperationsMovefileParams) (*OperationsMovefileOK, error)

	OperationsPurge(params *OperationsPurgeParams) (*OperationsPurgeOK, error)

	Reload(params *ReloadParams) (*ReloadOK, error)

	SyncCopyDir(params *SyncCopyDirParams) (*SyncCopyDirOK, error)

	SyncCopyPaths(params *SyncCopyPathsParams) (*SyncCopyPathsOK, error)

	SyncMoveDir(params *SyncMoveDirParams) (*SyncMoveDirOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CoreBwlimit sets the bandwidth limit

This sets the bandwidth limit to that passed in
*/
func (a *Client) CoreBwlimit(params *CoreBwlimitParams) (*CoreBwlimitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoreBwlimitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoreBwlimit",
		Method:             "POST",
		PathPattern:        "/rclone/core/bwlimit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoreBwlimitReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CoreBwlimitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CoreBwlimitDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CoreStatsDelete deletes specific stats group

Delete stats
*/
func (a *Client) CoreStatsDelete(params *CoreStatsDeleteParams) (*CoreStatsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoreStatsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoreStatsDelete",
		Method:             "POST",
		PathPattern:        "/rclone/core/stats-delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoreStatsDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CoreStatsDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CoreStatsDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CoreStatsReset resets all or specific stats group

Resets stats
*/
func (a *Client) CoreStatsReset(params *CoreStatsResetParams) (*CoreStatsResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoreStatsResetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoreStatsReset",
		Method:             "POST",
		PathPattern:        "/rclone/core/stats-reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoreStatsResetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CoreStatsResetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CoreStatsResetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FreeOSMemory returns memory to o s

Run debug.FreeOSMemory on the agent
*/
func (a *Client) FreeOSMemory(params *FreeOSMemoryParams) (*FreeOSMemoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFreeOSMemoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FreeOSMemory",
		Method:             "POST",
		PathPattern:        "/free_os_memory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FreeOSMemoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FreeOSMemoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FreeOSMemoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
JobInfo transfers stats about the job

Returns current, completed transfers and job stats
*/
func (a *Client) JobInfo(params *JobInfoParams) (*JobInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JobInfo",
		Method:             "POST",
		PathPattern:        "/rclone/job/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &JobInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*JobInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*JobInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
JobProgress returns aggregated job stats

Returns aggregated job stats
*/
func (a *Client) JobProgress(params *JobProgressParams) (*JobProgressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobProgressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JobProgress",
		Method:             "POST",
		PathPattern:        "/rclone/job/progress",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &JobProgressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*JobProgressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*JobProgressDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
JobStop stops async job

Stops job with provided ID
*/
func (a *Client) JobStop(params *JobStopParams) (*JobStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JobStop",
		Method:             "POST",
		PathPattern:        "/rclone/job/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &JobStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*JobStopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*JobStopDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
NodeInfo gets information about scylla node

Get information about Scylla node
*/
func (a *Client) NodeInfo(params *NodeInfoParams) (*NodeInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NodeInfo",
		Method:             "GET",
		PathPattern:        "/node_info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NodeInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodeInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OperationsAbout abouts remote

Get usage information from the remote
*/
func (a *Client) OperationsAbout(params *OperationsAboutParams) (*OperationsAboutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsAboutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsAbout",
		Method:             "POST",
		PathPattern:        "/rclone/operations/about",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsAboutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OperationsAboutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OperationsAboutDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OperationsCheckPermissions checks fs

Check if the fs is fully accessible
*/
func (a *Client) OperationsCheckPermissions(params *OperationsCheckPermissionsParams) (*OperationsCheckPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsCheckPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsCheckPermissions",
		Method:             "POST",
		PathPattern:        "/rclone/operations/check-permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsCheckPermissionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OperationsCheckPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OperationsCheckPermissionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OperationsDeletefile deletes file

Remove the single file pointed to
*/
func (a *Client) OperationsDeletefile(params *OperationsDeletefileParams) (*OperationsDeletefileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsDeletefileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsDeletefile",
		Method:             "POST",
		PathPattern:        "/rclone/operations/deletefile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsDeletefileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OperationsDeletefileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OperationsDeletefileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OperationsFileInfo objects info

Get basic file information
*/
func (a *Client) OperationsFileInfo(params *OperationsFileInfoParams) (*OperationsFileInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsFileInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsFileInfo",
		Method:             "POST",
		PathPattern:        "/rclone/operations/fileinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsFileInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OperationsFileInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OperationsFileInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OperationsList lists remote

List the given remote and path
*/
func (a *Client) OperationsList(params *OperationsListParams) (*OperationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsList",
		Method:             "POST",
		PathPattern:        "/rclone/operations/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OperationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OperationsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OperationsMovefile moves a file

Move a file from source remote to destination remote
*/
func (a *Client) OperationsMovefile(params *OperationsMovefileParams) (*OperationsMovefileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsMovefileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsMovefile",
		Method:             "POST",
		PathPattern:        "/rclone/operations/movefile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsMovefileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OperationsMovefileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OperationsMovefileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
OperationsPurge purges container

Remove a directory or container and all of its contents
*/
func (a *Client) OperationsPurge(params *OperationsPurgeParams) (*OperationsPurgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperationsPurgeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperationsPurge",
		Method:             "POST",
		PathPattern:        "/rclone/operations/purge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperationsPurgeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OperationsPurgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*OperationsPurgeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Reload reloads agent config

Reload agent config
*/
func (a *Client) Reload(params *ReloadParams) (*ReloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReloadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Reload",
		Method:             "POST",
		PathPattern:        "/terminate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReloadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReloadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SyncCopyDir copies dir contents to directory

Copy contents from path on source fs to path on destination fs
*/
func (a *Client) SyncCopyDir(params *SyncCopyDirParams) (*SyncCopyDirOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncCopyDirParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SyncCopyDir",
		Method:             "POST",
		PathPattern:        "/rclone/sync/copydir",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SyncCopyDirReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SyncCopyDirOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SyncCopyDirDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SyncCopyPaths copies paths from fsrc remote src paths to fdst remote dst paths

Copy provided list of paths from directory on source fs to directory on destination fs
*/
func (a *Client) SyncCopyPaths(params *SyncCopyPathsParams) (*SyncCopyPathsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncCopyPathsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SyncCopyPaths",
		Method:             "POST",
		PathPattern:        "/rclone/sync/copypaths",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SyncCopyPathsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SyncCopyPathsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SyncCopyPathsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SyncMoveDir moves dir contents to directory

Move contents from path on source fs to path on destination fs
*/
func (a *Client) SyncMoveDir(params *SyncMoveDirParams) (*SyncMoveDirOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncMoveDirParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SyncMoveDir",
		Method:             "POST",
		PathPattern:        "/rclone/sync/movedir",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SyncMoveDirReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SyncMoveDirOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SyncMoveDirDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
