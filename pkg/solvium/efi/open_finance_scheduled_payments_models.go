package efi


type OpenFinanceScheduledPaymentInfo struct {
	Value         string `json:"valor"`
	PayerInfo     string `json:"infoPagador,omitempty"`
	OwnID         string `json:"idProprio,omitempty"`
	ScheduledDate string `json:"dataAgendamento"`
	TransactionID string `json:"identificadorTransacao,omitempty"`
}

type OpenFinanceScheduledPaymentRequest struct {
	Payer     OpenFinancePaymentPayer         `json:"pagador"`
	Recipient OpenFinanceRecipient            `json:"favorecido"`
	Payment   OpenFinanceScheduledPaymentInfo `json:"pagamento"`
}





type OpenFinanceScheduledPayment struct {
	PaymentID     string                   `json:"identificadorPagamento"`
	EndToEndID    string                   `json:"endToEndId"`
	Value         string                   `json:"valor"`
	Status        OpenFinancePaymentStatus `json:"status"`
	OperationDate string                   `json:"dataOperacao"`
	CreatedAt     string                   `json:"dataCriacao"`
	OwnID         string                   `json:"idProprio,omitempty"`
	Refunds       []OpenFinanceRefund      `json:"devolucoes,omitempty"`
}

type OpenFinanceScheduledPaymentList struct {
	Payments []OpenFinanceScheduledPayment `json:"pagamentos"`
	Total    int                           `json:"total"`
	PerPage  int                           `json:"porPagina"`
	Last     string                        `json:"ultimo"`
	Next     string                        `json:"proximo"`
	Previous string                        `json:"anterior"`
	Current  string                        `json:"atual"`
}


type OpenFinanceScheduledCancellationResponse struct {
	PaymentID        string                   `json:"identificadorPagamento"`
	Status           OpenFinancePaymentStatus `json:"status"`
	CancellationDate string                   `json:"dataCancelamento"`
}


type OpenFinanceScheduledRefundRequest struct {
	EndToEndID string `json:"endToEndId"`
	Value      string `json:"valor"`
}


