// Code generated by Aiven. DO NOT EDIT.

package staticip

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// Create create static IP address
	// StaticIPCreate POST /project/{project}/static-ips
	// https://api.aiven.io/doc/#tag/StaticIP/operation/StaticIPCreate
	Create(ctx context.Context, project string, in *CreateIn) (*CreateOut, error)

	// List list static IP addresses
	// StaticIPList GET /project/{project}/static-ips
	// https://api.aiven.io/doc/#tag/StaticIP/operation/StaticIPList
	List(ctx context.Context, project string) ([]StaticIp, error)

	// ProjectStaticIPAssociate associate a static IP address with a service
	// ProjectStaticIPAssociate POST /project/{project}/static-ips/{static_ip_address_id}/association
	// https://api.aiven.io/doc/#tag/StaticIP/operation/ProjectStaticIPAssociate
	ProjectStaticIPAssociate(ctx context.Context, project string, addressId string, in *ProjectStaticIpassociateIn) (*ProjectStaticIpassociateOut, error)

	// ProjectStaticIPAvailabilityList list static IP address cloud availability and prices for a project
	// ProjectStaticIPAvailabilityList GET /project/{project}/static-ip-availability
	// https://api.aiven.io/doc/#tag/StaticIP/operation/ProjectStaticIPAvailabilityList
	ProjectStaticIPAvailabilityList(ctx context.Context, project string) ([]StaticIpAddressAvailability, error)

	// ProjectStaticIPDissociate dissociate a static IP address from a service
	// ProjectStaticIPDissociate DELETE /project/{project}/static-ips/{static_ip_address_id}/association
	// https://api.aiven.io/doc/#tag/StaticIP/operation/ProjectStaticIPDissociate
	ProjectStaticIPDissociate(ctx context.Context, project string, addressId string) (*ProjectStaticIpdissociateOut, error)

	// ProjectStaticIPPatch update a static IP address configuration
	// ProjectStaticIPPatch PATCH /project/{project}/static-ips/{static_ip_address_id}
	// https://api.aiven.io/doc/#tag/StaticIP/operation/ProjectStaticIPPatch
	ProjectStaticIPPatch(ctx context.Context, project string, addressId string, in *ProjectStaticIppatchIn) (*ProjectStaticIppatchOut, error)
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

func (h *handler) Create(ctx context.Context, project string, in *CreateIn) (*CreateOut, error) {
	path := fmt.Sprintf("/project/%s/static-ips", project)
	b, err := h.doer.Do(ctx, "StaticIPCreate", "POST", path, in)
	out := new(CreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) List(ctx context.Context, project string) ([]StaticIp, error) {
	path := fmt.Sprintf("/project/%s/static-ips", project)
	b, err := h.doer.Do(ctx, "StaticIPList", "GET", path, nil)
	out := new(ListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.StaticIps, nil
}
func (h *handler) ProjectStaticIPAssociate(ctx context.Context, project string, addressId string, in *ProjectStaticIpassociateIn) (*ProjectStaticIpassociateOut, error) {
	path := fmt.Sprintf("/project/%s/static-ips/%s/association", project, addressId)
	b, err := h.doer.Do(ctx, "ProjectStaticIPAssociate", "POST", path, in)
	out := new(ProjectStaticIpassociateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) ProjectStaticIPAvailabilityList(ctx context.Context, project string) ([]StaticIpAddressAvailability, error) {
	path := fmt.Sprintf("/project/%s/static-ip-availability", project)
	b, err := h.doer.Do(ctx, "ProjectStaticIPAvailabilityList", "GET", path, nil)
	out := new(ProjectStaticIpavailabilityListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.StaticIpAddressAvailability, nil
}
func (h *handler) ProjectStaticIPDissociate(ctx context.Context, project string, addressId string) (*ProjectStaticIpdissociateOut, error) {
	path := fmt.Sprintf("/project/%s/static-ips/%s/association", project, addressId)
	b, err := h.doer.Do(ctx, "ProjectStaticIPDissociate", "DELETE", path, nil)
	out := new(ProjectStaticIpdissociateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) ProjectStaticIPPatch(ctx context.Context, project string, addressId string, in *ProjectStaticIppatchIn) (*ProjectStaticIppatchOut, error) {
	path := fmt.Sprintf("/project/%s/static-ips/%s", project, addressId)
	b, err := h.doer.Do(ctx, "ProjectStaticIPPatch", "PATCH", path, in)
	out := new(ProjectStaticIppatchOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type CreateIn struct {
	CloudName             string `json:"cloud_name"`
	TerminationProtection *bool  `json:"termination_protection,omitempty"`
}
type CreateOut struct {
	CloudName             string    `json:"cloud_name"`
	IpAddress             string    `json:"ip_address"`
	ServiceName           string    `json:"service_name"`
	State                 StateType `json:"state"`
	StaticIpAddressId     string    `json:"static_ip_address_id"`
	TerminationProtection bool      `json:"termination_protection"`
}
type ListOut struct {
	StaticIps []StaticIp `json:"static_ips"`
}
type ProjectStaticIpassociateIn struct {
	ServiceName string `json:"service_name"`
}
type ProjectStaticIpassociateOut struct {
	CloudName             string    `json:"cloud_name"`
	IpAddress             string    `json:"ip_address"`
	ServiceName           string    `json:"service_name"`
	State                 StateType `json:"state"`
	StaticIpAddressId     string    `json:"static_ip_address_id"`
	TerminationProtection bool      `json:"termination_protection"`
}
type ProjectStaticIpavailabilityListOut struct {
	StaticIpAddressAvailability []StaticIpAddressAvailability `json:"static_ip_address_availability"`
}
type ProjectStaticIpdissociateOut struct {
	CloudName             string    `json:"cloud_name"`
	IpAddress             string    `json:"ip_address"`
	ServiceName           string    `json:"service_name"`
	State                 StateType `json:"state"`
	StaticIpAddressId     string    `json:"static_ip_address_id"`
	TerminationProtection bool      `json:"termination_protection"`
}
type ProjectStaticIppatchIn struct {
	TerminationProtection *bool `json:"termination_protection,omitempty"`
}
type ProjectStaticIppatchOut struct {
	CloudName             string    `json:"cloud_name"`
	IpAddress             string    `json:"ip_address"`
	ServiceName           string    `json:"service_name"`
	State                 StateType `json:"state"`
	StaticIpAddressId     string    `json:"static_ip_address_id"`
	TerminationProtection bool      `json:"termination_protection"`
}
type StateType string

const (
	StateTypeCreating  StateType = "creating"
	StateTypeCreated   StateType = "created"
	StateTypeAvailable StateType = "available"
	StateTypeAssigned  StateType = "assigned"
	StateTypeDeleting  StateType = "deleting"
	StateTypeDeleted   StateType = "deleted"
)

type StaticIp struct {
	StaticIpAddressId     string    `json:"static_ip_address_id"`
	CloudName             string    `json:"cloud_name"`
	IpAddress             string    `json:"ip_address"`
	ServiceName           string    `json:"service_name"`
	State                 StateType `json:"state"`
	TerminationProtection bool      `json:"termination_protection"`
}
type StaticIpAddressAvailability struct {
	CloudName string `json:"cloud_name"`
	PriceUsd  string `json:"price_usd"`
}
