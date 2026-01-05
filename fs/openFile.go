package fs

import (
	"fmt"
	"os"
)

func OpenFile(
	path string,
	name_file string,
) (*os.File, error) {

	var pathFull string = fmt.Sprintf("./%s/%s", path, name_file)

	fileForRead, errForRead := os.Open(pathFull)
	if errForRead != nil {

		return nil, fmt.Errorf("Error al leer el archivo, %q", errForRead.Error())
	}

	return fileForRead, nil
}
