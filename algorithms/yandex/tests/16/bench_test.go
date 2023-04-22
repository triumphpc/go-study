package main

import (
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("8\npush 10\npush 1\npush 9\npop\nget_max\npush 3\npop\nget_max")

	for i := 0; i < b.N; i++ {
		task(in)

	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("8\npush 10\npush 1\npush 9\npop\nget_max\npush 3\npop\nget_max")

	for i := 0; i < b.N; i++ {
		task2(in)
	}
}
