package main

import (
	"fmt"
)

type parser struct {
	lexer          []Token
	token          Token
	intvariables   map[string]int64
	floatvariables map[string]float64
}

func newparser(lexer []Token) *parser {
	p := &parser{
		lexer:          lexer,
		intvariables:   make(map[string]int64),
		floatvariables: make(map[string]float64),
	}
	return p
}

func (p *parser) Parse() *parser {
	//	Takes input file coverts to string then runs tokenizer

	// for loop to show which tokens are scanned (THIS WILL BE YOUR SAVIOR LATER IN PARSING)
	for _, token := range p.lexer {
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)

		if token.Type == TokenKeyword {
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

	}
	return nil
}

// Takes values of keywords, outputs an action if actionable keyword otherwise initalizes
func (p *parser) keywordAction(tokenType TokenType) *parser {
	fmt.Printf("Action-, %s\n")

	return nil
}

func (p *parser) identifyVar(TokenType) *parser {
	return nil
}

// Saves a given Token value to an idetified token
func (p *parser) SaveVar(tokenType TokenType) *parser {
	return nil
}

func (p *parser) identifySymbol(TokenType) *parser {
	return nil
}
