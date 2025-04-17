package efi

type OpenFinanceParticipantRequest struct {
	Organization bool   `json:"organizacao,omitempty"`
	Name         string `json:"nome,omitempty"`
	Modality     string `json:"modalidade,omitempty"`
}

type OpenFinanceParticipantResponse struct {
	Participants []OpenFinanceParticipant `json:"participantes"`
}

type OpenFinanceParticipant struct {
	ID            string                    `json:"identificador"`
	Name          string                    `json:"nome"`
	Description   string                    `json:"descricao"`
	Portal        string                    `json:"portal"`
	Logo          string                    `json:"logo"`
	Organizations []OpenFinanceOrganization `json:"organizacoes"`
}

type OpenFinanceOrganization struct {
	Name   string `json:"nome"`
	CNPJ   string `json:"cnpj"`
	Status string `json:"status"`
}
