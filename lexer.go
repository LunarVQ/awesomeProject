package main

import (
	"unicode"
)

// Lexer represents a lexical analyzer
type Lexer struct {
	input string
	pos   int
	ch    byte
}

// NewLexer creates a new lexer instance
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character from the input
func (l *Lexer) readChar() {
	if l.pos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.pos]
	}
	l.pos++
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() Token {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}

	switch l.ch {
	case '+':
		l.readChar()
		return Token{Type: TokenPlus, Value: "+"}
	case '-':
		l.readChar()
		return Token{Type: TokenMinus, Value: "-"}
	case '*':
		l.readChar()
		return Token{Type: TokenStar, Value: "*"}
	case '/':
		l.readChar()
		return Token{Type: TokenSlash, Value: "/"}
	case '=':
		l.readChar()
		return Token{Type: TokenAssign, Value: "="}
	case 0:
		return Token{Type: TokenEOF, Value: ""}
	default:
		if unicode.IsLetter(rune(l.ch)) {
			return l.readVariable()
		}
		if unicode.IsDigit(rune(l.ch)) {
			return l.readNumber()
		}
		return Token{Type: TokenEOF, Value: ""}
	}
}

// readNumber reads a number from the input
func (l *Lexer) readNumber() Token {
	startPos := l.pos - 1
	for unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}
	return Token{Type: TokenNumber, Value: l.input[startPos : l.pos-1]}
}

// readVariable reads a variable from the input
func (l *Lexer) readVariable() Token {
	startPos := l.pos - 1
	for unicode.IsLetter(rune(l.ch)) || unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}
	return Token{Type: TokenVariable, Value: l.input[startPos : l.pos-1]}
}
