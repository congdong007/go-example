package main
import "fmt"

type testInt func(int) bool //
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	pl := fmt.Println
	slice := []int{1, 2, 3, 4, 5, 7}
	pl("slice = ", slice)

	odd : = filter(slice,isOdd)
	pl("Odd elements of slice are:",odd)

	even := filter(slice,isEven)
	pl("Even elements of slice are:",even)
}
