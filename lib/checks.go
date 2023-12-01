package lib

import (
	"os"
	"path/filepath"
)

func CheckNextjsProjectDir() bool {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	_, errEx := os.Stat(exPath + "/src/app")
	return !os.IsNotExist(errEx)
}
