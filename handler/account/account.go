// Code generated by Aiven. DO NOT EDIT.

package account

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// AccountAttachPaymentMethod attach payment method for account
	// POST /account/{account_id}/payment_methods
	// https://api.aiven.io/doc/#tag/Account/operation/AccountAttachPaymentMethod
	AccountAttachPaymentMethod(ctx context.Context, accountId string, in *AccountAttachPaymentMethodIn) (*Card, error)

	// AccountBillingGroupList list account billing groups
	// GET /account/{account_id}/billing-group
	// https://api.aiven.io/doc/#tag/Account/operation/AccountBillingGroupList
	AccountBillingGroupList(ctx context.Context, accountId string) ([]AccountBillingGroup, error)

	// AccountCreate create a new account
	// POST /account
	// https://api.aiven.io/doc/#tag/Account/operation/AccountCreate
	AccountCreate(ctx context.Context, in *AccountCreateIn) (*Account, error)

	// AccountDelete delete empty account
	// DELETE /account/{account_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountDelete
	AccountDelete(ctx context.Context, accountId string) error

	// AccountEventList list account events
	// GET /account/{account_id}/events
	// https://api.aiven.io/doc/#tag/Account/operation/AccountEventList
	AccountEventList(ctx context.Context, accountId string) ([]Event, error)

	// AccountGet get account details
	// GET /account/{account_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountGet
	AccountGet(ctx context.Context, accountId string) (*Account, error)

	// AccountList list accounts you have access to
	// GET /account
	// https://api.aiven.io/doc/#tag/Account/operation/AccountList
	AccountList(ctx context.Context) ([]Account, error)

	// AccountMove move an existing organization unitself
	// PUT /account/{account_id}/parent_account
	// https://api.aiven.io/doc/#tag/Account/operation/AccountMove
	AccountMove(ctx context.Context, accountId string, in *AccountMoveIn) (*Account, error)

	// AccountPaymentMethodDelete delete credit card attached to the account as a payment method
	// DELETE /account/{account_id}/payment_method/{card_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountPaymentMethodDelete
	AccountPaymentMethodDelete(ctx context.Context, accountId string, cardId string) error

	// AccountPaymentMethodsList list credit cards attached as a payment method to the account
	// GET /account/{account_id}/payment_methods
	// https://api.aiven.io/doc/#tag/Account/operation/AccountPaymentMethodsList
	AccountPaymentMethodsList(ctx context.Context, accountId string) ([]CardItem, error)

	// AccountProjectsList list projects belonging to account
	// GET /account/{account_id}/projects
	// https://api.aiven.io/doc/#tag/Account/operation/AccountProjectsList
	AccountProjectsList(ctx context.Context, accountId string) (*AccountProjectsListOut, error)

	// AccountUpdate update existing account
	// PUT /account/{account_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountUpdate
	AccountUpdate(ctx context.Context, accountId string, in *AccountUpdateIn) (*Account, error)

	// AccountUserProjectsList list projects associated with this account that user has access to
	// GET /account/{account_id}/user/{user_id}/projects
	// https://api.aiven.io/doc/#tag/Account/operation/AccountUserProjectsList
	AccountUserProjectsList(ctx context.Context, accountId string, userId string) ([]UserProject, error)

	// AccountUsersSearch list/search users who are members of any team on this account
	// POST /account/{account_id}/users/search
	// https://api.aiven.io/doc/#tag/Account/operation/AccountUsersSearch
	AccountUsersSearch(ctx context.Context, accountId string, in *AccountUsersSearchIn) ([]User, error)
}

func NewHandler(doer doer) AccountHandler {
	return AccountHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type AccountHandler struct {
	doer doer
}

func (h *AccountHandler) AccountAttachPaymentMethod(ctx context.Context, accountId string, in *AccountAttachPaymentMethodIn) (*Card, error) {
	path := fmt.Sprintf("/account/%s/payment_methods", accountId)
	b, err := h.doer.Do(ctx, "AccountAttachPaymentMethod", "POST", path, in)
	out := new(accountAttachPaymentMethodOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Card, nil
}
func (h *AccountHandler) AccountBillingGroupList(ctx context.Context, accountId string) ([]AccountBillingGroup, error) {
	path := fmt.Sprintf("/account/%s/billing-group", accountId)
	b, err := h.doer.Do(ctx, "AccountBillingGroupList", "GET", path, nil)
	out := new(accountBillingGroupListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AccountBillingGroups, nil
}
func (h *AccountHandler) AccountCreate(ctx context.Context, in *AccountCreateIn) (*Account, error) {
	path := fmt.Sprintf("/account")
	b, err := h.doer.Do(ctx, "AccountCreate", "POST", path, in)
	out := new(accountCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Account, nil
}
func (h *AccountHandler) AccountDelete(ctx context.Context, accountId string) error {
	path := fmt.Sprintf("/account/%s", accountId)
	_, err := h.doer.Do(ctx, "AccountDelete", "DELETE", path, nil)
	return err
}
func (h *AccountHandler) AccountEventList(ctx context.Context, accountId string) ([]Event, error) {
	path := fmt.Sprintf("/account/%s/events", accountId)
	b, err := h.doer.Do(ctx, "AccountEventList", "GET", path, nil)
	out := new(accountEventListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Events, nil
}
func (h *AccountHandler) AccountGet(ctx context.Context, accountId string) (*Account, error) {
	path := fmt.Sprintf("/account/%s", accountId)
	b, err := h.doer.Do(ctx, "AccountGet", "GET", path, nil)
	out := new(accountGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Account, nil
}
func (h *AccountHandler) AccountList(ctx context.Context) ([]Account, error) {
	path := fmt.Sprintf("/account")
	b, err := h.doer.Do(ctx, "AccountList", "GET", path, nil)
	out := new(accountListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Accounts, nil
}
func (h *AccountHandler) AccountMove(ctx context.Context, accountId string, in *AccountMoveIn) (*Account, error) {
	path := fmt.Sprintf("/account/%s/parent_account", accountId)
	b, err := h.doer.Do(ctx, "AccountMove", "PUT", path, in)
	out := new(accountMoveOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Account, nil
}
func (h *AccountHandler) AccountPaymentMethodDelete(ctx context.Context, accountId string, cardId string) error {
	path := fmt.Sprintf("/account/%s/payment_method/%s", accountId, cardId)
	_, err := h.doer.Do(ctx, "AccountPaymentMethodDelete", "DELETE", path, nil)
	return err
}
func (h *AccountHandler) AccountPaymentMethodsList(ctx context.Context, accountId string) ([]CardItem, error) {
	path := fmt.Sprintf("/account/%s/payment_methods", accountId)
	b, err := h.doer.Do(ctx, "AccountPaymentMethodsList", "GET", path, nil)
	out := new(accountPaymentMethodsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Cards, nil
}
func (h *AccountHandler) AccountProjectsList(ctx context.Context, accountId string) (*AccountProjectsListOut, error) {
	path := fmt.Sprintf("/account/%s/projects", accountId)
	b, err := h.doer.Do(ctx, "AccountProjectsList", "GET", path, nil)
	out := new(AccountProjectsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *AccountHandler) AccountUpdate(ctx context.Context, accountId string, in *AccountUpdateIn) (*Account, error) {
	path := fmt.Sprintf("/account/%s", accountId)
	b, err := h.doer.Do(ctx, "AccountUpdate", "PUT", path, in)
	out := new(accountUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Account, nil
}
func (h *AccountHandler) AccountUserProjectsList(ctx context.Context, accountId string, userId string) ([]UserProject, error) {
	path := fmt.Sprintf("/account/%s/user/%s/projects", accountId, userId)
	b, err := h.doer.Do(ctx, "AccountUserProjectsList", "GET", path, nil)
	out := new(accountUserProjectsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.UserProjects, nil
}
func (h *AccountHandler) AccountUsersSearch(ctx context.Context, accountId string, in *AccountUsersSearchIn) ([]User, error) {
	path := fmt.Sprintf("/account/%s/users/search", accountId)
	b, err := h.doer.Do(ctx, "AccountUsersSearch", "POST", path, in)
	out := new(accountUsersSearchOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Users, nil
}

type AccessSourceType string

const (
	AccessSourceTypeDescendantMembership   AccessSourceType = "descendant_membership"
	AccessSourceTypeOrganizationMembership AccessSourceType = "organization_membership"
	AccessSourceTypeProjectMembership      AccessSourceType = "project_membership"
	AccessSourceTypeTeamMembership         AccessSourceType = "team_membership"
)

type Account struct {
	AccessSource          AccessSourceType `json:"access_source,omitempty"`
	AccountId             string           `json:"account_id"`
	AccountName           string           `json:"account_name"`
	AccountOwnerTeamId    string           `json:"account_owner_team_id"`
	CreateTime            time.Time        `json:"create_time"`
	Features              map[string]any   `json:"features,omitempty"`
	IsAccountMember       *bool            `json:"is_account_member,omitempty"`
	IsAccountOwner        bool             `json:"is_account_owner"`
	OrganizationId        string           `json:"organization_id"`
	ParentAccountId       string           `json:"parent_account_id,omitempty"`
	PrimaryBillingGroupId string           `json:"primary_billing_group_id"`
	RootAccountId         string           `json:"root_account_id"`
	TenantId              string           `json:"tenant_id,omitempty"`
	UpdateTime            time.Time        `json:"update_time"`
}
type AccountAttachPaymentMethodIn struct {
	PaymentMethodId string `json:"payment_method_id"`
}
type accountAttachPaymentMethodOut struct {
	Card *Card `json:"card"`
}
type AccountBillingGroup struct {
	AccountId             string              `json:"account_id"`
	AccountName           string              `json:"account_name"`
	AddressLines          []string            `json:"address_lines"`
	BillingAddress        string              `json:"billing_address,omitempty"`
	BillingCurrency       BillingCurrencyType `json:"billing_currency"`
	BillingEmails         []BillingEmail      `json:"billing_emails"`
	BillingExtraText      string              `json:"billing_extra_text"`
	BillingGroupId        string              `json:"billing_group_id"`
	BillingGroupName      string              `json:"billing_group_name"`
	CardInfo              *CardInfo           `json:"card_info"`
	City                  string              `json:"city"`
	Company               string              `json:"company"`
	Country               string              `json:"country"`
	CountryCode           string              `json:"country_code"`
	EstimatedBalanceLocal string              `json:"estimated_balance_local"`
	EstimatedBalanceUsd   string              `json:"estimated_balance_usd"`
	PaymentMethod         PaymentMethodType   `json:"payment_method"`
	State                 string              `json:"state"`
	VatId                 string              `json:"vat_id"`
	ZipCode               string              `json:"zip_code"`
}
type accountBillingGroupListOut struct {
	AccountBillingGroups []AccountBillingGroup `json:"account_billing_groups"`
}
type AccountCreateIn struct {
	AccountName           string `json:"account_name"`
	ParentAccountId       string `json:"parent_account_id,omitempty"`
	PrimaryBillingGroupId string `json:"primary_billing_group_id,omitempty"`
}
type accountCreateOut struct {
	Account *Account `json:"account"`
}
type accountEventListOut struct {
	Events []Event `json:"events"`
}
type accountGetOut struct {
	Account *Account `json:"account"`
}
type accountListOut struct {
	Accounts []Account `json:"accounts"`
}
type AccountMoveIn struct {
	ParentAccountId string `json:"parent_account_id"`
}
type accountMoveOut struct {
	Account *Account `json:"account"`
}
type accountPaymentMethodsListOut struct {
	Cards []CardItem `json:"cards"`
}
type AccountProjectsListOut struct {
	Projects          []Project `json:"projects"`
	TotalProjectCount *int      `json:"total_project_count,omitempty"`
}
type AccountUpdateIn struct {
	AccountName           string `json:"account_name,omitempty"`
	PrimaryBillingGroupId string `json:"primary_billing_group_id,omitempty"`
}
type accountUpdateOut struct {
	Account *Account `json:"account"`
}
type accountUserProjectsListOut struct {
	UserProjects []UserProject `json:"user_projects"`
}
type AccountUsersSearchIn struct {
	Limit   *int        `json:"limit,omitempty"`
	OrderBy OrderByType `json:"order_by,omitempty"`
	Query   string      `json:"query,omitempty"`
}
type accountUsersSearchOut struct {
	Users []User `json:"users"`
}
type BillingCurrencyType string

const (
	BillingCurrencyTypeAud BillingCurrencyType = "AUD"
	BillingCurrencyTypeCad BillingCurrencyType = "CAD"
	BillingCurrencyTypeChf BillingCurrencyType = "CHF"
	BillingCurrencyTypeDkk BillingCurrencyType = "DKK"
	BillingCurrencyTypeEur BillingCurrencyType = "EUR"
	BillingCurrencyTypeGbp BillingCurrencyType = "GBP"
	BillingCurrencyTypeJpy BillingCurrencyType = "JPY"
	BillingCurrencyTypeNok BillingCurrencyType = "NOK"
	BillingCurrencyTypeNzd BillingCurrencyType = "NZD"
	BillingCurrencyTypeSek BillingCurrencyType = "SEK"
	BillingCurrencyTypeSgd BillingCurrencyType = "SGD"
	BillingCurrencyTypeUsd BillingCurrencyType = "USD"
)

type BillingEmail struct {
	Email string `json:"email"`
}
type Card struct {
	Brand          string   `json:"brand"`
	CardId         string   `json:"card_id"`
	Country        string   `json:"country"`
	CountryCode    string   `json:"country_code"`
	ExpMonth       int      `json:"exp_month"`
	ExpYear        int      `json:"exp_year"`
	Last4          string   `json:"last4"`
	Name           string   `json:"name"`
	OrganizationId string   `json:"organization_id,omitempty"`
	Projects       []string `json:"projects"`
}
type CardInfo struct {
	Brand       string `json:"brand"`
	CardId      string `json:"card_id"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	ExpMonth    int    `json:"exp_month"`
	ExpYear     int    `json:"exp_year"`
	Last4       string `json:"last4"`
	Name        string `json:"name"`
	UserEmail   string `json:"user_email"`
}
type CardItem struct {
	Brand       string `json:"brand"`
	CardId      string `json:"card_id"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	ExpMonth    int    `json:"exp_month"`
	ExpYear     int    `json:"exp_year"`
	Last4       string `json:"last4"`
	Name        string `json:"name"`
}
type Elasticsearch struct {
	EolDate string `json:"eol_date"`
	Version string `json:"version"`
}
type EndOfLifeExtension struct {
	Elasticsearch *Elasticsearch `json:"elasticsearch,omitempty"`
}
type Event struct {
	AccountId         string    `json:"account_id"`
	ActionDescription string    `json:"action_description"`
	ActionType        string    `json:"action_type"`
	Actor             string    `json:"actor"`
	ActorUserId       string    `json:"actor_user_id"`
	CreateTime        time.Time `json:"create_time"`
	LogEntryId        int       `json:"log_entry_id"`
	TeamId            string    `json:"team_id"`
}
type MemberType string

const (
	MemberTypeAdmin     MemberType = "admin"
	MemberTypeDeveloper MemberType = "developer"
	MemberTypeOperator  MemberType = "operator"
	MemberTypeReadOnly  MemberType = "read_only"
)

type OrderByType string

const (
	OrderByTypeUserEmailasc  OrderByType = "user_email:asc"
	OrderByTypeUserEmaildesc OrderByType = "user_email:desc"
	OrderByTypeUserIdasc     OrderByType = "user_id:asc"
	OrderByTypeUserIddesc    OrderByType = "user_id:desc"
	OrderByTypeRealNameasc   OrderByType = "real_name:asc"
	OrderByTypeRealNamedesc  OrderByType = "real_name:desc"
)

func OrderByTypeChoices() []string {
	return []string{"user_email:asc", "user_email:desc", "user_id:asc", "user_id:desc", "real_name:asc", "real_name:desc"}
}

type PaymentMethodType string

const (
	PaymentMethodTypeAccrual           PaymentMethodType = "accrual"
	PaymentMethodTypeCard              PaymentMethodType = "card"
	PaymentMethodTypeDisabled          PaymentMethodType = "disabled"
	PaymentMethodTypeEmail             PaymentMethodType = "email"
	PaymentMethodTypeNoPaymentExpected PaymentMethodType = "no_payment_expected"
	PaymentMethodTypePartner           PaymentMethodType = "partner"
)

type Project struct {
	AccountId             string              `json:"account_id"`
	AccountName           string              `json:"account_name,omitempty"`
	AddressLines          []string            `json:"address_lines"`
	AvailableCredits      string              `json:"available_credits,omitempty"`
	BillingAddress        string              `json:"billing_address"`
	BillingCurrency       BillingCurrencyType `json:"billing_currency"`
	BillingEmails         []BillingEmail      `json:"billing_emails"`
	BillingExtraText      string              `json:"billing_extra_text,omitempty"`
	BillingGroupId        string              `json:"billing_group_id"`
	BillingGroupName      string              `json:"billing_group_name"`
	CardInfo              *CardInfo           `json:"card_info"`
	City                  string              `json:"city,omitempty"`
	Company               string              `json:"company,omitempty"`
	Country               string              `json:"country"`
	CountryCode           string              `json:"country_code"`
	DefaultCloud          string              `json:"default_cloud"`
	EndOfLifeExtension    *EndOfLifeExtension `json:"end_of_life_extension,omitempty"`
	EstimatedBalance      string              `json:"estimated_balance"`
	EstimatedBalanceLocal string              `json:"estimated_balance_local,omitempty"`
	Features              map[string]any      `json:"features,omitempty"`
	OrganizationId        string              `json:"organization_id"`
	PaymentMethod         string              `json:"payment_method"`
	ProjectName           string              `json:"project_name"`
	State                 string              `json:"state,omitempty"`
	Tags                  map[string]string   `json:"tags,omitempty"`
	TechEmails            []BillingEmail      `json:"tech_emails"`
	TenantId              string              `json:"tenant_id,omitempty"`
	TrialExpirationTime   *time.Time          `json:"trial_expiration_time,omitempty"`
	VatId                 string              `json:"vat_id"`
	ZipCode               string              `json:"zip_code,omitempty"`
}
type User struct {
	RealName  string `json:"real_name"`
	UserEmail string `json:"user_email"`
	UserId    string `json:"user_id"`
}
type UserProject struct {
	AccessType  string     `json:"access_type,omitempty"`
	AccountId   string     `json:"account_id"`
	CreateTime  time.Time  `json:"create_time"`
	MemberType  MemberType `json:"member_type"`
	ProjectName string     `json:"project_name"`
	RealName    string     `json:"real_name"`
	TeamId      string     `json:"team_id"`
	TeamName    string     `json:"team_name"`
	UserEmail   string     `json:"user_email"`
}
