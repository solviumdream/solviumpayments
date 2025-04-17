package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

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

	

	
	

	
	

	
	

	
	

	
	demonstrateScheduledPayments(client)
}

func demonstratePayments(client *efi.Client) {
	
	fmt.Println("Initiating a payment with bank account...")

	paymentRequest := createSampleBankAccountPayment()

	paymentResponse, err := client.OpenFinance().InitiatePayment(paymentRequest)
	if err != nil {
		log.Fatalf("Failed to initiate payment: %v", err)
	}

	fmt.Printf("Payment initiated successfully!\n")
	fmt.Printf("Payment ID: %s\n", paymentResponse.PaymentID)
	fmt.Printf("Redirect URI: %s\n", paymentResponse.RedirectURI)

	
	fmt.Println("\nListing payments for the last 30 days...")

	
	endDate := time.Now().Format("2006-01-02")
	startDate := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	payments, err := client.OpenFinance().ListPayments(startDate, endDate, 1, 10)
	if err != nil {
		log.Fatalf("Failed to list payments: %v", err)
	}

	fmt.Printf("Found %d payments (total: %d)\n", len(payments.Payments), payments.Total)

	if len(payments.Payments) > 0 {
		fmt.Println("\nRecent payments:")
		for i, payment := range payments.Payments {
			fmt.Printf("%d. Payment ID: %s\n", i+1, payment.PaymentID)
			fmt.Printf("   End-to-End ID: %s\n", payment.EndToEndID)
			fmt.Printf("   Value: %s\n", payment.Value)
			fmt.Printf("   Status: %s\n", payment.Status)
			fmt.Printf("   Created: %s\n", payment.CreatedAt)
			if len(payment.Refunds) > 0 {
				fmt.Printf("   Refunds: %d\n", len(payment.Refunds))
			}
			fmt.Println()
		}

		
		if len(payments.Payments) > 0 {
			paymentToRefund := payments.Payments[0]

			
			if paymentToRefund.Status == efi.PaymentStatusAccepted {
				fmt.Printf("Refunding payment %s...\n", paymentToRefund.PaymentID)

				
				refundResponse, err := client.OpenFinance().RefundPayment(paymentToRefund.PaymentID, paymentToRefund.Value)
				if err != nil {
					fmt.Printf("Failed to refund payment: %v\n", err)
				} else {
					fmt.Printf("Refund initiated successfully!\n")
					fmt.Printf("Refund Status: %s\n", refundResponse.Status)
					fmt.Printf("Refund Value: %s\n", refundResponse.Value)
					fmt.Printf("End-to-End ID: %s\n", refundResponse.EndToEndID)
				}
			} else {
				fmt.Printf("Payment %s is not eligible for refund (status: %s)\n", paymentToRefund.PaymentID, paymentToRefund.Status)
			}
		}
	} else {
		fmt.Println("No payments found in the specified period.")
	}
}

func createSampleBankAccountPayment() *efi.OpenFinancePaymentRequest {
	return &efi.OpenFinancePaymentRequest{
		Payer: efi.OpenFinancePaymentPayer{
			ParticipantID: "9f4cd202-8f2b-11ec-b909-0242ac120002", 
			CPF:           "51687574219",                          
		},
		Recipient: efi.OpenFinanceRecipient{
			BankAccount: &efi.OpenFinanceBankAccount{
				Name:        "Lucas Silva",
				Document:    "17558266300",
				BankCode:    "09089356",
				Branch:      "0001",
				Account:     "654984",
				AccountType: efi.AccountTypePayment,
			},
		},
		Payment: efi.OpenFinancePaymentInfo{
			Value:         "9.99",
			PayerInfo:     "Churrasco",
			OwnID:         "6236574863254",
			TransactionID: "E00038166201907261559y6j6",
		},
	}
}

func createSamplePixKeyPayment() *efi.OpenFinancePaymentRequest {
	return &efi.OpenFinancePaymentRequest{
		Payer: efi.OpenFinancePaymentPayer{
			ParticipantID: "9f4cd202-8f2b-11ec-b909-0242ac120002", 
			CPF:           "51687574219",                          
		},
		Recipient: efi.OpenFinanceRecipient{
			PixKey: &efi.OpenFinancePixKey{
				KeyType: "EMAIL",
				Key:     "exemplo@email.com",
				Name:    "João da Silva",
			},
		},
		Payment: efi.OpenFinancePaymentInfo{
			Value:         "5.50",
			PayerInfo:     "Pagamento teste",
			OwnID:         "6236574863255",
			TransactionID: "E00038166201907261559y6j7",
		},
	}
}

func createSampleQRCodePayment() *efi.OpenFinancePaymentRequest {
	return &efi.OpenFinancePaymentRequest{
		Payer: efi.OpenFinancePaymentPayer{
			ParticipantID: "9f4cd202-8f2b-11ec-b909-0242ac120002", 
			CPF:           "51687574219",                          
		},
		Recipient: efi.OpenFinanceRecipient{
			QRCode: &efi.OpenFinanceQRCode{
				QRCode: "00020126580014br.gov.bcb.pix0136a629532e-7693-4846-b028-f142c345580252040000530398654041.005802BR5913Fulano de Tal6008BRASILIA62070503***63041D3D",
				Name:   "Maria das Dores",
			},
		},
		Payment: efi.OpenFinancePaymentInfo{
			Value:         "10.00",
			PayerInfo:     "QR Code teste",
			OwnID:         "6236574863256",
			TransactionID: "E00038166201907261559y6j8",
		},
	}
}

func listParticipants(client *efi.Client) {
	fmt.Println("Getting Open Finance participants...")

	
	fmt.Println("\nGetting all participants:")
	allParticipants, err := client.OpenFinance().GetParticipants(nil)
	if err != nil {
		log.Fatalf("Failed to get participants: %v", err)
	}

	fmt.Printf("Found %d participants\n", len(allParticipants.Participants))
	for i, p := range allParticipants.Participants {
		if i >= 3 {
			fmt.Printf("... and %d more participants\n", len(allParticipants.Participants)-3)
			break
		}
		fmt.Printf("- %s (%s)\n", p.Name, p.ID)
	}

	
	fmt.Println("\nGetting participants by name 'Efí':")
	namedParticipants, err := client.OpenFinance().GetParticipants(&efi.OpenFinanceParticipantRequest{
		Name: "Efí",
	})
	if err != nil {
		log.Fatalf("Failed to get participants by name: %v", err)
	}

	if len(namedParticipants.Participants) == 0 {
		fmt.Println("No participants found with the specified name.")
	} else {
		for _, p := range namedParticipants.Participants {
			fmt.Printf("ID: %s\n", p.ID)
			fmt.Printf("Name: %s\n", p.Name)
			fmt.Printf("Description: %s\n", p.Description)
			fmt.Printf("Portal: %s\n", p.Portal)
			fmt.Printf("Logo: %s\n", p.Logo)

			if len(p.Organizations) > 0 {
				fmt.Println("Organizations:")
				for _, org := range p.Organizations {
					fmt.Printf("  - %s (CNPJ: %s, Status: %s)\n", org.Name, org.CNPJ, org.Status)
				}
			}
		}
	}

	
	fmt.Println("\nGetting participants with scheduled payments modality:")
	modalityParticipants, err := client.OpenFinance().GetParticipants(&efi.OpenFinanceParticipantRequest{
		Modality: "agendado",
	})
	if err != nil {
		log.Fatalf("Failed to get participants by modality: %v", err)
	}

	fmt.Printf("Found %d participants with scheduled payments modality\n", len(modalityParticipants.Participants))
	for i, p := range modalityParticipants.Participants {
		if i >= 3 {
			fmt.Printf("... and %d more participants\n", len(modalityParticipants.Participants)-3)
			break
		}
		fmt.Printf("- %s (%s)\n", p.Name, p.ID)
	}
}

func configureApplication(client *efi.Client) {
	
	fmt.Println("Enabling receive without key for Open Finance...")
	err := client.OpenFinance().EnableReceiveWithoutKey()
	if err != nil {
		log.Fatalf("Failed to enable receive without key: %v", err)
	}
	fmt.Println("Receive without key enabled successfully.")

	
	fmt.Println("\nConfiguring Open Finance application...")
	config := &efi.OpenFinanceConfig{
		RedirectURL: "https://your-app.com/efi/redirect",
		WebhookURL:  "https://your-app.com/efi/webhook",
		WebhookSecurity: efi.OpenFinanceWebhookSecurity{
			Type: efi.WebhookSecurityTypeMTLS,
		},
		ProcessPayment:      efi.ProcessPaymentTypeSync,
		GenerateTxIdForInic: true,
	}

	configResponse, err := client.OpenFinance().ConfigureApplication(config)
	if err != nil {
		log.Fatalf("Failed to configure Open Finance application: %v", err)
	}

	fmt.Println("Configuration successful:")
	fmt.Printf("  Redirect URL: %s\n", configResponse.RedirectURL)
	fmt.Printf("  Webhook URL: %s\n", configResponse.WebhookURL)
	fmt.Printf("  Webhook Security Type: %s\n", configResponse.WebhookSecurity.Type)
	fmt.Printf("  Process Payment: %s\n", configResponse.ProcessPayment)
	fmt.Printf("  Generate TxId For Inic: %t\n", configResponse.GenerateTxIdForInic)

	
	fmt.Println("\nGetting current Open Finance application settings...")
	settings, err := client.OpenFinance().GetApplicationSettings()
	if err != nil {
		log.Fatalf("Failed to get Open Finance settings: %v", err)
	}

	fmt.Println("Current settings:")
	fmt.Printf("  Redirect URL: %s\n", settings.RedirectURL)
	fmt.Printf("  Webhook URL: %s\n", settings.WebhookURL)
	fmt.Printf("  Webhook Security Type: %s\n", settings.WebhookSecurity.Type)
	fmt.Printf("  Process Payment: %s\n", settings.ProcessPayment)
	fmt.Printf("  Generate TxId For Inic: %t\n", settings.GenerateTxIdForInic)
}

func parseRedirectParams() {
	
	fmt.Println("\nExample of parsing redirect parameters...")
	
	redirectURL := "https://your-app.com/efi/redirect?identificadorPagamento=urn:bancoabc:1234567677"

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	queryParams := parsedURL.Query()
	paymentIdentifier := queryParams.Get("identificadorPagamento")
	errorCode := queryParams.Get("erro")

	fmt.Printf("Payment Identifier: %s\n", paymentIdentifier)
	if errorCode != "" {
		fmt.Printf("Error: %s\n", errorCode)
	} else {
		fmt.Println("No errors in redirect parameters.")
	}

	
	fmt.Println("\nExample of parsing redirect parameters with error...")
	redirectURLWithError := "https://your-app.com/efi/redirect?identificadorPagamento=urn:bancoabc:1234567677&erro=acesso_negado"

	parsedURLWithError, err := url.Parse(redirectURLWithError)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	queryParamsWithError := parsedURLWithError.Query()
	paymentIdentifierWithError := queryParamsWithError.Get("identificadorPagamento")
	errorCodeWithError := queryParamsWithError.Get("erro")

	fmt.Printf("Payment Identifier: %s\n", paymentIdentifierWithError)
	if errorCodeWithError != "" {
		fmt.Printf("Error: %s\n", errorCodeWithError)
	} else {
		fmt.Println("No errors in redirect parameters.")
	}
}

func demonstrateScheduledPayments(client *efi.Client) {
	
	fmt.Println("Initiating a scheduled payment with bank account...")

	
	scheduledDate := time.Now().AddDate(0, 0, 30).Format("2006-01-02")
	paymentRequest := createSampleScheduledBankAccountPayment(scheduledDate)

	paymentResponse, err := client.OpenFinance().InitiateScheduledPayment(paymentRequest)
	if err != nil {
		log.Fatalf("Failed to initiate scheduled payment: %v", err)
	}

	fmt.Printf("Scheduled payment initiated successfully!\n")
	fmt.Printf("Payment ID: %s\n", paymentResponse.PaymentID)
	fmt.Printf("Redirect URI: %s\n", paymentResponse.RedirectURI)

	
	fmt.Println("\nListing scheduled payments for the next 90 days...")

	
	startDate := time.Now().Format("2006-01-02")
	endDate := time.Now().AddDate(0, 0, 90).Format("2006-01-02")

	payments, err := client.OpenFinance().ListScheduledPayments(startDate, endDate, 1, 10)
	if err != nil {
		log.Fatalf("Failed to list scheduled payments: %v", err)
	}

	fmt.Printf("Found %d scheduled payments (total: %d)\n", len(payments.Payments), payments.Total)

	if len(payments.Payments) > 0 {
		fmt.Println("\nScheduled payments:")
		for i, payment := range payments.Payments {
			fmt.Printf("%d. Payment ID: %s\n", i+1, payment.PaymentID)
			fmt.Printf("   End-to-End ID: %s\n", payment.EndToEndID)
			fmt.Printf("   Value: %s\n", payment.Value)
			fmt.Printf("   Status: %s\n", payment.Status)
			fmt.Printf("   Operation Date: %s\n", payment.OperationDate)
			fmt.Printf("   Created: %s\n", payment.CreatedAt)
			if len(payment.Refunds) > 0 {
				fmt.Printf("   Refunds: %d\n", len(payment.Refunds))
			}
			fmt.Println()
		}

		
		if len(payments.Payments) > 0 {
			paymentToCancel := payments.Payments[0]

			fmt.Printf("Cancelling scheduled payment %s...\n", paymentToCancel.PaymentID)

			cancellationResponse, err := client.OpenFinance().CancelScheduledPayment(paymentToCancel.PaymentID)
			if err != nil {
				fmt.Printf("Failed to cancel scheduled payment: %v\n", err)
			} else {
				fmt.Printf("Scheduled payment cancelled successfully!\n")
				fmt.Printf("Status: %s\n", cancellationResponse.Status)
				fmt.Printf("Cancellation Date: %s\n", cancellationResponse.CancellationDate)
			}

			
			
			if paymentToCancel.Status == efi.PaymentStatusAccepted {
				fmt.Printf("\nRefunding scheduled payment %s...\n", paymentToCancel.PaymentID)

				refundResponse, err := client.OpenFinance().RefundScheduledPayment(
					paymentToCancel.PaymentID,
					paymentToCancel.EndToEndID,
					paymentToCancel.Value,
				)
				if err != nil {
					fmt.Printf("Failed to refund scheduled payment: %v\n", err)
				} else {
					fmt.Printf("Refund initiated successfully!\n")
					fmt.Printf("Refund Status: %s\n", refundResponse.Status)
					fmt.Printf("Refund Value: %s\n", refundResponse.Value)
					fmt.Printf("End-to-End ID: %s\n", refundResponse.EndToEndID)
				}
			} else {
				fmt.Printf("\nScheduled payment %s is not eligible for refund (status: %s)\n",
					paymentToCancel.PaymentID, paymentToCancel.Status)
			}
		}
	} else {
		fmt.Println("No scheduled payments found in the specified period.")
	}
}

func createSampleScheduledBankAccountPayment(scheduledDate string) *efi.OpenFinanceScheduledPaymentRequest {
	return &efi.OpenFinanceScheduledPaymentRequest{
		Payer: efi.OpenFinancePaymentPayer{
			ParticipantID: "9f4cd202-8f2b-11ec-b909-0242ac120002", 
			CPF:           "45204392050",                          
		},
		Recipient: efi.OpenFinanceRecipient{
			BankAccount: &efi.OpenFinanceBankAccount{
				Name:        "Lucas Silva",
				Document:    "17558266300",
				BankCode:    "09089356",
				Branch:      "0001",
				Account:     "654984",
				AccountType: efi.AccountTypePayment,
			},
		},
		Payment: efi.OpenFinanceScheduledPaymentInfo{
			Value:         "9.99",
			PayerInfo:     "Churrasco",
			OwnID:         "6236574863254",
			ScheduledDate: scheduledDate,
			TransactionID: "E00038166201907261559y6j6",
		},
	}
}

func createSampleScheduledPixKeyPayment(scheduledDate string) *efi.OpenFinanceScheduledPaymentRequest {
	return &efi.OpenFinanceScheduledPaymentRequest{
		Payer: efi.OpenFinancePaymentPayer{
			ParticipantID: "9f4cd202-8f2b-11ec-b909-0242ac120002", 
			CPF:           "45204392050",                          
		},
		Recipient: efi.OpenFinanceRecipient{
			PixKey: &efi.OpenFinancePixKey{
				KeyType: "EMAIL",
				Key:     "exemplo@email.com",
				Name:    "João da Silva",
			},
		},
		Payment: efi.OpenFinanceScheduledPaymentInfo{
			Value:         "5.50",
			PayerInfo:     "Pagamento teste",
			OwnID:         "6236574863255",
			ScheduledDate: scheduledDate,
			TransactionID: "E00038166201907261559y6j7",
		},
	}
}

func createSampleScheduledQRCodePayment(scheduledDate string) *efi.OpenFinanceScheduledPaymentRequest {
	return &efi.OpenFinanceScheduledPaymentRequest{
		Payer: efi.OpenFinancePaymentPayer{
			ParticipantID: "9f4cd202-8f2b-11ec-b909-0242ac120002", 
			CPF:           "45204392050",                          
		},
		Recipient: efi.OpenFinanceRecipient{
			QRCode: &efi.OpenFinanceQRCode{
				QRCode: "00020126580014br.gov.bcb.pix0136a629532e-7693-4846-b028-f142c345580252040000530398654041.005802BR5913Fulano de Tal6008BRASILIA62070503***63041D3D",
				Name:   "Maria das Dores",
			},
		},
		Payment: efi.OpenFinanceScheduledPaymentInfo{
			Value:         "10.00",
			PayerInfo:     "QR Code teste",
			OwnID:         "6236574863256",
			ScheduledDate: scheduledDate,
			TransactionID: "E00038166201907261559y6j8",
		},
	}
}
