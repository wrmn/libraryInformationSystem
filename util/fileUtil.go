package util

import (
	"errors"
	"os"
)

// check if file exist
func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
