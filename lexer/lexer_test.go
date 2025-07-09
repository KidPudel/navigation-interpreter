package lexer

import (
	"interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `= + - ! * / < > == != , ; ( ) { } [ ] : <- -> goto show if then`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.BANG, "!"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.COLON, ":"},
		{token.PASS_LEFT, "<-"},
		{token.PASS_RIGHT, "->"},

		// keywords
		{token.GOTO, "goto"},
		{token.SHOW, "show"},

		// identifiers and literals
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.ExtractToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - wrong token type. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
