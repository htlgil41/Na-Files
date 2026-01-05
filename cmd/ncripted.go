package cmd

import (
	"fmt"
	"nafiles/fs"
	"os"
)

func InitNcripted() (string, error) {

	var opc int32
	var directorioNa string = fmt.Sprintf("./%s", "na_files")
	var aeskey string = "AES.key"

	InitDirs(directorioNa)

	fmt.Println("Ahora debera seleccionar las siguientes opcion para la accion")
	fmt.Println("la opcion que ingresara es invalida el programa se cerrara")

	fmt.Printf("1. Tengo una llave AES\n2. Crear llave AES\n")
	_, errScanOpt := fmt.Scanf("%d", &opc)
	if errScanOpt != nil {

		fmt.Printf("Valor invalido el programa se cerrara")
		os.Exit(0)

		return "", fmt.Errorf("Opcion invalida al crear o leer la llave AES")
	}

	switch opc {

	case 1:
		{
			fmt.Println("Usted posee una llave AES que debera moverla en el direcotorio creado")
			fmt.Printf("Mueva la llave al directorio %q (el archivo debe ser llamado %q)\n", directorioNa, aeskey)

			fmt.Println("Presione cualquier tecla + ENTER valida si ya movio la llave AES")
			var anyKey string
			fmt.Scanf("%s", &anyKey)

			fmt.Println(aeskey)

			key, errFileAes := fs.ReadShortFile(
				directorioNa,
				aeskey,
			)
			if errFileAes != nil {

				fmt.Println("Erorr al crear el archivo con la llave", errFileAes.Error())
				os.Exit(0)
				return "", fmt.Errorf("Erorr al crear el archivo con la llave %q", errFileAes.Error())

			}

			return key, nil
		}
	case 2:
		{

			fmt.Println("Se creara la llave AES")
			fmt.Printf("Direcotrio donde estara la llave %q (el archivo se llama %q)\n", directorioNa, aeskey)

			key, errAes := GenerateKeyAESforNcripted()
			if errAes != nil {

				fmt.Println("error al generar la llave AES ", errAes)
				os.Exit(0)
				return "", fmt.Errorf("error al generar la llave AES %q", errAes)
			}

			fileAes, errFileAes := fs.CreateFile(
				directorioNa,
				aeskey,
			)
			if errFileAes != nil {

				fmt.Println("Erorr al crear el archivo con la llave", errFileAes.Error())
				os.Exit(0)
				return "", fmt.Errorf("Erorr al crear el archivo con la llave %q", errFileAes.Error())

			}
			defer fileAes.Close()
			fileAes.WriteString(key)

			return key, nil
		}
	default:
		{
			fmt.Println("Opcion invalida programa cerrado....")
			os.Exit(0)
			return "", fmt.Errorf("Opcion no disponible al crear o leer la llave AES")
		}
	}
}
