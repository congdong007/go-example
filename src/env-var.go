package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pl := fmt.Println
	os.Setenv("FOO", "1")

	pl("FOO:", os.Getenv("FOO"))
	pl("BAR:", os.Getenv("BAR"))
	pl()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		pl(pair[0] + ":" + os.Getenv(pair[0]))
	}
}
