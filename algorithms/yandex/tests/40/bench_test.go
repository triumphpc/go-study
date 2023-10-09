package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("6\n1 1 1 2 2 3\n1\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}
