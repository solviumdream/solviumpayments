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

type DueCharges struct {
	client *Client
}

func NewDueCharges(client *Client) *DueCharges {
	return &DueCharges{
		client: client,
	}
}

func (c *DueCharges) Create(txid string, req CreateDueChargeRequest) (*DueChargeResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.Request("PUT", fmt.Sprintf("/v2/cobv/%s", txid), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create due charge: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create due charge with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp DueChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}

func (c *DueCharges) Review(txid string, req ReviewDueChargeRequest) (*DueChargeResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.Request("PATCH", fmt.Sprintf("/v2/cobv/%s", txid), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to review due charge: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to review due charge with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp DueChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}

func (c *DueCharges) Get(txid string, revision int) (*DueChargeResponse, error) {
	path := fmt.Sprintf("/v2/cobv/%s", txid)
	if revision > 0 {
		path = fmt.Sprintf("%s?revisao=%d", path, revision)
	}

	resp, err := c.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get due charge: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get due charge with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp DueChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}

type ListDueChargesOptions struct {
	CPF              string
	CNPJ             string
	Status           string
	IncluirReciboPix bool
	PaginaAtual      int
	ItensPorPagina   int
}

func (c *DueCharges) List(startDate, endDate time.Time, options *ListDueChargesOptions) (*ListDueChargesResponse, error) {
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

	path := fmt.Sprintf("/v2/cobv?%s", query.Encode())
	resp, err := c.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list due charges: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list due charges with status %d: %s", resp.StatusCode, body)
	}

	var listResp ListDueChargesResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &listResp, nil
}
