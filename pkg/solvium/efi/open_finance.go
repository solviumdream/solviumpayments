package efi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type OpenFinance struct {
	client *Client
}

func NewOpenFinance(client *Client) *OpenFinance {
	return &OpenFinance{
		client: client,
	}
}

func (o *OpenFinance) ConfigureApplication(config *OpenFinanceConfig) (*OpenFinanceConfig, error) {
	payload, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal open finance config: %w", err)
	}

	resp, err := o.client.Request(http.MethodPut, "/v1/config", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to configure application: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to configure application with status %d: %s", resp.StatusCode, body)
	}

	var configResponse OpenFinanceConfig
	if err := json.NewDecoder(resp.Body).Decode(&configResponse); err != nil {
		return nil, fmt.Errorf("failed to decode configuration response: %w", err)
	}

	return &configResponse, nil
}

func (o *OpenFinance) GetApplicationSettings() (*OpenFinanceConfig, error) {
	resp, err := o.client.Request(http.MethodGet, "/v1/config", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get application settings: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get application settings with status %d: %s", resp.StatusCode, body)
	}

	var config OpenFinanceConfig
	if err := json.NewDecoder(resp.Body).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode application settings: %w", err)
	}

	return &config, nil
}

func (o *OpenFinance) ParseRedirectParams(redirectURL string) (*OpenFinanceRedirectParams, error) {
	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redirect URL: %w", err)
	}

	queryParams := parsedURL.Query()

	params := &OpenFinanceRedirectParams{
		PaymentIdentifier: queryParams.Get("identificadorPagamento"),
		Error:             queryParams.Get("erro"),
	}

	if params.PaymentIdentifier == "" {
		return nil, fmt.Errorf("missing payment identifier in redirect URL")
	}

	return params, nil
}

func (o *OpenFinance) EnableReceiveWithoutKey() error {
	
	payload := []byte(`{"receberSemChave": true}`)

	resp, err := o.client.Request(http.MethodPut, "/v2/gn/config", bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to enable receive without key: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to enable receive without key with status %d: %s", resp.StatusCode, body)
	}

	return nil
}

func (o *OpenFinance) GetParticipants(request *OpenFinanceParticipantRequest) (*OpenFinanceParticipantResponse, error) {
	query := url.Values{}

	if request != nil {
		if request.Organization {
			query.Add("organizacao", "true")
		}

		if request.Name != "" {
			query.Add("nome", request.Name)
		}

		if request.Modality != "" {
			query.Add("modalidade", request.Modality)
		}
	}

	path := "/v1/participantes"
	if len(query) > 0 {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

	resp, err := o.client.Request(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get participants: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get participants with status %d: %s", resp.StatusCode, body)
	}

	var participantResponse OpenFinanceParticipantResponse
	if err := json.NewDecoder(resp.Body).Decode(&participantResponse); err != nil {
		return nil, fmt.Errorf("failed to decode participant response: %w", err)
	}

	return &participantResponse, nil
}


func (o *OpenFinance) InitiatePayment(request *OpenFinancePaymentRequest) (*OpenFinancePaymentResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payment request: %w", err)
	}

	resp, err := o.client.Request(http.MethodPost, "/v1/pagamentos/pix", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to initiate payment: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to initiate payment with status %d: %s", resp.StatusCode, body)
	}

	var paymentResponse OpenFinancePaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to decode payment response: %w", err)
	}

	return &paymentResponse, nil
}

func (o *OpenFinance) ListPayments(startDate, endDate string, page, limit int) (*OpenFinancePaymentList, error) {
	query := url.Values{}
	query.Add("inicio", startDate)
	query.Add("fim", endDate)

	if page > 0 {
		query.Add("pagina", fmt.Sprintf("%d", page))
	}

	if limit > 0 {
		query.Add("quantidade", fmt.Sprintf("%d", limit))
	}

	path := fmt.Sprintf("/v1/pagamentos/pix?%s", query.Encode())

	resp, err := o.client.Request(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list payments: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list payments with status %d: %s", resp.StatusCode, body)
	}

	var paymentList OpenFinancePaymentList
	if err := json.NewDecoder(resp.Body).Decode(&paymentList); err != nil {
		return nil, fmt.Errorf("failed to decode payment list: %w", err)
	}

	return &paymentList, nil
}

func (o *OpenFinance) RefundPayment(paymentID string, value string) (*OpenFinanceRefundResponse, error) {
	request := &OpenFinanceRefundRequest{
		Value: value,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal refund request: %w", err)
	}

	path := fmt.Sprintf("/v1/pagamentos/pix/%s/devolver", paymentID)

	resp, err := o.client.Request(http.MethodPost, path, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to initiate refund: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to initiate refund with status %d: %s", resp.StatusCode, body)
	}

	var refundResponse OpenFinanceRefundResponse
	if err := json.NewDecoder(resp.Body).Decode(&refundResponse); err != nil {
		return nil, fmt.Errorf("failed to decode refund response: %w", err)
	}

	return &refundResponse, nil
}


func (o *OpenFinance) InitiateScheduledPayment(request *OpenFinanceScheduledPaymentRequest) (*OpenFinancePaymentResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal scheduled payment request: %w", err)
	}

	resp, err := o.client.Request(http.MethodPost, "/v1/pagamentos-agendados/pix", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to initiate scheduled payment: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to initiate scheduled payment with status %d: %s", resp.StatusCode, body)
	}

	var paymentResponse OpenFinancePaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to decode scheduled payment response: %w", err)
	}

	return &paymentResponse, nil
}

func (o *OpenFinance) ListScheduledPayments(startDate, endDate string, page, limit int) (*OpenFinanceScheduledPaymentList, error) {
	query := url.Values{}
	query.Add("inicio", startDate)
	query.Add("fim", endDate)

	if page > 0 {
		query.Add("pagina", fmt.Sprintf("%d", page))
	}

	if limit > 0 {
		query.Add("quantidade", fmt.Sprintf("%d", limit))
	}

	path := fmt.Sprintf("/v1/pagamentos-agendados/pix?%s", query.Encode())

	resp, err := o.client.Request(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list scheduled payments: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to list scheduled payments with status %d: %s", resp.StatusCode, body)
	}

	var paymentList OpenFinanceScheduledPaymentList
	if err := json.NewDecoder(resp.Body).Decode(&paymentList); err != nil {
		return nil, fmt.Errorf("failed to decode scheduled payment list: %w", err)
	}

	return &paymentList, nil
}

func (o *OpenFinance) CancelScheduledPayment(paymentID string) (*OpenFinanceScheduledCancellationResponse, error) {
	path := fmt.Sprintf("/v1/pagamentos-agendados/pix/%s/cancelar", paymentID)

	resp, err := o.client.Request(http.MethodPatch, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel scheduled payment: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to cancel scheduled payment with status %d: %s", resp.StatusCode, body)
	}

	var cancellationResponse OpenFinanceScheduledCancellationResponse
	if err := json.NewDecoder(resp.Body).Decode(&cancellationResponse); err != nil {
		return nil, fmt.Errorf("failed to decode scheduled payment cancellation response: %w", err)
	}

	return &cancellationResponse, nil
}

func (o *OpenFinance) RefundScheduledPayment(paymentID, endToEndID, value string) (*OpenFinanceRefundResponse, error) {
	request := &OpenFinanceScheduledRefundRequest{
		EndToEndID: endToEndID,
		Value:      value,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal scheduled payment refund request: %w", err)
	}

	path := fmt.Sprintf("/v1/pagamentos-agendados/pix/%s/devolver", paymentID)

	resp, err := o.client.Request(http.MethodPost, path, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to initiate scheduled payment refund: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to initiate scheduled payment refund with status %d: %s", resp.StatusCode, body)
	}

	var refundResponse OpenFinanceRefundResponse
	if err := json.NewDecoder(resp.Body).Decode(&refundResponse); err != nil {
		return nil, fmt.Errorf("failed to decode scheduled payment refund response: %w", err)
	}

	return &refundResponse, nil
}
