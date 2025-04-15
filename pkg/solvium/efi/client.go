package efi

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	ProductionBaseURL = "https://pix.api.efipay.com.br"
	SandboxBaseURL    = "https://pix-h.api.efipay.com.br"
)

type Environment string

const (
	Production Environment = "production"
	Sandbox    Environment = "sandbox"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	ExpiresAt   time.Time
}

type Client struct {
	ClientID         string
	ClientSecret     string
	Certificate      tls.Certificate
	Environment      Environment
	BaseURL          string
	Token            *Token
	HTTPClient       *http.Client
	immediateCharges *ImmediateCharges
	dueCharges       *DueCharges
	pixSend          *PixSend
	pixManagement    *PixManagement
	payloadLocation  *PayloadLocation
	batchDueCharges  *BatchDueCharges
	paymentSplit     *PaymentSplit
}

func NewClient(clientID, clientSecret string, certPath string, certPassword string, env Environment) (*Client, error) {
	cert, err := tls.LoadX509KeyPair(certPath, certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load certificate: %w", err)
	}

	baseURL := SandboxBaseURL
	if env == Production {
		baseURL = ProductionBaseURL
	}

	client := &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Certificate:  cert,
		Environment:  env,
		BaseURL:      baseURL,
		Token:        nil,
	}

	
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	client.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: time.Second * 30,
	}

	return client, nil
}


func NewClientFromP12(clientID, clientSecret string, p12Path string, p12Password string, env Environment) (*Client, error) {
	cert, err := LoadCertificateFromP12(p12Path, p12Password)
	if err != nil {
		return nil, fmt.Errorf("failed to load P12 certificate: %w", err)
	}

	baseURL := SandboxBaseURL
	if env == Production {
		baseURL = ProductionBaseURL
	}

	client := &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Certificate:  cert,
		Environment:  env,
		BaseURL:      baseURL,
		Token:        nil,
	}

	
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	client.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: time.Second * 30,
	}

	return client, nil
}

func (c *Client) IsTokenValid() bool {
	if c.Token == nil {
		return false
	}
	return time.Now().Before(c.Token.ExpiresAt)
}

func (c *Client) Authenticate() error {
	if c.IsTokenValid() {
		return nil
	}

	authHeader := base64.StdEncoding.EncodeToString([]byte(c.ClientID + ":" + c.ClientSecret))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token", c.BaseURL), strings.NewReader(`{"grant_type": "client_credentials"}`))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Basic "+authHeader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("authentication request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("authentication failed with status %d: %s", resp.StatusCode, body)
	}

	var token Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return fmt.Errorf("failed to decode token response: %w", err)
	}

	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	c.Token = &token

	return nil
}

func (c *Client) Request(method, path string, body io.Reader) (*http.Response, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, path), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("%s %s", c.Token.TokenType, c.Token.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	return c.HTTPClient.Do(req)
}


func (c *Client) ImmediateCharge() *ImmediateCharges {
	if c.immediateCharges == nil {
		c.immediateCharges = NewImmediateCharges(c)
	}
	return c.immediateCharges
}


func (c *Client) DueCharge() *DueCharges {
	if c.dueCharges == nil {
		c.dueCharges = NewDueCharges(c)
	}
	return c.dueCharges
}


func (c *Client) PixSend() *PixSend {
	if c.pixSend == nil {
		c.pixSend = NewPixSend(c)
	}
	return c.pixSend
}


func (c *Client) PixManagement() *PixManagement {
	if c.pixManagement == nil {
		c.pixManagement = NewPixManagement(c)
	}
	return c.pixManagement
}


func (c *Client) PayloadLocation() *PayloadLocation {
	if c.payloadLocation == nil {
		c.payloadLocation = NewPayloadLocation(c)
	}
	return c.payloadLocation
}


func (c *Client) BatchDueCharges() *BatchDueCharges {
	if c.batchDueCharges == nil {
		c.batchDueCharges = NewBatchDueCharges(c)
	}
	return c.batchDueCharges
}


func (c *Client) PaymentSplit() *PaymentSplit {
	if c.paymentSplit == nil {
		c.paymentSplit = NewPaymentSplit(c)
	}
	return c.paymentSplit
}


func (c *Client) VerifyStatus(id string, txType TransactionType) (*TransactionStatus, error) {
	status := &TransactionStatus{
		ID:   id,
		Type: txType,
	}

	var err error

	switch txType {
	case TransactionTypeCharge:
		err = c.verifyChargeStatus(status)
	case TransactionTypeDueCharge:
		err = c.verifyDueChargeStatus(status)
	case TransactionTypePixSend:
		err = c.verifyPixSendStatus(status)
	case TransactionTypeRefund:
		err = c.verifyRefundStatus(status)
	default:
		return nil, fmt.Errorf("unsupported transaction type: %s", txType)
	}

	if err != nil {
		return nil, err
	}

	
	status.setMessage()

	return status, nil
}
