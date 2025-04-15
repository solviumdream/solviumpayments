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


type PixSend struct {
	client *Client
}


func NewPixSend(client *Client) *PixSend {
	return &PixSend{
		client: client,
	}
}


func (p *PixSend) Send(idEnvio string, req PixSendRequest) (*PixSendResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := p.client.Request("PUT", fmt.Sprintf("/v3/gn/pix/%s", idEnvio), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to send Pix: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to send Pix with status %d: %s", resp.StatusCode, body)
	}

	var pixResp PixSendResponse
	if err := json.NewDecoder(resp.Body).Decode(&pixResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	
	if bucketSize := resp.Header.Get("Bucket-Size"); bucketSize != "" {
		fmt.Printf("Pix sending bucket size: %s\n", bucketSize)
	}

	
	if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
		fmt.Printf("Retry sending after %s seconds\n", retryAfter)
	}

	return &pixResp, nil
}


func (p *PixSend) GetByE2EID(e2eID string) (*PixSentDetail, error) {
	resp, err := p.client.Request("GET", fmt.Sprintf("/v2/gn/pix/enviados/%s", e2eID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get sent Pix: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get sent Pix with status %d: %s", resp.StatusCode, body)
	}

	var pixDetail PixSentDetail
	if err := json.NewDecoder(resp.Body).Decode(&pixDetail); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &pixDetail, nil
}


func (p *PixSend) GetByIDEnvio(idEnvio string) (*PixSentDetail, error) {
	resp, err := p.client.Request("GET", fmt.Sprintf("/v2/gn/pix/enviados/id-envio/%s", idEnvio), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get sent Pix: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get sent Pix with status %d: %s", resp.StatusCode, body)
	}

	var pixDetail PixSentDetail
	if err := json.NewDecoder(resp.Body).Decode(&pixDetail); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &pixDetail, nil
}


type ListSentOptions struct {
	Status      string
	InfoPagador string
	CPF         string
	CNPJ        string
}


func (p *PixSend) List(startDate, endDate time.Time, options *ListSentOptions) (*PixSentListResponse, error) {
	query := url.Values{}
	query.Add("inicio", startDate.Format(time.RFC3339))
	query.Add("fim", endDate.Format(time.RFC3339))

	if options != nil {
		if options.Status != "" {
			query.Add("status", options.Status)
		}
		if options.InfoPagador != "" {
			query.Add("infoPagador", options.InfoPagador)
		}
		if options.CPF != "" {
			query.Add("cpf", options.CPF)
		}
		if options.CNPJ != "" {
			query.Add("cnpj", options.CNPJ)
		}
	}

	path := fmt.Sprintf("/v2/gn/pix/enviados?%s", query.Encode())
	resp, err := p.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list sent Pix: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list sent Pix with status %d: %s", resp.StatusCode, body)
	}

	var listResp PixSentListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &listResp, nil
}


func (p *PixSend) DetailQRCode(req DetailQRCodeRequest) (*QRCodeDetail, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := p.client.Request("POST", "/v2/gn/qrcodes/detalhar", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to detail QR code: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to detail QR code with status %d: %s", resp.StatusCode, body)
	}

	var qrDetail QRCodeDetail
	if err := json.NewDecoder(resp.Body).Decode(&qrDetail); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &qrDetail, nil
}


func (p *PixSend) PayQRCode(idEnvio string, req PayQRCodeRequest) (*PixSendResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := p.client.Request("PUT", fmt.Sprintf("/v2/gn/pix/%s/qrcode", idEnvio), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to pay QR code: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to pay QR code with status %d: %s", resp.StatusCode, body)
	}

	var pixResp PixSendResponse
	if err := json.NewDecoder(resp.Body).Decode(&pixResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &pixResp, nil
}
