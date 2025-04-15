package main

import (
	"fmt"
	"log"

	"github.com/solviumdream/solviumpayments/pkg/solvium/efi"
)

func main() {
	
	client, err := efi.NewClientFromP12(
		"YOUR_CLIENT_ID",
		"YOUR_CLIENT_SECRET",
		"path/to/certificate.p12",
		"",          
		efi.Sandbox, 
	)
	if err != nil {
		log.Fatalf("Failed to create Efi client: %v", err)
	}

	
	chargeTxID := "your-charge-txid"
	chargeStatus, err := client.VerifyStatus(chargeTxID, efi.TransactionTypeCharge)
	if err != nil {
		log.Printf("Failed to verify charge status: %v", err)
	} else {
		fmt.Printf("Charge Status: %s\n", chargeStatus.Status)
		fmt.Printf("Message: %s\n", chargeStatus.Message)
		fmt.Printf("Is Completed: %t\n", chargeStatus.IsCompleted)
		fmt.Printf("Is Failed: %t\n\n", chargeStatus.IsFailed)
	}

	
	pixSendID := "your-pix-send-id" 
	sendStatus, err := client.VerifyStatus(pixSendID, efi.TransactionTypePixSend)
	if err != nil {
		log.Printf("Failed to verify Pix send status: %v", err)
	} else {
		fmt.Printf("Pix Send Status: %s\n", sendStatus.Status)
		fmt.Printf("Message: %s\n", sendStatus.Message)
		fmt.Printf("Is Completed: %t\n", sendStatus.IsCompleted)
		fmt.Printf("Is Failed: %t\n\n", sendStatus.IsFailed)
	}

	
	
	e2eID := "E12345678901234567890123456789012"
	refundID := "refund-123"
	combinedID := fmt.Sprintf("%s:%s", e2eID, refundID)

	refundStatus, err := client.VerifyStatus(combinedID, efi.TransactionTypeRefund)
	if err != nil {
		log.Printf("Failed to verify refund status: %v", err)
	} else {
		fmt.Printf("Refund Status: %s\n", refundStatus.Status)
		fmt.Printf("Message: %s\n", refundStatus.Message)
		fmt.Printf("Is Completed: %t\n", refundStatus.IsCompleted)
		fmt.Printf("Is Failed: %t\n\n", refundStatus.IsFailed)
	}

	
	dueTxID := "your-due-charge-txid"
	dueStatus, err := client.VerifyStatus(dueTxID, efi.TransactionTypeDueCharge)
	if err != nil {
		log.Printf("Failed to verify due charge status: %v", err)
	} else {
		if dueStatus.IsCompleted {
			fmt.Println("Due charge has been paid!")
			
		} else if dueStatus.IsFailed {
			fmt.Println("Due charge failed or was removed")
			
		} else {
			fmt.Println("Due charge is still awaiting payment")
			
		}
	}

	
	fmt.Println("Example of polling for status (simplified):")
	fmt.Println("In a real application, you would implement this with proper intervals and timeouts")

	
	for i := 0; i < 3; i++ {
		fmt.Printf("Polling attempt %d...\n", i+1)

		status, err := client.VerifyStatus(chargeTxID, efi.TransactionTypeCharge)
		if err != nil {
			log.Printf("Error checking status: %v", err)
			continue
		}

		fmt.Printf("Current status: %s\n", status.Status)

		if status.IsCompleted {
			fmt.Println("Transaction completed successfully!")
			break
		} else if status.IsFailed {
			fmt.Println("Transaction failed!")
			break
		} else {
			fmt.Println("Transaction still in progress, waiting...")
			
		}
	}
}
