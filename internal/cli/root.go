package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ai",
	Short: "AI CLI - Asisten AI Hemat token yang diciptakan untuk terminalmu",
	Long:  "Sebuah perkakas CLI berbasis Golang untuk berinteraksi dengan LLM secara efisien.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
