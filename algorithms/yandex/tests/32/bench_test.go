package main

import (
	"io"
	"strings"
	"testing"
)

// BenchmarkBasic-16     	 1531215	      2377 ns/op	    8351 B/op	       7 allocs/op
// BenchmarkBasic2-16    	 1526366	      2369 ns/op	    8351 B/op	       7 allocs/op
func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("10\n3 5 6 4 3 2 3 7 6 100\n")

	for i := 0; i < b.N; i++ {
		task(in, io.Discard)

	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("10\n3 5 6 4 3 2 3 7 6 100\n")

	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)

	}
}
