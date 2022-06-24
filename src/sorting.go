package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"banana", "apple", "coffee"}

	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 3, 6}

	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorts:", s)
}
