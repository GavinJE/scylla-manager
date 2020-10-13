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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClusterClusterIDTaskRepairTaskIDRunIDParams creates a new GetClusterClusterIDTaskRepairTaskIDRunIDParams object
// with the default values initialized.
func NewGetClusterClusterIDTaskRepairTaskIDRunIDParams() *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	var ()
	return &GetClusterClusterIDTaskRepairTaskIDRunIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterClusterIDTaskRepairTaskIDRunIDParamsWithTimeout creates a new GetClusterClusterIDTaskRepairTaskIDRunIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterClusterIDTaskRepairTaskIDRunIDParamsWithTimeout(timeout time.Duration) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	var ()
	return &GetClusterClusterIDTaskRepairTaskIDRunIDParams{

		timeout: timeout,
	}
}

// NewGetClusterClusterIDTaskRepairTaskIDRunIDParamsWithContext creates a new GetClusterClusterIDTaskRepairTaskIDRunIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterClusterIDTaskRepairTaskIDRunIDParamsWithContext(ctx context.Context) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	var ()
	return &GetClusterClusterIDTaskRepairTaskIDRunIDParams{

		Context: ctx,
	}
}

// NewGetClusterClusterIDTaskRepairTaskIDRunIDParamsWithHTTPClient creates a new GetClusterClusterIDTaskRepairTaskIDRunIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterClusterIDTaskRepairTaskIDRunIDParamsWithHTTPClient(client *http.Client) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	var ()
	return &GetClusterClusterIDTaskRepairTaskIDRunIDParams{
		HTTPClient: client,
	}
}

/*GetClusterClusterIDTaskRepairTaskIDRunIDParams contains all the parameters to send to the API endpoint
for the get cluster cluster ID task repair task ID run ID operation typically these are written to a http.Request
*/
type GetClusterClusterIDTaskRepairTaskIDRunIDParams struct {

	/*ClusterID*/
	ClusterID string
	/*RunID*/
	RunID string
	/*TaskID*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) WithTimeout(timeout time.Duration) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) WithContext(ctx context.Context) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) WithHTTPClient(client *http.Client) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) WithClusterID(clusterID string) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithRunID adds the runID to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) WithRunID(runID string) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) SetRunID(runID string) {
	o.RunID = runID
}

// WithTaskID adds the taskID to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) WithTaskID(taskID string) *GetClusterClusterIDTaskRepairTaskIDRunIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get cluster cluster ID task repair task ID run ID params
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param run_id
	if err := r.SetPathParam("run_id", o.RunID); err != nil {
		return err
	}

	// path param task_id
	if err := r.SetPathParam("task_id", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}