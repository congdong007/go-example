package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	//++
	i++
	fmt.Println("i = ", i)
	for j := 7; j <= 100; j++ {
		fmt.Println(j)
	}
	//单入库单出口
	n := 1
	for {
		fmt.Println("loop")
		n++
		if n == 10 {
			break
		}
	}
	i++
	fmt.Println("i = ", i)
	for n := 0; n <= 19; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	i++
	fmt.Println("i = ", i)
}
