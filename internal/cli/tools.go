package cli

import (
	"fmt"
	"go-CLI/internal/tools"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(toolsCmd)
}

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Tools for AI CLI",
	Long:  "Tools for AI CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tools for AI CLI")
		agentTools := tools.Tools

		// Contoh kalau AI ngirim perintah lewat JSON:
		commandFromAI := "WriteFile"
		paramsFromAI := map[string]string{
			"path":    "internal/test.txt",
			"content": "Perubahan kedua!",
		}

		// Panggil tool sesuai perintah AI
		if tool, ok := agentTools[commandFromAI]; ok {
			tool.Func(paramsFromAI)
		}
	},
}
