package main

import (
	"io"
	"strings"
	"testing"
)

// BenchmarkBasic-16    	  578454	      1927 ns/op
func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("6789\n")

	for i := 0; i < b.N; i++ {
		task(in, io.Discard)

	}
}

// BenchmarkBasic2-16    	 1268401	       955.5 ns/op
// BenchmarkBasic2-16    	 1240272	       965.6 ns/op

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("6789\n")

	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)

	}
}
