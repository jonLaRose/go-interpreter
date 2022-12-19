package parser

import (
	"github.com/jonLaRose/go-interpreter/ast"
	"github.com/jonLaRose/go-interpreter/lexer"
	"github.com/jonLaRose/go-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Set both 'curToken' and 'peekToken'
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}