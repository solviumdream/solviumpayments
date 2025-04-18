package mercadopago

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("MERCADO_PAGO_ACCESS_TOKEN", "TEST-YOUR-ACCESS-TOKEN-HERE")
}

func TestCreatePayment(t *testing.T) {
	client := NewClientFromEnv(Sandbox)

	response, err := client.Payment().Create(PaymentRequest{
		ExternalReference: "test-payment-001",
		Items: []Item{
			{
				Title:     "Test Product",
				Quantity:  1,
				UnitPrice: 100.00,
			},
		},
		Payer: Payer{
			Name:    "Test",
			Surname: "User",
			Email:   "test.user@example.com",
			Identification: PaymentIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
		},
	})

	if err != nil {
		if mercadopagoErr, ok := err.(*ErrorResponse); ok {
			t.Errorf("MercadoPago error: %s (status: %d)", mercadopagoErr.Message, mercadopagoErr.Status)
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
		return
	}

	t.Logf("Payment created successfully. ID: %s, InitPoint: %s", response.ID, response.InitPoint)
}

func TestGetPayment(t *testing.T) {
	client := NewClientFromEnv(Sandbox)

	
	createResponse, err := client.Payment().Create(PaymentRequest{
		ExternalReference: "test-payment-get",
		Items: []Item{
			{
				Title:     "Test Product",
				Quantity:  1,
				UnitPrice: 100.00,
			},
		},
		Payer: Payer{
			Name:    "Test",
			Surname: "User",
			Email:   "test.user@example.com",
			Identification: PaymentIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
		},
	})

	if err != nil {
		t.Fatalf("Failed to create payment for test: %v", err)
	}

	
	response, err := client.Payment().Get(createResponse.ID)

	if err != nil {
		if mercadopagoErr, ok := err.(*ErrorResponse); ok {
			t.Errorf("MercadoPago error: %s (status: %d)", mercadopagoErr.Message, mercadopagoErr.Status)
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
		return
	}

	if response.ID != createResponse.ID {
		t.Errorf("Expected payment ID %s, got %s", createResponse.ID, response.ID)
	}

	t.Logf("Payment retrieved successfully. ID: %s", response.ID)
}

func TestSearchPayments(t *testing.T) {
	client := NewClientFromEnv(Sandbox)

	
	externalRef := "test-search-unique-ref"
	_, err := client.Payment().Create(PaymentRequest{
		ExternalReference: externalRef,
		Items: []Item{
			{
				Title:     "Test Search Product",
				Quantity:  1,
				UnitPrice: 100.00,
			},
		},
		Payer: Payer{
			Name:    "Test",
			Surname: "User",
			Email:   "test.user@example.com",
			Identification: PaymentIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
		},
	})

	if err != nil {
		t.Fatalf("Failed to create payment for search test: %v", err)
	}

	
	response, err := client.Payment().Search(PaymentSearchParams{
		"external_reference": externalRef,
	})

	if err != nil {
		if mercadopagoErr, ok := err.(*ErrorResponse); ok {
			t.Errorf("MercadoPago error: %s (status: %d)", mercadopagoErr.Message, mercadopagoErr.Status)
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
		return
	}

	if response.Paging.Total == 0 {
		t.Errorf("No payments found with external reference %s", externalRef)
		return
	}

	t.Logf("Found %d payments with external reference %s", response.Paging.Total, externalRef)
}
