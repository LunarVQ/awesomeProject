package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <program_file>")
		return
	}

	programFile := os.Args[1]
	content, err := os.ReadFile(programFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(content)
	lexer := NewLexer(input)
	parser := NewParser(lexer)

	parser.Parse()
}
