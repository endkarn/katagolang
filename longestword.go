package main

import (
	"fmt"
	"strings"
)

func LongestWord(sentence string) string {
	listWord := strings.Split(sentence, " ")
	result := ""
	longest := 0
	for _, value := range listWord {
		if len(value) > longest {
			longest = len(value)
			result = value
		}
	}
	return result
}

func LongestWordv2(sentence string) string {
	listWord := strings.Split(sentence, " ")

	//trim special char
	for i := 0; i < len(listWord); i++ {
		listWord[i] = strings.ReplaceAll(listWord[i], "&", "")
		listWord[i] = strings.ReplaceAll(listWord[i], "!", "")
	}

	//find result
	result := ""
	longest := 0
	for _, value := range listWord {
		if len(value) > longest {
			longest = len(value)
			result = value
		}
	}

	return result
}

func main() {
	fmt.Println(LongestWord("fun&!! time"))
	fmt.Println(LongestWordv2("fun&!! time"))
	fmt.Println(LongestWord("I love dogs"))
}
