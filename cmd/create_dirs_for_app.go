package cmd

import (
	"fmt"
	"os"
	"time"
)

func InitDirsByApp(
	directorioNa string,
) {

	var opc int32
	fmt.Println("Debe seguir cada paso como se explica para lograr encriptar de forma segura y eficiente el archivo deseado")
	fmt.Printf("Se creara un directorio llamado %q en el path de la aplicacion\n", directorioNa)

chooseOpc:
	for {

		fmt.Printf("1. (Si) crear el directorio de trabajo\n2. (No) crear el directorio de trabajo\n")
		_, errScanOpt := fmt.Scanf("%d", &opc)
		if errScanOpt != nil {

			fmt.Printf("Ingrese un valor como se indica anteriormente")
			continue chooseOpc
		}

		switch opc {

		case 1:
			{
				os.Mkdir(directorioNa, 0750)
				fmt.Println("Se creo el directorio correctamente")
				fmt.Println("..............")

				break chooseOpc
			}
		case 2:
			{
				fmt.Println("Se ha abortado el programa porque nafile necesita crear un directorio")
				fmt.Println("..............")
				time.Sleep(2 * time.Second)

				os.Exit(0)
				break chooseOpc
			}
		default:
			{
				fmt.Println("Debe ingresar un valor valido segun la opcion")
				fmt.Println("..............")
				continue chooseOpc
			}
		}
	}
}
