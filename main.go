package main

import (
	"interpreter/repl"
	"os"
)

func main() {
	repl.StartInteraction(os.Stdin, os.Stdout)
}
