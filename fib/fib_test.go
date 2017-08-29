package fib_test

import (
	"fmt"

	"github.com/baljanak/gogogo/fib"
)

func ExampleFibonacci() {
	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	// Output: 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// 55
	// 89
}
