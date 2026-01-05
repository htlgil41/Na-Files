package cmd

import (
	"fmt"
	"os"
)

func InitDesencripted() (string, error) {

	var aes_key string = "AES.key"
	var directorioNa string = fmt.Sprintf("./%s", "na_files")

	InitDirs(directorioNa)

	fmt.Printf("Mueva la llave AES dentro del directorio creado (especificamente con el nombre %q)\n", aes_key)
	fmt.Println("La llave debe estaer codigocada en hex....")

	fmt.Println("Presione ENTER si ya movio la llave AES")
	var anyKey string
	fmt.Scanf("%s", &anyKey)

	keyAes, errKeyAes := os.ReadFile(fmt.Sprintf("./%s/%s", "na_files", aes_key))
	if errKeyAes != nil {
		fmt.Println("No se pudo leer el archivo que contiene la llave, asegurece de mover la llave al directorio")
		return "", fmt.Errorf("No se pudo leer el archivo que contiene la llave, asegurece de mover la llave al directorio ", errKeyAes.Error())
	}

	return string(keyAes), nil

}
