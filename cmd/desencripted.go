package cmd

import "fmt"

func InitDesencripted() {

	//var opc int32
	var directorioNa string = fmt.Sprintf("./%s", "na_files")
	//var aeskey string = "AES.key"

	InitDirs(directorioNa)
}
