package cli

import (
	"encoding/json"
	"fmt"
	"go-CLI/internal/ai"
	"go-CLI/internal/prompt"
	"go-CLI/internal/structs"
	"go-CLI/internal/tools"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(toolsCmd)
}

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Tools for AI CLI",
	Long:  "Tools for AI CLI",
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args , " ")
		folderStructure, err := prompt.ScanFolderTree(".")
		if err != nil {
			fmt.Println(err)
			return
		}

		rawPrompt := prompt.LoadSystemPrompt("tools")

		fmt.Println("Tools for AI CLI")
		// cek tools yang ada
		agentTools := tools.Tools

		apiKey := viper.GetString("gemini_api_key")
		if apiKey == "" {
			fmt.Println("❌ Error: API Key belum diatur. Jalankan 'ai config' terlebih dahulu.")
			os.Exit(1)
		}

		aiClient, err := ai.NewGeminiClient(apiKey)
		if err != nil {
			fmt.Printf("❌ Gagal terhubung: %v\n", err)
			os.Exit(1)
		}

		// Contoh kalau AI ngirim perintah lewat JSON:

			payload := map[string]interface{}{
				"text":             text,
				"folder_structure": folderStructure,
				"tools":            agentTools,
			}

			bytesResult, err := json.Marshal(payload)
			if err != nil {
				fmt.Println("❌ Gagal convert ke JSON:", err)
				return
			}

			jsonString := string(bytesResult)

			prompt := strings.ReplaceAll(rawPrompt, "{{payload_json}}", jsonString)

			aiResponse, err := aiClient.SendMessage(prompt)
			if err != nil {
				fmt.Printf("❌ Gagal terhubung: %v\n", err)
				os.Exit(1)
			}

			// Panggil tool sesuai perintah AI
			fmt.Println("ai response: ", *aiResponse )
			var jsonResponse structs.AIToolsResponse
			err = json.Unmarshal([]byte(*aiResponse), &jsonResponse)
			if err != nil {
				fmt.Println("❌ Gagal convert ke JSON:", err)
				return
			}

			if tool, ok := agentTools[jsonResponse.ToolName]; ok {
				tool.Func(jsonResponse.Params)
			}

			fmt.Println(jsonResponse.Suggestion)
	},
}
