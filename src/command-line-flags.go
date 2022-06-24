package main

import (
	"flag"
	"fmt"
)

func main() {
	pl := fmt.Println

	wordPtr := flag.String("word", "foo", "a string")

	numberPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var strVar string
	flag.StringVar(&strVar, "svar", "bar", "a string var")

	flag.Parse()

	pl("word:", *wordPtr)
	pl("numb:", *numberPtr)
	pl("fork:", *forkPtr)
	pl("svar:", strVar)
	pl("tail:", flag.Args())
}
