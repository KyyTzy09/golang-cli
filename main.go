package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input string

	for {
		fmt.Print("You: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Bye!")
			break
		}

		// Process user input
		fmt.Println("AI:", input)
	}
}
