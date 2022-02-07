package main

import "testing"

func Benchmark2Max(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		if -1 > i {
			r = -1
		} else {
			r = i
		}
	}
	Result = r
}

//Benchmark2Max-16    	1000000000	         0.3981 ns/op
//PASS
