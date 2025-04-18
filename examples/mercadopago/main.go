package main

import (
	"fmt"
	"log"
	"os"

	"github.com/solviumdream/solviumpayments/pkg/solvium/mercadopago"
)

func main() {
	
	os.Setenv("MERCADO_PAGO_ACCESS_TOKEN", "TEST-YOUR-ACCESS-TOKEN-HERE")

	
	client := mercadopago.NewClientFromEnv(mercadopago.Sandbox)

	
	paymentResponse, err := createPayment(client)
	if err != nil {
		log.Fatalf("Failed to create payment: %v", err)
	}
	fmt.Printf("Payment created with ID: %s\n", paymentResponse.ID)
	fmt.Printf("Payment URL: %s\n", paymentResponse.InitPoint)

	
	getPayment(client, paymentResponse.ID)

	
	getIdentificationTypes(client)

	
	getPaymentMethods(client)
}

func createPayment(client *mercadopago.Client) (*mercadopago.PaymentResponse, error) {
	paymentRequest := mercadopago.PaymentRequest{
		ExternalReference: "test-payment-001",
		Items: []mercadopago.Item{
			{
				Title:     "Test Product",
				Quantity:  1,
				UnitPrice: 100.00,
			},
		},
		Payer: mercadopago.Payer{
			Name:    "Test",
			Surname: "User",
			Email:   "test.user@example.com",
			Identification: mercadopago.PaymentIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
		},
		BackURLs: &mercadopago.BackURLs{
			Success: "https://www.example.com/success",
			Pending: "https://www.example.com/pending",
			Failure: "https://www.example.com/failure",
		},
		NotificationURL: "https://www.example.com/webhook",
	}

	return client.Payment().Create(paymentRequest)
}

func getPayment(client *mercadopago.Client, paymentID string) {
	payment, err := client.Payment().Get(paymentID)
	if err != nil {
		fmt.Printf("Error getting payment: %v\n", err)
		return
	}

	fmt.Printf("Payment details - ID: %s, InitPoint: %s\n", payment.ID, payment.InitPoint)
}

func getIdentificationTypes(client *mercadopago.Client) {
	types, err := client.Identification().GetTypes()
	if err != nil {
		fmt.Printf("Error getting identification types: %v\n", err)
		return
	}

	fmt.Println("Available identification types:")
	for _, idType := range types {
		fmt.Printf("- %s (%s)\n", idType.Name, idType.ID)
	}
}

func getPaymentMethods(client *mercadopago.Client) {
	methods, err := client.PaymentMethods().GetAll()
	if err != nil {
		fmt.Printf("Error getting payment methods: %v\n", err)
		return
	}

	fmt.Println("Available payment methods:")
	for _, method := range methods {
		fmt.Printf("- %s (%s)\n", method.Name, method.ID)
	}
}
