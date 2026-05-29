package ai

import (
	"context"
	"fmt"
	"go-CLI/internal/prompt"

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

func (g *GeminiClient) SendMessage(msg string) (*string, error) {
	ctx := context.Background()
	systemIntruction := prompt.LoadSystemPrompt("kyy-agent-prompt")
	resp, err := g.client.Models.GenerateContent(
		ctx,
		g.model,
		genai.Text(msg),
		&genai.GenerateContentConfig{
			SystemInstruction: &genai.Content{
				Role: "system",
				Parts: []*genai.Part{
					{Text: systemIntruction},
				},
			},
			MaxOutputTokens: int32(g.maxOutputTokens),
			Temperature:     &g.temperature,
		},
	)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("Tidak ada response")
	}

	responseText := resp.Text()
	if responseText != "" {
		return &responseText, nil
	}

	return nil, fmt.Errorf("Tidak ada response")

}
