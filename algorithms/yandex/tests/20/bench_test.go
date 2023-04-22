package main

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("10\nput 1\nput 2\nget\nsize\nget\nsize\nget\nget\nput 80\nsize\n")

	for i := 0; i < b.N; i++ {
		task(in)

	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("10\nput 1\nput 2\nget\nsize\nget\nsize\nget\nget\nput 80\nsize\n")

	for i := 0; i < b.N; i++ {
		task2(in, os.Stdout)
	}
}
