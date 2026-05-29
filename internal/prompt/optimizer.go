package prompt

import "strings"

func OptimizePrompt(prompt string) string {
	var result string
	lowercase := strings.ToLower(prompt)
	result = strings.TrimSpace(lowercase)

	words := strings.Fields(result)
	// Hilangkan kata kata yang tidak berguna
	junkWords := map[string]bool{
		"kan":     true,
		"tolong":  true,
		"makasih": true,
		"dong":    true,
		"saya":    true,
	}

	filteredWords := []string{}
	for _, word := range words {
		if !junkWords[word] {
			filteredWords = append(filteredWords, word)
		}
	}

	return strings.Join(filteredWords, " ")
}
