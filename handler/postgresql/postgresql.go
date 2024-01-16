// Code generated by Aiven. DO NOT EDIT.

package postgresql

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// PGServiceAvailableExtensions list PostgreSQL extensions that can be loaded with CREATE EXTENSION in this service
	// GET /project/{project}/service/{service_name}/pg/available-extensions
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/PGServiceAvailableExtensions
	PGServiceAvailableExtensions(ctx context.Context, project string, serviceName string) ([]Extension, error)

	// PGServiceQueryStatistics fetch PostgreSQL service query statistics
	// POST /project/{project}/service/{service_name}/pg/query/stats
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/PGServiceQueryStatistics
	PGServiceQueryStatistics(ctx context.Context, project string, serviceName string, in *PgserviceQueryStatisticsIn) ([]Query, error)

	// PgAvailableExtensions list PostgreSQL extensions available for this tenant grouped by PG version
	// GET /tenants/{tenant}/pg-available-extensions
	// https://api.aiven.io/doc/#tag/Service/operation/PgAvailableExtensions
	PgAvailableExtensions(ctx context.Context, tenant string) ([]Pg, error)

	// ServicePGBouncerCreate create a new connection pool for service
	// POST /project/{project}/service/{service_name}/connection_pool
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/ServicePGBouncerCreate
	ServicePGBouncerCreate(ctx context.Context, project string, serviceName string, in *ServicePgbouncerCreateIn) error

	// ServicePGBouncerDelete delete a connection pool
	// DELETE /project/{project}/service/{service_name}/connection_pool/{pool_name}
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/ServicePGBouncerDelete
	ServicePGBouncerDelete(ctx context.Context, project string, serviceName string, poolName string) error

	// ServicePGBouncerUpdate update a connection pool
	// PUT /project/{project}/service/{service_name}/connection_pool/{pool_name}
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/ServicePGBouncerUpdate
	ServicePGBouncerUpdate(ctx context.Context, project string, serviceName string, poolName string, in *ServicePgbouncerUpdateIn) error
}

func NewHandler(doer doer) PostgreSQLHandler {
	return PostgreSQLHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type PostgreSQLHandler struct {
	doer doer
}

func (h *PostgreSQLHandler) PGServiceAvailableExtensions(ctx context.Context, project string, serviceName string) ([]Extension, error) {
	path := fmt.Sprintf("/project/%s/service/%s/pg/available-extensions", project, serviceName)
	b, err := h.doer.Do(ctx, "PGServiceAvailableExtensions", "GET", path, nil)
	out := new(pgserviceAvailableExtensionsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Extensions, nil
}
func (h *PostgreSQLHandler) PGServiceQueryStatistics(ctx context.Context, project string, serviceName string, in *PgserviceQueryStatisticsIn) ([]Query, error) {
	path := fmt.Sprintf("/project/%s/service/%s/pg/query/stats", project, serviceName)
	b, err := h.doer.Do(ctx, "PGServiceQueryStatistics", "POST", path, in)
	out := new(pgserviceQueryStatisticsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Queries, nil
}
func (h *PostgreSQLHandler) PgAvailableExtensions(ctx context.Context, tenant string) ([]Pg, error) {
	path := fmt.Sprintf("/tenants/%s/pg-available-extensions", tenant)
	b, err := h.doer.Do(ctx, "PgAvailableExtensions", "GET", path, nil)
	out := new(pgAvailableExtensionsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Pg, nil
}
func (h *PostgreSQLHandler) ServicePGBouncerCreate(ctx context.Context, project string, serviceName string, in *ServicePgbouncerCreateIn) error {
	path := fmt.Sprintf("/project/%s/service/%s/connection_pool", project, serviceName)
	_, err := h.doer.Do(ctx, "ServicePGBouncerCreate", "POST", path, in)
	return err
}
func (h *PostgreSQLHandler) ServicePGBouncerDelete(ctx context.Context, project string, serviceName string, poolName string) error {
	path := fmt.Sprintf("/project/%s/service/%s/connection_pool/%s", project, serviceName, poolName)
	_, err := h.doer.Do(ctx, "ServicePGBouncerDelete", "DELETE", path, nil)
	return err
}
func (h *PostgreSQLHandler) ServicePGBouncerUpdate(ctx context.Context, project string, serviceName string, poolName string, in *ServicePgbouncerUpdateIn) error {
	path := fmt.Sprintf("/project/%s/service/%s/connection_pool/%s", project, serviceName, poolName)
	_, err := h.doer.Do(ctx, "ServicePGBouncerUpdate", "PUT", path, in)
	return err
}

type Extension struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}
type Pg struct {
	Extensions []Extension `json:"extensions"`
	Version    string      `json:"version"`
}
type pgAvailableExtensionsOut struct {
	Pg []Pg `json:"pg"`
}
type pgserviceAvailableExtensionsOut struct {
	Extensions []Extension `json:"extensions"`
}
type PgserviceQueryStatisticsIn struct {
	Limit   *int   `json:"limit,omitempty"`
	Offset  *int   `json:"offset,omitempty"`
	OrderBy string `json:"order_by,omitempty"`
}
type pgserviceQueryStatisticsOut struct {
	Queries []Query `json:"queries"`
}
type PoolModeType string

const (
	PoolModeTypeSession     PoolModeType = "session"
	PoolModeTypeTransaction PoolModeType = "transaction"
	PoolModeTypeStatement   PoolModeType = "statement"
)

func PoolModeTypeChoices() []string {
	return []string{"session", "transaction", "statement"}
}

type Query struct {
	BlkReadTime       *float64 `json:"blk_read_time,omitempty"`
	BlkWriteTime      *float64 `json:"blk_write_time,omitempty"`
	Calls             *float64 `json:"calls,omitempty"`
	DatabaseName      string   `json:"database_name,omitempty"`
	LocalBlksDirtied  *float64 `json:"local_blks_dirtied,omitempty"`
	LocalBlksHit      *float64 `json:"local_blks_hit,omitempty"`
	LocalBlksRead     *float64 `json:"local_blks_read,omitempty"`
	LocalBlksWritten  *float64 `json:"local_blks_written,omitempty"`
	MaxExecTime       *float64 `json:"max_exec_time,omitempty"`
	MaxPlanTime       *float64 `json:"max_plan_time,omitempty"`
	MaxTime           *float64 `json:"max_time,omitempty"`
	MeanExecTime      *float64 `json:"mean_exec_time,omitempty"`
	MeanPlanTime      *float64 `json:"mean_plan_time,omitempty"`
	MeanTime          *float64 `json:"mean_time,omitempty"`
	MinExecTime       *float64 `json:"min_exec_time,omitempty"`
	MinPlanTime       *float64 `json:"min_plan_time,omitempty"`
	MinTime           *float64 `json:"min_time,omitempty"`
	Query             string   `json:"query,omitempty"`
	Queryid           *float64 `json:"queryid,omitempty"`
	Rows              *float64 `json:"rows,omitempty"`
	SharedBlksDirtied *float64 `json:"shared_blks_dirtied,omitempty"`
	SharedBlksHit     *float64 `json:"shared_blks_hit,omitempty"`
	SharedBlksRead    *float64 `json:"shared_blks_read,omitempty"`
	SharedBlksWritten *float64 `json:"shared_blks_written,omitempty"`
	StddevExecTime    *float64 `json:"stddev_exec_time,omitempty"`
	StddevPlanTime    *float64 `json:"stddev_plan_time,omitempty"`
	StddevTime        *float64 `json:"stddev_time,omitempty"`
	TempBlksRead      *float64 `json:"temp_blks_read,omitempty"`
	TempBlksWritten   *float64 `json:"temp_blks_written,omitempty"`
	TotalExecTime     *float64 `json:"total_exec_time,omitempty"`
	TotalPlanTime     *float64 `json:"total_plan_time,omitempty"`
	TotalTime         *float64 `json:"total_time,omitempty"`
	UserName          string   `json:"user_name,omitempty"`
	WalBytes          string   `json:"wal_bytes,omitempty"`
	WalFpi            *float64 `json:"wal_fpi,omitempty"`
	WalRecords        *float64 `json:"wal_records,omitempty"`
}
type ServicePgbouncerCreateIn struct {
	Database string       `json:"database"`
	PoolMode PoolModeType `json:"pool_mode,omitempty"`
	PoolName string       `json:"pool_name"`
	PoolSize *int         `json:"pool_size,omitempty"`
	Username string       `json:"username,omitempty"`
}
type ServicePgbouncerUpdateIn struct {
	Database string       `json:"database,omitempty"`
	PoolMode PoolModeType `json:"pool_mode,omitempty"`
	PoolSize *int         `json:"pool_size,omitempty"`
	Username string       `json:"username,omitempty"`
}
