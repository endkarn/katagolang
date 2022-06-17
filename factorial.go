package main

import "fmt"

func FirstFactorial(num int) int {
	result := 1
	for i := num; i > 1; i-- {
		result = result * i
	}
	return result
}

func main() {
	fmt.Println(FirstFactorial(4))
}
