package tools

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteFileContent(args map[string]string) (any, error) {
	rawPath := args["path"]
	content := args["content"]

	rawPath = strings.TrimPrefix(rawPath, "/")
	rawPath = strings.TrimPrefix(rawPath, "\\")

	path := filepath.Clean(rawPath)
	if path == "" || content == "" {
		msg := "Gagal: Argumen 'path' atau 'content' tidak boleh kosong."
		return nil,errors.New(msg) 
	}

	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		msg := fmt.Sprintf("❌ Gagal membuat folder di dalam project: %v\n", err)
		return nil, errors.New(msg)
	}

	// buka file
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// masukan content ke file
	_, err = file.WriteString(content)
	if err != nil {
		msg := fmt.Sprintf("❌ Gagal menulis ke file: %v\n", err)
		return nil, errors.New(msg)
	}

	return nil, nil
}
