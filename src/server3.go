package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s : %s : %s \n", r.Method, r.URL, r.Proto)
	for name, headers := range r.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "%v : %v\n", name, header)
		}
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Address ： %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	fmt.Print("localhost:8000 listening ... \n")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
