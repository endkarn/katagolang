package main

import "fmt"

func GetSize(w, h, d int) [2]int {
	return [2]int{calArea(w, d, h), calVol(w, d, h)}
}

func calVol(w, h, d int) int {
	return w * h * d
}

func calArea(w, h, d int) int {
	return 2*(w*d) + 2*(w*h) + 2*(d*h)
}

func main() {
	fmt.Println(GetSize(1, 1, 1))
	fmt.Println(GetSize(1, 2, 1))
	fmt.Println(GetSize(4, 2, 6))
}
