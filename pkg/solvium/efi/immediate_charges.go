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


type ImmediateCharges struct {
	client *Client
}


func NewImmediateCharges(client *Client) *ImmediateCharges {
	return &ImmediateCharges{
		client: client,
	}
}


func (c *ImmediateCharges) CreateWithoutTxid(req CreateImmediateChargeRequest) (*ImmediateChargeResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.Request("POST", "/v2/cob", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create immediate charge: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create immediate charge with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp ImmediateChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}


func (c *ImmediateCharges) CreateWithTxid(txid string, req CreateImmediateChargeRequest) (*ImmediateChargeResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.Request("PUT", fmt.Sprintf("/v2/cob/%s", txid), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create immediate charge: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create immediate charge with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp ImmediateChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}


func (c *ImmediateCharges) ReviewCharge(txid string, req ReviewChargeRequest) (*ImmediateChargeResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.Request("PATCH", fmt.Sprintf("/v2/cob/%s", txid), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to review charge: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to review charge with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp ImmediateChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}


func (c *ImmediateCharges) GetCharge(txid string, revision int) (*ImmediateChargeResponse, error) {
	path := fmt.Sprintf("/v2/cob/%s", txid)
	if revision > 0 {
		path = fmt.Sprintf("%s?revisao=%d", path, revision)
	}

	resp, err := c.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get charge: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get charge with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp ImmediateChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}


type ListChargesOptions struct {
	CPF              string
	CNPJ             string
	Status           string
	IncluirReciboPix bool
	PaginaAtual      int
	ItensPorPagina   int
}


func (c *ImmediateCharges) ListCharges(startDate, endDate time.Time, options *ListChargesOptions) (*ListChargesResponse, error) {
	query := url.Values{}
	query.Add("inicio", startDate.Format(time.RFC3339))
	query.Add("fim", endDate.Format(time.RFC3339))

	if options != nil {
		if options.CPF != "" {
			query.Add("cpf", options.CPF)
		}
		if options.CNPJ != "" {
			query.Add("cnpj", options.CNPJ)
		}
		if options.Status != "" {
			query.Add("status", options.Status)
		}
		if options.IncluirReciboPix {
			query.Add("paginacao.incluirReciboPix", "true")
		}
		if options.PaginaAtual > 0 {
			query.Add("paginacao.paginaAtual", fmt.Sprintf("%d", options.PaginaAtual))
		}
		if options.ItensPorPagina > 0 {
			query.Add("paginacao.itensPorPagina", fmt.Sprintf("%d", options.ItensPorPagina))
		}
	}

	path := fmt.Sprintf("/v2/cob?%s", query.Encode())
	resp, err := c.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list charges: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list charges with status %d: %s", resp.StatusCode, body)
	}

	var listResp ListChargesResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &listResp, nil
}
