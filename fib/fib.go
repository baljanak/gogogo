package fib

func fibonacci() func(ch chan int) {
	x, y := 0, 1
	return func(ch chan int) {
		defer func() {
			x, y = y, x+y
		}()
		ch <- x
		return
	}
}

func CallFib(n int, ch chan int) {
	f := fibonacci()
	for i := 0; i < n; i++ {
		f(ch)
	}
	close(ch)
}
