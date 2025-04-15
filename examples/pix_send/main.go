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

	idEnvio := "YOUR_UNIQUE_ID"
	sendReq := efi.PixSendRequest{
		Valor: "9.99",
		Pagador: efi.PagadorSend{
			Chave:       "YOUR_PIX_KEY",
			InfoPagador: "Pix payment for services",
		},
		Favorecido: efi.Favorecido{
			Chave: "efipay@sejaefi.com.br",
		},
	}

	sendResp, err := client.PixSend().Send(idEnvio, sendReq)
	if err != nil {
		log.Fatalf("Failed to send Pix: %v", err)
	}

	fmt.Printf("Pix sent with ID: %s\n", sendResp.IDEnvio)
	fmt.Printf("E2E ID: %s\n", sendResp.E2EID)
	fmt.Printf("Status: %s\n", sendResp.Status)

	if sendResp.E2EID != "" {
		pixDetail, err := client.PixSend().GetByE2EID(sendResp.E2EID)
		if err != nil {
			log.Printf("Failed to get Pix details: %v", err)
		} else {
			fmt.Printf("Pix details - Status: %s, Value: %s\n",
				pixDetail.Status, pixDetail.Valor)
		}
	}

	if sendResp.IDEnvio != "" {
		pixDetail, err := client.PixSend().GetByIDEnvio(sendResp.IDEnvio)
		if err != nil {
			log.Printf("Failed to get Pix details: %v", err)
		} else {
			fmt.Printf("Pix details - Status: %s, Value: %s\n",
				pixDetail.Status, pixDetail.Valor)
		}
	}

	startDate := time.Now().Add(-24 * time.Hour)
	endDate := time.Now()
	options := &efi.ListSentOptions{
		Status: "REALIZADO",
	}

	sentList, err := client.PixSend().List(startDate, endDate, options)
	if err != nil {
		log.Printf("Failed to list sent Pix: %v", err)
	} else {
		fmt.Printf("Found %d sent Pix transactions\n", len(sentList.Pix))
		for i, pix := range sentList.Pix {
			fmt.Printf("%d. E2E ID: %s, Status: %s, Value: %s\n",
				i+1, pix.EndToEndID, pix.Status, pix.Valor)
		}
	}

	qrCodeReq := efi.DetailQRCodeRequest{
		PixCopiaECola: "YOUR_PIX_COPY_PASTE_CODE",
	}

	qrDetail, err := client.PixSend().DetailQRCode(qrCodeReq)
	if err != nil {
		log.Printf("Failed to get QR code details: %v", err)
	} else {
		fmt.Printf("QR Code details - TxID: %s, Status: %s, Value: %s\n",
			qrDetail.TxID, qrDetail.Status, qrDetail.Valor.Original)
	}

	payQRReq := efi.PayQRCodeRequest{
		Pagador: efi.PagadorSend{
			Chave:       "YOUR_PIX_KEY",
			InfoPagador: "Payment for QR code",
		},
		PixCopiaECola: "YOUR_PIX_COPY_PASTE_CODE",
	}

	qrPayResp, err := client.PixSend().PayQRCode("UNIQUE_ID_FOR_QR_PAYMENT", payQRReq)
	if err != nil {
		log.Printf("Failed to pay QR code: %v", err)
	} else {
		fmt.Printf("QR code payment - ID: %s, Status: %s, Value: %s\n",
			qrPayResp.IDEnvio, qrPayResp.Status, qrPayResp.Valor)
	}
}
