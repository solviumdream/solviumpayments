package efi

type EnderecoDevedor struct {
	Logradouro string `json:"logradouro,omitempty"`
	Cidade     string `json:"cidade,omitempty"`
	UF         string `json:"uf,omitempty"`
	CEP        string `json:"cep,omitempty"`
}

type DevedorDueCharge struct {
	EnderecoDevedor
	CPF   string `json:"cpf,omitempty"`
	CNPJ  string `json:"cnpj,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Email string `json:"email,omitempty"`
}

type Recebedor struct {
	Logradouro string `json:"logradouro,omitempty"`
	Cidade     string `json:"cidade,omitempty"`
	UF         string `json:"uf,omitempty"`
	CEP        string `json:"cep,omitempty"`
	CNPJ       string `json:"cnpj,omitempty"`
	Nome       string `json:"nome,omitempty"`
}

type CalendarioDueCharge struct {
	Criacao                string `json:"criacao,omitempty"`
	DataDeVencimento       string `json:"dataDeVencimento,omitempty"`
	ValidadeAposVencimento int    `json:"validadeAposVencimento,omitempty"`
}

type Multa struct {
	Modalidade int    `json:"modalidade,omitempty"`
	ValorPerc  string `json:"valorPerc,omitempty"`
}

type Juros struct {
	Modalidade int    `json:"modalidade,omitempty"`
	ValorPerc  string `json:"valorPerc,omitempty"`
}

type DescontoDataFixa struct {
	Data      string `json:"data,omitempty"`
	ValorPerc string `json:"valorPerc,omitempty"`
}

type Desconto struct {
	Modalidade       int                `json:"modalidade,omitempty"`
	DescontoDataFixa []DescontoDataFixa `json:"descontoDataFixa,omitempty"`
	ValorPerc        string             `json:"valorPerc,omitempty"`
}

type Abatimento struct {
	Modalidade int    `json:"modalidade,omitempty"`
	ValorPerc  string `json:"valorPerc,omitempty"`
}

type ValorDueCharge struct {
	Original   string     `json:"original,omitempty"`
	Multa      Multa      `json:"multa,omitempty"`
	Juros      Juros      `json:"juros,omitempty"`
	Desconto   Desconto   `json:"desconto,omitempty"`
	Abatimento Abatimento `json:"abatimento,omitempty"`
}

type CreateDueChargeRequest struct {
	Calendario         CalendarioDueCharge `json:"calendario,omitempty"`
	Devedor            DevedorDueCharge    `json:"devedor,omitempty"`
	Valor              ValorDueCharge      `json:"valor,omitempty"`
	Chave              string              `json:"chave,omitempty"`
	SolicitacaoPagador string              `json:"solicitacaoPagador,omitempty"`
	InfoAdicionais     []InfoAdicional     `json:"infoAdicionais,omitempty"`
}

type ReviewDueChargeRequest struct {
	Loc                LocInfo          `json:"loc,omitempty"`
	Devedor            DevedorDueCharge `json:"devedor,omitempty"`
	Valor              ValorDueCharge   `json:"valor,omitempty"`
	SolicitacaoPagador string           `json:"solicitacaoPagador,omitempty"`
	InfoAdicionais     []InfoAdicional  `json:"infoAdicionais,omitempty"`
}

type DueChargeResponse struct {
	Calendario         CalendarioDueCharge `json:"calendario,omitempty"`
	TxID               string              `json:"txid,omitempty"`
	Revisao            int                 `json:"revisao,omitempty"`
	Loc                LocInfo             `json:"loc,omitempty"`
	Status             string              `json:"status,omitempty"`
	Devedor            DevedorDueCharge    `json:"devedor,omitempty"`
	Recebedor          Recebedor           `json:"recebedor,omitempty"`
	Valor              ValorDueCharge      `json:"valor,omitempty"`
	Chave              string              `json:"chave,omitempty"`
	SolicitacaoPagador string              `json:"solicitacaoPagador,omitempty"`
	PixCopiaECola      string              `json:"pixCopiaECola,omitempty"`
	Pix                []PixInfo           `json:"pix,omitempty"`
	InfoAdicionais     []InfoAdicional     `json:"infoAdicionais,omitempty"`
}

type ListDueChargesResponse struct {
	Parametros Parametros          `json:"parametros,omitempty"`
	Cobs       []DueChargeResponse `json:"cobs,omitempty"`
}
