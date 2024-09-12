package main

import (
	"unicode"
)

// Lexer represents a lexical analyzer
type lexer struct {
	input string
	pos   int
	ch    byte
}

// NewLexer creates a new lexer instance
func Newlexer(input string) *lexer {
	L := &lexer{input: input}
	L.readChar()
	return L
}

// readChar reads the next character from the input
func (L *lexer) readChar() {
	if L.pos >= len(L.input) {
		L.ch = 0
	} else {
		L.ch = L.input[L.pos]
	}
	L.pos++
}

// NextToken returns the next token from the input
func (L *lexer) Nexttoken() Token {
	for unicode.IsSpace(rune(L.ch)) {
		L.readChar()
	}

	switch L.ch {
	case '+':
		L.readChar()
		return Token{Type: TokenPlus, Value: "+"}
	case '-':
		L.readChar()
		return Token{Type: TokenMinus, Value: "-"}
	case '*':
		L.readChar()
		return Token{Type: TokenStar, Value: "*"}
	case '/':
		L.readChar()
		return Token{Type: TokenSlash, Value: "/"}
	case '=':
		L.readChar()
		return Token{Type: TokenAssign, Value: "="}
	case 0:
		return Token{Type: TokenEOF, Value: ""}
	default:
		if unicode.IsLetter(rune(L.ch)) {
			return L.readVariable()
		}
		if unicode.IsDigit(rune(L.ch)) {
			return L.readNumber()
		}
		return Token{Type: TokenEOF, Value: ""}
	}
}

// readNumber reads a number from the input
func (L *lexer) readNumber() Token {
	startPos := L.pos - 1
	for unicode.IsDigit(rune(L.ch)) {
		L.readChar()
	}
	return Token{Type: TokenNumber, Value: L.input[startPos : L.pos-1]}
}

// readVariable reads a variable from the input
func (L *lexer) readVariable() Token {
	startPos := L.pos - 1
	for unicode.IsLetter(rune(L.ch)) || unicode.IsDigit(rune(L.ch)) {
		L.readChar()
	}
	return Token{Type: TokenVariable, Value: L.input[startPos : L.pos-1]}
}
