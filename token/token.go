package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

// tokens for navigation interface

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	IDENT TokenType = "IDENT" // add, foobar, x, y, ...
	INT   TokenType = "INT"   // 123456

	ASSIGN     TokenType = "="
	PLUS       TokenType = "+"
	MINUS      TokenType = "-"
	BANG       TokenType = "!"
	ASTERISK   TokenType = "*"
	SLASH      TokenType = "/"
	LT         TokenType = "<"
	GT         TokenType = ">"
	EQ         TokenType = "=="
	NOT_EQ     TokenType = "!="
	COMMA      TokenType = ","
	SEMICOLON  TokenType = ";"
	LPAREN     TokenType = "("
	RPAREN     TokenType = ")"
	LBRACE     TokenType = "{"
	RBRACE     TokenType = "}"
	LBRACKET   TokenType = "["
	RBRACKET   TokenType = "]"
	COLON      TokenType = ":"
	PASS_LEFT  TokenType = "<-" // indicates that we want to insert result from subcommand to the "base of the current command"
	PASS_RIGHT TokenType = "->" // indicates that we want to insert result from subcommand to the "top of the current command"

	// keywords
	GOTO TokenType = "GOTO"
	SHOW TokenType = "SHOW"
	IF   TokenType = "IF"
	THEN TokenType = "THEN"
)

var keywords = map[string]TokenType{
	"goto": GOTO,
	"show": SHOW,
	"if":   IF,
	"then": THEN,
}

func IsKeyword(ident string) (TokenType, bool) {
	tok, ok := keywords[ident]
	return tok, ok
}
