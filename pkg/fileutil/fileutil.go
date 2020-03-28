package fileutil

import (
	"errors"
	"os"
	"path"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func SearchUp(file string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for p := cwd; p != "/"; p = path.Dir(p) {
		cp := path.Join(p, file)

		if Exists(cp) {
			return cp, nil
		}
	}

	return "", errors.New("file does not exist in the current tree")
}
