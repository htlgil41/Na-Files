package cmd

import (
	"fmt"
	"os"
	"strings"
)

func InitFileNcripted() (*os.File, *os.File, error) {

	var path_full string = fmt.Sprintf("%s", "./na_files")
	var fileNcripted string

	fmt.Println("..............")
	fmt.Printf("Ahora bien, necesito que muevas el archivo que deseas encriptar al directorio\n")

	for {

		fmt.Println("Ingrese el nombre del archivo que movio al directorio junto con su extencion. Ejeplo File.zip")

		_, errScan := fmt.Scanf("%s", &fileNcripted)
		if errScan != nil {

			fmt.Println("Error al leer el nombre del archivo que desea encriptar")
			continue

		}

		filePlain, errverifyFilePlain := os.Open(
			fmt.Sprintf("%s/%s", path_full, fileNcripted),
		)
		if errverifyFilePlain != nil {

			fmt.Println("No se pudo leer la informacion del archivo, asegurece de moverlo al direcotrio y que este escrito correctamente al ingresar su nombre")
			continue
		}

		filenameOnly := strings.Split(fileNcripted, ".")[0]
		fileForNcripted, errFile := os.OpenFile(
			fmt.Sprintf("%s/%s", path_full, fmt.Sprintf("ncripted-%s.enc", filenameOnly)),
			os.O_CREATE|
				os.O_RDONLY|
				os.O_WRONLY,
			0750,
		)
		if errFile != nil {

			fmt.Println("Error al crear el archivo que tendra los datos encriptados")
			continue
		}

		return fileForNcripted, filePlain, nil
	}
}
