package fib

func fibonacci() func(ch, sig chan int) bool {
	x, y := 0, 1
	return func(ch, sig chan int) bool {
		select {
		case ch <- x:
			x, y = y, x+y
			return true
		case <-sig:
			close(ch)
			return false
		}
	}
}

func CallFib(ch, sig chan int) {
	f := fibonacci()
	for {
		if !f(ch, sig) {
			return
		}
	}
}
