package main

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("50\n")

	for i := 0; i < b.N; i++ {
		task(in)

	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("50\n")

	for i := 0; i < b.N; i++ {
		task2(in, os.Stdout)
	}
}

func BenchmarkBasic3(b *testing.B) {
	in := strings.NewReader("50\n")

	for i := 0; i < b.N; i++ {
		task3(in, os.Stdout)
	}
}
