package ast

import "github.com/jonLaRose/go-interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node // A statement is also a Node
	statementNode()
}

type Expression interface {
	Node // An expression is also a node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct { 
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) expressionNode() {} // This is because in some other parts of the Monkey language Identifiers do produce values (like a function)

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier  // Although it doesn't produce any value it is convinient to keep all identifiers alike, and because identifiers can produce values in some scenarios, we keep it this way and not using another representation for an identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) statementNode() {}