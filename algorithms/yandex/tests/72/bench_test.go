package main

import (
	"testing"
)

// BenchmarkBasic
// BenchmarkBasic-16     	 1244456	       963.0 ns/op
// BenchmarkBasic2
// BenchmarkBasic2-16    	 2351632	       502.7 ns/op
func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		paths(5, 5)
	}
}

func BenchmarkBasic2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		paths2(5, 5)
	}
}
