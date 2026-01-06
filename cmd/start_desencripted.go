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

	var chuckFilePlain int = (1024 * 1024) * 10
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
	encryptedChunkSize := nonceSize + chuckFilePlain + gcm.Overhead()
	buffer := make([]byte, encryptedChunkSize)

	infoFile, _ := fileEncripted.Stat()
	percentageTotal := float64(infoFile.Size()) / (float64(1024) * float64(1024))

	p := float64(0)
	for {
		n, errRead := fileEncripted.Read(buffer)
		if errRead == io.EOF {
			fmt.Println("Se ha leido todo el archivo")
			break
		}

		if errRead != nil {
			fmt.Println("Error inesperado al leer el archivo el progama se cerrara ", errRead.Error())
			break
		}

		nonceExtracr, ciphertext := buffer[:nonceSize], buffer[nonceSize:n]
		textDesen, errTextDesen := gcm.Open(
			nil,
			nonceExtracr,
			ciphertext,
			nil,
		)
		if errTextDesen != nil {

			fmt.Println("Error al abrir el GCM para desencriptar parte del chuck ", errTextDesen.Error())
			break
		}

		fileDesencripted.Write(textDesen)
		p = p + float64(n)/(float64(1024)*float64(1024))
		fmt.Printf("%d%s\n", int64((p/percentageTotal)*100), "%")
	}
}
