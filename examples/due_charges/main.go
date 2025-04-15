package main

import (
	"fmt"
	"log"
	"time"

	"github.com/solviumdream/solviumpayments/pkg/solvium/efi"
)


func RunExample() {
	
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

	
	txid := "YOUR_TXID" 

	chargeReq := efi.CreateDueChargeRequest{
		Calendario: efi.CalendarioDueCharge{
			DataDeVencimento:       "2023-12-31",
			ValidadeAposVencimento: 30,
		},
		Devedor: efi.DevedorDueCharge{
			CPF:  "12345678909",
			Nome: "Francisco da Silva",
			EnderecoDevedor: efi.EnderecoDevedor{
				Logradouro: "Alameda Souza, Numero 80, Bairro Braz",
				Cidade:     "Recife",
				UF:         "PE",
				CEP:        "70011750",
			},
		},
		Valor: efi.ValorDueCharge{
			Original: "123.45",
			Multa: efi.Multa{
				Modalidade: 2,
				ValorPerc:  "15.00",
			},
			Juros: efi.Juros{
				Modalidade: 2,
				ValorPerc:  "2.00",
			},
			Desconto: efi.Desconto{
				Modalidade: 1,
				DescontoDataFixa: []efi.DescontoDataFixa{
					{
						Data:      "2023-12-15",
						ValorPerc: "30.00",
					},
				},
			},
		},
		Chave:              "YOUR_PIX_KEY",
		SolicitacaoPagador: "Cobrança dos serviços prestados.",
	}

	charge, err := client.DueCharge().Create(txid, chargeReq)
	if err != nil {
		log.Fatalf("Failed to create due charge: %v", err)
	}

	fmt.Printf("Due charge created with txid: %s\n", charge.TxID)
	fmt.Printf("Copy and paste code: %s\n", charge.PixCopiaECola)

	
	fetchedCharge, err := client.DueCharge().Get(txid, 0)
	if err != nil {
		log.Fatalf("Failed to get due charge: %v", err)
	}

	fmt.Printf("Due charge status: %s\n", fetchedCharge.Status)

	
	reviewReq := efi.ReviewDueChargeRequest{
		Valor: efi.ValorDueCharge{
			Original: "150.00",
		},
		SolicitacaoPagador: "Valor atualizado da cobrança.",
	}

	updatedCharge, err := client.DueCharge().Review(txid, reviewReq)
	if err != nil {
		log.Fatalf("Failed to review due charge: %v", err)
	}

	fmt.Printf("Updated due charge value: %s\n", updatedCharge.Valor.Original)

	
	startDate := time.Now().Add(-30 * 24 * time.Hour) 
	endDate := time.Now().Add(30 * 24 * time.Hour)    
	options := &efi.ListDueChargesOptions{
		Status:         "ATIVA",
		PaginaAtual:    0,
		ItensPorPagina: 10,
	}

	charges, err := client.DueCharge().List(startDate, endDate, options)
	if err != nil {
		log.Fatalf("Failed to list due charges: %v", err)
	}

	fmt.Printf("Found %d due charges\n", len(charges.Cobs))
	for i, cob := range charges.Cobs {
		fmt.Printf("%d. TxID: %s, Status: %s, Value: %s, Due Date: %s\n",
			i+1, cob.TxID, cob.Status, cob.Valor.Original, cob.Calendario.DataDeVencimento)
	}
}
