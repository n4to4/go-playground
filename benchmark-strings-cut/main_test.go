// original: https://zenn.dev/mattn/articles/01f258a5127ef8

package main

import (
	"strings"
	"testing"
)

// gotip test -bench=.
/*
goos: darwin
goarch: amd64
pkg: benchmark
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkCut-16         130403316                9.117 ns/op
BenchmarkSplitN-16      19267407                56.52 ns/op
PASS
ok      benchmark       3.282s
*/

func BenchmarkCut(b *testing.B) {
	s := "FooBarBaz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lhs, rhs, ok := strings.Cut(s, "Bar")
		if !ok || lhs != "Foo" || rhs != "Baz" {
			b.Fatal("something wrong")
		}
	}
}

func BenchmarkSplitN(b *testing.B) {
	s := "FooBarBaz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tok := strings.SplitN(s, "Bar", 2)
		if len(tok) != 2 || tok[0] != "Foo" || tok[1] != "Baz" {
			b.Fatal("something wrong")
		}
	}
}
