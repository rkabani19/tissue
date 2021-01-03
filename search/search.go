package search

import (
	"fmt"
	"os"
	"path/filepath"
)

func TraverseFiles() ([]string, error) {
	searchDir := "."

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if f.Mode().IsRegular() {
			fileList = append(fileList, path)
		}
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	return fileList, nil
}
