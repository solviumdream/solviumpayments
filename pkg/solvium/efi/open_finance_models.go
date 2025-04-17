package efi

type WebhookSecurityType string

const (
	WebhookSecurityTypeMTLS WebhookSecurityType = "mtls"
	WebhookSecurityTypeHMAC WebhookSecurityType = "hmac"
)

type ProcessPaymentType string

const (
	ProcessPaymentTypeSync  ProcessPaymentType = "sync"
	ProcessPaymentTypeAsync ProcessPaymentType = "async"
)

type OpenFinanceConfig struct {
	RedirectURL         string                     `json:"redirectURL"`
	WebhookURL          string                     `json:"webhookURL"`
	WebhookSecurity     OpenFinanceWebhookSecurity `json:"webhookSecurity"`
	ProcessPayment      ProcessPaymentType         `json:"processPayment"`
	GenerateTxIdForInic bool                       `json:"generateTxIdForInic"`
}

type OpenFinanceWebhookSecurity struct {
	Type WebhookSecurityType `json:"type"`
	Key  string              `json:"key,omitempty"` 
}

type OpenFinanceRedirectParams struct {
	PaymentIdentifier string `json:"identificadorPagamento"`
	Error             string `json:"erro,omitempty"`
}
