package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 23
	fmt.Println("map:", m)

	v1 := m["key1"]

	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "key2")
	fmt.Println("delete map:", m)

	_, isExist := m["key2"]
	fmt.Println("prs:", isExist)

	value1, isExist := m["key1"]
	fmt.Println("key1 value:", value1, ",key1 status:", isExist)

	n := map[string]int{"foo": 1, "bar": 2, "abs": 3}
	fmt.Println("map:", n)
}
