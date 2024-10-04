package main

import (
	"fmt"
)

type parsing struct {
	Type        TokenType
	input       string
	position    int
	currentChar rune
}

func newparser(input string) *parsing {
	p := &parsing{
		input:    input,
		position: 0,
	}
	//	Takes input file coverts to string then runs tokenizer
	tokenizer := NewTokenizer(p.input)
	tokens := tokenizer.Tokenize()

	// for loop to show which tokens are scanned (THIS WILL BE YOUR SAVIOR LATER IN PARSING)
	for _, token := range tokens {
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
	}
	return nil
}
