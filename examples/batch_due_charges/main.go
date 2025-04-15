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

	batchID := fmt.Sprintf("batch-%d", time.Now().Unix())

	batchRequest := efi.BatchDueChargesRequest{
		Descricao: "Cobranças dos alunos do turno vespertino",
		CobsV: []efi.CreateDueChargeRequest{
			{
				Calendario: efi.CalendarioDueCharge{
					DataDeVencimento:       "2024-12-31",
					ValidadeAposVencimento: 30,
				},
				Devedor: efi.DevedorDueCharge{
					CPF:  "08577095428",
					Nome: "João Souza",
				},
				Valor: efi.ValorDueCharge{
					Original: "100.00",
				},
				Chave:              "7c084cd4-54af-4172-a516-a7d1a12b75cc",
				SolicitacaoPagador: "Informar matrícula",
			},
			{
				Calendario: efi.CalendarioDueCharge{
					DataDeVencimento:       "2024-12-31",
					ValidadeAposVencimento: 30,
				},
				Devedor: efi.DevedorDueCharge{
					CPF:  "15311295449",
					Nome: "Manoel Silva",
				},
				Valor: efi.ValorDueCharge{
					Original: "100.00",
				},
				Chave:              "7c084cd4-54af-4172-a516-a7d1a12b75cc",
				SolicitacaoPagador: "Informar matrícula",
			},
		},
	}

	_, err = client.BatchDueCharges().CreateOrUpdate(batchID, batchRequest)
	if err != nil {
		log.Printf("Failed to create batch due charges: %v", err)
	} else {
		log.Printf("Batch due charges created with ID: %s", batchID)

		batchDetails, err := client.BatchDueCharges().GetByID(batchID)
		if err != nil {
			log.Printf("Failed to get batch due charges: %v", err)
		} else {
			fmt.Printf("Batch Description: %s, Created: %s\n",
				batchDetails.Descricao, batchDetails.Criacao)

			fmt.Printf("Charges in batch:\n")
			for i, charge := range batchDetails.CobsV {
				fmt.Printf("%d. TxID: %s, Status: %s\n", i+1, charge.TxID, charge.Status)
				if charge.Problema != nil {
					fmt.Printf("   Problem: %s - %s\n", charge.Problema.Title, charge.Problema.Detail)
				}
			}
		}

		reviewRequest := efi.BatchDueChargesReviewRequest{
			CobsV: []efi.BatchDueChargesReviewItem{
				{
					TxID: "fb2761260e554ad593c7226beb5cb650",
					Calendario: efi.CalendarioDueCharge{
						DataDeVencimento: "2025-01-15",
					},
					Valor: efi.Valor{
						Original: "110.00",
					},
				},
				{
					TxID: "7978c0c97ea847e78e8849634473c1f1",
					Calendario: efi.CalendarioDueCharge{
						DataDeVencimento: "2025-01-15",
					},
					Valor: efi.Valor{
						Original: "110.00",
					},
				},
			},
		}

		_, err = client.BatchDueCharges().ReviewBatch(batchID, reviewRequest)
		if err != nil {
			log.Printf("Failed to review batch due charges: %v", err)
		} else {
			log.Printf("Batch due charges reviewed successfully")
		}
	}

	startDate := time.Now().Add(-30 * 24 * time.Hour)
	endDate := time.Now()
	listOptions := &efi.ListBatchDueChargesOptions{
		PaginaAtual:    0,
		ItensPorPagina: 10,
	}

	batchList, err := client.BatchDueCharges().List(startDate, endDate, listOptions)
	if err != nil {
		log.Printf("Failed to list batch due charges: %v", err)
	} else {
		fmt.Printf("Found %d batches of due charges\n", len(batchList.Lotes))
		for i, batch := range batchList.Lotes {
			fmt.Printf("%d. Description: %s, Created: %s\n",
				i+1, batch.Descricao, batch.Criacao)
			fmt.Printf("   Contains %d charges\n", len(batch.CobsV))
		}
	}
}
