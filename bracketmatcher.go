package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(BracketMatcher("(coder)(byte))"))
	fmt.Println(BracketMatcher("(c(oder)) b(yte)"))
}

func BracketMatcher(s string) string {
	openBracket := strings.Count(s, "(")
	closeBracket := strings.Count(s, ")")
	if openBracket != closeBracket {
		return "0"
	}
	return "1"
}
