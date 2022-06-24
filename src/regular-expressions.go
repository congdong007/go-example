package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println("FindStringIndex:", r.FindStringIndex("peach punch"))

	fmt.Println("FindStringSubmatch:", r.FindStringSubmatch("peach punch"))

	fmt.Println("FindStringSubmatchIndex:", r.FindStringSubmatchIndex("peach punch"))

	fmt.Println("FindAllString(-1):", r.FindAllString("peach punch pinch", -1))

	fmt.Println("FindAllStringSubmatchIndex:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println("FindAllString(2):", r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
