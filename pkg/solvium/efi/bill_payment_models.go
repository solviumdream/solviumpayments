package efi

type BillPaymentStatus string

const (
	BillPaymentStatusSettled    BillPaymentStatus = "LIQUIDADO"
	BillPaymentStatusUnsettled  BillPaymentStatus = "NAO_LIQUIDADO"
	BillPaymentStatusFailed     BillPaymentStatus = "FALHA"
	BillPaymentStatusCanceled   BillPaymentStatus = "CANCELADO"
	BillPaymentStatusProcessing BillPaymentStatus = "PROCESSANDO"
)

type BillPaymentRequest struct {
	Value       float64 `json:"valor"`
	PaymentDate string  `json:"dataPagamento"`
	Description string  `json:"descricao,omitempty"`
}

type BillPaymentResponse struct {
	PaymentID    string            `json:"idPagamento"`
	AmountPaid   float64           `json:"valorPago"`
	Status       BillPaymentStatus `json:"status"`
	RejectReason string            `json:"motivoRecusa,omitempty"`
	Data         BillPaymentData   `json:"data"`
}

type BillPaymentData struct {
	RequestDate string `json:"solicitacao"`
	PaymentDate string `json:"pagamento,omitempty"`
}

type BillPaymentSummaryRequest struct {
	StartDate string `json:"dataInicial"`
	EndDate   string `json:"dataFinal"`
}

type BillPaymentSummary struct {
	Dates     BillPaymentSummaryDates    `json:"datas"`
	Requests  BillPaymentSummaryRequests `json:"solicitacoes"`
	FailedIds []string                   `json:"solicitacoesFalhas"`
}

type BillPaymentSummaryDates struct {
	StartDate string `json:"inicial"`
	EndDate   string `json:"final"`
}

type BillPaymentSummaryRequests struct {
	Total      int `json:"total"`
	Processing int `json:"processando"`
	Success    int `json:"sucesso"`
	Failed     int `json:"falha"`
	Canceled   int `json:"cancelada"`
}

type BillDetails struct {
	Barcode       string  `json:"codigoDeBarras"`
	Type          string  `json:"tipo"`
	Value         float64 `json:"valor"`
	DueDate       string  `json:"dataVencimento,omitempty"`
	Beneficiary   string  `json:"beneficiario,omitempty"`
	DocumentType  string  `json:"tipoDocumento,omitempty"`
	DocumentValue string  `json:"valorDocumento,omitempty"`
	IssueDate     string  `json:"dataEmissao,omitempty"`
	Discounts     float64 `json:"descontos,omitempty"`
	Interest      float64 `json:"juros,omitempty"`
	Fine          float64 `json:"multa,omitempty"`
	CIP           string  `json:"cip,omitempty"`
}
