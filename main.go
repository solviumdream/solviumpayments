package main

import (
	"fmt"

	"github.com/solviumdream/solviumpayments/pkg/solvium/efi"
)


const Version = "0.1.0"

func main() {
	fmt.Println("SolviumPayments - Version", Version)
	fmt.Println("A modular payment solution for Go projects")
	fmt.Println("Currently supported payment providers:")
	fmt.Println("- Efi Payments Pix API")
	fmt.Println("\nFor usage examples, see the examples directory")
}



func GetEfiClient(clientID, clientSecret, certPath, certPassword string, isSandbox bool) (*efi.Client, error) {
	env := efi.Production
	if isSandbox {
		env = efi.Sandbox
	}

	return efi.NewClientFromP12(clientID, clientSecret, certPath, certPassword, env)
}
