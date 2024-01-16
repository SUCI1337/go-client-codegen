// Code generated by Aiven. DO NOT EDIT.

package flinkapplication

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// ServiceFlinkCreateApplication create a Flink Application
	// POST /project/{project}/service/{service_name}/flink/application
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCreateApplication
	ServiceFlinkCreateApplication(ctx context.Context, project string, serviceName string, in *ServiceFlinkCreateApplicationIn) (*ServiceFlinkCreateApplicationOut, error)

	// ServiceFlinkCreateApplicationVersion create a Flink ApplicationVersion
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/version
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCreateApplicationVersion
	ServiceFlinkCreateApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkCreateApplicationVersionIn) (*ServiceFlinkCreateApplicationVersionOut, error)

	// ServiceFlinkDeleteApplication delete a Flink Application
	// DELETE /project/{project}/service/{service_name}/flink/application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkDeleteApplication
	ServiceFlinkDeleteApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkDeleteApplicationOut, error)

	// ServiceFlinkDeleteApplicationVersion delete a Flink ApplicationVersion
	// DELETE /project/{project}/service/{service_name}/flink/application/{application_id}/version/{application_version_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkDeleteApplicationVersion
	ServiceFlinkDeleteApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, applicationVersionId string) (*ServiceFlinkDeleteApplicationVersionOut, error)

	// ServiceFlinkGetApplication get a Flink Application
	// GET /project/{project}/service/{service_name}/flink/application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkGetApplication
	ServiceFlinkGetApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkGetApplicationOut, error)

	// ServiceFlinkGetApplicationVersion get a Flink ApplicationVersion
	// GET /project/{project}/service/{service_name}/flink/application/{application_id}/version/{application_version_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkGetApplicationVersion
	ServiceFlinkGetApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, applicationVersionId string) (*ServiceFlinkGetApplicationVersionOut, error)

	// ServiceFlinkListApplications get all Flink Applications
	// GET /project/{project}/service/{service_name}/flink/application
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkListApplications
	ServiceFlinkListApplications(ctx context.Context, project string, serviceName string) ([]Application, error)

	// ServiceFlinkUpdateApplication update a Flink Application
	// PUT /project/{project}/service/{service_name}/flink/application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkUpdateApplication
	ServiceFlinkUpdateApplication(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkUpdateApplicationIn) (*ServiceFlinkUpdateApplicationOut, error)

	// ServiceFlinkValidateApplicationVersion validate a Flink ApplicationVersion
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/version/validate
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkValidateApplicationVersion
	ServiceFlinkValidateApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkValidateApplicationVersionIn) (*ServiceFlinkValidateApplicationVersionOut, error)
}

func NewHandler(doer doer) FlinkApplicationHandler {
	return FlinkApplicationHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type FlinkApplicationHandler struct {
	doer doer
}

func (h *FlinkApplicationHandler) ServiceFlinkCreateApplication(ctx context.Context, project string, serviceName string, in *ServiceFlinkCreateApplicationIn) (*ServiceFlinkCreateApplicationOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceFlinkCreateApplication", "POST", path, in)
	out := new(ServiceFlinkCreateApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkCreateApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkCreateApplicationVersionIn) (*ServiceFlinkCreateApplicationVersionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/version", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkCreateApplicationVersion", "POST", path, in)
	out := new(ServiceFlinkCreateApplicationVersionOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkDeleteApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkDeleteApplicationOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkDeleteApplication", "DELETE", path, nil)
	out := new(ServiceFlinkDeleteApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkDeleteApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, applicationVersionId string) (*ServiceFlinkDeleteApplicationVersionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/version/%s", project, serviceName, applicationId, applicationVersionId)
	b, err := h.doer.Do(ctx, "ServiceFlinkDeleteApplicationVersion", "DELETE", path, nil)
	out := new(ServiceFlinkDeleteApplicationVersionOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkGetApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkGetApplicationOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkGetApplication", "GET", path, nil)
	out := new(ServiceFlinkGetApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkGetApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, applicationVersionId string) (*ServiceFlinkGetApplicationVersionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/version/%s", project, serviceName, applicationId, applicationVersionId)
	b, err := h.doer.Do(ctx, "ServiceFlinkGetApplicationVersion", "GET", path, nil)
	out := new(ServiceFlinkGetApplicationVersionOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkListApplications(ctx context.Context, project string, serviceName string) ([]Application, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceFlinkListApplications", "GET", path, nil)
	out := new(serviceFlinkListApplicationsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Applications, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkUpdateApplication(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkUpdateApplicationIn) (*ServiceFlinkUpdateApplicationOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkUpdateApplication", "PUT", path, in)
	out := new(ServiceFlinkUpdateApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkValidateApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkValidateApplicationVersionIn) (*ServiceFlinkValidateApplicationVersionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/version/validate", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkValidateApplicationVersion", "POST", path, in)
	out := new(ServiceFlinkValidateApplicationVersionOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type Application struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	CreatedBy string     `json:"created_by,omitempty"`
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy string     `json:"updated_by,omitempty"`
}
type ApplicationVersion struct {
	Sinks     []Sink   `json:"sinks"`
	Sources   []Source `json:"sources"`
	Statement string   `json:"statement"`
}
type ApplicationVersionItem struct {
	CreatedAt time.Time    `json:"created_at"`
	CreatedBy string       `json:"created_by"`
	Id        string       `json:"id"`
	Sinks     []SinkItem   `json:"sinks"`
	Sources   []SourceItem `json:"sources"`
	Statement string       `json:"statement"`
	Version   int          `json:"version"`
}
type Column struct {
	DataType  string `json:"data_type"`
	Extras    string `json:"extras,omitempty"`
	Key       string `json:"key,omitempty"`
	Name      string `json:"name"`
	Nullable  bool   `json:"nullable"`
	Watermark string `json:"watermark,omitempty"`
}
type CurrentDeployment struct {
	CreatedAt         time.Time  `json:"created_at"`
	CreatedBy         string     `json:"created_by"`
	ErrorMsg          string     `json:"error_msg,omitempty"`
	Id                string     `json:"id"`
	JobId             string     `json:"job_id,omitempty"`
	LastSavepoint     string     `json:"last_savepoint,omitempty"`
	Parallelism       int        `json:"parallelism"`
	RestartEnabled    bool       `json:"restart_enabled"`
	StartingSavepoint string     `json:"starting_savepoint,omitempty"`
	Status            StatusType `json:"status"`
	VersionId         string     `json:"version_id"`
}
type Position struct {
	CharacterNumber    int `json:"character_number"`
	EndCharacterNumber int `json:"end_character_number"`
	EndLineNumber      int `json:"end_line_number"`
	LineNumber         int `json:"line_number"`
}
type ServiceFlinkCreateApplicationIn struct {
	ApplicationVersion *ApplicationVersion `json:"application_version,omitempty"`
	Name               string              `json:"name"`
}
type ServiceFlinkCreateApplicationOut struct {
	ApplicationVersions []ApplicationVersionItem `json:"application_versions"`
	CreatedAt           time.Time                `json:"created_at"`
	CreatedBy           string                   `json:"created_by"`
	CurrentDeployment   *CurrentDeployment       `json:"current_deployment,omitempty"`
	Id                  string                   `json:"id"`
	Name                string                   `json:"name"`
	UpdatedAt           time.Time                `json:"updated_at"`
	UpdatedBy           string                   `json:"updated_by"`
}
type ServiceFlinkCreateApplicationVersionIn struct {
	Sinks     []Sink   `json:"sinks"`
	Sources   []Source `json:"sources"`
	Statement string   `json:"statement"`
}
type ServiceFlinkCreateApplicationVersionOut struct {
	CreatedAt time.Time    `json:"created_at"`
	CreatedBy string       `json:"created_by"`
	Id        string       `json:"id"`
	Sinks     []SinkItem   `json:"sinks"`
	Sources   []SourceItem `json:"sources"`
	Statement string       `json:"statement"`
	Version   int          `json:"version"`
}
type ServiceFlinkDeleteApplicationOut struct {
	ApplicationVersions []ApplicationVersionItem `json:"application_versions"`
	CreatedAt           time.Time                `json:"created_at"`
	CreatedBy           string                   `json:"created_by"`
	CurrentDeployment   *CurrentDeployment       `json:"current_deployment,omitempty"`
	Id                  string                   `json:"id"`
	Name                string                   `json:"name"`
	UpdatedAt           time.Time                `json:"updated_at"`
	UpdatedBy           string                   `json:"updated_by"`
}
type ServiceFlinkDeleteApplicationVersionOut struct {
	CreatedAt time.Time    `json:"created_at"`
	CreatedBy string       `json:"created_by"`
	Id        string       `json:"id"`
	Sinks     []SinkItem   `json:"sinks"`
	Sources   []SourceItem `json:"sources"`
	Statement string       `json:"statement"`
	Version   int          `json:"version"`
}
type ServiceFlinkGetApplicationOut struct {
	ApplicationVersions []ApplicationVersionItem `json:"application_versions"`
	CreatedAt           time.Time                `json:"created_at"`
	CreatedBy           string                   `json:"created_by"`
	CurrentDeployment   *CurrentDeployment       `json:"current_deployment,omitempty"`
	Id                  string                   `json:"id"`
	Name                string                   `json:"name"`
	UpdatedAt           time.Time                `json:"updated_at"`
	UpdatedBy           string                   `json:"updated_by"`
}
type ServiceFlinkGetApplicationVersionOut struct {
	CreatedAt time.Time    `json:"created_at"`
	CreatedBy string       `json:"created_by"`
	Id        string       `json:"id"`
	Sinks     []SinkItem   `json:"sinks"`
	Sources   []SourceItem `json:"sources"`
	Statement string       `json:"statement"`
	Version   int          `json:"version"`
}
type serviceFlinkListApplicationsOut struct {
	Applications []Application `json:"applications"`
}
type ServiceFlinkUpdateApplicationIn struct {
	Name string `json:"name"`
}
type ServiceFlinkUpdateApplicationOut struct {
	ApplicationVersions []ApplicationVersionItem `json:"application_versions"`
	CreatedAt           time.Time                `json:"created_at"`
	CreatedBy           string                   `json:"created_by"`
	CurrentDeployment   *CurrentDeployment       `json:"current_deployment,omitempty"`
	Id                  string                   `json:"id"`
	Name                string                   `json:"name"`
	UpdatedAt           time.Time                `json:"updated_at"`
	UpdatedBy           string                   `json:"updated_by"`
}
type ServiceFlinkValidateApplicationVersionIn struct {
	Sinks     []Sink   `json:"sinks"`
	Sources   []Source `json:"sources"`
	Statement string   `json:"statement,omitempty"`
}
type ServiceFlinkValidateApplicationVersionOut struct {
	Sinks          []ServiceFlinkValidateApplicationVersionOutSinkItem   `json:"sinks"`
	Sources        []ServiceFlinkValidateApplicationVersionOutSourceItem `json:"sources"`
	Statement      string                                                `json:"statement,omitempty"`
	StatementError *StatementError                                       `json:"statement_error,omitempty"`
}
type ServiceFlinkValidateApplicationVersionOutSinkItem struct {
	Columns       []Column       `json:"columns"`
	CreateTable   string         `json:"create_table"`
	IntegrationId string         `json:"integration_id,omitempty"`
	Message       string         `json:"message,omitempty"`
	Options       map[string]any `json:"options,omitempty"`
	Position      *Position      `json:"position,omitempty"`
	TableName     string         `json:"table_name,omitempty"`
}
type ServiceFlinkValidateApplicationVersionOutSourceItem struct {
	Columns       []Column       `json:"columns"`
	CreateTable   string         `json:"create_table"`
	IntegrationId string         `json:"integration_id,omitempty"`
	Message       string         `json:"message,omitempty"`
	Options       map[string]any `json:"options,omitempty"`
	Position      *Position      `json:"position,omitempty"`
	TableName     string         `json:"table_name,omitempty"`
}
type Sink struct {
	CreateTable   string `json:"create_table"`
	IntegrationId string `json:"integration_id,omitempty"`
}
type SinkItem struct {
	Columns       []Column       `json:"columns"`
	CreateTable   string         `json:"create_table"`
	IntegrationId string         `json:"integration_id,omitempty"`
	Options       map[string]any `json:"options"`
	TableId       string         `json:"table_id"`
	TableName     string         `json:"table_name"`
}
type Source struct {
	CreateTable   string `json:"create_table"`
	IntegrationId string `json:"integration_id,omitempty"`
}
type SourceItem struct {
	Columns       []Column       `json:"columns"`
	CreateTable   string         `json:"create_table"`
	IntegrationId string         `json:"integration_id,omitempty"`
	Options       map[string]any `json:"options"`
	TableId       string         `json:"table_id"`
	TableName     string         `json:"table_name"`
}
type StatementError struct {
	Message  string    `json:"message"`
	Position *Position `json:"position,omitempty"`
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
