package main

import "fmt"

func FirstReverse(str string) string {
	result := ""
	for _, v := range str {
		result = string(v) + result
		//println(result)
	}
	return result
}

func SecondReverse(str string) string {
	result := ""

	//println(string(str[0]))
	//println(string(str[1]))
	//println(string(str[2]))
	//println(string(str[3]))
	//println(len(str))

	for i := len(str) - 1; i >= 0; i-- {
		result = result + string(str[i])
	}
	return result
}

func main() {
	fmt.Println(FirstReverse("karn"))
	fmt.Println(SecondReverse("karn"))
}
