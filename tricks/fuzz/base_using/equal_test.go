package main

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'}) {
		t.Error("expected true, got false")
	}
}

// It's fuzz tests
// gotip test -fuzz=. -fuzztime=5s .                                                                  triumphpc@MacBook-Pro-triumphpc
//fuzz: elapsed: 0s, gathering baseline coverage: 0/1 completed
//fuzz: elapsed: 0s, gathering baseline coverage: 1/1 completed, now fuzzing with 16 workers
//fuzz: elapsed: 3s, execs: 728742 (242911/sec), new interesting: 7 (total: 8)
//fuzz: elapsed: 5s, execs: 1212952 (229300/sec), new interesting: 7 (total: 8)
//PASS
//ok  	github.com/triumphpc/go-study/tricks/fuzz/1	5.323s

func FuzzEqual(f *testing.F) {
	// target, can be only one per test
	// values of a and b will be auto-generated
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
