package mercadopago

import (
	"fmt"
)

type Payment struct {
	client *Client
}

func NewPayment(client *Client) *Payment {
	return &Payment{
		client: client,
	}
}

func (p *Payment) Create(request PaymentRequest) (*PaymentResponse, error) {
	resp, err := p.client.Request("POST", "/checkout/preferences", request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	var result PaymentResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *Payment) Update(paymentID string, request PaymentRequest) (*PaymentResponse, error) {
	path := fmt.Sprintf("/checkout/preferences/%s", paymentID)
	resp, err := p.client.Request("PUT", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update payment: %w", err)
	}

	var result PaymentResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *Payment) Get(paymentID string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/checkout/preferences/%s", paymentID)
	resp, err := p.client.Request("GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}

	var result PaymentResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *Payment) Search(params PaymentSearchParams) (*PaymentSearchResponse, error) {
	queryParams := make(map[string]string)
	for key, value := range params {
		queryParams[key] = value
	}

	resp, err := p.client.Request("GET", "/checkout/preferences/search", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to search payments: %w", err)
	}

	var result PaymentSearchResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *Payment) Consult(paymentID string) (*PaymentConsultResponse, error) {
	path := fmt.Sprintf("/v1/payments/%s", paymentID)
	resp, err := p.client.Request("GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to consult payment: %w", err)
	}

	var result PaymentConsultResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
