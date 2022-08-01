package main

// Замеры работы Mutex и Atomic

//

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func BenchmarkMutexMultipleReaders(b *testing.B) {
	var lastValue uint64
	var lock sync.RWMutex

	cfg := Config{
		a: []int{1, 2, 3, 4, 5, 6},
	}

	var wg sync.WaitGroup

	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				lock.RLock()
				atomic.SwapUint64(&lastValue, uint64(cfg.a[0]))
				lock.RUnlock()
			}
			wg.Done()
		}()
	}
}
