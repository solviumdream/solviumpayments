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

type PixManagement struct {
	client *Client
}

func NewPixManagement(client *Client) *PixManagement {
	return &PixManagement{
		client: client,
	}
}

func (p *PixManagement) GetByE2EID(e2eID string) (*PixDetail, error) {
	resp, err := p.client.Request("GET", fmt.Sprintf("/v2/pix/%s", e2eID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get Pix: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get Pix with status %d: %s", resp.StatusCode, body)
	}

	var pixDetail PixDetail
	if err := json.NewDecoder(resp.Body).Decode(&pixDetail); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &pixDetail, nil
}

type ListReceivedOptions struct {
	TxID        string
	CPF         string
	CNPJ        string
	Status      string
	PaginaAtual int
	ItensPagina int
}

func (p *PixManagement) ListReceived(startDate, endDate time.Time, options *ListReceivedOptions) (*PixListResponse, error) {
	query := url.Values{}
	query.Add("inicio", startDate.Format(time.RFC3339))
	query.Add("fim", endDate.Format(time.RFC3339))

	if options != nil {
		if options.TxID != "" {
			query.Add("txid", options.TxID)
		}
		if options.CPF != "" {
			query.Add("cpf", options.CPF)
		}
		if options.CNPJ != "" {
			query.Add("cnpj", options.CNPJ)
		}
		if options.Status != "" {
			query.Add("status", options.Status)
		}
		if options.PaginaAtual > 0 {
			query.Add("paginacao.paginaAtual", fmt.Sprintf("%d", options.PaginaAtual))
		}
		if options.ItensPagina > 0 {
			query.Add("paginacao.itensPorPagina", fmt.Sprintf("%d", options.ItensPagina))
		}
	}

	path := fmt.Sprintf("/v2/pix?%s", query.Encode())
	resp, err := p.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list received Pix: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list received Pix with status %d: %s", resp.StatusCode, body)
	}

	var listResp PixListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &listResp, nil
}

func (p *PixManagement) RequestRefund(e2eID, refundID string, req RefundRequest) (*RefundResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := p.client.Request("PUT", fmt.Sprintf("/v2/pix/%s/devolucao/%s", e2eID, refundID), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to request refund: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to request refund with status %d: %s", resp.StatusCode, body)
	}

	var refundResp RefundResponse
	if err := json.NewDecoder(resp.Body).Decode(&refundResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &refundResp, nil
}

func (p *PixManagement) GetRefund(e2eID, refundID string) (*RefundResponse, error) {
	resp, err := p.client.Request("GET", fmt.Sprintf("/v2/pix/%s/devolucao/%s", e2eID, refundID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get refund: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get refund with status %d: %s", resp.StatusCode, body)
	}

	var refundResp RefundResponse
	if err := json.NewDecoder(resp.Body).Decode(&refundResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &refundResp, nil
}
