// Code generated by Aiven. DO NOT EDIT.

package kafkatopic

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// Create create a Kafka topic
	// ServiceKafkaTopicCreate POST /project/{project}/service/{service_name}/topic
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicCreate
	Create(ctx context.Context, project string, serviceName string, in *CreateIn) error

	// Delete delete a Kafka topic
	// ServiceKafkaTopicDelete DELETE /project/{project}/service/{service_name}/topic/{topic_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicDelete
	Delete(ctx context.Context, project string, serviceName string, topicName string) error

	// Get get Kafka topic info
	// ServiceKafkaTopicGet GET /project/{project}/service/{service_name}/topic/{topic_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicGet
	Get(ctx context.Context, project string, serviceName string, topicName string) (*Topic, error)

	// List get Kafka topic list
	// ServiceKafkaTopicList GET /project/{project}/service/{service_name}/topic
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicList
	List(ctx context.Context, project string, serviceName string) ([]TopicItem, error)

	// MessageList list kafka topic messages
	// ServiceKafkaTopicMessageList POST /project/{project}/service/{service_name}/kafka/rest/topics/{topic_name}/messages
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicMessageList
	MessageList(ctx context.Context, project string, serviceName string, topicName string, in *MessageListIn) ([]Message, error)

	// MessageProduce produce message into a kafka topic
	// ServiceKafkaTopicMessageProduce POST /project/{project}/service/{service_name}/kafka/rest/topics/{topic_name}/produce
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicMessageProduce
	MessageProduce(ctx context.Context, project string, serviceName string, topicName string, in *MessageProduceIn) (*MessageProduceOut, error)

	// Update update a Kafka topic
	// ServiceKafkaTopicUpdate PUT /project/{project}/service/{service_name}/topic/{topic_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicUpdate
	Update(ctx context.Context, project string, serviceName string, topicName string, in *UpdateIn) error
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

func (h *handler) Create(ctx context.Context, project string, serviceName string, in *CreateIn) error {
	path := fmt.Sprintf("/project/%s/service/%s/topic", project, serviceName)
	_, err := h.doer.Do(ctx, "ServiceKafkaTopicCreate", "POST", path, in)
	return err
}
func (h *handler) Delete(ctx context.Context, project string, serviceName string, topicName string) error {
	path := fmt.Sprintf("/project/%s/service/%s/topic/%s", project, serviceName, topicName)
	_, err := h.doer.Do(ctx, "ServiceKafkaTopicDelete", "DELETE", path, nil)
	return err
}
func (h *handler) Get(ctx context.Context, project string, serviceName string, topicName string) (*Topic, error) {
	path := fmt.Sprintf("/project/%s/service/%s/topic/%s", project, serviceName, topicName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTopicGet", "GET", path, nil)
	out := new(GetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Topic, nil
}
func (h *handler) List(ctx context.Context, project string, serviceName string) ([]TopicItem, error) {
	path := fmt.Sprintf("/project/%s/service/%s/topic", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTopicList", "GET", path, nil)
	out := new(ListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Topics, nil
}
func (h *handler) MessageList(ctx context.Context, project string, serviceName string, topicName string, in *MessageListIn) ([]Message, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/rest/topics/%s/messages", project, serviceName, topicName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTopicMessageList", "POST", path, in)
	out := new(MessageListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Messages, nil
}
func (h *handler) MessageProduce(ctx context.Context, project string, serviceName string, topicName string, in *MessageProduceIn) (*MessageProduceOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/rest/topics/%s/produce", project, serviceName, topicName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTopicMessageProduce", "POST", path, in)
	out := new(MessageProduceOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) Update(ctx context.Context, project string, serviceName string, topicName string, in *UpdateIn) error {
	path := fmt.Sprintf("/project/%s/service/%s/topic/%s", project, serviceName, topicName)
	_, err := h.doer.Do(ctx, "ServiceKafkaTopicUpdate", "PUT", path, in)
	return err
}

type CleanupPolicy struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    string     `json:"value,omitempty"`
}
type CleanupPolicyType string

const (
	CleanupPolicyTypeDelete        CleanupPolicyType = "delete"
	CleanupPolicyTypeCompact       CleanupPolicyType = "compact"
	CleanupPolicyTypeCompactdelete CleanupPolicyType = "compact,delete"
)

func CleanupPolicyTypeChoices() []string {
	return []string{"delete", "compact", "compact,delete"}
}

type CompressionType string

const (
	CompressionTypeSnappy       CompressionType = "snappy"
	CompressionTypeGzip         CompressionType = "gzip"
	CompressionTypeLz4          CompressionType = "lz4"
	CompressionTypeProducer     CompressionType = "producer"
	CompressionTypeUncompressed CompressionType = "uncompressed"
	CompressionTypeZstd         CompressionType = "zstd"
)

func CompressionTypeChoices() []string {
	return []string{"snappy", "gzip", "lz4", "producer", "uncompressed", "zstd"}
}

type Config struct {
	CleanupPolicy                   CleanupPolicyType        `json:"cleanup_policy,omitempty"`
	CompressionType                 CompressionType          `json:"compression_type,omitempty"`
	DeleteRetentionMs               *int                     `json:"delete_retention_ms,omitempty"`
	FileDeleteDelayMs               *int                     `json:"file_delete_delay_ms,omitempty"`
	FlushMessages                   *int                     `json:"flush_messages,omitempty"`
	FlushMs                         *int                     `json:"flush_ms,omitempty"`
	IndexIntervalBytes              *int                     `json:"index_interval_bytes,omitempty"`
	LocalRetentionBytes             *int                     `json:"local_retention_bytes,omitempty"`
	LocalRetentionMs                *int                     `json:"local_retention_ms,omitempty"`
	MaxCompactionLagMs              *int                     `json:"max_compaction_lag_ms,omitempty"`
	MaxMessageBytes                 *int                     `json:"max_message_bytes,omitempty"`
	MessageDownconversionEnable     *bool                    `json:"message_downconversion_enable,omitempty"`
	MessageFormatVersion            MessageFormatVersionType `json:"message_format_version,omitempty"`
	MessageTimestampDifferenceMaxMs *int                     `json:"message_timestamp_difference_max_ms,omitempty"`
	MessageTimestampType            MessageTimestampType     `json:"message_timestamp_type,omitempty"`
	MinCleanableDirtyRatio          *float64                 `json:"min_cleanable_dirty_ratio,omitempty"`
	MinCompactionLagMs              *int                     `json:"min_compaction_lag_ms,omitempty"`
	MinInsyncReplicas               *int                     `json:"min_insync_replicas,omitempty"`
	Preallocate                     *bool                    `json:"preallocate,omitempty"`
	RemoteStorageEnable             *bool                    `json:"remote_storage_enable,omitempty"`
	RetentionBytes                  *int                     `json:"retention_bytes,omitempty"`
	RetentionMs                     *int                     `json:"retention_ms,omitempty"`
	SegmentBytes                    *int                     `json:"segment_bytes,omitempty"`
	SegmentIndexBytes               *int                     `json:"segment_index_bytes,omitempty"`
	SegmentJitterMs                 *int                     `json:"segment_jitter_ms,omitempty"`
	SegmentMs                       *int                     `json:"segment_ms,omitempty"`
}
type ConsumerGroup struct {
	GroupName string `json:"group_name"`
	Offset    int    `json:"offset"`
}
type CreateIn struct {
	CleanupPolicy     CleanupPolicyType `json:"cleanup_policy,omitempty"`
	Config            *Config           `json:"config,omitempty"`
	MinInsyncReplicas *int              `json:"min_insync_replicas,omitempty"`
	Partitions        *int              `json:"partitions,omitempty"`
	Replication       *int              `json:"replication,omitempty"`
	RetentionBytes    *int              `json:"retention_bytes,omitempty"`
	RetentionHours    *int              `json:"retention_hours,omitempty"`
	Tags              []Tag             `json:"tags,omitempty"`
	TopicName         string            `json:"topic_name"`
}
type DeleteRetentionMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type FileDeleteDelayMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type FlushMessages struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type FlushMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type FormatType string

const (
	FormatTypeBinary     FormatType = "binary"
	FormatTypeJson       FormatType = "json"
	FormatTypeAvro       FormatType = "avro"
	FormatTypeProtobuf   FormatType = "protobuf"
	FormatTypeJsonschema FormatType = "jsonschema"
)

func FormatTypeChoices() []string {
	return []string{"binary", "json", "avro", "protobuf", "jsonschema"}
}

type GetOut struct {
	Topic *Topic `json:"topic"`
}
type IndexIntervalBytes struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type ListOut struct {
	Topics []TopicItem `json:"topics"`
}
type LocalRetentionBytes struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type LocalRetentionMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type MaxCompactionLagMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type MaxMessageBytes struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type Message struct {
	Key       map[string]any `json:"key,omitempty"`
	Offset    *int           `json:"offset,omitempty"`
	Partition *int           `json:"partition,omitempty"`
	Topic     string         `json:"topic,omitempty"`
	Value     map[string]any `json:"value,omitempty"`
}
type MessageDownconversionEnable struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *bool      `json:"value,omitempty"`
}
type MessageFormatVersion struct {
	Value    MessageFormatVersionValueType `json:"value,omitempty"`
	Source   SourceType                    `json:"source,omitempty"`
	Synonyms []Synonym                     `json:"synonyms,omitempty"`
}
type MessageFormatVersionType string

const (
	MessageFormatVersionType080     MessageFormatVersionType = "0.8.0"
	MessageFormatVersionType081     MessageFormatVersionType = "0.8.1"
	MessageFormatVersionType082     MessageFormatVersionType = "0.8.2"
	MessageFormatVersionType090     MessageFormatVersionType = "0.9.0"
	MessageFormatVersionType0100    MessageFormatVersionType = "0.10.0"
	MessageFormatVersionType0100Iv0 MessageFormatVersionType = "0.10.0-IV0"
	MessageFormatVersionType0100Iv1 MessageFormatVersionType = "0.10.0-IV1"
	MessageFormatVersionType0101    MessageFormatVersionType = "0.10.1"
	MessageFormatVersionType0101Iv0 MessageFormatVersionType = "0.10.1-IV0"
	MessageFormatVersionType0101Iv1 MessageFormatVersionType = "0.10.1-IV1"
	MessageFormatVersionType0101Iv2 MessageFormatVersionType = "0.10.1-IV2"
	MessageFormatVersionType0102    MessageFormatVersionType = "0.10.2"
	MessageFormatVersionType0102Iv0 MessageFormatVersionType = "0.10.2-IV0"
	MessageFormatVersionType0110    MessageFormatVersionType = "0.11.0"
	MessageFormatVersionType0110Iv0 MessageFormatVersionType = "0.11.0-IV0"
	MessageFormatVersionType0110Iv1 MessageFormatVersionType = "0.11.0-IV1"
	MessageFormatVersionType0110Iv2 MessageFormatVersionType = "0.11.0-IV2"
	MessageFormatVersionType10      MessageFormatVersionType = "1.0"
	MessageFormatVersionType10Iv0   MessageFormatVersionType = "1.0-IV0"
	MessageFormatVersionType11      MessageFormatVersionType = "1.1"
	MessageFormatVersionType11Iv0   MessageFormatVersionType = "1.1-IV0"
	MessageFormatVersionType20      MessageFormatVersionType = "2.0"
	MessageFormatVersionType20Iv0   MessageFormatVersionType = "2.0-IV0"
	MessageFormatVersionType20Iv1   MessageFormatVersionType = "2.0-IV1"
	MessageFormatVersionType21      MessageFormatVersionType = "2.1"
	MessageFormatVersionType21Iv0   MessageFormatVersionType = "2.1-IV0"
	MessageFormatVersionType21Iv1   MessageFormatVersionType = "2.1-IV1"
	MessageFormatVersionType21Iv2   MessageFormatVersionType = "2.1-IV2"
	MessageFormatVersionType22      MessageFormatVersionType = "2.2"
	MessageFormatVersionType22Iv0   MessageFormatVersionType = "2.2-IV0"
	MessageFormatVersionType22Iv1   MessageFormatVersionType = "2.2-IV1"
	MessageFormatVersionType23      MessageFormatVersionType = "2.3"
	MessageFormatVersionType23Iv0   MessageFormatVersionType = "2.3-IV0"
	MessageFormatVersionType23Iv1   MessageFormatVersionType = "2.3-IV1"
	MessageFormatVersionType24      MessageFormatVersionType = "2.4"
	MessageFormatVersionType24Iv0   MessageFormatVersionType = "2.4-IV0"
	MessageFormatVersionType24Iv1   MessageFormatVersionType = "2.4-IV1"
	MessageFormatVersionType25      MessageFormatVersionType = "2.5"
	MessageFormatVersionType25Iv0   MessageFormatVersionType = "2.5-IV0"
	MessageFormatVersionType26      MessageFormatVersionType = "2.6"
	MessageFormatVersionType26Iv0   MessageFormatVersionType = "2.6-IV0"
	MessageFormatVersionType27      MessageFormatVersionType = "2.7"
	MessageFormatVersionType27Iv0   MessageFormatVersionType = "2.7-IV0"
	MessageFormatVersionType27Iv1   MessageFormatVersionType = "2.7-IV1"
	MessageFormatVersionType27Iv2   MessageFormatVersionType = "2.7-IV2"
	MessageFormatVersionType28      MessageFormatVersionType = "2.8"
	MessageFormatVersionType28Iv0   MessageFormatVersionType = "2.8-IV0"
	MessageFormatVersionType28Iv1   MessageFormatVersionType = "2.8-IV1"
	MessageFormatVersionType30      MessageFormatVersionType = "3.0"
	MessageFormatVersionType30Iv0   MessageFormatVersionType = "3.0-IV0"
	MessageFormatVersionType30Iv1   MessageFormatVersionType = "3.0-IV1"
	MessageFormatVersionType31      MessageFormatVersionType = "3.1"
	MessageFormatVersionType31Iv0   MessageFormatVersionType = "3.1-IV0"
	MessageFormatVersionType32      MessageFormatVersionType = "3.2"
	MessageFormatVersionType32Iv0   MessageFormatVersionType = "3.2-IV0"
	MessageFormatVersionType33      MessageFormatVersionType = "3.3"
	MessageFormatVersionType33Iv0   MessageFormatVersionType = "3.3-IV0"
	MessageFormatVersionType33Iv1   MessageFormatVersionType = "3.3-IV1"
	MessageFormatVersionType33Iv2   MessageFormatVersionType = "3.3-IV2"
	MessageFormatVersionType33Iv3   MessageFormatVersionType = "3.3-IV3"
	MessageFormatVersionType34      MessageFormatVersionType = "3.4"
	MessageFormatVersionType34Iv0   MessageFormatVersionType = "3.4-IV0"
	MessageFormatVersionType35      MessageFormatVersionType = "3.5"
	MessageFormatVersionType35Iv0   MessageFormatVersionType = "3.5-IV0"
	MessageFormatVersionType35Iv1   MessageFormatVersionType = "3.5-IV1"
	MessageFormatVersionType35Iv2   MessageFormatVersionType = "3.5-IV2"
	MessageFormatVersionType36      MessageFormatVersionType = "3.6"
	MessageFormatVersionType36Iv0   MessageFormatVersionType = "3.6-IV0"
	MessageFormatVersionType36Iv1   MessageFormatVersionType = "3.6-IV1"
	MessageFormatVersionType36Iv2   MessageFormatVersionType = "3.6-IV2"
)

func MessageFormatVersionTypeChoices() []string {
	return []string{"0.8.0", "0.8.1", "0.8.2", "0.9.0", "0.10.0", "0.10.0-IV0", "0.10.0-IV1", "0.10.1", "0.10.1-IV0", "0.10.1-IV1", "0.10.1-IV2", "0.10.2", "0.10.2-IV0", "0.11.0", "0.11.0-IV0", "0.11.0-IV1", "0.11.0-IV2", "1.0", "1.0-IV0", "1.1", "1.1-IV0", "2.0", "2.0-IV0", "2.0-IV1", "2.1", "2.1-IV0", "2.1-IV1", "2.1-IV2", "2.2", "2.2-IV0", "2.2-IV1", "2.3", "2.3-IV0", "2.3-IV1", "2.4", "2.4-IV0", "2.4-IV1", "2.5", "2.5-IV0", "2.6", "2.6-IV0", "2.7", "2.7-IV0", "2.7-IV1", "2.7-IV2", "2.8", "2.8-IV0", "2.8-IV1", "3.0", "3.0-IV0", "3.0-IV1", "3.1", "3.1-IV0", "3.2", "3.2-IV0", "3.3", "3.3-IV0", "3.3-IV1", "3.3-IV2", "3.3-IV3", "3.4", "3.4-IV0", "3.5", "3.5-IV0", "3.5-IV1", "3.5-IV2", "3.6", "3.6-IV0", "3.6-IV1", "3.6-IV2"}
}

type MessageFormatVersionValueType string

const (
	MessageFormatVersionValueType080     MessageFormatVersionValueType = "0.8.0"
	MessageFormatVersionValueType081     MessageFormatVersionValueType = "0.8.1"
	MessageFormatVersionValueType082     MessageFormatVersionValueType = "0.8.2"
	MessageFormatVersionValueType090     MessageFormatVersionValueType = "0.9.0"
	MessageFormatVersionValueType0100    MessageFormatVersionValueType = "0.10.0"
	MessageFormatVersionValueType0100Iv0 MessageFormatVersionValueType = "0.10.0-IV0"
	MessageFormatVersionValueType0100Iv1 MessageFormatVersionValueType = "0.10.0-IV1"
	MessageFormatVersionValueType0101    MessageFormatVersionValueType = "0.10.1"
	MessageFormatVersionValueType0101Iv0 MessageFormatVersionValueType = "0.10.1-IV0"
	MessageFormatVersionValueType0101Iv1 MessageFormatVersionValueType = "0.10.1-IV1"
	MessageFormatVersionValueType0101Iv2 MessageFormatVersionValueType = "0.10.1-IV2"
	MessageFormatVersionValueType0102    MessageFormatVersionValueType = "0.10.2"
	MessageFormatVersionValueType0102Iv0 MessageFormatVersionValueType = "0.10.2-IV0"
	MessageFormatVersionValueType0110    MessageFormatVersionValueType = "0.11.0"
	MessageFormatVersionValueType0110Iv0 MessageFormatVersionValueType = "0.11.0-IV0"
	MessageFormatVersionValueType0110Iv1 MessageFormatVersionValueType = "0.11.0-IV1"
	MessageFormatVersionValueType0110Iv2 MessageFormatVersionValueType = "0.11.0-IV2"
	MessageFormatVersionValueType10      MessageFormatVersionValueType = "1.0"
	MessageFormatVersionValueType10Iv0   MessageFormatVersionValueType = "1.0-IV0"
	MessageFormatVersionValueType11      MessageFormatVersionValueType = "1.1"
	MessageFormatVersionValueType11Iv0   MessageFormatVersionValueType = "1.1-IV0"
	MessageFormatVersionValueType20      MessageFormatVersionValueType = "2.0"
	MessageFormatVersionValueType20Iv0   MessageFormatVersionValueType = "2.0-IV0"
	MessageFormatVersionValueType20Iv1   MessageFormatVersionValueType = "2.0-IV1"
	MessageFormatVersionValueType21      MessageFormatVersionValueType = "2.1"
	MessageFormatVersionValueType21Iv0   MessageFormatVersionValueType = "2.1-IV0"
	MessageFormatVersionValueType21Iv1   MessageFormatVersionValueType = "2.1-IV1"
	MessageFormatVersionValueType21Iv2   MessageFormatVersionValueType = "2.1-IV2"
	MessageFormatVersionValueType22      MessageFormatVersionValueType = "2.2"
	MessageFormatVersionValueType22Iv0   MessageFormatVersionValueType = "2.2-IV0"
	MessageFormatVersionValueType22Iv1   MessageFormatVersionValueType = "2.2-IV1"
	MessageFormatVersionValueType23      MessageFormatVersionValueType = "2.3"
	MessageFormatVersionValueType23Iv0   MessageFormatVersionValueType = "2.3-IV0"
	MessageFormatVersionValueType23Iv1   MessageFormatVersionValueType = "2.3-IV1"
	MessageFormatVersionValueType24      MessageFormatVersionValueType = "2.4"
	MessageFormatVersionValueType24Iv0   MessageFormatVersionValueType = "2.4-IV0"
	MessageFormatVersionValueType24Iv1   MessageFormatVersionValueType = "2.4-IV1"
	MessageFormatVersionValueType25      MessageFormatVersionValueType = "2.5"
	MessageFormatVersionValueType25Iv0   MessageFormatVersionValueType = "2.5-IV0"
	MessageFormatVersionValueType26      MessageFormatVersionValueType = "2.6"
	MessageFormatVersionValueType26Iv0   MessageFormatVersionValueType = "2.6-IV0"
	MessageFormatVersionValueType27      MessageFormatVersionValueType = "2.7"
	MessageFormatVersionValueType27Iv0   MessageFormatVersionValueType = "2.7-IV0"
	MessageFormatVersionValueType27Iv1   MessageFormatVersionValueType = "2.7-IV1"
	MessageFormatVersionValueType27Iv2   MessageFormatVersionValueType = "2.7-IV2"
	MessageFormatVersionValueType28      MessageFormatVersionValueType = "2.8"
	MessageFormatVersionValueType28Iv0   MessageFormatVersionValueType = "2.8-IV0"
	MessageFormatVersionValueType28Iv1   MessageFormatVersionValueType = "2.8-IV1"
	MessageFormatVersionValueType30      MessageFormatVersionValueType = "3.0"
	MessageFormatVersionValueType30Iv0   MessageFormatVersionValueType = "3.0-IV0"
	MessageFormatVersionValueType30Iv1   MessageFormatVersionValueType = "3.0-IV1"
	MessageFormatVersionValueType31      MessageFormatVersionValueType = "3.1"
	MessageFormatVersionValueType31Iv0   MessageFormatVersionValueType = "3.1-IV0"
	MessageFormatVersionValueType32      MessageFormatVersionValueType = "3.2"
	MessageFormatVersionValueType32Iv0   MessageFormatVersionValueType = "3.2-IV0"
	MessageFormatVersionValueType33      MessageFormatVersionValueType = "3.3"
	MessageFormatVersionValueType33Iv0   MessageFormatVersionValueType = "3.3-IV0"
	MessageFormatVersionValueType33Iv1   MessageFormatVersionValueType = "3.3-IV1"
	MessageFormatVersionValueType33Iv2   MessageFormatVersionValueType = "3.3-IV2"
	MessageFormatVersionValueType33Iv3   MessageFormatVersionValueType = "3.3-IV3"
	MessageFormatVersionValueType34      MessageFormatVersionValueType = "3.4"
	MessageFormatVersionValueType34Iv0   MessageFormatVersionValueType = "3.4-IV0"
	MessageFormatVersionValueType35      MessageFormatVersionValueType = "3.5"
	MessageFormatVersionValueType35Iv0   MessageFormatVersionValueType = "3.5-IV0"
	MessageFormatVersionValueType35Iv1   MessageFormatVersionValueType = "3.5-IV1"
	MessageFormatVersionValueType35Iv2   MessageFormatVersionValueType = "3.5-IV2"
	MessageFormatVersionValueType36      MessageFormatVersionValueType = "3.6"
	MessageFormatVersionValueType36Iv0   MessageFormatVersionValueType = "3.6-IV0"
	MessageFormatVersionValueType36Iv1   MessageFormatVersionValueType = "3.6-IV1"
	MessageFormatVersionValueType36Iv2   MessageFormatVersionValueType = "3.6-IV2"
)

type MessageListIn struct {
	Format     FormatType     `json:"format,omitempty"`
	MaxBytes   *int           `json:"max_bytes,omitempty"`
	Partitions map[string]any `json:"partitions"`
	Timeout    *int           `json:"timeout,omitempty"`
}
type MessageListOut struct {
	Messages []Message `json:"messages,omitempty"`
}
type MessageProduceIn struct {
	Format        FormatType `json:"format"`
	KeySchema     string     `json:"key_schema,omitempty"`
	KeySchemaId   *int       `json:"key_schema_id,omitempty"`
	Records       []Record   `json:"records"`
	ValueSchema   string     `json:"value_schema,omitempty"`
	ValueSchemaId *int       `json:"value_schema_id,omitempty"`
}
type MessageProduceOut struct {
	KeySchemaId   *int     `json:"key_schema_id,omitempty"`
	Offsets       []Offset `json:"offsets,omitempty"`
	ValueSchemaId *int     `json:"value_schema_id,omitempty"`
}
type MessageTimestampDifferenceMaxMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type MessageTimestampType string

const (
	MessageTimestampTypeCreateTime    MessageTimestampType = "CreateTime"
	MessageTimestampTypeLogAppendTime MessageTimestampType = "LogAppendTime"
)

func MessageTimestampTypeChoices() []string {
	return []string{"CreateTime", "LogAppendTime"}
}

type MinCleanableDirtyRatio struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *float64   `json:"value,omitempty"`
}
type MinCompactionLagMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type MinInsyncReplicas struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type Offset struct {
	Error     string `json:"error,omitempty"`
	ErrorCode *int   `json:"error_code,omitempty"`
	Offset    *int   `json:"offset,omitempty"`
	Partition *int   `json:"partition,omitempty"`
}
type Partition struct {
	ConsumerGroups []ConsumerGroup `json:"consumer_groups"`
	EarliestOffset int             `json:"earliest_offset"`
	Isr            int             `json:"isr"`
	LatestOffset   int             `json:"latest_offset"`
	Partition      int             `json:"partition"`
	Size           int             `json:"size"`
}
type Preallocate struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *bool      `json:"value,omitempty"`
}
type Record struct {
	Key       map[string]any `json:"key,omitempty"`
	Partition *int           `json:"partition,omitempty"`
	Value     map[string]any `json:"value,omitempty"`
}
type RemoteStorageEnable struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *bool      `json:"value,omitempty"`
}
type RetentionBytes struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type RetentionMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type SegmentBytes struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type SegmentIndexBytes struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type SegmentJitterMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type SegmentMs struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *int       `json:"value,omitempty"`
}
type SourceType string

const (
	SourceTypeUnknownConfig              SourceType = "unknown_config"
	SourceTypeTopicConfig                SourceType = "topic_config"
	SourceTypeDynamicBrokerConfig        SourceType = "dynamic_broker_config"
	SourceTypeDynamicDefaultBrokerConfig SourceType = "dynamic_default_broker_config"
	SourceTypeStaticBrokerConfig         SourceType = "static_broker_config"
	SourceTypeDefaultConfig              SourceType = "default_config"
	SourceTypeDynamicBrokerLoggerConfig  SourceType = "dynamic_broker_logger_config"
)

type StateType string

const (
	StateTypeActive      StateType = "ACTIVE"
	StateTypeConfiguring StateType = "CONFIGURING"
	StateTypeDeleting    StateType = "DELETING"
)

type Synonym struct {
	Name   string     `json:"name,omitempty"`
	Source SourceType `json:"source,omitempty"`
	Value  *bool      `json:"value,omitempty"`
}
type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Topic struct {
	CleanupPolicy     string       `json:"cleanup_policy"`
	MinInsyncReplicas int          `json:"min_insync_replicas"`
	TopicName         string       `json:"topic_name"`
	Partitions        []Partition  `json:"partitions"`
	Replication       int          `json:"replication"`
	RetentionBytes    int          `json:"retention_bytes"`
	RetentionHours    int          `json:"retention_hours"`
	State             StateType    `json:"state"`
	Tags              []Tag        `json:"tags"`
	Config            *TopicConfig `json:"config"`
}
type TopicConfig struct {
	CleanupPolicy                   *CleanupPolicy                   `json:"cleanup_policy,omitempty"`
	DeleteRetentionMs               *DeleteRetentionMs               `json:"delete_retention_ms,omitempty"`
	FileDeleteDelayMs               *FileDeleteDelayMs               `json:"file_delete_delay_ms,omitempty"`
	FlushMessages                   *FlushMessages                   `json:"flush_messages,omitempty"`
	FlushMs                         *FlushMs                         `json:"flush_ms,omitempty"`
	IndexIntervalBytes              *IndexIntervalBytes              `json:"index_interval_bytes,omitempty"`
	LocalRetentionBytes             *LocalRetentionBytes             `json:"local_retention_bytes,omitempty"`
	LocalRetentionMs                *LocalRetentionMs                `json:"local_retention_ms,omitempty"`
	MaxCompactionLagMs              *MaxCompactionLagMs              `json:"max_compaction_lag_ms,omitempty"`
	MaxMessageBytes                 *MaxMessageBytes                 `json:"max_message_bytes,omitempty"`
	MessageDownconversionEnable     *MessageDownconversionEnable     `json:"message_downconversion_enable,omitempty"`
	MessageFormatVersion            *MessageFormatVersion            `json:"message_format_version,omitempty"`
	MessageTimestampDifferenceMaxMs *MessageTimestampDifferenceMaxMs `json:"message_timestamp_difference_max_ms,omitempty"`
	MinCleanableDirtyRatio          *MinCleanableDirtyRatio          `json:"min_cleanable_dirty_ratio,omitempty"`
	MinCompactionLagMs              *MinCompactionLagMs              `json:"min_compaction_lag_ms,omitempty"`
	MinInsyncReplicas               *MinInsyncReplicas               `json:"min_insync_replicas,omitempty"`
	Preallocate                     *Preallocate                     `json:"preallocate,omitempty"`
	RemoteStorageEnable             *RemoteStorageEnable             `json:"remote_storage_enable,omitempty"`
	RetentionBytes                  *RetentionBytes                  `json:"retention_bytes,omitempty"`
	RetentionMs                     *RetentionMs                     `json:"retention_ms,omitempty"`
	SegmentBytes                    *SegmentBytes                    `json:"segment_bytes,omitempty"`
	SegmentIndexBytes               *SegmentIndexBytes               `json:"segment_index_bytes,omitempty"`
	SegmentJitterMs                 *SegmentJitterMs                 `json:"segment_jitter_ms,omitempty"`
	SegmentMs                       *SegmentMs                       `json:"segment_ms,omitempty"`
	CompressionType                 *TopicConfigCompressionType      `json:"compression_type,omitempty"`
	MessageTimestampType            *TopicConfigMessageTimestampType `json:"message_timestamp_type,omitempty"`
	UncleanLeaderElectionEnable     *UncleanLeaderElectionEnable     `json:"unclean_leader_election_enable,omitempty"`
}
type TopicConfigCompressionType struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    ValueType  `json:"value,omitempty"`
}
type TopicConfigMessageTimestampType struct {
	Source   SourceType                               `json:"source,omitempty"`
	Synonyms []Synonym                                `json:"synonyms,omitempty"`
	Value    TopicConfigMessageTimestampTypeValueType `json:"value,omitempty"`
}
type TopicConfigMessageTimestampTypeValueType string

const (
	TopicConfigMessageTimestampTypeValueTypeCreateTime    TopicConfigMessageTimestampTypeValueType = "CreateTime"
	TopicConfigMessageTimestampTypeValueTypeLogAppendTime TopicConfigMessageTimestampTypeValueType = "LogAppendTime"
)

type TopicItem struct {
	CleanupPolicy       string    `json:"cleanup_policy"`
	MinInsyncReplicas   int       `json:"min_insync_replicas"`
	Partitions          int       `json:"partitions"`
	RemoteStorageEnable *bool     `json:"remote_storage_enable,omitempty"`
	Replication         int       `json:"replication"`
	RetentionBytes      int       `json:"retention_bytes"`
	RetentionHours      int       `json:"retention_hours"`
	State               StateType `json:"state"`
	Tags                []Tag     `json:"tags"`
	TopicName           string    `json:"topic_name"`
}
type UncleanLeaderElectionEnable struct {
	Source   SourceType `json:"source,omitempty"`
	Synonyms []Synonym  `json:"synonyms,omitempty"`
	Value    *bool      `json:"value,omitempty"`
}
type UpdateIn struct {
	Config            *Config `json:"config,omitempty"`
	MinInsyncReplicas *int    `json:"min_insync_replicas,omitempty"`
	Partitions        *int    `json:"partitions,omitempty"`
	Replication       *int    `json:"replication,omitempty"`
	RetentionBytes    *int    `json:"retention_bytes,omitempty"`
	RetentionHours    *int    `json:"retention_hours,omitempty"`
	Tags              []Tag   `json:"tags,omitempty"`
}
type ValueType string

const (
	ValueTypeSnappy       ValueType = "snappy"
	ValueTypeGzip         ValueType = "gzip"
	ValueTypeLz4          ValueType = "lz4"
	ValueTypeProducer     ValueType = "producer"
	ValueTypeUncompressed ValueType = "uncompressed"
	ValueTypeZstd         ValueType = "zstd"
)
