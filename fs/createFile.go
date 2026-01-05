package fs

import (
	"fmt"
	"os"
)

func CreateFile(
	path string,
	name_file string,
) (*os.File, error) {

	var pathFull string = fmt.Sprintf("./%s/%s", path, name_file)
	file, errFile := os.Create(pathFull)
	if errFile != nil {

		return nil, fmt.Errorf("Error al crear el archivo %q", pathFull)
	}

	return file, nil
}
