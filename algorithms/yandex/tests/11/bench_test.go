package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("10\n6 7 8 9 1 11 35 12 34 5\n")

	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("10\n6 7 8 9 1 11 35 12 34 5\n")

	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)
	}
}
