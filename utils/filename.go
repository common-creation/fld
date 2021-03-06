package utils

import (
	"os"
	"path/filepath"
)

func MustGetName() string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Base(exe)
}

