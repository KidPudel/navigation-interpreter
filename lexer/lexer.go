package lexer

import (
	"fmt"
	"interpreter/token"
)

// 1st phase in optimizing command, it goes through the input and builds tokens accordingly (side effects are possible)
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.shift()
	return l
}

// shift moves the lexer to the next character.
// returns true if can shift further
func (l *Lexer) shift() bool {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		return false
	}
	l.position = l.readPosition
	l.readPosition++
	l.ch = l.input[l.position]

	return true
}

func (l *Lexer) peek() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// extracts token and shifts the lexer to the next character.
func (l *Lexer) ExtractToken() token.Token {
	var tok token.Token

	l.eatWhitespaces()

	fmt.Printf("Lexer state: %d, %d, %q\n", l.position, l.readPosition, string(l.ch))
	switch l.ch {
	case '=':
		// since we have `==` for equality, we need to check next char
		if l.peek() == '=' {
			previousCh := l.ch
			l.shift() // shift to next char
			tok = token.NewToken(token.EQ, string(previousCh)+string(l.ch))
		} else {
			tok = token.NewToken(token.ASSIGN, string(l.ch))
		}
	case '+':
		tok = token.NewToken(token.PLUS, string(l.ch))
	case '-':
		// since we have `->` for passing result to the top of the command,
		if l.peek() == '>' {
			previousCh := l.ch
			l.shift() // shift to next char
			tok = token.NewToken(token.PASS_RIGHT, string(previousCh)+string(l.ch))
		} else {
			tok = token.NewToken(token.MINUS, string(l.ch))
		}
	case '!':
		// since we have `!=` for inequality, we need to check next char
		if l.peek() == '=' {
			previousCh := l.ch
			l.shift() // shift to next char
			tok = token.NewToken(token.NOT_EQ, string(previousCh)+string(l.ch))
		} else {
			tok = token.NewToken(token.BANG, string(l.ch))
		}
	case '*':
		tok = token.NewToken(token.ASTERISK, string(l.ch))
	case '/':
		tok = token.NewToken(token.SLASH, string(l.ch))
	case '<':
		if l.peek() == '-' {
			previousCh := l.ch
			l.shift() // shift to next char
			tok = token.NewToken(token.PASS_LEFT, string(previousCh)+string(l.ch))
		} else {
			tok = token.NewToken(token.LT, string(l.ch))
		}
	case '>':
		tok = token.NewToken(token.GT, string(l.ch))
	case ',':
		tok = token.NewToken(token.COMMA, string(l.ch))
	case ';':
		tok = token.NewToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = token.NewToken(token.LPAREN, string(l.ch))
	case ')':
		tok = token.NewToken(token.RPAREN, string(l.ch))
	case '{':
		tok = token.NewToken(token.LBRACE, string(l.ch))
	case '}':
		tok = token.NewToken(token.RBRACE, string(l.ch))
	case '[':
		tok = token.NewToken(token.LBRACKET, string(l.ch))
	case ']':
		tok = token.NewToken(token.RBRACKET, string(l.ch))
	case ':':
		tok = token.NewToken(token.COLON, string(l.ch))
	case 0:
		tok = token.NewToken(token.EOF, "")
	default:
		// dynamic character -> keyword, identifier, number, string, etc.
		identifier := l.checkIdentifier()
		if tokenType, ok := token.IsKeyword(identifier); ok {
			tok = token.NewToken(tokenType, identifier)
		} else {
			tok = token.NewToken(token.IDENT, identifier)
		}
	}

	l.shift()

	return tok
}

func (l *Lexer) eatWhitespaces() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.shift()
	}
}

func (l *Lexer) checkIdentifier() string {
	position := l.position
	for l.isLetter() {
		l.shift()
	}
	return l.input[position:l.position]

}

func (l *Lexer) isLetter() bool {
	return (l.ch >= 'a' && l.ch <= 'z') || (l.ch >= 'A' && l.ch <= 'Z') || l.ch == '_'
}
