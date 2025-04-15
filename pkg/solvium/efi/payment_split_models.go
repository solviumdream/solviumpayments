package efi




type SplitLancamento struct {
	Imediato bool `json:"imediato"`
}


type SplitFavorecido struct {
	CPF   string `json:"cpf,omitempty"`
	CNPJ  string `json:"cnpj,omitempty"`
	Conta string `json:"conta"`
}


type SplitRepasse struct {
	Tipo       string          `json:"tipo"`
	Valor      string          `json:"valor"`
	Favorecido SplitFavorecido `json:"favorecido"`
}


type SplitMinhaParte struct {
	Tipo  string `json:"tipo"`
	Valor string `json:"valor"`
}


type SplitConfig struct {
	DivisaoTarifa string          `json:"divisaoTarifa"`
	MinhaParte    SplitMinhaParte `json:"minhaParte"`
	Repasses      []SplitRepasse  `json:"repasses"`
}


type PaymentSplitConfigRequest struct {
	Descricao  string          `json:"descricao"`
	Lancamento SplitLancamento `json:"lancamento"`
	Split      SplitConfig     `json:"split"`
}


type PaymentSplitConfigResponse struct {
	ID         string          `json:"id,omitempty"`
	Status     string          `json:"status,omitempty"`
	TxID       string          `json:"txid,omitempty"`
	Revisao    int             `json:"revisao,omitempty"`
	Descricao  string          `json:"descricao,omitempty"`
	Lancamento SplitLancamento `json:"lancamento,omitempty"`
	Split      SplitConfig     `json:"split,omitempty"`
}


type SplitConfigInfo struct {
	ID        string `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
	Descricao string `json:"descricao,omitempty"`
}


type SplitChargeResponse struct {
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
	Config             SplitConfigInfo     `json:"config,omitempty"`
}
