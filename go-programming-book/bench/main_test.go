package main

import "testing"

// go test -bench .
// go test -benchmem -bench .

func BenchmarkDoSomething(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoSomething()
	}
}
