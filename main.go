package main

import (
	"fmt"
	"os"
)

func main() {
	//  Error checks for File
	file, err := os.ReadFile("program.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//	Takes input file coverts to string then runs parser
	input := string(file)
	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()
	parsing := newparser(tokens)
	parsing.Parse()

}
