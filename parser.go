package main

import (
	"strconv"
)

// Parser represents a parser
type Parser struct {
	lexer        *Lexer
	currentToken Token
	variables    map[string]int
}

// NewParser creates a new parser instance
func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer, variables: make(map[string]int)}
	p.nextToken()
	return p
}

// nextToken advances to the next token
func (p *Parser) nextToken() {
	p.currentToken = p.lexer.NextToken()
}

// Parse parses the input and returns the result
func (p *Parser) Parse() {
	for p.currentToken.Type != TokenEOF {
		if p.currentToken.Type == TokenVariable {
			p.handleAssignment()
		} else {
			result := p.expr()
			println(result)
		}
	}
}

func (p *Parser) handleAssignment() {
	varName := p.currentToken.Value
	p.nextToken() // Skip variable
	if p.currentToken.Type != TokenAssign {
		panic("Expected '=' after variable")
	}
	p.nextToken() // Skip '='
	result := p.expr()
	p.variables[varName] = result
}

func (p *Parser) expr() int {
	result := p.term()
	for p.currentToken.Type == TokenPlus || p.currentToken.Type == TokenMinus {
		op := p.currentToken
		p.nextToken()
		if op.Type == TokenPlus {
			result += p.term()
		} else if op.Type == TokenMinus {
			result -= p.term()
		}
	}
	return result
}

func (p *Parser) term() int {
	result := p.factor()
	for p.currentToken.Type == TokenStar || p.currentToken.Type == TokenSlash {
		op := p.currentToken
		p.nextToken()
		if op.Type == TokenStar {
			result *= p.factor()
		} else if op.Type == TokenSlash {
			result /= p.factor()
		}
	}
	return result
}

func (p *Parser) factor() int {
	token := p.currentToken
	if token.Type == TokenNumber {
		value, _ := strconv.Atoi(token.Value)
		p.nextToken()
		return value
	}
	if token.Type == TokenVariable {
		value, exists := p.variables[token.Value]
		if !exists {
			panic("Undefined variable: " + token.Value)
		}
		p.nextToken()
		return value
	}
	panic("Unexpected token: " + token.Value)
}
