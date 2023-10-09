package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("")

	for i := 0; i < b.N; i++ {
		task(in, io.Discard)

	}
}
