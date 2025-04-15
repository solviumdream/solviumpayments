package efi

import "time"




type Calendario struct {
	Criacao   string `json:"criacao,omitempty"`
	Expiracao int    `json:"expiracao,omitempty"`
}


type Devedor struct {
	CPF   string `json:"cpf,omitempty"`
	CNPJ  string `json:"cnpj,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Email string `json:"email,omitempty"`
}


type Valor struct {
	Original   string `json:"original,omitempty"`
	Modalidade int    `json:"modalidade,omitempty"`
}


type LocInfo struct {
	ID       int    `json:"id,omitempty"`
	Location string `json:"location,omitempty"`
	TipoCob  string `json:"tipoCob,omitempty"`
}


type Pagador struct {
	CPF   string `json:"cpf,omitempty"`
	CNPJ  string `json:"cnpj,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Email string `json:"email,omitempty"`
}


type HorarioInfo struct {
	Solicitacao time.Time `json:"solicitacao,omitempty"`
	Liquidacao  time.Time `json:"liquidacao,omitempty"`
}


type Devolucao struct {
	ID      string      `json:"id,omitempty"`
	RtrID   string      `json:"rtrId,omitempty"`
	Valor   string      `json:"valor,omitempty"`
	Horario HorarioInfo `json:"horario,omitempty"`
	Status  string      `json:"status,omitempty"`
}


type PixInfo struct {
	EndToEndID  string      `json:"endToEndId,omitempty"`
	TxID        string      `json:"txid,omitempty"`
	Valor       string      `json:"valor,omitempty"`
	Horario     string      `json:"horario,omitempty"`
	Pagador     Pagador     `json:"pagador,omitempty"`
	InfoPagador string      `json:"infoPagador,omitempty"`
	Devolucoes  []Devolucao `json:"devolucoes,omitempty"`
}


type Paginacao struct {
	PaginaAtual            int `json:"paginaAtual,omitempty"`
	ItensPorPagina         int `json:"itensPorPagina,omitempty"`
	QuantidadeDePaginas    int `json:"quantidadeDePaginas,omitempty"`
	QuantidadeTotalDeItens int `json:"quantidadeTotalDeItens,omitempty"`
}


type Parametros struct {
	Inicio    string    `json:"inicio,omitempty"`
	Fim       string    `json:"fim,omitempty"`
	Paginacao Paginacao `json:"paginacao,omitempty"`
}




type CreateImmediateChargeRequest struct {
	Calendario         Calendario      `json:"calendario,omitempty"`
	Devedor            Devedor         `json:"devedor,omitempty"`
	Valor              Valor           `json:"valor,omitempty"`
	Chave              string          `json:"chave,omitempty"`
	SolicitacaoPagador string          `json:"solicitacaoPagador,omitempty"`
	InfoAdicionais     []InfoAdicional `json:"infoAdicionais,omitempty"`
}


type InfoAdicional struct {
	Nome  string `json:"nome,omitempty"`
	Valor string `json:"valor,omitempty"`
}


type ImmediateChargeResponse struct {
	Calendario         Calendario      `json:"calendario,omitempty"`
	TxID               string          `json:"txid,omitempty"`
	Revisao            int             `json:"revisao,omitempty"`
	Loc                LocInfo         `json:"loc,omitempty"`
	Location           string          `json:"location,omitempty"`
	Status             string          `json:"status,omitempty"`
	Devedor            Devedor         `json:"devedor,omitempty"`
	Valor              Valor           `json:"valor,omitempty"`
	Chave              string          `json:"chave,omitempty"`
	SolicitacaoPagador string          `json:"solicitacaoPagador,omitempty"`
	PixCopiaECola      string          `json:"pixCopiaECola,omitempty"`
	Pix                []PixInfo       `json:"pix,omitempty"`
	InfoAdicionais     []InfoAdicional `json:"infoAdicionais,omitempty"`
}


type ReviewChargeRequest struct {
	Loc                LocInfo         `json:"loc,omitempty"`
	Devedor            Devedor         `json:"devedor,omitempty"`
	Valor              Valor           `json:"valor,omitempty"`
	SolicitacaoPagador string          `json:"solicitacaoPagador,omitempty"`
	InfoAdicionais     []InfoAdicional `json:"infoAdicionais,omitempty"`
}


type ListChargesResponse struct {
	Parametros Parametros                `json:"parametros,omitempty"`
	Cobs       []ImmediateChargeResponse `json:"cobs,omitempty"`
}
