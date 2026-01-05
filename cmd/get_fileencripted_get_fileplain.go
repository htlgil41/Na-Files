package cmd

import (
	"fmt"
	"os"
)

func GetFileForEncriptedAndFilePlain() (*os.File, *os.File, error) {

	var path_full string = fmt.Sprintf("%s", "./na_files")
	var fileNcripted string
	var fileDesencripted string

	fmt.Println("..............")
	fmt.Printf("Ahora bien, necesito que muevas el archivo que deseas desencriptar al directorio\n")

	for {

		fmt.Println("Ingrese el nombre del archivo encriptado que movio al directorio junto con su extencion. Ejeplo File.enc")

		_, errScan := fmt.Scanf("%s", &fileNcripted)
		if errScan != nil {

			fmt.Println("Error al leer el nombre del archivo que desea encriptar")
			continue

		}

		fmt.Println("Ingrese el nombre del archivo real con su extension original. Ejeplo File.mkv")

		_, errfileDesencripted := fmt.Scanf("%s", &fileDesencripted)
		if errfileDesencripted != nil {

			fmt.Println("Error al leer el nombre del archivo que desea encriptar")
			continue

		}

		fileHasEncripted, errfileHasEncripted := os.Open(
			fmt.Sprintf("%s/%s", path_full, fileNcripted),
		)
		if errfileHasEncripted != nil {

			fmt.Println("No se pudo leer la informacion del archivo, asegurece de moverlo al direcotrio y que este escrito correctamente al ingresar su nombre")
			continue
		}

		fileForDesencripted, errFile := os.OpenFile(
			fmt.Sprintf("%s/%s", path_full, fileDesencripted),
			os.O_CREATE|
				os.O_RDONLY|
				os.O_WRONLY,
			0750,
		)
		if errFile != nil {

			fmt.Println("Error al crear el archivo que tendra los datos desencriptados")
			continue
		}

		return fileForDesencripted, fileHasEncripted, nil
	}
}
