package fib

func Fibonacci() func() uint64 {
	var n1, n2 uint64 = 1, 0
	return func() uint64 {
		defer func() {
			n1 = n1 + n2
			n2 = n1 - n2
		}()
		return n1 + n2
	}
}
