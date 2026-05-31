package helpers

import "strings"

func CleanJSONResponse(aiResponse string) string {
	// 1. Potong spasi atau enter tak berguna di awal dan akhir respon
	cleaned := strings.TrimSpace(aiResponse)

	// 2. Buang baris pembuka ```json atau ``` jika ada
	cleaned = strings.TrimPrefix(cleaned, "```json")
	cleaned = strings.TrimPrefix(cleaned, "```")

	// 3. Buang baris penutup ``` di paling bawah jika ada
	cleaned = strings.TrimSuffix(cleaned, "```")

	// 4. Bersihkan sekali lagi buat mastiin gak ada enter yang tertinggal
	return strings.TrimSpace(cleaned)
}