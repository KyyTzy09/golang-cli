package cli

import (
	"fmt"
	"go-CLI/internal/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var geminiApiKey string

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&geminiApiKey, "gemini", "g", "", "API key untuk AI CLI")
}

var configCmd = &cobra.Command{
	Use:   "config [option] --flags",
	Short: "Konfigurasi AI CLI",
	Long:  "Opsi ini digunakan untuk konfigurasi AI CLI Anda",
	Run: func(cmd *cobra.Command, args []string) {
		if geminiApiKey != "" {
			viper.Set("gemini_api_key", geminiApiKey)

			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("❌ Gagal menyimpan konfigurasi: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("✅ API Key Gemini berhasil disimpan!")
			return
		}

		currentKey := viper.GetString("gemini_api_key")
		if currentKey != "" {
			config.AppConfig.GeminiAPIKey = currentKey
			fmt.Printf("✅ API Key Gemini saat ini: %s\n", currentKey)
		} else {
			fmt.Println("❌ API Key Gemini belum diset.")
		}
	},
}
