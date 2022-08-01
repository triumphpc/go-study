package main

import (
	"sort"
	"strconv"
	"testing"
)

func Benchmark_bubbleSortGeneric(b *testing.B) {
	var ss []string
	for i := 0; i < 1000; i++ {
		ss = append(ss, "str"+strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		bubbleSortGeneric(ss)
	}

}

func Benchmark_bubbleSortInterface(b *testing.B) {
	var ss []string
	for i := 0; i < 1000; i++ {
		ss = append(ss, "str"+strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		bubbleSortInterface(sort.StringSlice(ss))
	}

}
