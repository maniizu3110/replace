package logic

import (
	"io/ioutil"
	"path/filepath"
)

func DirToStrings(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(path, file.Name())

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return nil, err
			}

			fileEntry := "PATH: " + filePath + "\nCONTENT:\n" + string(content) + "\n"
			result = append(result, fileEntry)
		}
	}

	return result, nil
}
