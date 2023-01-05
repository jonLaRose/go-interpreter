package parser

import (
	"fmt"

	"github.com/jonLaRose/go-interpreter/ast"
	"github.com/jonLaRose/go-interpreter/lexer"
	"github.com/jonLaRose/go-interpreter/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression // The ast.Expression parameter is the "left side" of the infix operator being parsed...
)

const (
	_ int = iota // The iota here is important because we want to set a precedence to the operations! 
	LOWEST
	EQUALS	// ==
	LESSGREATER // < or >
	SUM			// +
	PRODUCT		// *
	PREFIX		// -X or !X
	CALL 		// myfunction(X)
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token

	errors []error

	// the bellow maps are used to check if there's a parsing function associated with 'curToken.Type':
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []error{},
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns: make(map[token.TokenType]infixParseFn),
	}
	p.registerPrefix(token.IDENT, p.parseIdentifier)

	// Set both 'curToken' and 'peekToken'
	p.nextToken()
	p.nextToken()


	return p
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement{
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IDENT:
		return p.parseExpressionStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	
	// expectPeek will advance the current token to point to the identifier (let IDENT = ...) and the peek token to point to the assignment operator
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: currently we don't evaluate the expression. We'll skipo until reaching a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt;
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expr := p.parseExpression(LOWEST)
	stmt := &ast.ExpressionStatement{Token: p.curToken, Expression: expr}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////   Helper methods   ////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// type TokenType string
// type Token struct { Token TokenType, Literal string} 

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) addPeekError(t token.TokenType) {
	err := fmt.Errorf("expected next token to be '%s', got '%s' instead", t, p.peekToken.Type)
	p.errors = append(p.errors, err)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.addPeekError(t)
		return false
	}
}

func (p *Parser) registerPrefix(tokenType token.TokenType,  fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType,  fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}