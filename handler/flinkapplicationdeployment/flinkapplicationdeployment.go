// Code generated by Aiven. DO NOT EDIT.

package flinkapplicationdeployment

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// Cancel cancel an ApplicationDeployment
	// ServiceFlinkCancelApplicationDeployment POST /project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}/cancel
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCancelApplicationDeployment
	Cancel(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*CancelOut, error)

	// Create create an ApplicationDeployment
	// ServiceFlinkCreateApplicationDeployment POST /project/{project}/service/{service_name}/flink/application/{application_id}/deployment
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCreateApplicationDeployment
	Create(ctx context.Context, project string, serviceName string, applicationId string, in *CreateIn) (*CreateOut, error)

	// Delete delete an ApplicationDeployment
	// ServiceFlinkDeleteApplicationDeployment DELETE /project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkDeleteApplicationDeployment
	Delete(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*DeleteOut, error)

	// Get get an ApplicationDeployment
	// ServiceFlinkGetApplicationDeployment GET /project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkGetApplicationDeployment
	Get(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*GetOut, error)

	// List get all ApplicationDeployments
	// ServiceFlinkListApplicationDeployments GET /project/{project}/service/{service_name}/flink/application/{application_id}/deployment
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkListApplicationDeployments
	List(ctx context.Context, project string, serviceName string, applicationId string) ([]Deployment, error)

	// Stop stop an ApplicationDeployment
	// ServiceFlinkStopApplicationDeployment POST /project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}/stop
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkStopApplicationDeployment
	Stop(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*StopOut, error)
}

func NewHandler(doer doer) Handler {
	return &handler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type handler struct {
	doer doer
}

func (h *handler) Cancel(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*CancelOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment/%s/cancel", project, serviceName, applicationId, deploymentId)
	b, err := h.doer.Do(ctx, "ServiceFlinkCancelApplicationDeployment", "POST", path, nil)
	out := new(CancelOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) Create(ctx context.Context, project string, serviceName string, applicationId string, in *CreateIn) (*CreateOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkCreateApplicationDeployment", "POST", path, in)
	out := new(CreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) Delete(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*DeleteOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment/%s", project, serviceName, applicationId, deploymentId)
	b, err := h.doer.Do(ctx, "ServiceFlinkDeleteApplicationDeployment", "DELETE", path, nil)
	out := new(DeleteOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) Get(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*GetOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment/%s", project, serviceName, applicationId, deploymentId)
	b, err := h.doer.Do(ctx, "ServiceFlinkGetApplicationDeployment", "GET", path, nil)
	out := new(GetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) List(ctx context.Context, project string, serviceName string, applicationId string) ([]Deployment, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkListApplicationDeployments", "GET", path, nil)
	out := new(ListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Deployments, nil
}
func (h *handler) Stop(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*StopOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment/%s/stop", project, serviceName, applicationId, deploymentId)
	b, err := h.doer.Do(ctx, "ServiceFlinkStopApplicationDeployment", "POST", path, nil)
	out := new(StopOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type CancelOut struct {
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         string     `json:"created_by"`
	Id                string     `json:"id"`
	JobId             string     `json:"job_id,omitempty"`
	LastSavepoint     string     `json:"last_savepoint,omitempty"`
	Parallelism       int        `json:"parallelism"`
	RestartEnabled    bool       `json:"restart_enabled"`
	StartingSavepoint string     `json:"starting_savepoint,omitempty"`
	Status            StatusType `json:"status"`
	VersionId         string     `json:"version_id"`
}
type CreateIn struct {
	Parallelism       *int   `json:"parallelism,omitempty"`
	RestartEnabled    *bool  `json:"restart_enabled,omitempty"`
	StartingSavepoint string `json:"starting_savepoint,omitempty"`
	VersionId         string `json:"version_id"`
}
type CreateOut struct {
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         string     `json:"created_by"`
	Id                string     `json:"id"`
	JobId             string     `json:"job_id,omitempty"`
	LastSavepoint     string     `json:"last_savepoint,omitempty"`
	Parallelism       int        `json:"parallelism"`
	RestartEnabled    bool       `json:"restart_enabled"`
	StartingSavepoint string     `json:"starting_savepoint,omitempty"`
	Status            StatusType `json:"status"`
	VersionId         string     `json:"version_id"`
}
type DeleteOut struct {
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         string     `json:"created_by"`
	Id                string     `json:"id"`
	JobId             string     `json:"job_id,omitempty"`
	LastSavepoint     string     `json:"last_savepoint,omitempty"`
	Parallelism       int        `json:"parallelism"`
	RestartEnabled    bool       `json:"restart_enabled"`
	StartingSavepoint string     `json:"starting_savepoint,omitempty"`
	Status            StatusType `json:"status"`
	VersionId         string     `json:"version_id"`
}
type Deployment struct {
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         string     `json:"created_by"`
	Id                string     `json:"id"`
	JobId             string     `json:"job_id,omitempty"`
	LastSavepoint     string     `json:"last_savepoint,omitempty"`
	Parallelism       int        `json:"parallelism"`
	RestartEnabled    bool       `json:"restart_enabled"`
	StartingSavepoint string     `json:"starting_savepoint,omitempty"`
	Status            StatusType `json:"status"`
	VersionId         string     `json:"version_id"`
}
type GetOut struct {
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         string     `json:"created_by"`
	Id                string     `json:"id"`
	JobId             string     `json:"job_id,omitempty"`
	LastSavepoint     string     `json:"last_savepoint,omitempty"`
	Parallelism       int        `json:"parallelism"`
	RestartEnabled    bool       `json:"restart_enabled"`
	StartingSavepoint string     `json:"starting_savepoint,omitempty"`
	Status            StatusType `json:"status"`
	VersionId         string     `json:"version_id"`
}
type ListOut struct {
	Deployments []Deployment `json:"deployments"`
}
type StatusType string

const (
	StatusTypeInitializing           StatusType = "INITIALIZING"
	StatusTypeCreated                StatusType = "CREATED"
	StatusTypeRunning                StatusType = "RUNNING"
	StatusTypeFailing                StatusType = "FAILING"
	StatusTypeFailed                 StatusType = "FAILED"
	StatusTypeSaving                 StatusType = "SAVING"
	StatusTypeCancellingRequested    StatusType = "CANCELLING_REQUESTED"
	StatusTypeCancelling             StatusType = "CANCELLING"
	StatusTypeCanceled               StatusType = "CANCELED"
	StatusTypeSavingAndStopRequested StatusType = "SAVING_AND_STOP_REQUESTED"
	StatusTypeSavingAndStop          StatusType = "SAVING_AND_STOP"
	StatusTypeFinished               StatusType = "FINISHED"
	StatusTypeRestarting             StatusType = "RESTARTING"
	StatusTypeSuspended              StatusType = "SUSPENDED"
	StatusTypeDeleteRequested        StatusType = "DELETE_REQUESTED"
	StatusTypeDeleting               StatusType = "DELETING"
	StatusTypeReconciling            StatusType = "RECONCILING"
)

type StopOut struct {
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         string     `json:"created_by"`
	Id                string     `json:"id"`
	JobId             string     `json:"job_id,omitempty"`
	LastSavepoint     string     `json:"last_savepoint,omitempty"`
	Parallelism       int        `json:"parallelism"`
	RestartEnabled    bool       `json:"restart_enabled"`
	StartingSavepoint string     `json:"starting_savepoint,omitempty"`
	Status            StatusType `json:"status"`
	VersionId         string     `json:"version_id"`
}
