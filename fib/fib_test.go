package fib_test

import (
	"fmt"

	"github.com/baljanak/gogogo/fib"
)

func ExampleFibonacci() {

	ch := make(chan int, 12)
	go fib.CallFib(ch)

	for i := range ch {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 1
	// 1
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
