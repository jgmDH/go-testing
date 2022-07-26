package fibo

/* func Fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
*/

func Fibonacci(n uint) uint {
	if n < 2 {
		return n
	}

	var a, b uint
	b = 1

	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}
