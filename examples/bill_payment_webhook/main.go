package main

import (
	"fmt"
	"log"
	"os"

	"github.com/solviumdream/solviumpayments/pkg/solvium/efi"
)

func main() {
	clientID := os.Getenv("EFI_CLIENT_ID")
	clientSecret := os.Getenv("EFI_CLIENT_SECRET")
	certPath := os.Getenv("EFI_CERT_PATH")
	certPassword := os.Getenv("EFI_CERT_PASSWORD")

	if clientID == "" || clientSecret == "" || certPath == "" {
		log.Fatal("Missing required environment variables. Please set EFI_CLIENT_ID, EFI_CLIENT_SECRET, EFI_CERT_PATH")
	}

	client, err := efi.NewClientFromP12(clientID, clientSecret, certPath, certPassword, efi.Sandbox)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	
	webhookURL := "https://your-webhook-server.com/efi/payment-callback"

	
	fmt.Println("Creating payment webhook...")
	webhook, err := client.BillPaymentWebhook().Create(webhookURL)
	if err != nil {
		log.Fatalf("Failed to create webhook: %v", err)
	}
	fmt.Printf("Webhook created with URL: %s\n", webhook.URL)

	
	fmt.Println("\nListing payment webhooks for the last 30 days...")
	webhooks, err := client.BillPaymentWebhook().ListByDateRange(30)
	if err != nil {
		log.Fatalf("Failed to list webhooks: %v", err)
	}

	fmt.Printf("Date Range: %s to %s\n", webhooks.Parameters.Start, webhooks.Parameters.End)
	fmt.Printf("Total Webhooks: %d\n", webhooks.Parameters.Pagination.TotalItems)

	if len(webhooks.Webhooks) > 0 {
		fmt.Println("\nWebhooks:")
		for i, hook := range webhooks.Webhooks {
			fmt.Printf("  %d. URL: %s (Created: %s)\n", i+1, hook.URL, hook.CreatedAt)
		}
	} else {
		fmt.Println("No webhooks found.")
	}

	
	fmt.Println("\nDeleting payment webhook...")
	err = client.BillPaymentWebhook().Delete(webhookURL)
	if err != nil {
		log.Fatalf("Failed to delete webhook: %v", err)
	}
	fmt.Println("Webhook deleted successfully.")

	
	fmt.Println("\nExample of parsing a webhook callback:")
	callbackJSON := []byte(`{
		"identificador": "1013",
		"status": {
			"anterior": "CRIADO",
			"atual": "EM_PROCESSAMENTO"
		},
		"valor": "150.10",
		"horario": {
			"solicitacao": "2024-02-07T14:32:54.000Z"
		},
		"efiExtras": {
			"dataExecucao": "2024-02-07",
			"codigoBarras": "23797962400000213204150060000055503009010000",
			"linhaDigitavel": "23794150096000005550330090100006796240000021320"
		}
	}`)

	callback, err := efi.ParseBillPaymentWebhookCallback(callbackJSON)
	if err != nil {
		log.Fatalf("Failed to parse callback: %v", err)
	}

	fmt.Printf("Payment ID: %s\n", callback.Identifier)
	fmt.Printf("Status Change: %s -> %s\n", callback.Status.Previous, callback.Status.Current)
	fmt.Printf("Value: %s\n", callback.Value)
	fmt.Printf("Request Time: %s\n", callback.Timestamp.RequestTime)
	fmt.Printf("Execution Date: %s\n", callback.EfiExtras.ExecutionDate)
	fmt.Printf("Barcode: %s\n", callback.EfiExtras.Barcode)
	fmt.Printf("Line Code: %s\n", callback.EfiExtras.LineCode)
}
