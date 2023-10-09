package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("20\n1 2 3 4 5 5 5 5 5 5 6 6 7 7 8 8 8 9 9 9\n4\n")

	for i := 0; i < b.N; i++ {
		task(in, io.Discard)

	}
}

//func BenchmarkBasic2(b *testing.B) {
//	in := strings.NewReader("20\n1 2 3 4 5 5 5 5 5 5 6 6 7 7 8 8 8 9 9 9\n4\n")
//
//	for i := 0; i < b.N; i++ {
//		task2(in, io.Discard)
//
//	}
//}
