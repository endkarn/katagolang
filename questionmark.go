package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func QuestionsMarks(str string) string {

	for index, aRune := range str {

		// find first number of pair
		if unicode.IsDigit(aRune) {
			startIndexOfPair := index
			endIndexOfPair := 0

			for i := index + 1; i < len(str); i++ {
				value := rune(str[i])
				// find second number of pair
				if unicode.IsDigit(value) {
					endIndexOfPair = i
					break
				}
			}

			if isThePairContainQuestionMarks(str, startIndexOfPair, endIndexOfPair) && isThePairAddUpToTen(str, startIndexOfPair, endIndexOfPair) {
				return "true"
			} else {
				return "false"
			}
		}
	}
	return "false"
}

func isThePairAddUpToTen(str string, pair int, pair2 int) bool {
	firstNumber, _ := strconv.Atoi(string(str[pair]))
	secondNumber, _ := strconv.Atoi(string(str[pair2]))
	sum := firstNumber + secondNumber
	if sum == 10 {
		return true
	}
	return false
}

func isThePairContainQuestionMarks(str string, pair int, pair2 int) bool {
	wordBetween := str[pair+1 : pair2]
	totalQuestionMark := strings.Count(wordBetween, "?")
	if totalQuestionMark == 3 {
		return true
	}
	return false
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(QuestionsMarks("9???1???9??1???9"))

}
