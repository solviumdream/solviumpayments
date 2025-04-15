package efi

type Horario struct {
	Solicitacao string `json:"solicitacao,omitempty"`
	Liquidacao  string `json:"liquidacao,omitempty"`
}

type PagadorSend struct {
	Chave       string `json:"chave,omitempty"`
	InfoPagador string `json:"infoPagador,omitempty"`
}

type Favorecido struct {
	Chave         string        `json:"chave,omitempty"`
	Identificacao Identificacao `json:"identificacao,omitempty"`
}

type Identificacao struct {
	Nome string `json:"nome,omitempty"`
	CPF  string `json:"cpf,omitempty"`
	CNPJ string `json:"cnpj,omitempty"`
}

type PixSendRequest struct {
	Valor      string      `json:"valor,omitempty"`
	Pagador    PagadorSend `json:"pagador,omitempty"`
	Favorecido Favorecido  `json:"favorecido,omitempty"`
}

type PixSendResponse struct {
	IDEnvio string  `json:"idEnvio,omitempty"`
	E2EID   string  `json:"e2eId,omitempty"`
	Valor   string  `json:"valor,omitempty"`
	Horario Horario `json:"horario,omitempty"`
	Status  string  `json:"status,omitempty"`
}

type PixSentDetail struct {
	EndToEndID  string     `json:"endToEndId,omitempty"`
	IDEnvio     string     `json:"idEnvio,omitempty"`
	Valor       string     `json:"valor,omitempty"`
	Chave       string     `json:"chave,omitempty"`
	Status      string     `json:"status,omitempty"`
	InfoPagador string     `json:"infoPagador,omitempty"`
	Horario     Horario    `json:"horario,omitempty"`
	Favorecido  Favorecido `json:"favorecido,omitempty"`
}

type PixSentListResponse struct {
	Pix        []PixSentDetail `json:"pix,omitempty"`
	Parametros Parametros      `json:"parametros,omitempty"`
}

type DetailQRCodeRequest struct {
	PixCopiaECola string `json:"pixCopiaECola,omitempty"`
}

type QRCodeDetail struct {
	TipoCob            string     `json:"tipoCob,omitempty"`
	TxID               string     `json:"txid,omitempty"`
	Revisao            int        `json:"revisao,omitempty"`
	Calendario         Calendario `json:"calendario,omitempty"`
	Status             string     `json:"status,omitempty"`
	Devedor            Devedor    `json:"devedor,omitempty"`
	Recebedor          Recebedor  `json:"recebedor,omitempty"`
	Valor              Valor      `json:"valor,omitempty"`
	Chave              string     `json:"chave,omitempty"`
	SolicitacaoPagador string     `json:"solicitacaoPagador,omitempty"`
}

type PayQRCodeRequest struct {
	Pagador       PagadorSend `json:"pagador,omitempty"`
	PixCopiaECola string      `json:"pixCopiaECola,omitempty"`
}
