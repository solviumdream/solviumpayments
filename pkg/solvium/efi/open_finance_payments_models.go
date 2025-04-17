package efi


type OpenFinanceAccountType string

const (
	AccountTypeChecking OpenFinanceAccountType = "CACC" 
	AccountTypeSavings  OpenFinanceAccountType = "SVGS" 
	AccountTypePayment  OpenFinanceAccountType = "TRAN" 
	AccountTypeSalary   OpenFinanceAccountType = "SLRY" 
)


type OpenFinancePaymentStatus string

const (
	PaymentStatusAccepted  OpenFinancePaymentStatus = "aceito"
	PaymentStatusRejected  OpenFinancePaymentStatus = "rejeitado"
	PaymentStatusPending   OpenFinancePaymentStatus = "pendente"
	PaymentStatusCompleted OpenFinancePaymentStatus = "concluido"
)



type OpenFinancePaymentPayer struct {
	ParticipantID string `json:"idParticipante"`
	CPF           string `json:"cpf,omitempty"`
	CNPJ          string `json:"cnpj,omitempty"`
}

type OpenFinanceBankAccount struct {
	Name        string                 `json:"nome"`
	Document    string                 `json:"documento"`
	BankCode    string                 `json:"codigoBanco"`
	Branch      string                 `json:"agencia"`
	Account     string                 `json:"conta"`
	AccountType OpenFinanceAccountType `json:"tipoConta"`
}

type OpenFinancePixKey struct {
	KeyType string `json:"tipoChave"` 
	Key     string `json:"chave"`
	Name    string `json:"nome,omitempty"`
}

type OpenFinanceQRCode struct {
	QRCode string `json:"qrCode"`
	Name   string `json:"nome,omitempty"`
}

type OpenFinanceRecipient struct {
	BankAccount *OpenFinanceBankAccount `json:"contaBanco,omitempty"`
	PixKey      *OpenFinancePixKey      `json:"chave,omitempty"`
	QRCode      *OpenFinanceQRCode      `json:"qrCode,omitempty"`
}

type OpenFinancePaymentInfo struct {
	Value         string `json:"valor"`
	PayerInfo     string `json:"infoPagador,omitempty"`
	OwnID         string `json:"idProprio,omitempty"`
	TransactionID string `json:"identificadorTransacao,omitempty"`
}

type OpenFinancePaymentRequest struct {
	Payer     OpenFinancePaymentPayer `json:"pagador"`
	Recipient OpenFinanceRecipient    `json:"favorecido"`
	Payment   OpenFinancePaymentInfo  `json:"pagamento"`
}



type OpenFinancePaymentResponse struct {
	PaymentID   string `json:"identificadorPagamento"`
	RedirectURI string `json:"redirectURI"`
}



type OpenFinanceRefund struct {
	RefundID  string                   `json:"identificadorDevolucao"`
	Value     string                   `json:"valor"`
	Status    OpenFinancePaymentStatus `json:"status"`
	CreatedAt string                   `json:"dataCriacao"`
}

type OpenFinancePayment struct {
	PaymentID  string                   `json:"identificadorPagamento"`
	EndToEndID string                   `json:"endToEndId"`
	Value      string                   `json:"valor"`
	Status     OpenFinancePaymentStatus `json:"status"`
	CreatedAt  string                   `json:"dataCriacao"`
	Refunds    []OpenFinanceRefund      `json:"devolucoes,omitempty"`
	OwnID      string                   `json:"idProprio,omitempty"`
}

type OpenFinancePaymentList struct {
	Payments []OpenFinancePayment `json:"pagamentos"`
	Total    int                  `json:"total"`
	PerPage  int                  `json:"porPagina"`
	Last     string               `json:"ultimo"`
	Next     string               `json:"proximo"`
	Previous string               `json:"anterior"`
	Current  string               `json:"atual"`
}



type OpenFinanceRefundRequest struct {
	Value string `json:"valor"`
}

type OpenFinanceRefundResponse struct {
	PaymentID  string                   `json:"identificadorPagamento"`
	EndToEndID string                   `json:"endToEndId"`
	Value      string                   `json:"valor"`
	CreatedAt  string                   `json:"dataCriacao"`
	Status     OpenFinancePaymentStatus `json:"status"`
}
