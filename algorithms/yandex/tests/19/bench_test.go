package main

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("10\n1\npush 1\nsize\npush 3\nsize\npush 1\npop\npush 1\npop\npush 3\npush 3\n")

	for i := 0; i < b.N; i++ {
		task(in)

	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("10\n1\npush 1\nsize\npush 3\nsize\npush 1\npop\npush 1\npop\npush 3\npush 3\n")

	for i := 0; i < b.N; i++ {
		task2(in, os.Stdout)
	}
}
