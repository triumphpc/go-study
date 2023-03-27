package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
//BenchmarkBasic-16     	 4604036	       712.7 ns/op	    4147 B/op	       2 allocs/op
//BenchmarkBasic2-16    	 4523956	       721.7 ns/op	    4147 B/op	       2 allocs/op
//PASS

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("20\n1 2 0 4 5 6 7 8 0 10 11 12 13 14 15 16 17 0 0 20")
	var out bytes.Buffer
	foo := bufio.NewWriter(&out)

	for i := 0; i < b.N; i++ {
		task(in, foo)

	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("20\n1 2 0 4 5 6 7 8 0 10 11 12 13 14 15 16 17 0 0 20")
	var out bytes.Buffer
	foo := bufio.NewWriter(&out)

	for i := 0; i < b.N; i++ {
		task2(in, foo)

	}
}
