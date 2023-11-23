package main

import (
	"testing"
)

// BenchmarkBasic-16     	   32331	    112031 ns/op	   40025 B/op	     152 allocs/op
//BenchmarkBasic2-16    	     405	   8849337 ns/op	  825625 B/op	   14113 allocs/op
//BenchmarkBasic3-16    	  255759	     13862 ns/op	     637 B/op	       4 allocs/op

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring1("qwertyuiopsdfghjklxcvbnmwertyuiopqwertyuiopsdfghjklxcvbnmwertyuiopqwertyuiopsdfghjklxcvbnmwertyuiopxxxxx")

	}
}

func BenchmarkBasic2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring2("qwertyuiopsdfghjklxcvbnmwertyuiopqwertyuiopsdfghjklxcvbnmwertyuiopqwertyuiopsdfghjklxcvbnmwertyuiopxxxxx")

	}
}

func BenchmarkBasic3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring3("qwertyuiopsdfghjklxcvbnmwertyuiopqwertyuiopsdfghjklxcvbnmwertyuiopqwertyuiopsdfghjklxcvbnmwertyuiopxxxxx")

	}
}
