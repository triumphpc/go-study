package test

import (
	"testing"

	y "github.com/triumphpc/go-study/algorithms/yandex"
	tt2 "github.com/triumphpc/go-study/algorithms/yandex/tests/2"
)

var N = []float32{43, 43, 434, 546, 65, 76, 767, 767, 7567, 43, 43, 43, 434, 5, 67567, 567}
var K = 2
var runner = y.Runner{}

// BenchmarkSprintBasic-16    	    1621	   1155807 ns/op
func BenchmarkSprintBasic(b *testing.B) {
	tc := &tt2.TestRun{}
	for i := 0; i < b.N; i++ {
		runner.RunTestContestFloat(tc, K, N...)
	}
}

// BenchmarkSprintfBasic-16    	    1732	   1120682 ns/op
func BenchmarkSprintfBasic(b *testing.B) {
	tc := &TestRun{}
	for i := 0; i < b.N; i++ {
		runner.RunTestContestFloat(tc, K, N...)
	}

}
