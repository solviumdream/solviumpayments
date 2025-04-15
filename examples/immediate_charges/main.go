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

	
	chargeReq := efi.CreateImmediateChargeRequest{
		Calendario: efi.Calendario{
			Expiracao: 3600, 
		},
		Devedor: efi.Devedor{
			CPF:  "12345678909",
			Nome: "Francisco da Silva",
		},
		Valor: efi.Valor{
			Original: "9.99", 
		},
		Chave:              "YOUR_PIX_KEY",
		SolicitacaoPagador: "Cobrança dos serviços prestados.",
	}

	charge, err := client.ImmediateCharge().CreateWithoutTxid(chargeReq)
	if err != nil {
		log.Fatalf("Failed to create charge: %v", err)
	}

	fmt.Printf("Charge created with txid: %s\n", charge.TxID)
	fmt.Printf("Copy and paste code: %s\n", charge.PixCopiaECola)

	
	txid := charge.TxID
	fetchedCharge, err := client.ImmediateCharge().GetCharge(txid, 0)
	if err != nil {
		log.Fatalf("Failed to get charge: %v", err)
	}

	fmt.Printf("Charge status: %s\n", fetchedCharge.Status)

	
	startDate := time.Now().Add(-24 * time.Hour)
	endDate := time.Now()
	options := &efi.ListChargesOptions{
		Status:         "ATIVA",
		PaginaAtual:    0,
		ItensPorPagina: 10,
	}

	charges, err := client.ImmediateCharge().ListCharges(startDate, endDate, options)
	if err != nil {
		log.Fatalf("Failed to list charges: %v", err)
	}

	fmt.Printf("Found %d charges\n", len(charges.Cobs))
	for i, cob := range charges.Cobs {
		fmt.Printf("%d. TxID: %s, Status: %s, Value: %s\n", i+1, cob.TxID, cob.Status, cob.Valor.Original)
	}

	
	reviewReq := efi.ReviewChargeRequest{
		Valor: efi.Valor{
			Original: "5.99",
		},
		SolicitacaoPagador: "Valor atualizado da cobrança.",
	}

	updatedCharge, err := client.ImmediateCharge().ReviewCharge(txid, reviewReq)
	if err != nil {
		log.Fatalf("Failed to review charge: %v", err)
	}

	fmt.Printf("Updated charge value: %s\n", updatedCharge.Valor.Original)
}
