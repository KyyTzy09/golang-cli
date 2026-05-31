package cli

import (
	"fmt"
	"go-CLI/internal/ai"
	"go-CLI/internal/prompt"
	"go-CLI/internal/tools"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(scanCmd)
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

		aiClient, err := ai.NewGeminiClient(apiKey)
		if err != nil {
			fmt.Printf("❌ Gagal terhubung: %v\n", err)
			os.Exit(1)
		}

		input := args[1]
		path := args[0]
		
		raw, err := tools.ScanFileContent(map[string]string{"path": path})
		if err != nil {
			fmt.Println(err)
			return
		}

		scanned, err := prompt.CleanCode(raw.(string))
		if err != nil {
			fmt.Println(err)
			return
		}

		finalPrompt := fmt.Sprintf("%s\n\n%s", input, scanned)

		fmt.Println("AI Sedang berpikir")
		response, err := aiClient.SendMessage(finalPrompt)
		if err != nil {
			fmt.Printf("❌ Gagal terhubung: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(*response)
	},
}