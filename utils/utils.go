package utils

import (
	"log"

	"github.com/btcsuite/btcutil/base64"
)

func convertToBase64(address string) string {
	data := []byte(string)
	encoded := base64.Encode(data)

	log.Printf("Encoded Data: %v", encoded)

	return encoded
}
