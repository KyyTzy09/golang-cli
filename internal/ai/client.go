package ai

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type GeminiClient struct {
	client          *genai.Client
	model           string
	temperature     float32
	maxOutputTokens int
}

func NewGeminiClient(apiKey string) (*GeminiClient, error) {
	ctx := context.Background()

	// Init client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("Gagal inisialisasi gemini client:", err)
	}

	return &GeminiClient{
		client:          client,
		model:           "gemini-2.5-flash",
		temperature:     0.2,
		maxOutputTokens: 4000,
	}, nil
}

func (g *GeminiClient) SendMessage(prompt string) (*string, error) {
	ctx := context.Background()
	systemInstruction := "You are a strict CLI assistant. Give direct answers or code snippets. No conversational filler, no pleasantries, no explanations unless explicitly asked."
	resp, err := g.client.Models.GenerateContent(
		ctx,
		g.model,
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			SystemInstruction: &genai.Content{
				Role: "system",
				Parts: []*genai.Part{
					{Text: systemInstruction},
				},
			},
			MaxOutputTokens: int32(g.maxOutputTokens),
			Temperature:     &g.temperature,
		},
	)

	if err != nil {
		return nil, err
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		responseText := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

		return &responseText, nil
	}

	return nil, fmt.Errorf("Tidak ada response")

}
