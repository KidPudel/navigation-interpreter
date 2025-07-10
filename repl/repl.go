package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
	"io"
)

// REPL is the Read-Eval-Print Loop.

const PROMPT = ">>"

func StartInteraction(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		l := lexer.New(scanner.Text())

		for tok := l.ExtractToken(); tok.Type != token.EOF; tok = l.ExtractToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
