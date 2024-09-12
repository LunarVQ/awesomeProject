package main

import (
	"fmt"
	"os"
)

func main() {

	content, err := os.ReadFile("program.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(content)
	lexer := Newlexer(input)
	parser := Newparser(lexer)

	parser.Parse()
}
