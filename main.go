package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func encryptAESGCM(secretKey, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("new cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("new gcm: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("nonce: %w", err)
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)
	combined := append(nonce, ciphertext...) // nonce|ciphertext

	return base64.StdEncoding.EncodeToString(combined), nil
}

func decryptAESGCM(secretKey, cipherB64 string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		return "", fmt.Errorf("base64 decode: %w", err)
	}
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("new cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("new gcm: %w", err)
	}
	if len(raw) < gcm.NonceSize() {
		return "", fmt.Errorf("ciphertext too short")
	}
	nonce := raw[:gcm.NonceSize()]
	ct := raw[gcm.NonceSize():]
	pt, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return "", fmt.Errorf("gcm open: %w", err)
	}
	return string(pt), nil
}

func main() {
	secretKey := os.Getenv("MASTER_KEY")
	if len(secretKey) != 32 {
		fmt.Println("Error: MASTER_KEY must be exactly 32 characters (AES-256)")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  go run encryptor.go encrypt <plaintext>")
		fmt.Println("  go run encryptor.go decrypt <ciphertext_base64>")
		os.Exit(1)
	}

	mode := os.Args[1]
	value := os.Args[2]

	switch mode {
	case "encrypt":
		enc, err := encryptAESGCM(secretKey, value)
		if err != nil {
			fmt.Println("Encrypt error:", err)
			os.Exit(1)
		}
		fmt.Println(enc)

	case "decrypt":
		dec, err := decryptAESGCM(secretKey, value)
		if err != nil {
			fmt.Println("Decrypt error:", err)
			os.Exit(1)
		}
		fmt.Println(dec)

	default:
		fmt.Println("Invalid mode. Use 'encrypt' or 'decrypt'.")
		os.Exit(1)
	}
}
