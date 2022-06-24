package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {
	ret := plus(1, 3)
	fmt.Println("1+3 = ", ret)

	ret = plusplus(1, 3, 5)
	fmt.Println("1+3+5 = ", ret)
}
