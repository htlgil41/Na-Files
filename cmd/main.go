package cmd

import (
	"fmt"
	"nafiles/fs"
	"os"
	"time"
)

func InitToolsForNa() {

	var directorioNa string = fmt.Sprintf("./%s", "na_files")
	var aeskey string = "AES.key"
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

	fmt.Println("Ahora debera seleccionar las siguientes opcion para la accion")
	fmt.Println("la opcion que ingresara es invalida el programa se cerrara")

	fmt.Printf("1. Tengo una llave AES\n2. Crear llave AES\n")
	_, errScanOpt := fmt.Scanf("%d", &opc)
	if errScanOpt != nil {

		fmt.Printf("Valor invalido el programa se cerrara")
		os.Exit(0)

		return
	}

	switch opc {

	case 1:
		{
			fmt.Println("Usted posee una llave AES que debera moverla en el direcotorio creado")
			fmt.Printf("Mueva la llave al directorio %q (el archivo debe ser llamado %q)\n", directorioNa, aeskey)

			fmt.Println("Presione cualquier tecla valida si ya movio la llave AES")
			var anyKey string
			fmt.Scanf("%s", &anyKey)

			fmt.Println(aeskey)
		}
	case 2:
		{

			fmt.Println("Se creara la llave AES")
			fmt.Printf("Direcotrio donde estara la llave %q (el archivo se llama %q)\n", directorioNa, aeskey)

			key, errAes := GenerateKeyAESforNcripted()
			if errAes != nil {

				fmt.Println("error al generar la llave AES ", errAes)
				os.Exit(0)
				return
			}

			fileAes, errFileAes := fs.CreateFile(
				directorioNa,
				aeskey,
			)
			if errFileAes != nil {

				fmt.Println("Erorr al crear el archivo con la llave")
				os.Exit(0)
				return
			}
			defer fileAes.Close()

			fileAes.WriteString(key)
		}
	default:
		{
			fmt.Println("Opcion invalida programa cerrado....")
			os.Exit(0)
			return
		}
	}
}
