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


type BatchDueCharges struct {
	client *Client
}


func NewBatchDueCharges(client *Client) *BatchDueCharges {
	return &BatchDueCharges{
		client: client,
	}
}


func (b *BatchDueCharges) CreateOrUpdate(id string, request BatchDueChargesRequest) (*BatchDueChargesResponse, error) {
	bodyBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := b.client.Request("PUT", fmt.Sprintf("/v2/lotecobv/%s", id), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create/update batch due charges: %w", err)
	}
	defer resp.Body.Close()

	
	if resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create/update batch due charges with status %d: %s", resp.StatusCode, body)
	}

	var batchResp BatchDueChargesResponse
	if err := json.NewDecoder(resp.Body).Decode(&batchResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &batchResp, nil
}


func (b *BatchDueCharges) ReviewBatch(id string, request BatchDueChargesReviewRequest) (*BatchDueChargesResponse, error) {
	bodyBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := b.client.Request("PATCH", fmt.Sprintf("/v2/lotecobv/%s", id), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to review batch due charges: %w", err)
	}
	defer resp.Body.Close()

	
	if resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to review batch due charges with status %d: %s", resp.StatusCode, body)
	}

	var batchResp BatchDueChargesResponse
	if err := json.NewDecoder(resp.Body).Decode(&batchResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &batchResp, nil
}


func (b *BatchDueCharges) GetByID(id string) (*BatchDueChargesResponse, error) {
	resp, err := b.client.Request("GET", fmt.Sprintf("/v2/lotecobv/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get batch due charges: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get batch due charges with status %d: %s", resp.StatusCode, body)
	}

	var batchResp BatchDueChargesResponse
	if err := json.NewDecoder(resp.Body).Decode(&batchResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &batchResp, nil
}


func (b *BatchDueCharges) List(startDate, endDate time.Time, options *ListBatchDueChargesOptions) (*BatchDueChargesListResponse, error) {
	query := url.Values{}
	query.Add("inicio", startDate.Format(time.RFC3339))
	query.Add("fim", endDate.Format(time.RFC3339))

	if options != nil {
		if options.PaginaAtual > 0 {
			query.Add("paginacao.paginaAtual", fmt.Sprintf("%d", options.PaginaAtual))
		}
		if options.ItensPorPagina > 0 {
			query.Add("paginacao.itensPorPagina", fmt.Sprintf("%d", options.ItensPorPagina))
		}
	}

	path := fmt.Sprintf("/v2/lotecobv?%s", query.Encode())
	resp, err := b.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list batch due charges: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list batch due charges with status %d: %s", resp.StatusCode, body)
	}

	var listResp BatchDueChargesListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &listResp, nil
}
