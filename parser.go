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

		/*if token.Type == TokenKeyword {
			return p.keywordAction(TokenType(token.Value))
		}

		if token.Type == TokenIdentifier {
			return p.identifyVar(TokenType(token.Value))
		}

		// Detect character literals: sequences like 'a'
		if token.Type == TokenCharLiteral {
			return p.SaveVar(TokenType(token.Value))
		}

		if token.Type == TokenStringLiteral {
			return p.SaveVar(TokenType(token.Value))
		}

		if token.Type == TokenEOF {
			break
		}

		// Symbols handling
		if token.Type == TokenSymbol {
			return p.identifySymbol(TokenType(token.Value))
		}

		if token.Type == TokenNumber {
			return p.SaveVar(TokenType(token.Value))
		}
		*/
	}
	return nil
}

/*
// Takes values of keywords, outputs an action if actionable keyword otherwise initalizes
func (p *parsing) keywordAction(tokenType TokenType) *parsing {
	return nil
}

func (p *parsing) identifyVar(TokenType) *parsing {
	return nil
}

// Saves a given Token value to an idetified token
func (p *parsing) SaveVar(tokenType TokenType) *parsing {
	return nil
}

func (p *parsing) identifySymbol(TokenType) *parsing {
	return nil
}
*/
