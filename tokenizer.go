package main

// TokenType represents the type of token
type TokenType string

const (
	// Token types
	TokenNumber   TokenType = "NUMBER"
	TokenPlus     TokenType = "PLUS"
	TokenMinus    TokenType = "MINUS"
	TokenStar     TokenType = "STAR"
	TokenSlash    TokenType = "SLASH"
	TokenAssign   TokenType = "ASSIGN"
	TokenVariable TokenType = "VARIABLE"
	TokenEOF      TokenType = "EOF"
)

// Token represents a single token
type Token struct {
	Type  TokenType
	Value string
}
