package main

import (
	"io"
	"strings"
	"testing"
)

// BenchmarkBasic-16    	  460678	      2356 ns/op
// PASS
func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("6 7 8 9 1 11\n")

	for i := 0; i < b.N; i++ {
		task(in, io.Discard)

	}
}
