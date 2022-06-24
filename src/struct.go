package main

import "fmt"

type person struct {
	name string
	age  int
	addr string
}

func newPerson(name string) *person {
	p := person{name: name, addr: "ooooooooooo"}
	p.age = 34
	return &p
}

func main() {
	fmt.Println(person{"Bob", 30, "21 st-charles rue,kirkland,quebec"})
	fmt.Println(person{name: "Esson", age: 32, addr: "32 des-source rue,ddo,quebec"})
	fmt.Println(person{name: "Vicent"})

	fmt.Println(&person{name: "andy", age: 43, addr: "dddddd"})
	fmt.Println(newPerson("John"))
}
