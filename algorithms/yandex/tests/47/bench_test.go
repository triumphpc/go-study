package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("10\n0 0 1 0 1 1 1 0 0 0\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("10\n0 0 1 0 1 1 1 0 0 0\n")
	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)
	}
}
