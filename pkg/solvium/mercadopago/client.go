package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	ProductionBaseURL = "https://api.mercadopago.com"
	SandboxBaseURL    = "https://api.mercadopago.com"
)

type Environment string

const (
	Production Environment = "production"
	Sandbox    Environment = "sandbox"
)

type Client struct {
	AccessToken    string
	Environment    Environment
	BaseURL        string
	HTTPClient     *http.Client
	payment        *Payment
	paymentMethods *PaymentMethods
	identification *Identification
}

func NewClient(accessToken string, env Environment) *Client {
	baseURL := SandboxBaseURL
	if env == Production {
		baseURL = ProductionBaseURL
	}

	client := &Client{
		AccessToken: accessToken,
		Environment: env,
		BaseURL:     baseURL,
		HTTPClient:  &http.Client{Timeout: time.Second * 30},
	}

	return client
}

func NewClientFromEnv(env Environment) *Client {
	accessToken := os.Getenv("MERCADO_PAGO_ACCESS_TOKEN")
	return NewClient(accessToken, env)
}

func (c *Client) Request(method, path string, body interface{}, queryParams map[string]string) (*http.Response, error) {
	var bodyReader io.Reader

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	if queryParams != nil {
		q := req.URL.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	return c.HTTPClient.Do(req)
}

func (c *Client) Payment() *Payment {
	if c.payment == nil {
		c.payment = NewPayment(c)
	}
	return c.payment
}

func (c *Client) PaymentMethods() *PaymentMethods {
	if c.paymentMethods == nil {
		c.paymentMethods = NewPaymentMethods(c)
	}
	return c.paymentMethods
}

func (c *Client) Identification() *Identification {
	if c.identification == nil {
		c.identification = NewIdentification(c)
	}
	return c.identification
}

func parseResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errorResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return fmt.Errorf("request failed with status %d and couldn't decode error", resp.StatusCode)
		}
		return &errorResp
	}

	if result == nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
