package cli

import (
	"bufio"
	"fmt"
	"go-CLI/internal/ai"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(chatCmd)
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Masuk ke mode chat interaktif dengan AI",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("gemini_api_key")
		if apiKey == "" {
			fmt.Println("❌ Error: API Key belum diatur. Jalankan 'ai config' terlebih dahulu.")
			os.Exit(1)
		}

		// Inisialisasi AI Client
		aiClient, err := ai.NewGeminiClient(apiKey)
		if err != nil {
			fmt.Printf("❌ Gagal terhubung: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("🤖 Mode Chat Aktif! Ketik 'keluar' atau 'exit' untuk berhenti.")
		fmt.Println("----------------------------------------------------------------")

		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("\n👤 Kamu > ")
			
			// Ambil input setelah user pencet Enter
			if !scanner.Scan() {
				break
			}
			input := scanner.Text()

			// Validasi tombol keluar
			if input == "exit" || input == "keluar" {
				fmt.Println("👋 Sampai jumpa!")
				break
			}

			// Lewati jika input kosong
			if strings.TrimSpace(input) == "" {
				continue
			}

			fmt.Println("🤖 AI sedang berpikir...")

			response, err := aiClient.SendMessage(input)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
				continue
			}

			fmt.Println("\n📝 JAWABAN AI:")
			fmt.Println(*response)
		}
	},
}