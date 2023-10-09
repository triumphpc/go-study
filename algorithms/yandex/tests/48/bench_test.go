package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("mxyskaoghi\nqodfrgmslc\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("mxyskaoghi\nqodfrgmslc\n")
	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)
	}
}
