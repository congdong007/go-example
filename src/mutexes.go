package main

import (
	"fmt"
	"sync"
)

type Container struct {
	sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.Lock()
	defer c.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
		fmt.Printf("%s: done!\n\n", name)
	}

	wg.Add(3)

	doIncrement("a", 10000)
	doIncrement("c", 50000)
	doIncrement("b", 100000)

	wg.Wait()
	fmt.Println(c.counters)
}
