package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("14\n5 3 74 2 8 34 6 5 3 74 2 8 34 6\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("14\n5 3 74 2 8 34 6 5 3 74 2 8 34 6\n")
	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)
	}
}
