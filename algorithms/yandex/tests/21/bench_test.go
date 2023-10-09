package main

import (
	"strings"
	"testing"
)

// BenchmarkBasic-16     	 1889005	      1946 ns/op	    8240 B/op	       3 allocs/op
// BenchmarkBasic4-16    	 1856695	      1948 ns/op	    8240 B/op	       3 allocs/op
func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("50\n")

	for i := 0; i < b.N; i++ {
		task(in)

	}
}

//
//func BenchmarkBasic2(b *testing.B) {
//	in := strings.NewReader("50\n")
//
//	for i := 0; i < b.N; i++ {
//		task2(in, os.Stdout)
//	}
//}
//
//func BenchmarkBasic3(b *testing.B) {
//	in := strings.NewReader("50\n")
//
//	for i := 0; i < b.N; i++ {
//		task3(in, os.Stdout)
//	}
//}

func BenchmarkBasic4(b *testing.B) {
	in := strings.NewReader("50\n")

	for i := 0; i < b.N; i++ {
		task4(in)

	}
}
