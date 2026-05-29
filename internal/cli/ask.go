package cli

import (
	"fmt"
	"go-CLI/internal/prompt"
	"strings"

	"github.com/spf13/cobra"
)

var isRaw bool

func init() {
	RootCmd.AddCommand(askCmd)
	askCmd.Flags().BoolVarP(&isRaw, "raw", "r", false, "Tampilkan pertanyaan mentah")
}

var askCmd = &cobra.Command{
	Use:   "ask [pertanyaan] --[flags]",
	Short: "Tanya sesuatu ke AI",
	Args:  cobra.MinimumNArgs(1), // Memaksa user untuk memasukkan minimal 1 argumen (pertanyaan)
	Run: func(cmd *cobra.Command, args []string) {
		// Menggabungkan argumen menjadi satu string pertanyaan
		pertanyaan := strings.Join(args, " ")
		if isRaw {
			fmt.Printf("🤖 Pertanyaan mentah: '%s'\n", pertanyaan)
			return // Langsung stop di sini, gak usah di-optimize
		}

		fmt.Println("🤖 AI sedang berpikir... ")
		pertanyaan = prompt.OptimizePrompt(pertanyaan)
		fmt.Printf("🔍 Kamu bertanya: '%s'\n", pertanyaan)
	},
}
