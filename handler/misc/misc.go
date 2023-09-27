// Code generated by Aiven. DO NOT EDIT.

package misc

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// ListClouds list available cloud platforms
	// ListClouds GET /clouds
	// https://api.aiven.io/doc/#operation/ListClouds
	ListClouds(ctx context.Context) ([]Cloud, error)

	// ListProjectClouds list available cloud platforms for a project
	// ListProjectClouds GET /project/{project}/clouds
	// https://api.aiven.io/doc/#operation/ListProjectClouds
	ListProjectClouds(ctx context.Context, project string) ([]Cloud, error)

	// PublicStaticIPAvailabilityList list static IP address cloud availability and prices for a tenant
	// PublicStaticIPAvailabilityList GET /tenants/{tenant}/static-ip-availability
	// https://api.aiven.io/doc/#operation/PublicStaticIPAvailabilityList
	PublicStaticIPAvailabilityList(ctx context.Context, tenant string) ([]StaticIpAddressAvailability, error)
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

func (h *handler) ListClouds(ctx context.Context) ([]Cloud, error) {
	path := fmt.Sprintf("/clouds")
	b, err := h.doer.Do(ctx, "ListClouds", "GET", path, nil)
	out := new(ListCloudsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Clouds, nil
}
func (h *handler) ListProjectClouds(ctx context.Context, project string) ([]Cloud, error) {
	path := fmt.Sprintf("/project/%s/clouds", project)
	b, err := h.doer.Do(ctx, "ListProjectClouds", "GET", path, nil)
	out := new(ListProjectCloudsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Clouds, nil
}
func (h *handler) PublicStaticIPAvailabilityList(ctx context.Context, tenant string) ([]StaticIpAddressAvailability, error) {
	path := fmt.Sprintf("/tenants/%s/static-ip-availability", tenant)
	b, err := h.doer.Do(ctx, "PublicStaticIPAvailabilityList", "GET", path, nil)
	out := new(PublicStaticIpavailabilityListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.StaticIpAddressAvailability, nil
}

type Cloud struct {
	CloudDescription    string   `json:"cloud_description,omitempty"`
	GeoLatitude         *float64 `json:"geo_latitude,omitempty"`
	GeoLongitude        *float64 `json:"geo_longitude,omitempty"`
	GeoRegion           string   `json:"geo_region"`
	CloudName           string   `json:"cloud_name"`
	Provider            string   `json:"provider,omitempty"`
	ProviderDescription string   `json:"provider_description,omitempty"`
}
type ListCloudsOut struct {
	Clouds []Cloud `json:"clouds"`
}
type ListProjectCloudsOut struct {
	Clouds []Cloud `json:"clouds"`
}
type PublicStaticIpavailabilityListOut struct {
	StaticIpAddressAvailability []StaticIpAddressAvailability `json:"static_ip_address_availability"`
}
type StaticIpAddressAvailability struct {
	CloudName string `json:"cloud_name"`
	PriceUsd  string `json:"price_usd"`
}
