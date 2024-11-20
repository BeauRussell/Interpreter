package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"Interpreter/token"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer}

	// Read the first 2 tokens to set both
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
