package main

import (
	"strings"
	"unicode"
)

// TokenType represents different types of tokens
type TokenType string

const (
	TokenKeyword    TokenType = "KEYWORD"
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenNumber     TokenType = "NUMBER"
	TokenSymbol     TokenType = "SYMBOL"
	TokenEOF        TokenType = "EOF"
)

// Token struct with type and value
type Token struct {
	Type  TokenType
	Value string
}

// Define KeyWords type and a set of keywords
type KeyWord string

const (
	KeyWordPrint   KeyWord = "print"
	KeyWordIf      KeyWord = "if"
	KeyWordElse    KeyWord = "else"
	KeyWordGoto    KeyWord = "goto"
	KeyWordReturn  KeyWord = "return"
	KeyWordFor     KeyWord = "for"
	KeyWordInt     KeyWord = "int"
	KeyWordFloat   KeyWord = "float"
	KeyWordDouble  KeyWord = "double"
	KeyWordString  KeyWord = "string"
	KeyWordBreak   KeyWord = "break"
	KeyWordDefault KeyWord = "default"
	KeyWordSwitch  KeyWord = "switch"
	KeyWordCase    KeyWord = "case"
)

// A map of keywords for easy lookup
var keywords = map[string]KeyWord{
	"print":   KeyWordPrint,
	"if":      KeyWordIf,
	"else":    KeyWordElse,
	"goto":    KeyWordGoto,
	"return":  KeyWordReturn,
	"for":     KeyWordFor,
	"int":     KeyWordInt,
	"float":   KeyWordFloat,
	"double":  KeyWordDouble,
	"string":  KeyWordString,
	"break":   KeyWordBreak,
	"default": KeyWordDefault,
	"switch":  KeyWordSwitch,
	"case":    KeyWordCase,
}

// Tokenizer structure with input and current position
type Tokenizer struct {
	input       string
	position    int
	currentChar rune
}

// NewTokenizer initializes the tokenizer
func NewTokenizer(input string) *Tokenizer {
	t := &Tokenizer{
		input:    input,
		position: 0,
	}
	t.advance()
	return t
}

// Advance the tokenizer to the next character
func (t *Tokenizer) advance() {
	if t.position < len(t.input) {
		t.currentChar = rune(t.input[t.position])
	} else {
		t.currentChar = 0 // Null char represents EOF
	}
	t.position++
}

// Peek the next character without advancing
func (t *Tokenizer) peek() rune {
	if t.position < len(t.input) {
		return rune(t.input[t.position])
	}
	return 0
}

// GetNextToken reads the next token from the input
func (t *Tokenizer) GetNextToken() Token {
	// Skip whitespace
	for unicode.IsSpace(t.currentChar) {
		t.advance()
	}

	if unicode.IsLetter(t.currentChar) {
		return t.readIdentifierOrKeyword()
	}

	if unicode.IsDigit(t.currentChar) {
		return t.readNumber()
	}

	if t.currentChar == 0 {
		return Token{Type: TokenEOF, Value: ""}
	}

	// Symbols handling
	if strings.ContainsRune("(){};", t.currentChar) {
		ch := t.currentChar
		t.advance()
		return Token{Type: TokenSymbol, Value: string(ch)}
	}

	// Default to unknown single character
	ch := t.currentChar
	t.advance()
	return Token{Type: TokenSymbol, Value: string(ch)}
}

// Read identifier or keyword
func (t *Tokenizer) readIdentifierOrKeyword() Token {
	start := t.position - 1
	for unicode.IsLetter(t.currentChar) || unicode.IsDigit(t.currentChar) {
		t.advance()
	}
	value := t.input[start : t.position-1]

	// Check if it's a keyword
	if keyword, found := keywords[value]; found {
		return Token{Type: TokenKeyword, Value: string(keyword)}
	}

	// Otherwise, it's an identifier
	return Token{Type: TokenIdentifier, Value: value}
}

// Read number token
func (t *Tokenizer) readNumber() Token {
	start := t.position - 1
	for unicode.IsDigit(t.currentChar) {
		t.advance()
	}
	return Token{Type: TokenNumber, Value: t.input[start : t.position-1]}
}

// Tokenize the entire input
func (t *Tokenizer) Tokenize() []Token {
	var tokens []Token
	for {
		token := t.GetNextToken()
		tokens = append(tokens, token)
		if token.Type == TokenEOF {
			break
		}
	}
	return tokens
}
