package parser

import (
	"testing"

	"github.com/jonLaRose/go-interpreter/ast"
	"github.com/jonLaRose/go-interpreter/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatal("ParserProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got =%d", len(program.Statements))
	}

	tests := []struct{
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	literal := stmt.TokenLiteral()
	if literal != "let" {
		t.Errorf("s.TokenLiteral not 'let', got=%s", literal)
		return false
	}

	letStatement, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not *astLetStatemnt. got=%T", stmt)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("stmt.Name not '%s'. got=%s", name, letStatement.Name)
		return false
	}
	return true
}