package main

import (
	"fmt"
	"log"
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

	
	barcode := "83660000001266700201006504401576111983794981768"

	
	fmt.Println("Detailing barcode...")
	details, err := client.BillPayment().DetailBarcode(barcode)
	if err != nil {
		log.Fatalf("Failed to detail barcode: %v", err)
	}

	fmt.Printf("Barcode Type: %s\n", details.Type)
	fmt.Printf("Value: %.2f\n", details.Value)
	if details.DueDate != "" {
		fmt.Printf("Due Date: %s\n", details.DueDate)
	}
	if details.Beneficiary != "" {
		fmt.Printf("Beneficiary: %s\n", details.Beneficiary)
	}

	
	fmt.Println("\nRequesting payment...")
	paymentRequest := &efi.BillPaymentRequest{
		Value:       details.Value,
		PaymentDate: time.Now().Format("2006-01-02"),
		Description: "Example bill payment",
	}

	payment, err := client.BillPayment().RequestPayment(barcode, paymentRequest)
	if err != nil {
		log.Fatalf("Failed to request payment: %v", err)
	}

	fmt.Printf("Payment ID: %s\n", payment.PaymentID)
	fmt.Printf("Amount Paid: %.2f\n", payment.AmountPaid)
	fmt.Printf("Status: %s\n", payment.Status)
	fmt.Printf("Request Date: %s\n", payment.Data.RequestDate)
	if payment.Data.PaymentDate != "" {
		fmt.Printf("Payment Date: %s\n", payment.Data.PaymentDate)
	}

	
	fmt.Println("\nChecking payment status...")
	paymentStatus, err := client.BillPayment().GetPaymentStatus(payment.PaymentID)
	if err != nil {
		log.Fatalf("Failed to get payment status: %v", err)
	}

	fmt.Printf("Payment ID: %s\n", paymentStatus.PaymentID)
	fmt.Printf("Status: %s\n", paymentStatus.Status)
	if paymentStatus.RejectReason != "" {
		fmt.Printf("Reject Reason: %s\n", paymentStatus.RejectReason)
	}

	
	fmt.Println("\nGetting payment summary for the last 30 days...")
	summary, err := client.BillPayment().GetPaymentSummaryByDateRange(30)
	if err != nil {
		log.Fatalf("Failed to get payment summary: %v", err)
	}

	fmt.Printf("Date Range: %s to %s\n", summary.Dates.StartDate, summary.Dates.EndDate)
	fmt.Printf("Total Requests: %d\n", summary.Requests.Total)
	fmt.Printf("Processing: %d\n", summary.Requests.Processing)
	fmt.Printf("Success: %d\n", summary.Requests.Success)
	fmt.Printf("Failed: %d\n", summary.Requests.Failed)
	fmt.Printf("Canceled: %d\n", summary.Requests.Canceled)
}
