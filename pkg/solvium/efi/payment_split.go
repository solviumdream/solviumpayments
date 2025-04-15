package efi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type PaymentSplit struct {
	client *Client
}


func NewPaymentSplit(client *Client) *PaymentSplit {
	return &PaymentSplit{
		client: client,
	}
}


func (p *PaymentSplit) CreateConfig(request PaymentSplitConfigRequest) (*PaymentSplitConfigResponse, error) {
	bodyBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := p.client.Request("POST", "/v2/gn/split/config", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create payment split config: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create payment split config with status %d: %s", resp.StatusCode, body)
	}

	var configResp PaymentSplitConfigResponse
	if err := json.NewDecoder(resp.Body).Decode(&configResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &configResp, nil
}


func (p *PaymentSplit) CreateConfigWithID(id string, request PaymentSplitConfigRequest) (*PaymentSplitConfigResponse, error) {
	bodyBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := p.client.Request("PUT", fmt.Sprintf("/v2/gn/split/config/%s", id), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create/update payment split config: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create/update payment split config with status %d: %s", resp.StatusCode, body)
	}

	var configResp PaymentSplitConfigResponse
	if err := json.NewDecoder(resp.Body).Decode(&configResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &configResp, nil
}


func (p *PaymentSplit) GetConfig(id string, revision int) (*PaymentSplitConfigResponse, error) {
	path := fmt.Sprintf("/v2/gn/split/config/%s", id)
	if revision > 0 {
		path = fmt.Sprintf("%s?revisao=%d", path, revision)
	}

	resp, err := p.client.Request("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment split config: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get payment split config with status %d: %s", resp.StatusCode, body)
	}

	var configResp PaymentSplitConfigResponse
	if err := json.NewDecoder(resp.Body).Decode(&configResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &configResp, nil
}


func (p *PaymentSplit) LinkImmediateCharge(txid, splitConfigID string) error {
	resp, err := p.client.Request("PUT", fmt.Sprintf("/v2/gn/split/cob/%s/vinculo/%s", txid, splitConfigID), nil)
	if err != nil {
		return fmt.Errorf("failed to link immediate charge to payment split: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to link immediate charge to payment split with status %d: %s", resp.StatusCode, body)
	}

	return nil
}


func (p *PaymentSplit) GetImmediateChargeWithSplit(txid string) (*SplitChargeResponse, error) {
	resp, err := p.client.Request("GET", fmt.Sprintf("/v2/gn/split/cob/%s", txid), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get immediate charge with split: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get immediate charge with split with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp SplitChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}


func (p *PaymentSplit) UnlinkImmediateCharge(txid string) error {
	resp, err := p.client.Request("DELETE", fmt.Sprintf("/v2/gn/split/cob/%s/vinculo", txid), nil)
	if err != nil {
		return fmt.Errorf("failed to unlink immediate charge from payment split: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to unlink immediate charge from payment split with status %d: %s", resp.StatusCode, body)
	}

	return nil
}


func (p *PaymentSplit) LinkDueCharge(txid, splitConfigID string) error {
	resp, err := p.client.Request("PUT", fmt.Sprintf("/v2/gn/split/cobv/%s/vinculo/%s", txid, splitConfigID), nil)
	if err != nil {
		return fmt.Errorf("failed to link due charge to payment split: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to link due charge to payment split with status %d: %s", resp.StatusCode, body)
	}

	return nil
}


func (p *PaymentSplit) GetDueChargeWithSplit(txid string) (*SplitChargeResponse, error) {
	resp, err := p.client.Request("GET", fmt.Sprintf("/v2/gn/split/cobv/%s", txid), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get due charge with split: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get due charge with split with status %d: %s", resp.StatusCode, body)
	}

	var chargeResp SplitChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chargeResp, nil
}


func (p *PaymentSplit) UnlinkDueCharge(txid string) error {
	resp, err := p.client.Request("DELETE", fmt.Sprintf("/v2/gn/split/cobv/%s/vinculo", txid), nil)
	if err != nil {
		return fmt.Errorf("failed to unlink due charge from payment split: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to unlink due charge from payment split with status %d: %s", resp.StatusCode, body)
	}

	return nil
}
