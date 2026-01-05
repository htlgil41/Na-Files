package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"nafiles/cmd"
)

func main() {

	var chuckFilePlain int64 = (1024 * 1024) * 10

	fmt.Println("Bienvendio!!, Na-files")
	keyAES, errKeyAES := cmd.InitToolsForNa()
	if errKeyAES != nil {
		log.Fatal(errKeyAES.Error())
	}

	keyAESAtByte, errkeyAESAtBte := hex.DecodeString(keyAES)
	if errkeyAESAtBte != nil {
		log.Fatal("Error al Decodificar la llave AES")
	}

	bufferFile := make([]byte, chuckFilePlain)
	fileForNcripted, filePlain, errGetFiles := cmd.InitFileNcripted()
	if errGetFiles != nil {

		log.Fatal(errGetFiles.Error())
	}
	defer filePlain.Close()
	defer fileForNcripted.Close()

	block, errBlock := aes.NewCipher(keyAESAtByte)
	if errBlock != nil {

		log.Fatal("Error al crear el bloque del cipher")
	}

	gcm, erGcm := cipher.NewGCM(block)
	if erGcm != nil {

		log.Fatal("Error al crar el modelo GCM")
	}

	var p int32 = 1
	for {
		nonce := make([]byte, gcm.NonceSize())
		io.ReadFull(rand.Reader, nonce)

		_, errRead := filePlain.Read(bufferFile)

		if errRead == io.EOF {
			fmt.Println("Se ha leido todo el archivo")
			break
		}

		if errRead != nil {
			fmt.Println("Error inesperado al leer el archivo el progama se cerrara")
			break
		}

		cipherText := gcm.Seal(
			nonce,
			nonce,
			bufferFile,
			nil,
		)

		fileForNcripted.Write(nonce)
		fileForNcripted.Write(cipherText)

		for range p {

			fmt.Print(".")
		}

		if p%10 == 0 {

			p = 1
		} else {

			p++
		}

		fmt.Println()
	}

	fmt.Printf("Llave AES %x\n", keyAESAtByte)
}
