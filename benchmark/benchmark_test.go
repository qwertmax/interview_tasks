package benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkStrConcat(b *testing.B) {
	var sa = "hello"
	var sb = "world"
	var sc string
	for i := 0; i < b.N; i++ {
		sc += sa + sb
	}
	sa = sc
}

func BenchmarkFmtConcat(b *testing.B) {
	var sa = "hello"
	var sb = "world"
	var sc string
	for i := 0; i < b.N; i++ {
		sc = fmt.Sprintf("%s%s%s", sc, sa, sb)
	}
	sa = sc
}
