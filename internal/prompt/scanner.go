package prompt

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// ScanFolderTree akan menghasilkan string bentuk 'tree' dari sebuah direktori
func ScanFolderTree(rootPath string ) (string, error) {
	var result strings.Builder

	// Daftar folder sampah yang HARUS di-ignore (Hemat Token!)
	ignoreList := map[string]bool{
		".git":         true,
		"node_modules": true,
		"vendor":       true,
		"dist":         true,
	}

	// Mulai berjalan menelusuri folder
	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Kembalikan error kalau folder gak bisa diakses
		}

		// Ambil nama file/folder-nya saja (misal: "internal/ai" -> "ai")
		name := d.Name()

		// 1. STRATEGI HEMAT: Cek apakah folder ini masuk daftar ignore?
		if d.IsDir() && ignoreList[name] {
			// Kalau masuk list, perintahkan Go untuk LEWATI folder ini beserta seluruh isinya
			return filepath.SkipDir 
		}

		// Jangan masukkan folder utama (".") ke dalam hasil agar rapi
		if path == rootPath {
			result.WriteString(fmt.Sprintf("📁 %s/\n", name))
			return nil
		}

		// 2. Hitung kedalaman folder untuk efek tabulasi (indendation)
		// Menghitung berapa banyak karakter slash '/' untuk menentukan jarak geser ke kanan
		relPath, _ := filepath.Rel(rootPath, path)
		depth := strings.Count(relPath, string(filepath.Separator))
		indent := strings.Repeat("  ", depth)

		// 3. Format tampilan pohonnya
		if d.IsDir() {
			result.WriteString(fmt.Sprintf("%s├── 📁 %s/\n", indent, name))
		} else {
			result.WriteString(fmt.Sprintf("%s├── 📄 %s\n", indent, name))
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return result.String(), nil
}

