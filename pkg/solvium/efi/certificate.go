package efi

import (
	"crypto/tls"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ReadFileBytes(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func ConvertP12ToPEM(p12Path, pemOutputPath, password string) error {
	
	if _, err := exec.LookPath("openssl"); err != nil {
		return errors.New("OpenSSL not found in PATH, required for P12 to PEM conversion")
	}

	args := []string{
		"pkcs12",
		"-in", p12Path,
		"-out", pemOutputPath,
		"-nodes",
	}

	if password != "" {
		args = append(args, "-password", "pass:"+password)
	} else {
		args = append(args, "-password", "pass:")
	}

	cmd := exec.Command("openssl", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to convert P12 to PEM: %w, output: %s", err, string(output))
	}

	return nil
}

func LoadCertificateFromP12(p12Path, password string) (tls.Certificate, error) {
	
	tempDir, err := ioutil.TempDir("", "efi-cert-")
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	pemPath := filepath.Join(tempDir, "cert.pem")

	
	if err := ConvertP12ToPEM(p12Path, pemPath, password); err != nil {
		return tls.Certificate{}, err
	}

	
	pemData, err := ioutil.ReadFile(pemPath)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to read PEM file: %w", err)
	}

	
	var certPEMBlock, keyPEMBlock []byte
	var certBlocks [][]byte

	for {
		var block *pem.Block
		block, pemData = pem.Decode(pemData)
		if block == nil {
			break
		}

		if block.Type == "CERTIFICATE" {
			certBlocks = append(certBlocks, pem.EncodeToMemory(block))
		} else if block.Type == "PRIVATE KEY" || strings.HasSuffix(block.Type, "PRIVATE KEY") {
			keyPEMBlock = pem.EncodeToMemory(block)
		}
	}

	if len(certBlocks) == 0 || keyPEMBlock == nil {
		return tls.Certificate{}, errors.New("failed to extract certificate and/or private key from PEM")
	}

	
	for _, certBlock := range certBlocks {
		certPEMBlock = append(certPEMBlock, certBlock...)
	}

	
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to load X509 key pair: %w", err)
	}

	return cert, nil
}
