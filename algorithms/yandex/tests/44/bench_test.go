package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("10\n4 6 4 3 2 8 7 5 4 3\n9\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("10\n4 6 4 3 2 8 7 5 4 3\n9\n")
	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)
	}
}
