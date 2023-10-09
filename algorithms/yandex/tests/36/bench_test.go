package main

import (
	"io"
	"strings"
	"testing"
)

// BenchmarkBasic-16     	 1643786	      2241 ns/op	    8325 B/op	       5 allocs/op
// BenchmarkBasic2-16    	 3200130	      1145 ns/op	    4146 B/op	       2 allocs/op
func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("4\n3 5 7 4\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("6\n5 3 7 2 8 3 6\n")
	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)
	}
}
