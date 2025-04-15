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

	splitConfig := efi.PaymentSplitConfigRequest{
		Descricao: "Batatinha frita 1, 2, 3",
		Lancamento: efi.SplitLancamento{
			Imediato: true,
		},
		Split: efi.SplitConfig{
			DivisaoTarifa: "assumir_total",
			MinhaParte: efi.SplitMinhaParte{
				Tipo:  "porcentagem",
				Valor: "60.00",
			},
			Repasses: []efi.SplitRepasse{
				{
					Tipo:  "porcentagem",
					Valor: "15.00",
					Favorecido: efi.SplitFavorecido{
						CPF:   "12345678909",
						Conta: "1234567",
					},
				},
				{
					Tipo:  "porcentagem",
					Valor: "25.00",
					Favorecido: efi.SplitFavorecido{
						CPF:   "94271564656",
						Conta: "7654321",
					},
				},
			},
		},
	}

	splitConfigResp, err := client.PaymentSplit().CreateConfig(splitConfig)
	if err != nil {
		log.Printf("Failed to create payment split config: %v", err)
	} else {
		fmt.Printf("Split config created with ID: %s, Status: %s\n",
			splitConfigResp.ID, splitConfigResp.Status)

		chargeReq := efi.CreateImmediateChargeRequest{
			Calendario: efi.Calendario{
				Expiracao: 3600,
			},
			Devedor: efi.Devedor{
				CPF:  "12345678909",
				Nome: "Francisco da Silva",
			},
			Valor: efi.Valor{
				Original: "100.00",
			},
			Chave:              "YOUR_PIX_KEY",
			SolicitacaoPagador: "Cobrança com split de pagamento",
		}

		txid := fmt.Sprintf("txid-%d", time.Now().Unix())

		charge, err := client.ImmediateCharge().CreateWithTxid(txid, chargeReq)
		if err != nil {
			log.Printf("Failed to create immediate charge: %v", err)
		} else {
			fmt.Printf("Charge created with txid: %s\n", charge.TxID)

			err = client.PaymentSplit().LinkImmediateCharge(charge.TxID, splitConfigResp.ID)
			if err != nil {
				log.Printf("Failed to link charge to split config: %v", err)
			} else {
				fmt.Println("Charge successfully linked to split config")

				splitCharge, err := client.PaymentSplit().GetImmediateChargeWithSplit(charge.TxID)
				if err != nil {
					log.Printf("Failed to get charge with split: %v", err)
				} else {
					fmt.Printf("Charge with split - TxID: %s, Split Config ID: %s\n",
						splitCharge.TxID, splitCharge.Config.ID)
				}

				err = client.PaymentSplit().UnlinkImmediateCharge(charge.TxID)
				if err != nil {
					log.Printf("Failed to unlink charge from split config: %v", err)
				} else {
					fmt.Println("Charge successfully unlinked from split config")
				}
			}
		}

		splitConfigID := fmt.Sprintf("split-%d", time.Now().Unix())
		configWithID, err := client.PaymentSplit().CreateConfigWithID(splitConfigID, splitConfig)
		if err != nil {
			log.Printf("Failed to create payment split config with ID: %v", err)
		} else {
			fmt.Printf("Split config created with specified ID: %s, Status: %s\n",
				configWithID.ID, configWithID.Status)
		}

		retrievedConfig, err := client.PaymentSplit().GetConfig(splitConfigResp.ID, 0)
		if err != nil {
			log.Printf("Failed to get payment split config: %v", err)
		} else {
			fmt.Printf("Retrieved split config - ID: %s, Description: %s\n",
				retrievedConfig.ID, retrievedConfig.Descricao)
			fmt.Printf("Main share: %s %s%%\n",
				retrievedConfig.Split.MinhaParte.Tipo, retrievedConfig.Split.MinhaParte.Valor)
			fmt.Printf("Splits: %d\n", len(retrievedConfig.Split.Repasses))
			for i, repasse := range retrievedConfig.Split.Repasses {
				fmt.Printf("  %d. Account: %s, Type: %s, Value: %s%%\n",
					i+1, repasse.Favorecido.Conta, repasse.Tipo, repasse.Valor)
			}
		}
	}

	dueTxid := fmt.Sprintf("due-txid-%d", time.Now().Unix())
	dueChargeReq := efi.CreateDueChargeRequest{
		Calendario: efi.CalendarioDueCharge{
			DataDeVencimento:       "2024-12-31",
			ValidadeAposVencimento: 30,
		},
		Devedor: efi.DevedorDueCharge{
			CPF:  "12345678909",
			Nome: "João Souza",
		},
		Valor: efi.ValorDueCharge{
			Original: "150.00",
		},
		Chave:              "YOUR_PIX_KEY",
		SolicitacaoPagador: "Cobrança com vencimento e split",
	}

	dueCharge, err := client.DueCharge().Create(dueTxid, dueChargeReq)
	if err != nil {
		log.Printf("Failed to create due charge: %v", err)
	} else {
		fmt.Printf("Due charge created with txid: %s\n", dueCharge.TxID)

		if splitConfigResp != nil && splitConfigResp.ID != "" {

			err = client.PaymentSplit().LinkDueCharge(dueCharge.TxID, splitConfigResp.ID)
			if err != nil {
				log.Printf("Failed to link due charge to split config: %v", err)
			} else {
				fmt.Println("Due charge successfully linked to split config")

				dueChargeSplit, err := client.PaymentSplit().GetDueChargeWithSplit(dueCharge.TxID)
				if err != nil {
					log.Printf("Failed to get due charge with split: %v", err)
				} else {
					fmt.Printf("Due charge with split - TxID: %s, Split Config ID: %s\n",
						dueChargeSplit.TxID, dueChargeSplit.Config.ID)
				}

				err = client.PaymentSplit().UnlinkDueCharge(dueCharge.TxID)
				if err != nil {
					log.Printf("Failed to unlink due charge from split config: %v", err)
				} else {
					fmt.Println("Due charge successfully unlinked from split config")
				}
			}
		}
	}
}
