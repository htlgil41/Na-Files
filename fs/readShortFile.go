package fs

import (
	"fmt"
	"os"
)

func ReadShortFile(
	path string,
	name_file string,
) (string, error) {

	var pathFull string = fmt.Sprintf("./%s/%s", path, name_file)

	fileForRead, errForRead := os.ReadFile(pathFull)
	if errForRead != nil {

		return "", fmt.Errorf("Error al leer el archivo, %q", errForRead.Error())
	}

	return string(fileForRead), nil
}
