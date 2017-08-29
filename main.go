package main

import (
	"fmt"
	"net/http"

	"github.com/baljanak/gogogo/fib"
)

func testFib() {
	f := fib.Fibonacci()
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
