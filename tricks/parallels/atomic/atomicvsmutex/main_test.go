package main

import (
	"sync"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	var v uint32 = 42
	var wg sync.WaitGroup
	wg.Add(2)
	go f(&v, &wg)
	go f(&v, &wg)
	wg.Wait()
}

func BenchmarkBasic2(b *testing.B) {
	var v uint32 = 42
	var wg sync.WaitGroup
	wg.Add(2)
	mx := sync.Mutex{}

	go f2(&v, &wg, &mx)
	go f2(&v, &wg, &mx)
	wg.Wait()
}
