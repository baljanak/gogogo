package fib_test

import (
	"fmt"

	"github.com/baljanak/gogogo/fib"
)

func ExampleFibonacci() {

	ch := make(chan int)

	sig := make(chan int)
	defer close(sig)

	go fib.CallFib(ch, sig)

	for i := 0; i < 12; i++ {
		fmt.Println(<-ch)
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
