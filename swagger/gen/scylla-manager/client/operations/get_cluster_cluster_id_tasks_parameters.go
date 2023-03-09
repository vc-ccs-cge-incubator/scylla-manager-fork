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
	"github.com/go-openapi/swag"
)

// NewGetClusterClusterIDTasksParams creates a new GetClusterClusterIDTasksParams object
// with the default values initialized.
func NewGetClusterClusterIDTasksParams() *GetClusterClusterIDTasksParams {
	var ()
	return &GetClusterClusterIDTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterClusterIDTasksParamsWithTimeout creates a new GetClusterClusterIDTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterClusterIDTasksParamsWithTimeout(timeout time.Duration) *GetClusterClusterIDTasksParams {
	var ()
	return &GetClusterClusterIDTasksParams{

		timeout: timeout,
	}
}

// NewGetClusterClusterIDTasksParamsWithContext creates a new GetClusterClusterIDTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterClusterIDTasksParamsWithContext(ctx context.Context) *GetClusterClusterIDTasksParams {
	var ()
	return &GetClusterClusterIDTasksParams{

		Context: ctx,
	}
}

// NewGetClusterClusterIDTasksParamsWithHTTPClient creates a new GetClusterClusterIDTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterClusterIDTasksParamsWithHTTPClient(client *http.Client) *GetClusterClusterIDTasksParams {
	var ()
	return &GetClusterClusterIDTasksParams{
		HTTPClient: client,
	}
}

/*
GetClusterClusterIDTasksParams contains all the parameters to send to the API endpoint
for the get cluster cluster ID tasks operation typically these are written to a http.Request
*/
type GetClusterClusterIDTasksParams struct {

	/*All*/
	All *bool
	/*ClusterID*/
	ClusterID string
	/*Short*/
	Short *bool
	/*Status*/
	Status *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) WithTimeout(timeout time.Duration) *GetClusterClusterIDTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) WithContext(ctx context.Context) *GetClusterClusterIDTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) WithHTTPClient(client *http.Client) *GetClusterClusterIDTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) WithAll(all *bool) *GetClusterClusterIDTasksParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) SetAll(all *bool) {
	o.All = all
}

// WithClusterID adds the clusterID to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) WithClusterID(clusterID string) *GetClusterClusterIDTasksParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithShort adds the short to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) WithShort(short *bool) *GetClusterClusterIDTasksParams {
	o.SetShort(short)
	return o
}

// SetShort adds the short to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) SetShort(short *bool) {
	o.Short = short
}

// WithStatus adds the status to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) WithStatus(status *string) *GetClusterClusterIDTasksParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) SetStatus(status *string) {
	o.Status = status
}

// WithType adds the typeVar to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) WithType(typeVar *string) *GetClusterClusterIDTasksParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get cluster cluster ID tasks params
func (o *GetClusterClusterIDTasksParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterClusterIDTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool
		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {
			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}

	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.Short != nil {

		// query param short
		var qrShort bool
		if o.Short != nil {
			qrShort = *o.Short
		}
		qShort := swag.FormatBool(qrShort)
		if qShort != "" {
			if err := r.SetQueryParam("short", qShort); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
