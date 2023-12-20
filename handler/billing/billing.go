// Code generated by Aiven. DO NOT EDIT.

package billing

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// InvoiceGet get a single invoice
	// InvoiceGet GET /invoices/{invoice_number}
	// https://api.aiven.io/doc/#tag/Billing/operation/InvoiceGet
	InvoiceGet(ctx context.Context, invoiceNumber string) (*Invoice, error)
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

func (h *handler) InvoiceGet(ctx context.Context, invoiceNumber string) (*Invoice, error) {
	path := fmt.Sprintf("/invoices/%s", invoiceNumber)
	b, err := h.doer.Do(ctx, "InvoiceGet", "GET", path, nil)
	out := new(invoiceGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Invoice, nil
}

type Invoice struct {
	BillingGroupId    string    `json:"billing_group_id"`
	BillingGroupName  string    `json:"billing_group_name"`
	BillingGroupState string    `json:"billing_group_state"`
	Currency          string    `json:"currency"`
	DownloadCookie    string    `json:"download_cookie,omitempty"`
	GeneratedAt       time.Time `json:"generated_at"`
	InvoiceNumber     string    `json:"invoice_number"`
	PeriodBegin       string    `json:"period_begin"`
	PeriodEnd         string    `json:"period_end"`
	State             string    `json:"state"`
	TotalIncVat       string    `json:"total_inc_vat"`
	TotalVatZero      string    `json:"total_vat_zero"`
}
type invoiceGetOut struct {
	Invoice *Invoice `json:"invoice"`
}