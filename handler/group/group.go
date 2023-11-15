// Code generated by Aiven. DO NOT EDIT.

package group

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// UserGroupCreate create a group
	// UserGroupCreate POST /organization/{organization_id}/user-groups
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupCreate
	UserGroupCreate(ctx context.Context, organizationId string, in *UserGroupCreateIn) (*UserGroupCreateOut, error)

	// UserGroupDelete delete a group
	// UserGroupDelete DELETE /organization/{organization_id}/user-groups/{user_group_id}
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupDelete
	UserGroupDelete(ctx context.Context, organizationId string, userGroupId string) error

	// UserGroupGet show group details
	// UserGroupGet GET /organization/{organization_id}/user-groups/{user_group_id}
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupGet
	UserGroupGet(ctx context.Context, organizationId string, userGroupId string) (*UserGroupGetOut, error)

	// UserGroupMemberList list group members
	// UserGroupMemberList GET /organization/{organization_id}/user-groups/{user_group_id}/members
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupMemberList
	UserGroupMemberList(ctx context.Context, organizationId string, userGroupId string) ([]Member, error)

	// UserGroupMembersUpdate add or remove group members
	// UserGroupMembersUpdate PATCH /organization/{organization_id}/user-groups/{user_group_id}/members
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupMembersUpdate
	UserGroupMembersUpdate(ctx context.Context, organizationId string, userGroupId string, in *UserGroupMembersUpdateIn) error

	// UserGroupUpdate update a group
	// UserGroupUpdate PATCH /organization/{organization_id}/user-groups/{user_group_id}
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupUpdate
	UserGroupUpdate(ctx context.Context, organizationId string, userGroupId string, in *UserGroupUpdateIn) (*UserGroupUpdateOut, error)

	// UserGroupsList list groups
	// UserGroupsList GET /organization/{organization_id}/user-groups
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupsList
	UserGroupsList(ctx context.Context, organizationId string) ([]UserGroup, error)
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

func (h *handler) UserGroupCreate(ctx context.Context, organizationId string, in *UserGroupCreateIn) (*UserGroupCreateOut, error) {
	path := fmt.Sprintf("/organization/%s/user-groups", organizationId)
	b, err := h.doer.Do(ctx, "UserGroupCreate", "POST", path, in)
	out := new(UserGroupCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) UserGroupDelete(ctx context.Context, organizationId string, userGroupId string) error {
	path := fmt.Sprintf("/organization/%s/user-groups/%s", organizationId, userGroupId)
	_, err := h.doer.Do(ctx, "UserGroupDelete", "DELETE", path, nil)
	return err
}
func (h *handler) UserGroupGet(ctx context.Context, organizationId string, userGroupId string) (*UserGroupGetOut, error) {
	path := fmt.Sprintf("/organization/%s/user-groups/%s", organizationId, userGroupId)
	b, err := h.doer.Do(ctx, "UserGroupGet", "GET", path, nil)
	out := new(UserGroupGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) UserGroupMemberList(ctx context.Context, organizationId string, userGroupId string) ([]Member, error) {
	path := fmt.Sprintf("/organization/%s/user-groups/%s/members", organizationId, userGroupId)
	b, err := h.doer.Do(ctx, "UserGroupMemberList", "GET", path, nil)
	out := new(UserGroupMemberListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Members, nil
}
func (h *handler) UserGroupMembersUpdate(ctx context.Context, organizationId string, userGroupId string, in *UserGroupMembersUpdateIn) error {
	path := fmt.Sprintf("/organization/%s/user-groups/%s/members", organizationId, userGroupId)
	_, err := h.doer.Do(ctx, "UserGroupMembersUpdate", "PATCH", path, in)
	return err
}
func (h *handler) UserGroupUpdate(ctx context.Context, organizationId string, userGroupId string, in *UserGroupUpdateIn) (*UserGroupUpdateOut, error) {
	path := fmt.Sprintf("/organization/%s/user-groups/%s", organizationId, userGroupId)
	b, err := h.doer.Do(ctx, "UserGroupUpdate", "PATCH", path, in)
	out := new(UserGroupUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) UserGroupsList(ctx context.Context, organizationId string) ([]UserGroup, error) {
	path := fmt.Sprintf("/organization/%s/user-groups", organizationId)
	b, err := h.doer.Do(ctx, "UserGroupsList", "GET", path, nil)
	out := new(UserGroupsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.UserGroups, nil
}

type Member struct {
	LastActivityTime *time.Time `json:"last_activity_time,omitempty"`
	UserId           string     `json:"user_id"`
	UserInfo         *UserInfo  `json:"user_info,omitempty"`
}
type OperationType string

const (
	OperationTypeAddMembers    OperationType = "add_members"
	OperationTypeRemoveMembers OperationType = "remove_members"
)

func OperationTypeChoices() []string {
	return []string{"add_members", "remove_members"}
}

type UserGroup struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	UserGroupId   string    `json:"user_group_id"`
	MemberCount   int       `json:"member_count"`
	UserGroupName string    `json:"user_group_name"`
	UpdateTime    time.Time `json:"update_time"`
}
type UserGroupCreateIn struct {
	Description   string `json:"description"`
	UserGroupName string `json:"user_group_name"`
}
type UserGroupCreateOut struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	UpdateTime    time.Time `json:"update_time"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
}
type UserGroupGetOut struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	UpdateTime    time.Time `json:"update_time"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
}
type UserGroupMemberListOut struct {
	Members []Member `json:"members"`
}
type UserGroupMembersUpdateIn struct {
	MemberIds []string      `json:"member_ids"`
	Operation OperationType `json:"operation"`
}
type UserGroupUpdateIn struct {
	Description   string `json:"description,omitempty"`
	UserGroupName string `json:"user_group_name,omitempty"`
}
type UserGroupUpdateOut struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	UpdateTime    time.Time `json:"update_time"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
}
type UserGroupsListOut struct {
	UserGroups []UserGroup `json:"user_groups"`
}
type UserInfo struct {
	City                   string    `json:"city,omitempty"`
	Country                string    `json:"country,omitempty"`
	CreateTime             time.Time `json:"create_time"`
	Department             string    `json:"department,omitempty"`
	JobTitle               string    `json:"job_title,omitempty"`
	ManagedByScim          *bool     `json:"managed_by_scim,omitempty"`
	ManagingOrganizationId string    `json:"managing_organization_id,omitempty"`
	RealName               string    `json:"real_name"`
	State                  string    `json:"state"`
	UserEmail              string    `json:"user_email"`
}
