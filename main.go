package main

import (
	"fmt"
	"log"
	"nafiles/cmd"
)

func main() {

	fmt.Println("Bienvendio!!, Na-files")
	keyAES, errKeyAES := cmd.InitToolsForNa()
	if errKeyAES != nil {
		log.Fatal(errKeyAES.Error())
	}

	fmt.Println("Llave AES ", keyAES)
}
