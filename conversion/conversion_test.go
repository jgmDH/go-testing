package conversion

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFmtSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprint(100)
		_ = s
	}
}

func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(100)
		_ = s
	}
}
