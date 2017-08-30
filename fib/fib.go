package fib

func fibonacci() func(ch chan int) {
	x, y := 0, 1
	return func(ch chan int) {
		defer func() {
			x, y = y, x+y
		}()
		ch <- x
	}
}

func CallFib(ch chan int) {
	f := fibonacci()
	for i := 0; i < cap(ch); i++ {
		f(ch)
	}
	close(ch)
}
