package lexer

import (
	"unicode/utf8"

	"github.com/jonLaRose/go-interpreter/token"
)

type Lexer struct {
	input string
	position int // current position in input (points to current rune)
	readPosition int // current reading position in input (after current rune)
	r rune // current rune under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token{
	var tok token.Token

	l.skipWhitespace()

	switch l.r {
	case '=':
		if l.peekChar() == '=' {
			r := l.r
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(r)+string(l.r)}
		} else {
			tok = newToken(token.ASSIGN, l.r)
		}
	case '+':
		tok = newToken(token.PLUS, l.r)
	case ';':
		tok = newToken(token.SEMICOLON, l.r)
	case '(':
		tok = newToken(token.LPAREN, l.r)
	case ')':
		tok = newToken(token.RPAREN, l.r)
	case ',':
		tok = newToken(token.COMMA, l.r)
	case '{':
		tok = newToken(token.LBRACE, l.r)
	case '}':
		tok = newToken(token.RBRACE, l.r)
	case '>':
		tok = newToken(token.GT, l.r)
	case '!':
		if l.peekChar() == '=' {
			r := l.r
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(r)+string(l.r)}
		} else {
			tok = newToken(token.BANG, l.r)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.r)
	case '<':
		tok = newToken(token.LT, l.r)
	case '-':
		tok = newToken(token.MINUS, l.r)
	case '/':
		tok = newToken(token.SLASH, l.r)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.r) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.r) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.r)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.r == ' ' || l.r == '\t' || l.r == '\n' || l.r == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.r) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.r) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_'
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func newToken(tokenType token.TokenType, r rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(r)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.position = l.readPosition
		l.r = 0
	} else {
		runeValue, width := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.position = l.readPosition
		l.readPosition += width
		l.r = runeValue
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
		return r
	}
}