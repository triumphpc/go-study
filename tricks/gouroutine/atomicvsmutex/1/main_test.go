package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Замеры работы Mutex и Atomic
//RWMutex
//  go test -run none -bench . -benchmem -benchtime 3s
// 1000000000	         0.0000130 ns/op	       0 B/op	       0 allocs/op
// vs atomic
// 1000000000	         0.0000142 ns/op	       0 B/op	       0 allocs/op

type Config struct {
	a []int
}

func BenchmarkMutexMultipleReaders(b *testing.B) {
	var lastValue uint64
	//var lock sync.RWMutex

	cfg := Config{
		a: []int{1, 2, 3, 4, 5, 6},
	}

	var wg sync.WaitGroup

	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				//lock.RLock()
				//lastValue = uint64(cfg.a[0])

				atomic.SwapUint64(&lastValue, uint64(cfg.a[0]))
				//lock.RUnlock()
			}
			wg.Done()
		}()
	}

	fmt.Println(lastValue)
}
