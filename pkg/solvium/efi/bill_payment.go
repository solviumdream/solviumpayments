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

type BillPayment struct {
	client *Client
}

func NewBillPayment(client *Client) *BillPayment {
	return &BillPayment{
		client: client,
	}
}

func (b *BillPayment) DetailBarcode(barcode string) (*BillDetails, error) {
	path := fmt.Sprintf("/v1/codBarras/%s", barcode)

	resp, err := b.client.Request(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to detail barcode: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to detail barcode with status %d: %s", resp.StatusCode, body)
	}

	var details BillDetails
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, fmt.Errorf("failed to decode barcode details: %w", err)
	}

	return &details, nil
}

func (b *BillPayment) RequestPayment(barcode string, request *BillPaymentRequest) (*BillPaymentResponse, error) {
	path := fmt.Sprintf("/v1/codBarras/%s", barcode)

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payment request: %w", err)
	}

	resp, err := b.client.Request(http.MethodPost, path, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to request payment: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to request payment with status %d: %s", resp.StatusCode, body)
	}

	var paymentResponse BillPaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to decode payment response: %w", err)
	}

	return &paymentResponse, nil
}

func (b *BillPayment) GetPaymentStatus(paymentID string) (*BillPaymentResponse, error) {
	path := fmt.Sprintf("/v1/%s", paymentID)

	resp, err := b.client.Request(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment status: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get payment status with status %d: %s", resp.StatusCode, body)
	}

	var paymentResponse BillPaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to decode payment status: %w", err)
	}

	return &paymentResponse, nil
}

func (b *BillPayment) GetPaymentSummary(startDate, endDate string) (*BillPaymentSummary, error) {
	query := url.Values{}
	query.Add("dataInicial", startDate)
	query.Add("dataFinal", endDate)

	path := fmt.Sprintf("/v1/resumo?%s", query.Encode())

	resp, err := b.client.Request(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment summary: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get payment summary with status %d: %s", resp.StatusCode, body)
	}

	var summary BillPaymentSummary
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return nil, fmt.Errorf("failed to decode payment summary: %w", err)
	}

	return &summary, nil
}

func (b *BillPayment) GetPaymentSummaryByDateRange(days int) (*BillPaymentSummary, error) {
	endDate := time.Now().Format("2006-01-02")
	startDate := time.Now().AddDate(0, 0, -days).Format("2006-01-02")

	return b.GetPaymentSummary(startDate, endDate)
}
