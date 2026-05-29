package prompt

import (
	"fmt"
	"os"
	"strings"
)

func LoadSystemPrompt(filename string) string {
	// Tentukan jalur file-nya
	filePath := fmt.Sprintf("internal/prompt/%s.prompt", filename)

	// Baca file pake os.ReadFile
	data, err := os.ReadFile(filePath)
	if err != nil {
		// Fallback prompt kalau filenya gak ketemu atau kehapus
		return "You are a helpful CLI assistant."
	}

	// Ubah bytes menjadi string dan bersihkan spasi/enter berlebih di ujung file
	return strings.TrimSpace(string(data))
}