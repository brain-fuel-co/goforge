package inrootdir

import "os"

func InRootDir() bool {
	_, statusErr := os.Stat("go.mod")
	return !os.IsNotExist(statusErr)
}
