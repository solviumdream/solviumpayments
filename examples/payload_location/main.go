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

	locationReq := efi.CreatePayloadLocationRequest{
		TipoCob: efi.PayloadLocationTypeCOB,
	}

	location, err := client.PayloadLocation().Create(locationReq)
	if err != nil {
		log.Printf("Failed to create payload location: %v", err)
	} else {
		fmt.Printf("Created payload location - ID: %d, Location: %s, Type: %s\n",
			location.ID, location.Location, location.TipoCob)

		qrCode, err := client.PayloadLocation().GenerateQRCode(location.ID)
		if err != nil {
			log.Printf("Failed to generate QR code: %v", err)
		} else {
			fmt.Printf("QR Code: %s\n", qrCode.QRCode)
			fmt.Printf("Visualization Link: %s\n", qrCode.LinkVisualizacao)
		}
	}

	locationCOBVReq := efi.CreatePayloadLocationRequest{
		TipoCob: efi.PayloadLocationTypeCOBV,
	}

	locationCOBV, err := client.PayloadLocation().Create(locationCOBVReq)
	if err != nil {
		log.Printf("Failed to create COBV payload location: %v", err)
	} else {
		fmt.Printf("Created COBV payload location - ID: %d, Location: %s, Type: %s\n",
			locationCOBV.ID, locationCOBV.Location, locationCOBV.TipoCob)
	}

	startDate := time.Now().Add(-30 * 24 * time.Hour)
	endDate := time.Now()
	listOptions := &efi.ListPayloadLocationsOptions{
		PaginaAtual:    0,
		ItensPorPagina: 10,
	}

	locations, err := client.PayloadLocation().List(startDate, endDate, listOptions)
	if err != nil {
		log.Printf("Failed to list payload locations: %v", err)
	} else {
		fmt.Printf("Found %d payload locations\n", len(locations.Loc))
		for i, loc := range locations.Loc {
			fmt.Printf("%d. ID: %d, Type: %s, Location: %s\n",
				i+1, loc.ID, loc.TipoCob, loc.Location)
			if loc.TxID != "" {
				fmt.Printf("   Linked to TxID: %s\n", loc.TxID)
			}
		}
	}

	if len(locations.Loc) > 0 {
		firstLocation := locations.Loc[0]
		locationDetail, err := client.PayloadLocation().GetByID(firstLocation.ID)
		if err != nil {
			log.Printf("Failed to get payload location: %v", err)
		} else {
			fmt.Printf("Location details - ID: %d, Type: %s, Created: %s\n",
				locationDetail.ID, locationDetail.TipoCob, locationDetail.Criacao)
			if locationDetail.TxID != "" {
				fmt.Printf("Linked to TxID: %s\n", locationDetail.TxID)

				unlinkedLocation, err := client.PayloadLocation().UnlinkTxID(locationDetail.ID)
				if err != nil {
					log.Printf("Failed to unlink txid: %v", err)
				} else {
					fmt.Printf("Successfully unlinked TxID from location ID: %d\n", unlinkedLocation.ID)
				}
			}
		}
	}
}
