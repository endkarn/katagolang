package main

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	message := "Hello, <NAME> how are you doing today?"
	return strings.ReplaceAll(message, "<NAME>", name)
}

func main() {
	fmt.Println(Greet("Ryan"))
}
