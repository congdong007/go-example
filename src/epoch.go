package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	seconds := now.Unix()
	nanos := now.UnixNano()

	p(now)

	millis := nanos / 1000000

	p(seconds)
	p(millis)
	p(nanos)

	p(time.Unix(seconds, 0))

	p(time.Unix(0, nanos))
}
