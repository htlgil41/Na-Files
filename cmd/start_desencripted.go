package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
)

func DesencripteStart() {

	var chuckFilePlain int64 = (1024 * 1024) * 10
	keyAe, errkeyAes := InitDesencripted()
	if errkeyAes != nil {

		fmt.Println(errkeyAes.Error())
		return
	}

	keyAesAtByte, errkeyAesAtByte := hex.DecodeString(keyAe)
	if errkeyAesAtByte != nil {

		fmt.Println("Error al decodificar llave AES")
		return
	}

	block, errBlock := aes.NewCipher(keyAesAtByte)
	if errBlock != nil {

		log.Fatal("Error al crear el bloque del cipher")
	}
	gcm, errGcm := cipher.NewGCM(block)
	if errGcm != nil {

		log.Fatal("Error al crar el modelo GCM")
	}

	nonceSize := gcm.NonceSize()
	encryptedChunkSize := nonceSize + int(chuckFilePlain) + gcm.Overhead()
	buffer := make([]byte, encryptedChunkSize)
}
