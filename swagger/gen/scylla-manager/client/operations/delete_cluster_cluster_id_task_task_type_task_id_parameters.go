// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDParams creates a new DeleteClusterClusterIDTaskTaskTypeTaskIDParams object
// with the default values initialized.
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDParams() *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	var ()
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDParamsWithTimeout creates a new DeleteClusterClusterIDTaskTaskTypeTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDParamsWithTimeout(timeout time.Duration) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	var ()
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDParams{

		timeout: timeout,
	}
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDParamsWithContext creates a new DeleteClusterClusterIDTaskTaskTypeTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDParamsWithContext(ctx context.Context) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	var ()
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDParams{

		Context: ctx,
	}
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDParamsWithHTTPClient creates a new DeleteClusterClusterIDTaskTaskTypeTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDParamsWithHTTPClient(client *http.Client) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	var ()
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDParams{
		HTTPClient: client,
	}
}

/*
DeleteClusterClusterIDTaskTaskTypeTaskIDParams contains all the parameters to send to the API endpoint
for the delete cluster cluster ID task task type task ID operation typically these are written to a http.Request
*/
type DeleteClusterClusterIDTaskTaskTypeTaskIDParams struct {

	/*ClusterID*/
	ClusterID string
	/*TaskID*/
	TaskID string
	/*TaskType*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) WithTimeout(timeout time.Duration) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) WithContext(ctx context.Context) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) WithHTTPClient(client *http.Client) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) WithClusterID(clusterID string) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithTaskID adds the taskID to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) WithTaskID(taskID string) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WithTaskType adds the taskType to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) WithTaskType(taskType string) *DeleteClusterClusterIDTaskTaskTypeTaskIDParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the delete cluster cluster ID task task type task ID params
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param task_id
	if err := r.SetPathParam("task_id", o.TaskID); err != nil {
		return err
	}

	// path param task_type
	if err := r.SetPathParam("task_type", o.TaskType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
