package main

import (
	"fmt"
	"monkey_lang/token"
)

func main() {
	var tok token.Token

	tok.Literal = "teasd"
	tok.Type = "asd"

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", tok.Literal)
}
