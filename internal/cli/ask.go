package cli

import (
	"fmt"
	"go-CLI/internal/ai"
	"go-CLI/internal/prompt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var isRaw bool

func init() {
	RootCmd.AddCommand(askCmd)
	askCmd.Flags().BoolVarP(&isRaw, "raw", "r", false, "Tampilkan pertanyaan mentah")
}

var askCmd = &cobra.Command{
	Use:   "ask [question] --[flags]",
	Short: "Tanya sesuatu ke AI",
	Args:  cobra.MinimumNArgs(1), // Memaksa user untuk memasukkan minimal 1 argumen (pertanyaan)
	Run: func(cmd *cobra.Command, args []string) {
		// Menggabungkan argumen menjadi satu string pertanyaan
		question := strings.Join(args, " ")
		fmt.Printf("🔍 Kamu bertanya: '%s'\n", question)

		fmt.Println("🤖 AI sedang berpikir... ")
		// Mendapatkan api key
		apikey := viper.GetString("gemini_api_key")
		if apikey == "" {
			fmt.Println("❌ API Key Gemini belum diset.")
			return
		}

		// init client
		client, err := ai.NewGeminiClient(apikey)
		if err != nil {
			fmt.Printf("❌ Gagal inisialisasi gemini client: %v\n", err)
			return
		}
		
		// Normalize pertanyaan
		question = prompt.OptimizePrompt(question)
		
		// Kirim pertanyaan ke AI
		response, err := client.SendMessage(question)
		if err != nil {
			fmt.Printf("❌ Gagal mengirim pertanyaan ke AI: %v\n", err)
			return
		}
		fmt.Printf("💡 AI menjawab: '%s'\n", *response)
	},
}
