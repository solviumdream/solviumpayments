package main

import (
	"fmt"
	"log"
	"time"

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

	e2eID := "E12345678901234567890123456789012"
	receivedPix, err := client.PixManagement().GetByE2EID(e2eID)
	if err != nil {
		log.Printf("Failed to get received Pix: %v", err)
	} else {
		fmt.Printf("Received Pix - E2E ID: %s, TxID: %s, Value: %s\n",
			receivedPix.EndToEndID, receivedPix.TxID, receivedPix.Valor)

		if len(receivedPix.Devolucoes) > 0 {
			fmt.Printf("This Pix has %d refunds\n", len(receivedPix.Devolucoes))
			for i, refund := range receivedPix.Devolucoes {
				fmt.Printf("  %d. Refund ID: %s, Status: %s, Value: %s\n",
					i+1, refund.ID, refund.Status, refund.Valor)
			}
		}
	}

	startDate := time.Now().Add(-30 * 24 * time.Hour)
	endDate := time.Now()
	listOptions := &efi.ListReceivedOptions{
		PaginaAtual: 0,
		ItensPagina: 10,
	}

	pixList, err := client.PixManagement().ListReceived(startDate, endDate, listOptions)
	if err != nil {
		log.Printf("Failed to list received Pix: %v", err)
	} else {
		fmt.Printf("Found %d received Pix transactions\n", len(pixList.Pix))
		for i, pix := range pixList.Pix {
			fmt.Printf("%d. E2E ID: %s, TxID: %s, Value: %s\n",
				i+1, pix.EndToEndID, pix.TxID, pix.Valor)
		}
	}

	if receivedPix != nil && receivedPix.EndToEndID != "" {
		refundID := "YOUR_REFUND_ID"
		refundReq := efi.RefundRequest{
			Valor: "1.00",
		}

		refundResp, err := client.PixManagement().RequestRefund(receivedPix.EndToEndID, refundID, refundReq)
		if err != nil {
			log.Printf("Failed to request refund: %v", err)
		} else {
			fmt.Printf("Refund requested - ID: %s, Status: %s, Value: %s\n",
				refundResp.ID, refundResp.Status, refundResp.Valor)
		}
	}

	refundIDToCheck := "123456"
	e2eIDForRefund := "E12345678901234567890123456789012"

	refundDetails, err := client.PixManagement().GetRefund(e2eIDForRefund, refundIDToCheck)
	if err != nil {
		log.Printf("Failed to get refund details: %v", err)
	} else {
		fmt.Printf("Refund details - ID: %s, Status: %s, Value: %s\n",
			refundDetails.ID, refundDetails.Status, refundDetails.Valor)

		if refundDetails.Horario.Solicitacao != "" {
			fmt.Printf("Requested at: %s\n", refundDetails.Horario.Solicitacao)
		}

		if refundDetails.Horario.Liquidacao != "" {
			fmt.Printf("Settled at: %s\n", refundDetails.Horario.Liquidacao)
		}
	}

	refundReq := efi.RefundRequest{
		Valor: "0.01",
	}

	testRefundResp, err := client.PixManagement().RequestRefund(e2eIDForRefund, "test_refund_id", refundReq)
	if err != nil {
		log.Printf("Failed to request test refund: %v", err)
	} else {
		fmt.Printf("Test refund requested - ID: %s, Status: %s\n",
			testRefundResp.ID, testRefundResp.Status)
		fmt.Println("This refund should be rejected in the sandbox environment")
	}
}
