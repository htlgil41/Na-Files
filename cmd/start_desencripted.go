package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func StartDesencripted() {

	var chuckFilePlain int64 = (1024 * 1024) * 10
	keyAe, errkeyAes := GetAesKey()
	if errkeyAes != nil {

		fmt.Println(errkeyAes.Error())
		return
	}

	keyAesAtByte, errkeyAesAtByte := hex.DecodeString(keyAe)
	if errkeyAesAtByte != nil {

		fmt.Println("Error al decodificar llave AES")
		return
	}

	fileDesencripted, fileEncripted, errGetFiles := GetFileForEncriptedAndFilePlain()
	if errGetFiles != nil {

		log.Fatal(errGetFiles.Error())
	}
	defer fileDesencripted.Close()
	defer fileEncripted.Close()

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

	for {

		n, errRead := fileDesencripted.Read(buffer)
		if errRead == io.EOF {
			fmt.Println("Se ha leido todo el archivo")
			break
		}

		if errRead != nil {
			fmt.Println("Error inesperado al leer el archivo el progama se cerrara")
			break
		}

		fmt.Println(string(buffer[:n]))
	}
}
