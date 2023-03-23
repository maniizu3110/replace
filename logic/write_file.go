package logic

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func WiteFile(path string, content string) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
