package main

import "fmt"

func main() {
	fmt.Println(Opposite(1))
}

func Opposite(i int) int {
	return i * -1
}
