package main

import (
	"strconv"
)

// parser represents a parser
type parser struct {
	lexer        *lexer
	currentToken Token
	variables    map[string]int
}

// Newparser creates a new parser instance
func Newparser(lexer *lexer) *parser {
	p := &parser{lexer: lexer, variables: make(map[string]int)}
	p.NextToken()
	return p
}

// NextToken advances to the next token
func (P *parser) NextToken() {
	P.currentToken = P.lexer.Nexttoken()
}

// Parse parses the input and returns the result
func (P *parser) Parse() {
	for P.currentToken.Type != TokenEOF {
		if P.currentToken.Type == TokenVariable {
			P.NextToken() // Skip variable
			if P.currentToken.Type == TokenPlus || P.currentToken.Type == TokenMinus {
				P.expr()

			} else if P.currentToken.Type == TokenAssign {
				P.handleAssignment()

			} else if P.currentToken.Type == TokenStar || P.currentToken.Type == TokenSlash {
				P.term()
			} else {
				panic("Expected '=' after variable")
			}

		} else {
			result := P.expr()
			println(result)
		}
	}
}

func (P *parser) handleAssignment() {

	varName := P.currentToken.Value

	P.NextToken() // Skip '='
	result := P.expr()
	P.variables[varName] = result
}

func (P *parser) expr() int {
	result := P.term()
	for P.currentToken.Type == TokenPlus || P.currentToken.Type == TokenMinus {
		op := P.currentToken
		P.NextToken()
		if op.Type == TokenPlus {
			result += P.term()
		} else if op.Type == TokenMinus {
			result -= P.term()
		}
	}
	return result
}

func (P *parser) term() int {
	result := P.factor()
	for P.currentToken.Type == TokenStar || P.currentToken.Type == TokenSlash {
		op := P.currentToken
		P.NextToken()
		if op.Type == TokenStar {
			result *= P.factor()
		} else if op.Type == TokenSlash {
			result /= P.factor()
		}
	}
	return result
}

func (P *parser) factor() int {
	token := P.currentToken
	if token.Type == TokenNumber {
		value, _ := strconv.Atoi(token.Value)
		P.NextToken()
		return value
	}
	if token.Type == TokenVariable {
		value, exists := P.variables[token.Value]
		if !exists {
			panic("Undefined variable: " + token.Value)
		}
		P.NextToken()
		return value
	}
	panic("Unexpected token: " + token.Value)
}
