// Code generated by Aiven. DO NOT EDIT.

package accountauthentication

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// AccountAuthenticationMethodCreate create a new authentication method
	// POST /account/{account_id}/authentication
	// https://api.aiven.io/doc/#tag/Account/operation/AccountAuthenticationMethodCreate
	AccountAuthenticationMethodCreate(ctx context.Context, accountId string, in *AccountAuthenticationMethodCreateIn) (*AuthenticationMethod, error)

	// AccountAuthenticationMethodDelete delete authentication method
	// DELETE /account/{account_id}/authentication/{account_authentication_method_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountAuthenticationMethodDelete
	AccountAuthenticationMethodDelete(ctx context.Context, accountId string, accountAuthenticationMethodId string) error

	// AccountAuthenticationMethodGet get details of a single authentication method
	// GET /account/{account_id}/authentication/{account_authentication_method_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountAuthenticationMethodGet
	AccountAuthenticationMethodGet(ctx context.Context, accountId string, accountAuthenticationMethodId string) (*AuthenticationMethod, error)

	// AccountAuthenticationMethodUpdate update authentication method
	// PUT /account/{account_id}/authentication/{account_authentication_method_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountAuthenticationMethodUpdate
	AccountAuthenticationMethodUpdate(ctx context.Context, accountId string, accountAuthenticationMethodId string, in *AccountAuthenticationMethodUpdateIn) (*AuthenticationMethod, error)

	// AccountAuthenticationMethodsList list authentication methods
	// GET /account/{account_id}/authentication
	// https://api.aiven.io/doc/#tag/Account/operation/AccountAuthenticationMethodsList
	AccountAuthenticationMethodsList(ctx context.Context, accountId string) ([]AuthenticationMethod, error)
}

func NewHandler(doer doer) AccountAuthenticationHandler {
	return AccountAuthenticationHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type AccountAuthenticationHandler struct {
	doer doer
}

func (h *AccountAuthenticationHandler) AccountAuthenticationMethodCreate(ctx context.Context, accountId string, in *AccountAuthenticationMethodCreateIn) (*AuthenticationMethod, error) {
	path := fmt.Sprintf("/account/%s/authentication", accountId)
	b, err := h.doer.Do(ctx, "AccountAuthenticationMethodCreate", "POST", path, in)
	out := new(accountAuthenticationMethodCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AuthenticationMethod, nil
}
func (h *AccountAuthenticationHandler) AccountAuthenticationMethodDelete(ctx context.Context, accountId string, accountAuthenticationMethodId string) error {
	path := fmt.Sprintf("/account/%s/authentication/%s", accountId, accountAuthenticationMethodId)
	_, err := h.doer.Do(ctx, "AccountAuthenticationMethodDelete", "DELETE", path, nil)
	return err
}
func (h *AccountAuthenticationHandler) AccountAuthenticationMethodGet(ctx context.Context, accountId string, accountAuthenticationMethodId string) (*AuthenticationMethod, error) {
	path := fmt.Sprintf("/account/%s/authentication/%s", accountId, accountAuthenticationMethodId)
	b, err := h.doer.Do(ctx, "AccountAuthenticationMethodGet", "GET", path, nil)
	out := new(accountAuthenticationMethodGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AuthenticationMethod, nil
}
func (h *AccountAuthenticationHandler) AccountAuthenticationMethodUpdate(ctx context.Context, accountId string, accountAuthenticationMethodId string, in *AccountAuthenticationMethodUpdateIn) (*AuthenticationMethod, error) {
	path := fmt.Sprintf("/account/%s/authentication/%s", accountId, accountAuthenticationMethodId)
	b, err := h.doer.Do(ctx, "AccountAuthenticationMethodUpdate", "PUT", path, in)
	out := new(accountAuthenticationMethodUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AuthenticationMethod, nil
}
func (h *AccountAuthenticationHandler) AccountAuthenticationMethodsList(ctx context.Context, accountId string) ([]AuthenticationMethod, error) {
	path := fmt.Sprintf("/account/%s/authentication", accountId)
	b, err := h.doer.Do(ctx, "AccountAuthenticationMethodsList", "GET", path, nil)
	out := new(accountAuthenticationMethodsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AuthenticationMethods, nil
}

type AccountAuthenticationMethodCreateIn struct {
	AuthenticationMethodName         string                     `json:"authentication_method_name"`
	AuthenticationMethodType         AuthenticationMethodType   `json:"authentication_method_type"`
	AutoJoinTeamId                   string                     `json:"auto_join_team_id,omitempty"`
	AutoJoinUserGroupId              string                     `json:"auto_join_user_group_id,omitempty"`
	LinkedDomains                    []LinkedDomain             `json:"linked_domains"`
	SamlAssertionSignedEnabled       *bool                      `json:"saml_assertion_signed_enabled,omitempty"`
	SamlAuthnRequestsSignedEnabled   *bool                      `json:"saml_authn_requests_signed_enabled,omitempty"`
	SamlCertificate                  string                     `json:"saml_certificate,omitempty"`
	SamlDigestAlgorithm              SamlDigestAlgorithmType    `json:"saml_digest_algorithm,omitempty"`
	SamlEntityId                     string                     `json:"saml_entity_id,omitempty"`
	SamlFieldMapping                 *SamlFieldMapping          `json:"saml_field_mapping,omitempty"`
	SamlIdpLoginAllowed              *bool                      `json:"saml_idp_login_allowed,omitempty"`
	SamlIdpUrl                       string                     `json:"saml_idp_url,omitempty"`
	SamlRequestedAuthnContextEnabled *bool                      `json:"saml_requested_authn_context_enabled,omitempty"`
	SamlSignatureAlgorithm           SamlSignatureAlgorithmType `json:"saml_signature_algorithm,omitempty"`
	SamlVariant                      SamlVariantType            `json:"saml_variant,omitempty"`
}
type accountAuthenticationMethodCreateOut struct {
	AuthenticationMethod *AuthenticationMethod `json:"authentication_method"`
}
type accountAuthenticationMethodGetOut struct {
	AuthenticationMethod *AuthenticationMethod `json:"authentication_method"`
}
type AccountAuthenticationMethodUpdateIn struct {
	AuthenticationMethodEnabled      *bool                      `json:"authentication_method_enabled,omitempty"`
	AuthenticationMethodName         string                     `json:"authentication_method_name,omitempty"`
	AutoJoinTeamId                   string                     `json:"auto_join_team_id,omitempty"`
	AutoJoinUserGroupId              string                     `json:"auto_join_user_group_id,omitempty"`
	SamlAssertionSignedEnabled       *bool                      `json:"saml_assertion_signed_enabled,omitempty"`
	SamlAuthnRequestsSignedEnabled   *bool                      `json:"saml_authn_requests_signed_enabled,omitempty"`
	SamlCertificate                  string                     `json:"saml_certificate,omitempty"`
	SamlDigestAlgorithm              SamlDigestAlgorithmType    `json:"saml_digest_algorithm,omitempty"`
	SamlEntityId                     string                     `json:"saml_entity_id,omitempty"`
	SamlFieldMapping                 *SamlFieldMapping          `json:"saml_field_mapping,omitempty"`
	SamlIdpLoginAllowed              *bool                      `json:"saml_idp_login_allowed,omitempty"`
	SamlIdpUrl                       string                     `json:"saml_idp_url,omitempty"`
	SamlRequestedAuthnContextEnabled *bool                      `json:"saml_requested_authn_context_enabled,omitempty"`
	SamlSignatureAlgorithm           SamlSignatureAlgorithmType `json:"saml_signature_algorithm,omitempty"`
	SamlVariant                      SamlVariantType            `json:"saml_variant,omitempty"`
}
type accountAuthenticationMethodUpdateOut struct {
	AuthenticationMethod *AuthenticationMethod `json:"authentication_method"`
}
type accountAuthenticationMethodsListOut struct {
	AuthenticationMethods []AuthenticationMethod `json:"authentication_methods"`
}
type AuthenticationMethod struct {
	AccountId                        string                     `json:"account_id"`
	AuthenticationMethodEnabled      bool                       `json:"authentication_method_enabled"`
	AuthenticationMethodId           string                     `json:"authentication_method_id"`
	AuthenticationMethodName         string                     `json:"authentication_method_name,omitempty"`
	AuthenticationMethodType         AuthenticationMethodType   `json:"authentication_method_type"`
	AutoJoinTeamId                   string                     `json:"auto_join_team_id"`
	AutoJoinUserGroupId              string                     `json:"auto_join_user_group_id"`
	CreateTime                       time.Time                  `json:"create_time"`
	DeleteTime                       time.Time                  `json:"delete_time"`
	OrganizationId                   string                     `json:"organization_id,omitempty"`
	SamlAcsUrl                       string                     `json:"saml_acs_url,omitempty"`
	SamlAssertionSignedEnabled       *bool                      `json:"saml_assertion_signed_enabled,omitempty"`
	SamlAuthnRequestsSignedEnabled   *bool                      `json:"saml_authn_requests_signed_enabled,omitempty"`
	SamlCert                         SamlCertType               `json:"saml_cert,omitempty"`
	SamlCertificate                  string                     `json:"saml_certificate,omitempty"`
	SamlCertificateIssuer            string                     `json:"saml_certificate_issuer,omitempty"`
	SamlCertificateNotValidAfter     string                     `json:"saml_certificate_not_valid_after,omitempty"`
	SamlCertificateNotValidBefore    string                     `json:"saml_certificate_not_valid_before,omitempty"`
	SamlCertificateSubject           string                     `json:"saml_certificate_subject,omitempty"`
	SamlDigestAlgorithm              SamlDigestAlgorithmType    `json:"saml_digest_algorithm,omitempty"`
	SamlEntityId                     string                     `json:"saml_entity_id,omitempty"`
	SamlFieldMapping                 *SamlFieldMapping          `json:"saml_field_mapping,omitempty"`
	SamlIdpLoginAllowed              *bool                      `json:"saml_idp_login_allowed,omitempty"`
	SamlIdpUrl                       string                     `json:"saml_idp_url,omitempty"`
	SamlMetadataUrl                  string                     `json:"saml_metadata_url,omitempty"`
	SamlRequestedAuthnContextEnabled *bool                      `json:"saml_requested_authn_context_enabled,omitempty"`
	SamlSignatureAlgorithm           SamlSignatureAlgorithmType `json:"saml_signature_algorithm,omitempty"`
	SamlSpCertificate                string                     `json:"saml_sp_certificate,omitempty"`
	SamlVariant                      SamlVariantType            `json:"saml_variant,omitempty"`
	State                            StateType                  `json:"state"`
	UpdateTime                       time.Time                  `json:"update_time"`
}
type AuthenticationMethodType string

const (
	AuthenticationMethodTypeInternal AuthenticationMethodType = "internal"
	AuthenticationMethodTypeSaml     AuthenticationMethodType = "saml"
)

type LinkedDomain struct {
	DomainId string `json:"domain_id"`
}
type SamlCertType string

const (
	SamlCertTypeAdfs SamlCertType = "adfs"
)

type SamlDigestAlgorithmType string

const (
	SamlDigestAlgorithmTypeSha1   SamlDigestAlgorithmType = "sha1"
	SamlDigestAlgorithmTypeSha256 SamlDigestAlgorithmType = "sha256"
	SamlDigestAlgorithmTypeSha384 SamlDigestAlgorithmType = "sha384"
	SamlDigestAlgorithmTypeSha512 SamlDigestAlgorithmType = "sha512"
)

type SamlFieldMapping struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Identity  string `json:"identity,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	RealName  string `json:"real_name,omitempty"`
}
type SamlSignatureAlgorithmType string

const (
	SamlSignatureAlgorithmTypeRsaSha1   SamlSignatureAlgorithmType = "rsa-sha1"
	SamlSignatureAlgorithmTypeDsaSha1   SamlSignatureAlgorithmType = "dsa-sha1"
	SamlSignatureAlgorithmTypeRsaSha256 SamlSignatureAlgorithmType = "rsa-sha256"
	SamlSignatureAlgorithmTypeRsaSha384 SamlSignatureAlgorithmType = "rsa-sha384"
	SamlSignatureAlgorithmTypeRsaSha512 SamlSignatureAlgorithmType = "rsa-sha512"
)

type SamlVariantType string

const (
	SamlVariantTypeAdfs SamlVariantType = "adfs"
)

type StateType string

const (
	StateTypeActive               StateType = "active"
	StateTypeDeleted              StateType = "deleted"
	StateTypePendingConfiguration StateType = "pending_configuration"
)
