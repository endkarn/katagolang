package main

import "fmt"

func RemoveChar(word string) string {
	charCount := len(word)
	result := word[1 : charCount-1]
	return result
}

func main() {
	fmt.Println(RemoveChar("eloquent"))
}
