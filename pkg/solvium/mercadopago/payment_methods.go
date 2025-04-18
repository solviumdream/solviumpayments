package mercadopago

import (
	"fmt"
)

type PaymentMethods struct {
	client *Client
}

func NewPaymentMethods(client *Client) *PaymentMethods {
	return &PaymentMethods{
		client: client,
	}
}

func (pm *PaymentMethods) GetAll() ([]PaymentMethod, error) {
	resp, err := pm.client.Request("GET", "/v1/payment_methods", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment methods: %w", err)
	}

	var result []PaymentMethod
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}
