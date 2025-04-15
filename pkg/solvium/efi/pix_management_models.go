package efi

type HorarioRefund struct {
	Solicitacao string `json:"solicitacao,omitempty"`
	Liquidacao  string `json:"liquidacao,omitempty"`
}

type DevolucaoRefund struct {
	ID      string        `json:"id,omitempty"`
	RtrID   string        `json:"rtrId,omitempty"`
	Valor   string        `json:"valor,omitempty"`
	Horario HorarioRefund `json:"horario,omitempty"`
	Status  string        `json:"status,omitempty"`
}

type PixDetail struct {
	EndToEndID  string            `json:"endToEndId,omitempty"`
	TxID        string            `json:"txid,omitempty"`
	Valor       string            `json:"valor,omitempty"`
	Chave       string            `json:"chave,omitempty"`
	Horario     string            `json:"horario,omitempty"`
	InfoPagador string            `json:"infoPagador,omitempty"`
	Devolucoes  []DevolucaoRefund `json:"devolucoes,omitempty"`
}

type PixListResponse struct {
	Parametros Parametros  `json:"parametros,omitempty"`
	Pix        []PixDetail `json:"pix,omitempty"`
}

type RefundRequest struct {
	Valor string `json:"valor,omitempty"`
}

type RefundResponse struct {
	ID      string        `json:"id,omitempty"`
	RtrID   string        `json:"rtrId,omitempty"`
	Valor   string        `json:"valor,omitempty"`
	Horario HorarioRefund `json:"horario,omitempty"`
	Status  string        `json:"status,omitempty"`
}
