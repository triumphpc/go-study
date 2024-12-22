package main

import (
	"testing"
)

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman(3)
	}
}

func BenchmarkIntToRoman2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman2(3)
	}
}
