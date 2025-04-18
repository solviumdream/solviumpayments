package mercadopago

import (
	"fmt"
)

type Identification struct {
	client *Client
}

func NewIdentification(client *Client) *Identification {
	return &Identification{
		client: client,
	}
}

func (i *Identification) GetTypes() ([]IdentificationType, error) {
	resp, err := i.client.Request("GET", "/v1/identification_types", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get identification types: %w", err)
	}

	var result []IdentificationType
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}
