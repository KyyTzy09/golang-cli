package tools

import (
	"fmt"
	"os"
)

func ScanFileContent(args map[string]string) (any, error) {
	filePath := args["path"]
	if filePath == "" {
		return "", fmt.Errorf("argumen 'path' tidak boleh kosong")
	}
	
	wd, _ := os.Getwd()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("terminal di '%s' tidak menemukan file '%s' (%w)", wd, filePath, err)
	}

	return string(data), nil
}
