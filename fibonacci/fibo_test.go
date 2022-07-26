package fibo

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	lote := []uint{1, 2, 5, 3, 4}

	for _, value := range lote {
		for n := 0; n < b.N; n++ {
			Fibonacci(value)
		}
	}
}
