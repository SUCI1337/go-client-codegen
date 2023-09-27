// Code generated by Aiven. DO NOT EDIT.

package projectbilling

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// ProjectCreditsClaim claim a credit code
	// ProjectCreditsClaim POST /project/{project}/credits
	// https://api.aiven.io/doc/#tag/Project_Billing/operation/ProjectCreditsClaim
	ProjectCreditsClaim(ctx context.Context, project string, in *ProjectCreditsClaimIn) (*Credit, error)

	// ProjectCreditsList list project credits
	// ProjectCreditsList GET /project/{project}/credits
	// https://api.aiven.io/doc/#tag/Project_Billing/operation/ProjectCreditsList
	ProjectCreditsList(ctx context.Context, project string) ([]Credit, error)

	// ProjectInvoiceList list project invoices
	// ProjectInvoiceList GET /project/{project}/invoice
	// https://api.aiven.io/doc/#tag/Project_Billing/operation/ProjectInvoiceList
	ProjectInvoiceList(ctx context.Context, project string) ([]Invoice, error)
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

func (h *handler) ProjectCreditsClaim(ctx context.Context, project string, in *ProjectCreditsClaimIn) (*Credit, error) {
	path := fmt.Sprintf("/project/%s/credits", project)
	b, err := h.doer.Do(ctx, "ProjectCreditsClaim", "POST", path, in)
	out := new(ProjectCreditsClaimOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Credit, nil
}
func (h *handler) ProjectCreditsList(ctx context.Context, project string) ([]Credit, error) {
	path := fmt.Sprintf("/project/%s/credits", project)
	b, err := h.doer.Do(ctx, "ProjectCreditsList", "GET", path, nil)
	out := new(ProjectCreditsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Credits, nil
}
func (h *handler) ProjectInvoiceList(ctx context.Context, project string) ([]Invoice, error) {
	path := fmt.Sprintf("/project/%s/invoice", project)
	b, err := h.doer.Do(ctx, "ProjectInvoiceList", "GET", path, nil)
	out := new(ProjectInvoiceListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Invoices, nil
}

type BillingGroupStateType string

const (
	BillingGroupStateTypeActive  BillingGroupStateType = "active"
	BillingGroupStateTypeDeleted BillingGroupStateType = "deleted"
)

type Credit struct {
	Code           string     `json:"code,omitempty"`
	ExpireTime     *time.Time `json:"expire_time,omitempty"`
	RemainingValue string     `json:"remaining_value,omitempty"`
	StartTime      *time.Time `json:"start_time,omitempty"`
	Type           Type       `json:"type,omitempty"`
	Value          string     `json:"value,omitempty"`
}
type CurrencyType string

const (
	CurrencyTypeAud CurrencyType = "AUD"
	CurrencyTypeCad CurrencyType = "CAD"
	CurrencyTypeChf CurrencyType = "CHF"
	CurrencyTypeDkk CurrencyType = "DKK"
	CurrencyTypeEur CurrencyType = "EUR"
	CurrencyTypeGbp CurrencyType = "GBP"
	CurrencyTypeJpy CurrencyType = "JPY"
	CurrencyTypeNok CurrencyType = "NOK"
	CurrencyTypeNzd CurrencyType = "NZD"
	CurrencyTypeSek CurrencyType = "SEK"
	CurrencyTypeSgd CurrencyType = "SGD"
	CurrencyTypeUsd CurrencyType = "USD"
)

type Invoice struct {
	BillingGroupId    string                `json:"billing_group_id"`
	BillingGroupName  string                `json:"billing_group_name"`
	BillingGroupState BillingGroupStateType `json:"billing_group_state"`
	Currency          CurrencyType          `json:"currency"`
	DownloadCookie    string                `json:"download_cookie"`
	GeneratedAt       *time.Time            `json:"generated_at,omitempty"`
	InvoiceNumber     string                `json:"invoice_number"`
	PeriodBegin       string                `json:"period_begin"`
	PeriodEnd         string                `json:"period_end"`
	State             StateType             `json:"state"`
	TotalIncVat       string                `json:"total_inc_vat"`
	TotalVatZero      string                `json:"total_vat_zero"`
}
type ProjectCreditsClaimIn struct {
	Code string `json:"code"`
}
type ProjectCreditsClaimOut struct {
	Credit *Credit `json:"credit"`
}
type ProjectCreditsListOut struct {
	Credits []Credit `json:"credits"`
}
type ProjectInvoiceListOut struct {
	Invoices []Invoice `json:"invoices"`
}
type StateType string

const (
	StateTypeAccrual                StateType = "accrual"
	StateTypeConsolidated           StateType = "consolidated"
	StateTypeDue                    StateType = "due"
	StateTypeEstimate               StateType = "estimate"
	StateTypeFailedCreditCardCharge StateType = "failed_credit_card_charge"
	StateTypeFailedNoCreditCard     StateType = "failed_no_credit_card"
	StateTypeMailed                 StateType = "mailed"
	StateTypeNoPaymentExpected      StateType = "no_payment_expected"
	StateTypePaid                   StateType = "paid"
	StateTypePartnerMetering        StateType = "partner_metering"
	StateTypeUncollectible          StateType = "uncollectible"
	StateTypeWaived                 StateType = "waived"
)

type Type string

const (
	TypeDiscount    Type = "discount"
	TypeEmployee    Type = "employee"
	TypeEvaluation  Type = "evaluation"
	TypeInternal    Type = "internal"
	TypeOther       Type = "other"
	TypeOutage      Type = "outage"
	TypePartner     Type = "partner"
	TypePromotion   Type = "promotion"
	TypePurchase    Type = "purchase"
	TypeReferral    Type = "referral"
	TypeSponsorship Type = "sponsorship"
	TypeTrial       Type = "trial"
	TypeTrialOver   Type = "trial_over"
)
