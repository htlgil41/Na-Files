package fs

import (
	"fmt"
	"os"
)

func ReadFileForChuck(
	chuck_mg int64,
	path string,
	name_file string,
) (*[]byte, error) {

	var pathFull string = fmt.Sprintf("./%s/%s", path, name_file)
	chuckBuffer := make([]byte, chuck_mg)
	chuck_mg = (1024 * 1024) * chuck_mg

	if chuck_mg > 10 {

		fmt.Println("Por cuestion de rendimiento el chuck del file se establecio en 10Mb y no lo ingresado")
		chuck_mg = (1024 * 1024) * 10
	}

	fileForRead, errForRead := os.Open(pathFull)
	if errForRead != nil {

		return nil, fmt.Errorf("Error al leer el archivo, %q", errForRead.Error())
	}
	defer fileForRead.Close()

	fileForRead.Read(chuckBuffer)
	return &chuckBuffer, nil
}
