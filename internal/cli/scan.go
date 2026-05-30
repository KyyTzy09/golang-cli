package cli

import (
	"fmt"
	"go-CLI/internal/ai"
	"go-CLI/internal/prompt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var path string = "."
func init() {
	RootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&path, "path", "p", "", "Root folder")
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan file",
	Long:  "Scan file with AI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Scan file...")
		apiKey := viper.GetString("gemini_api_key")
		if apiKey == "" {
			fmt.Println("❌ Error: API Key belum diatur. Jalankan 'ai config' terlebih dahulu.")
			os.Exit(1)
		}

		// Inisialisasi AI Client
		fmt.Println("AI Sedang berpikir")

		aiClient, err := ai.NewGeminiClient(apiKey)
		if err != nil {
			fmt.Printf("❌ Gagal terhubung: %v\n", err)
			os.Exit(1)
		}

		input := strings.Join(args, " ")
		raw, err := prompt.ScanFileContent(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		scanned, err := prompt.CleanCode(raw)
		if err != nil {
			fmt.Println(err)
			return
		}

		finalPrompt := fmt.Sprintf("%s\n\n%s", input, scanned)

		response, err := aiClient.SendMessage(finalPrompt)
		if err != nil {
			fmt.Printf("❌ Gagal terhubung: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(*response)
	},
}