package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

var keyword = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keyword[ident]; ok {
		return tok
	}
	return IDENT
}

// a constant list of TokenTypes
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y,...
	INT = "INT" // 1234522

	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords 
	FUNCTION = "FUNCTION"
	LET = "LET"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
)