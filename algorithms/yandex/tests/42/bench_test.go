package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("8\n10\n0 0 0 1 3 3 5 10\n4 4 5 7 7 7 8 9 9 10")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}
