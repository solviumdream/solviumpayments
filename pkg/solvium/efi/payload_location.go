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


type PayloadLocation struct {
	client *Client
}


func NewPayloadLocation(client *Client) *PayloadLocation {
	return &PayloadLocation{
		client: client,
	}
}


func (p *PayloadLocation) Create(request CreatePayloadLocationRequest) (*PayloadLocationResponse, error) {
	bodyBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := p.client.Request("POST", "/v2/loc", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create payload location: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create payload location with status %d: %s", resp.StatusCode, body)
	}

	var location PayloadLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &location, nil
}


func (p *PayloadLocation) List(startDate, endDate time.Time, options *ListPayloadLocationsOptions) (*PayloadLocationListResponse, error) {
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

	path := fmt.Sprintf("/v2/loc?%s", query.Encode())
	resp, err := p.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list payload locations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list payload locations with status %d: %s", resp.StatusCode, body)
	}

	var listResp PayloadLocationListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &listResp, nil
}


func (p *PayloadLocation) GetByID(id int64) (*PayloadLocationResponse, error) {
	resp, err := p.client.Request("GET", fmt.Sprintf("/v2/loc/%d", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get payload location: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get payload location with status %d: %s", resp.StatusCode, body)
	}

	var location PayloadLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &location, nil
}


func (p *PayloadLocation) GenerateQRCode(id int64) (*PayloadLocationQRCodeResponse, error) {
	resp, err := p.client.Request("GET", fmt.Sprintf("/v2/loc/%d/qrcode", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to generate QR code with status %d: %s", resp.StatusCode, body)
	}

	var qrCode PayloadLocationQRCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&qrCode); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &qrCode, nil
}


func (p *PayloadLocation) UnlinkTxID(id int64) (*PayloadLocationResponse, error) {
	resp, err := p.client.Request("DELETE", fmt.Sprintf("/v2/loc/%d/txid", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to unlink txid: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to unlink txid with status %d: %s", resp.StatusCode, body)
	}

	var location PayloadLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &location, nil
}
