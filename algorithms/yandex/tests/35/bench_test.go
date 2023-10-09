package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("10\n1 1 5 7 3 5 6 8 8 9\n5\n1 2 4 1 2\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)

	}
}
