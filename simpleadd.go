package main

import "fmt"

func SimpleAdding(num int) int {
	// code goes here
	result := 0
	if num >= 1 && num <= 1000 {
		for i := 0; i <= num; i++ {
			result += i
		}
	}
	return result
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(SimpleAdding(12))
	fmt.Println(SimpleAdding(4))
	fmt.Println(SimpleAdding(4999))
	fmt.Println(SimpleAdding(140))

}
