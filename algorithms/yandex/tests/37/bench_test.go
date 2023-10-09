package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("7\n5 3 7 2 8 3 6\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}
