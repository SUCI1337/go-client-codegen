// Code generated by Aiven. DO NOT EDIT.

package kafka

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// ServiceKafkaAclAdd add a Kafka ACL entry
	// POST /project/{project}/service/{service_name}/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaAclAdd
	ServiceKafkaAclAdd(ctx context.Context, project string, serviceName string, in *ServiceKafkaAclAddIn) ([]Acl, error)

	// ServiceKafkaAclDelete delete a Kafka ACL entry
	// DELETE /project/{project}/service/{service_name}/acl/{kafka_acl_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaAclDelete
	ServiceKafkaAclDelete(ctx context.Context, project string, serviceName string, kafkaAclId string) ([]Acl, error)

	// ServiceKafkaAclList list Kafka ACL entries
	// GET /project/{project}/service/{service_name}/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaAclList
	ServiceKafkaAclList(ctx context.Context, project string, serviceName string) ([]Acl, error)

	// ServiceKafkaQuotaCreate create Kafka quota
	// POST /project/{project}/service/{service_name}/quota
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaQuotaCreate
	ServiceKafkaQuotaCreate(ctx context.Context, project string, serviceName string, in *ServiceKafkaQuotaCreateIn) error

	// ServiceKafkaQuotaDelete delete Kafka quota
	// DELETE /project/{project}/service/{service_name}/quota
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaQuotaDelete
	ServiceKafkaQuotaDelete(ctx context.Context, project string, serviceName string) error

	// ServiceKafkaQuotaDescribe describe Specific Kafka quotas
	// GET /project/{project}/service/{service_name}/quota/describe
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaQuotaDescribe
	ServiceKafkaQuotaDescribe(ctx context.Context, project string, serviceName string) (*Quota, error)

	// ServiceKafkaQuotaList list Kafka quotas
	// GET /project/{project}/service/{service_name}/quota
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaQuotaList
	ServiceKafkaQuotaList(ctx context.Context, project string, serviceName string) ([]Quota, error)

	// ServiceKafkaTieredStorageStorageUsageByTopic get the Kafka tiered storage object storage usage by topic
	// GET /project/{project}/service/{service_name}/kafka/tiered-storage/storage-usage/by-topic
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTieredStorageStorageUsageByTopic
	ServiceKafkaTieredStorageStorageUsageByTopic(ctx context.Context, project string, serviceName string) (map[string]any, error)

	// ServiceKafkaTieredStorageStorageUsageTotal get the Kafka tiered storage total object storage usage
	// GET /project/{project}/service/{service_name}/kafka/tiered-storage/storage-usage/total
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTieredStorageStorageUsageTotal
	ServiceKafkaTieredStorageStorageUsageTotal(ctx context.Context, project string, serviceName string) (int, error)

	// ServiceKafkaTieredStorageSummary get the Kafka tiered storage summary
	// GET /project/{project}/service/{service_name}/kafka/tiered-storage/summary
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTieredStorageSummary
	ServiceKafkaTieredStorageSummary(ctx context.Context, project string, serviceName string) (*ServiceKafkaTieredStorageSummaryOut, error)
}

func NewHandler(doer doer) KafkaHandler {
	return KafkaHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type KafkaHandler struct {
	doer doer
}

func (h *KafkaHandler) ServiceKafkaAclAdd(ctx context.Context, project string, serviceName string, in *ServiceKafkaAclAddIn) ([]Acl, error) {
	path := fmt.Sprintf("/project/%s/service/%s/acl", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaAclAdd", "POST", path, in)
	out := new(serviceKafkaAclAddOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaHandler) ServiceKafkaAclDelete(ctx context.Context, project string, serviceName string, kafkaAclId string) ([]Acl, error) {
	path := fmt.Sprintf("/project/%s/service/%s/acl/%s", project, serviceName, kafkaAclId)
	b, err := h.doer.Do(ctx, "ServiceKafkaAclDelete", "DELETE", path, nil)
	out := new(serviceKafkaAclDeleteOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaHandler) ServiceKafkaAclList(ctx context.Context, project string, serviceName string) ([]Acl, error) {
	path := fmt.Sprintf("/project/%s/service/%s/acl", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaAclList", "GET", path, nil)
	out := new(serviceKafkaAclListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaHandler) ServiceKafkaQuotaCreate(ctx context.Context, project string, serviceName string, in *ServiceKafkaQuotaCreateIn) error {
	path := fmt.Sprintf("/project/%s/service/%s/quota", project, serviceName)
	_, err := h.doer.Do(ctx, "ServiceKafkaQuotaCreate", "POST", path, in)
	return err
}
func (h *KafkaHandler) ServiceKafkaQuotaDelete(ctx context.Context, project string, serviceName string) error {
	path := fmt.Sprintf("/project/%s/service/%s/quota", project, serviceName)
	_, err := h.doer.Do(ctx, "ServiceKafkaQuotaDelete", "DELETE", path, nil)
	return err
}
func (h *KafkaHandler) ServiceKafkaQuotaDescribe(ctx context.Context, project string, serviceName string) (*Quota, error) {
	path := fmt.Sprintf("/project/%s/service/%s/quota/describe", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaQuotaDescribe", "GET", path, nil)
	out := new(serviceKafkaQuotaDescribeOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Quota, nil
}
func (h *KafkaHandler) ServiceKafkaQuotaList(ctx context.Context, project string, serviceName string) ([]Quota, error) {
	path := fmt.Sprintf("/project/%s/service/%s/quota", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaQuotaList", "GET", path, nil)
	out := new(serviceKafkaQuotaListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Quotas, nil
}
func (h *KafkaHandler) ServiceKafkaTieredStorageStorageUsageByTopic(ctx context.Context, project string, serviceName string) (map[string]any, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/tiered-storage/storage-usage/by-topic", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTieredStorageStorageUsageByTopic", "GET", path, nil)
	out := new(serviceKafkaTieredStorageStorageUsageByTopicOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.StorageUsage, nil
}
func (h *KafkaHandler) ServiceKafkaTieredStorageStorageUsageTotal(ctx context.Context, project string, serviceName string) (int, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/tiered-storage/storage-usage/total", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTieredStorageStorageUsageTotal", "GET", path, nil)
	out := new(serviceKafkaTieredStorageStorageUsageTotalOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return 0, err
	}
	return out.TotalStorageUsage, nil
}
func (h *KafkaHandler) ServiceKafkaTieredStorageSummary(ctx context.Context, project string, serviceName string) (*ServiceKafkaTieredStorageSummaryOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/tiered-storage/summary", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTieredStorageSummary", "GET", path, nil)
	out := new(ServiceKafkaTieredStorageSummaryOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type Acl struct {
	Id         string         `json:"id,omitempty"`
	Permission PermissionType `json:"permission"`
	Topic      string         `json:"topic"`
	Username   string         `json:"username"`
}
type Hourly struct {
	EstimatedCost   string `json:"estimated_cost,omitempty"`
	HourStart       string `json:"hour_start"`
	PeakStoredBytes int    `json:"peak_stored_bytes"`
}
type PermissionType string

const (
	PermissionTypeAdmin     PermissionType = "admin"
	PermissionTypeRead      PermissionType = "read"
	PermissionTypeReadwrite PermissionType = "readwrite"
	PermissionTypeWrite     PermissionType = "write"
)

type Quota struct {
	ClientId          string  `json:"client-id,omitempty"`
	ConsumerByteRate  float64 `json:"consumer_byte_rate"`
	ProducerByteRate  float64 `json:"producer_byte_rate"`
	RequestPercentage float64 `json:"request_percentage"`
	User              string  `json:"user"`
}
type ServiceKafkaAclAddIn struct {
	Permission PermissionType `json:"permission"`
	Topic      string         `json:"topic"`
	Username   string         `json:"username"`
}
type serviceKafkaAclAddOut struct {
	Acl []Acl `json:"acl"`
}
type serviceKafkaAclDeleteOut struct {
	Acl []Acl `json:"acl"`
}
type serviceKafkaAclListOut struct {
	Acl []Acl `json:"acl"`
}
type ServiceKafkaQuotaCreateIn struct {
	ClientId          string   `json:"client-id,omitempty"`
	ConsumerByteRate  *float64 `json:"consumer_byte_rate,omitempty"`
	ProducerByteRate  *float64 `json:"producer_byte_rate,omitempty"`
	RequestPercentage *float64 `json:"request_percentage,omitempty"`
	User              string   `json:"user,omitempty"`
}
type serviceKafkaQuotaDescribeOut struct {
	Quota *Quota `json:"quota"`
}
type serviceKafkaQuotaListOut struct {
	Quotas []Quota `json:"quotas"`
}
type serviceKafkaTieredStorageStorageUsageByTopicOut struct {
	StorageUsage map[string]any `json:"storage_usage"`
}
type serviceKafkaTieredStorageStorageUsageTotalOut struct {
	TotalStorageUsage int `json:"total_storage_usage"`
}
type ServiceKafkaTieredStorageSummaryOut struct {
	CurrentCost         string               `json:"current_cost"`
	ForecastedCost      string               `json:"forecasted_cost"`
	StorageUsageHistory *StorageUsageHistory `json:"storage_usage_history"`
	TotalStorageUsage   int                  `json:"total_storage_usage"`
}
type StorageUsageHistory struct {
	Hourly []Hourly `json:"hourly"`
}
