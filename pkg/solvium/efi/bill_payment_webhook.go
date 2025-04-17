package efi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type BillPaymentWebhookClient struct {
	client *Client
}

func NewBillPaymentWebhook(client *Client) *BillPaymentWebhookClient {
	return &BillPaymentWebhookClient{
		client: client,
	}
}

func (w *BillPaymentWebhookClient) Create(webhookURL string) (*BillPaymentWebhookResponse, error) {
	request := BillPaymentWebhookRequest{
		URL: webhookURL,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal webhook request: %w", err)
	}

	resp, err := w.client.Request(http.MethodPut, "/v1/webhook", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create webhook with status %d: %s", resp.StatusCode, body)
	}

	var webhookResponse BillPaymentWebhookResponse
	if err := json.NewDecoder(resp.Body).Decode(&webhookResponse); err != nil {
		return nil, fmt.Errorf("failed to decode webhook response: %w", err)
	}

	return &webhookResponse, nil
}

func (w *BillPaymentWebhookClient) List(startDate, endDate time.Time) (*BillPaymentWebhookListResponse, error) {
	query := url.Values{}
	query.Add("dataInicio", startDate.Format(time.RFC3339))
	query.Add("dataFim", endDate.Format(time.RFC3339))

	path := fmt.Sprintf("/v1/webhook?%s", query.Encode())

	resp, err := w.client.Request(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list webhooks: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list webhooks with status %d: %s", resp.StatusCode, body)
	}

	var listResponse BillPaymentWebhookListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResponse); err != nil {
		return nil, fmt.Errorf("failed to decode webhook list response: %w", err)
	}

	return &listResponse, nil
}

func (w *BillPaymentWebhookClient) Delete(webhookURL string) error {
	request := BillPaymentWebhookRequest{
		URL: webhookURL,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook delete request: %w", err)
	}

	resp, err := w.client.Request(http.MethodDelete, "/v1/webhook", bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete webhook with status %d: %s", resp.StatusCode, body)
	}

	return nil
}

func (w *BillPaymentWebhookClient) ListByDateRange(days int) (*BillPaymentWebhookListResponse, error) {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -days)

	return w.List(startDate, endDate)
}

func ParseBillPaymentWebhookCallback(payload []byte) (*BillPaymentWebhookCallback, error) {
	var callback BillPaymentWebhookCallback
	if err := json.Unmarshal(payload, &callback); err != nil {
		return nil, fmt.Errorf("failed to parse webhook callback: %w", err)
	}

	return &callback, nil
}
