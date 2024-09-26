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

	//	Takes input file coverts to string then runs tokenizer
	input := string(file)
	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()

	// for loop to show which tokens are scanned (THIS WILL BE YOUR SAVIOR LATER IN PARSING)
	for _, token := range tokens {
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
	}
}
