package main

import (
	"fmt"
	"go-CLI/internal/cli"
	"go-CLI/internal/config"
	"os"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Printf("error load config: %v\n", err)
		os.Exit(1)
	}

	cli.Execute()
}
