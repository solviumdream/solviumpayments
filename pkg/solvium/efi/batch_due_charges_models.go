package efi

type BatchDueChargesRequest struct {
	Descricao string                   `json:"descricao,omitempty"`
	CobsV     []CreateDueChargeRequest `json:"cobsv"`
}

type BatchDueChargesItem struct {
	Criacao  string    `json:"criacao,omitempty"`
	TxID     string    `json:"txid,omitempty"`
	Status   string    `json:"status,omitempty"`
	Problema *Problema `json:"problema,omitempty"`
}

type Problema struct {
	Type      string     `json:"type,omitempty"`
	Title     string     `json:"title,omitempty"`
	Status    int        `json:"status,omitempty"`
	Detail    string     `json:"detail,omitempty"`
	Violacoes []Violacao `json:"violacoes,omitempty"`
}

type Violacao struct {
	Razao       string `json:"razao,omitempty"`
	Propriedade string `json:"propriedade,omitempty"`
}

type BatchDueChargesResponse struct {
	Descricao string                `json:"descricao,omitempty"`
	Criacao   string                `json:"criacao,omitempty"`
	CobsV     []BatchDueChargesItem `json:"cobsv,omitempty"`
}

type BatchDueChargesListResponse struct {
	Parametros Parametros                `json:"parametros,omitempty"`
	Lotes      []BatchDueChargesResponse `json:"lotes,omitempty"`
}

type BatchDueChargesReviewRequest struct {
	CobsV []BatchDueChargesReviewItem `json:"cobsv"`
}

type BatchDueChargesReviewItem struct {
	Calendario CalendarioDueCharge `json:"calendario,omitempty"`
	TxID       string              `json:"txid"`
	Valor      Valor               `json:"valor,omitempty"`
}

type ListBatchDueChargesOptions struct {
	PaginaAtual    int
	ItensPorPagina int
}
