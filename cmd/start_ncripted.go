package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func StartEncripted() {

	var chuckFilePlain int64 = (1024 * 1024) * 10
	keyAES, errKeyAES := CreateOrGetAes()
	if errKeyAES != nil {
		log.Fatal(errKeyAES.Error())
	}

	keyAESAtByte, errkeyAESAtBte := hex.DecodeString(keyAES)
	if errkeyAESAtBte != nil {
		log.Fatal("Error al Decodificar la llave AES")
	}

	bufferFile := make([]byte, chuckFilePlain)
	fileForNcripted, filePlain, errGetFiles := CreateFileForEncriptedAndFilePlain()
	if errGetFiles != nil {

		log.Fatal(errGetFiles.Error())
	}
	defer filePlain.Close()
	defer fileForNcripted.Close()

	block, errBlock := aes.NewCipher(keyAESAtByte)
	if errBlock != nil {

		log.Fatal("Error al crear el bloque del cipher")
	}

	gcm, errGcm := cipher.NewGCM(block)
	if errGcm != nil {

		log.Fatal("Error al crar el modelo GCM")
	}

	infoFile, _ := filePlain.Stat()
	percentageTotal := float64(infoFile.Size()) / (float64(1024) * float64(1024))

	p := float64(0)
	for {
		n, errRead := filePlain.Read(bufferFile)
		if errRead == io.EOF {
			fmt.Println("Se ha leido todo el archivo")
			break
		}
		if errRead != nil {
			fmt.Println("Error inesperado al leer el archivo el progama se cerrara")
			break
		}

		nonce := make([]byte, gcm.NonceSize())
		io.ReadFull(rand.Reader, nonce)

		cipherText := gcm.Seal(
			nil,
			nonce,
			bufferFile[:n],
			nil,
		)

		_, errWriteNonce := fileForNcripted.Write(nonce)
		if errWriteNonce != nil {

			fmt.Println("Error al escribir en el achivo encriptado [nonce]")
			return
		}
		_, errWriteCihper := fileForNcripted.Write(cipherText)
		if errWriteCihper != nil {

			fmt.Println("Error al escribir en el achivo encriptado [cipher]")
			return
		}

		p = p + float64(n)/(float64(1024)*float64(1024))
		fmt.Printf("%d%s\n", int64((p/percentageTotal)*100), "%")
	}

	fmt.Printf("Debere proteger su llave ya que si la pierde no hay manera de recuperar su archivo '%xw'\n", keyAESAtByte)
}
