package main

import "fmt"

//返回 A+B 和 A*B
func SumAndProduct(A, B, C int) (int, int, int) {
	return A + B, A * B, A * C
}

func main() {
	x := 3
	y := 4
	z := 5

	xPLUSy, xTIMESy, zTIMESy := SumAndProduct(x, y, z)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
	fmt.Printf("%d * %d = %d\n", x, z, zTIMESy)
}
