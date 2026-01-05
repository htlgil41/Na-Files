package fs

import (
	"fmt"
	"os"
)

func GetFile(
	path string,
	name_file string,
) (*os.File, error) {

	var pathFull string = fmt.Sprintf("./%s/%s", path, name_file)
	file, errFile := os.OpenFile(
		pathFull,
		os.O_CREATE|
			os.O_RDONLY|
			os.O_WRONLY,
		0750,
	)
	if errFile != nil {

		return nil, fmt.Errorf("Error al lee el archivo, %q", errFile.Error())
	}

	return file, errFile
}
