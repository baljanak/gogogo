package main

import (
	"fmt"
	"net/http"
)

func fibonacci() func() uint64 {
	var n1, n2 uint64 = 1, 0
	return func() uint64 {
		defer func() {
			n1 = n1 + n2
			n2 = n1 - n2
		}()
		return n1 + n2
	}
}

func testFib() {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	testFib()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
