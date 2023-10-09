package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("5\n4 3 9 2 1\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("5\n4 3 9 2 1\n")
	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)
	}
}
