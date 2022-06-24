// the telephone alphabet:
package main

import (
    "fmt"
    "sort"
)

var (
    barVal = map[string]int{"alpha": 1, "bravo": 2, "charlie": 3,
        "delta": 4, "echo": 5, "foxtrot": 6,
        "golf": 7, "hotel": 8, "indio": 9,
        "juliet": 10, "kili": 11, "lima": 12}
)

func main() {
    fmt.Println("*************:")
    fmt.Println()
    fmt.Println("map:", barVal)
    fmt.Println()
    fmt.Println("unsorted:")
    fmt.Println()
    for k, v := range barVal {
        fmt.Printf("Key: %v, Value: %v / ", k, v)
    }

    keys := make([]string, len(barVal))
    i := 0
    for k := range barVal {
        keys[i] = k
        i++
    }
    sort.Strings(keys)
    fmt.Println()
    fmt.Println()
    fmt.Println("sorted:")
    fmt.Println()
    for _, k := range keys {
        fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
    }
}
