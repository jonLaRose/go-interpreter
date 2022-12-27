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
	checkParserErrors(t, p)

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

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 9919199;
`
	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()

	checkParserErrors(t, parser)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statemnts does not contain 3 statements. got %d", len(program.Statements))
	}

	for _, statement := range program.Statements {
		retStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got = %T", retStatement)
			continue
		}
		if retStatement.TokenLiteral() != "return" {
			t.Errorf("retStatement.TokenLiter not 'return', got %q", retStatement.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser){
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))

	for _, err := range errors {
		t.Errorf("parser error: %q", err.Error())
	}
	t.FailNow()
}