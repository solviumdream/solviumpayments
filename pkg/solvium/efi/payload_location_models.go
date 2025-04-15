package efi

type PayloadLocationType string

const (
	PayloadLocationTypeCOB  PayloadLocationType = "cob"
	PayloadLocationTypeCOBV PayloadLocationType = "cobv"
)

type CreatePayloadLocationRequest struct {
	TipoCob PayloadLocationType `json:"tipoCob"`
}

type PayloadLocationResponse struct {
	ID       int64               `json:"id,omitempty"`
	Location string              `json:"location,omitempty"`
	TipoCob  PayloadLocationType `json:"tipoCob,omitempty"`
	Criacao  string              `json:"criacao,omitempty"`
	TxID     string              `json:"txid,omitempty"`
}

type PayloadLocationQRCodeResponse struct {
	QRCode           string `json:"qrcode,omitempty"`
	ImagemQRCode     string `json:"imagemQrcode,omitempty"`
	LinkVisualizacao string `json:"linkVisualizacao,omitempty"`
}

type PayloadLocationListResponse struct {
	Parametros Parametros                `json:"parametros,omitempty"`
	Loc        []PayloadLocationResponse `json:"loc,omitempty"`
}

type ListPayloadLocationsOptions struct {
	PaginaAtual    int
	ItensPorPagina int
}
