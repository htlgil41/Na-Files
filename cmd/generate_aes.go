package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateKeyAES() (string, error) {

	byesForSecuen := make([]byte, 32)
	_, errRand := rand.Read(byesForSecuen)
	if errRand != nil {

		return "", fmt.Errorf("Error al leer la scuencia de bytes")
	}

	return hex.EncodeToString(byesForSecuen), nil
}
