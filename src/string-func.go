package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:    ", s.Contains("test", "es"))
	p("Count:    ", s.Count("test", "t"))
	p("HasPrefix:    ", s.HasPrefix("test", "te"))
	p("HasSuffix:    ", s.HasSuffix("test", "st"))
	p("Index:    ", s.Index("test", "es"))
	p("Join:    ", s.Join([]string{"test", "ing"}, "-"))
	p("Repeat:    ", s.Repeat("test", 3))
	p("Replace:    ", s.Replace("test", "es", "oo", -1))
	p("Replace:    ", s.Replace("test", "es", "oo", 1))
	p("Split:    ", s.Split("t-e-s-t", "-"))
	p("ToLower:    ", s.ToLower("Test"))
	p("ToUpper:    ", s.ToUpper("test"))

	p()

	p("Len:    ", len("test"))
	p("Char:    ", "test"[1])
}
