package main

import (
	"fmt"
)

// ASTNode represents any node in the AST.
type ASTNode interface {
	String() string
}

// VariableDeclarationNode represents a variable declaration with a type, name, and optional value.
type VariableDeclarationNode struct {
	Type  string
	Name  string
	Value ASTNode
}

func (n *VariableDeclarationNode) String() string {
	return fmt.Sprintf("%s %s = %s", n.Type, n.Name, n.Value)
}

// IdentifierNode represents an identifier.
type IdentifierNode struct {
	Name string
}

func (n *IdentifierNode) String() string {
	return n.Name
}

// NumberNode represents a numeric literal.
type NumberNode struct {
	Value string
}

func (n *NumberNode) String() string {
	return n.Value
}

// CharNode represents a character literal.
type CharNode struct {
	Value string
}

func (n *CharNode) String() string {
	return fmt.Sprintf("'%s'", n.Value)
}

// StringNode represents a string literal.
type StringNode struct {
	Value string
}

func (n *StringNode) String() string {
	return fmt.Sprintf("\"%s\"", n.Value)
}

// BinaryOpNode represents a binary operation.
type BinaryOpNode struct {
	Left     ASTNode
	Operator string
	Right    ASTNode
}

func (n *BinaryOpNode) String() string {
	return fmt.Sprintf("(%s %s %s)", n.Left, n.Operator, n.Right)
}

// PrintStmtNode represents a print statement.
type PrintStmtNode struct {
	Expression ASTNode
}

func (n *PrintStmtNode) String() string {
	if n.Expression != nil {
		return fmt.Sprintf("print(%s)", n.Expression.String())
	}
	return "print(<nil>)"
}

// IfStmtNode represents an if statement with an optional else clause.
type IfStmtNode struct {
	Condition ASTNode
	Body      []ASTNode
}

func (n *IfStmtNode) String() string {
	// Convert the body nodes to a string representation
	bodyStr := ""
	for _, stmt := range n.Body {
		bodyStr += stmt.String() + "; "
	}
	return fmt.Sprintf("if (%s) { %s }", n.Condition, bodyStr)
}

// Parser struct to manage the token stream.
type Parser struct {
	tokens  []Token
	current int
}

// NewParser initializes a new parser.
func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

// currentToken returns the current token.
func (p *Parser) currentToken() Token {
	if p.current < len(p.tokens) {
		return p.tokens[p.current]
	}
	return Token{Type: TokenEOF}
}

// advance moves to the next token.
func (p *Parser) advance() {
	if p.current < len(p.tokens) {
		p.current++
	}
}

// match checks if the current token matches the given type and value, then advances.
func (p *Parser) match(tokenType TokenType, value string) bool {
	if p.currentToken().Type == tokenType && p.currentToken().Value == value {
		p.advance()
		return true
	}
	return false
}

// Parse parses the entire input into a list of statements.
func (p *Parser) Parse() []ASTNode {
	var statements []ASTNode

	for p.currentToken().Type != TokenEOF {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
	}

	return statements
}

// parseStatement parses a statement (variable declaration, print, if, or expression).
func (p *Parser) parseStatement() ASTNode {
	switch p.currentToken().Type {
	case TokenKeyword:
		switch p.currentToken().Value {
		case "char", "string", "int":
			return p.parseVariableDeclaration()
		case "print":
			return p.parsePrintStatement()
		case "if":
			return p.parseIfStatement()
		}
	default:
		return p.parseExpression()
	}
	return nil
}

// parseVariableDeclaration parses a variable declaration statement.
func (p *Parser) parseVariableDeclaration() ASTNode {
	varType := p.currentToken().Value
	p.advance()

	if p.currentToken().Type != TokenIdentifier {
		return nil
	}
	varName := p.currentToken().Value
	p.advance()

	if !p.match(TokenSymbol, "=") {
		return nil
	}

	value := p.parseExpression()
	return &VariableDeclarationNode{Type: varType, Name: varName, Value: value}
}

// parsePrintStatement parses a print statement.
func (p *Parser) parsePrintStatement() ASTNode {
	p.advance() // Skip the "print" keyword.
	expr := p.parseExpression()
	return &PrintStmtNode{Expression: expr}
}

// parseIfStatement parses an if statement.
func (p *Parser) parseIfStatement() ASTNode {
	p.advance() // Skip "if"
	condition := p.parseExpression()
	if condition == nil {
		fmt.Println("Error: Invalid condition in if statement")
		return nil
	}

	body := p.parseBlock()
	return &IfStmtNode{Condition: condition, Body: body}
}

// parseBlock parses a block of statements inside braces.
func (p *Parser) parseBlock() []ASTNode {
	var body []ASTNode
	if p.match(TokenSymbol, "{") {
		for !p.match(TokenSymbol, "}") && p.currentToken().Type != TokenEOF {
			stmt := p.parseStatement()
			if stmt != nil {
				body = append(body, stmt)
			}
		}
	}
	return body
}

// parseExpression parses an expression (identifiers, numbers, or binary operations).
func (p *Parser) parseExpression() ASTNode {
	left := p.parsePrimary()

	for p.currentToken().Type == TokenSymbol && (p.currentToken().Value == "+" || p.currentToken().Value == "-" || p.currentToken().Value == ">") {
		op := p.currentToken().Value
		p.advance()
		right := p.parsePrimary()
		left = &BinaryOpNode{Left: left, Operator: op, Right: right}
	}

	return left
}

// parsePrimary parses primary expressions (identifiers, numbers, char, or string).
func (p *Parser) parsePrimary() ASTNode {
	token := p.currentToken()
	switch token.Type {
	case TokenIdentifier:
		p.advance()
		return &IdentifierNode{Name: token.Value}
	case TokenNumber:
		p.advance()
		return &NumberNode{Value: token.Value}
	case TokenCharLiteral:
		p.advance()
		return &CharNode{Value: token.Value}
	case TokenStringLiteral:
		p.advance()
		return &StringNode{Value: token.Value}
	default:
		p.advance()
		return nil
	}
}
