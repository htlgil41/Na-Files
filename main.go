package main

import (
	"fmt"
	"nafiles/cmd"
	"os"
	"time"
)

func main() {

	var opc int64

	fmt.Println("Bienvendio!!, Na-files")
	fmt.Println("Ingrese una opcion para seguir con el programa")

chooseOpc:
	for {
		fmt.Printf("1. Encriptar un archivo\n2. Desencriptar mi archivo\n3. Salir\n")
		_, errScanOpc := fmt.Scanf("%d", &opc)
		if errScanOpc != nil {

			fmt.Println("Error al leer la respuesta asegurece de ingresar el dato correspondiente")
			continue
		}

		switch opc {
		case 1:
			{

				fmt.Println("Encriptar mi archivo")
				cmd.NcripteStart()
				break chooseOpc
			}
		case 2:
			{
				fmt.Println("Desencriptar mi archivo")
				break chooseOpc
			}
		case 3:
			{
				fmt.Println("Cerrando programa ......")
				time.Sleep(2 * time.Second)
				os.Exit(0)
			}
		default:
			{
				fmt.Println("Debera ingresar solo los datos que se muestran a continuacion")
				continue
			}
		}
	}
}
