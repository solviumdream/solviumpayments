package efi

import "time"

type BillPaymentWebhookRequest struct {
	URL string `json:"url"`
}

type BillPaymentWebhookResponse struct {
	URL string `json:"url"`
}

type BillPaymentWebhookListRequest struct {
	StartDate time.Time
	EndDate   time.Time
}

type BillPaymentWebhookListResponse struct {
	Parameters BillPaymentWebhookListParameters `json:"parametros"`
	Webhooks   []BillPaymentWebhook             `json:"webhooks"`
}

type BillPaymentWebhookListParameters struct {
	Start      string                       `json:"inicio"`
	End        string                       `json:"fim"`
	Pagination BillPaymentWebhookPagination `json:"paginacao"`
}

type BillPaymentWebhookPagination struct {
	CurrentPage  int `json:"paginaAtual"`
	ItemsPerPage int `json:"itensPorPagina"`
	TotalPages   int `json:"quantidadeDePaginas"`
	TotalItems   int `json:"quantidadeTotalDeItens"`
}

type BillPaymentWebhook struct {
	URL       string `json:"url"`
	CreatedAt string `json:"criacao"`
}

type BillPaymentWebhookCallback struct {
	Identifier string                    `json:"identificador"`
	Status     BillPaymentStatusChange   `json:"status"`
	Value      string                    `json:"valor"`
	Timestamp  BillPaymentCallbackTime   `json:"horario"`
	EfiExtras  BillPaymentCallbackExtras `json:"efiExtras"`
}

type BillPaymentStatusChange struct {
	Previous string `json:"anterior"`
	Current  string `json:"atual"`
}

type BillPaymentCallbackTime struct {
	RequestTime string `json:"solicitacao"`
}

type BillPaymentCallbackExtras struct {
	ExecutionDate string `json:"dataExecucao"`
	Barcode       string `json:"codigoBarras"`
	LineCode      string `json:"linhaDigitavel"`
}
