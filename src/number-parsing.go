package main

import (
	"fmt"
	"strconv"
)

func main() {
	pl := fmt.Println

	f, _ := strconv.ParseFloat("1.234", 64)
	pl(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	pl(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	pl(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	pl(u)

	k, _ := strconv.Atoi("135")
	pl(k)

	_, e := strconv.Atoi("wat")
	pl(e)
}
