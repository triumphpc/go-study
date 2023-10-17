package main

import (
	"testing"
)

// BenchmarkWithoutPool
//BenchmarkWithoutPool-16    	    3169	    361842 ns/op	  160001 B/op	   10000 allocs/op
//BenchmarkWithPool
//BenchmarkWithPool-16       	    6285	    193857 ns/op	       0 B/op	       0 allocs/op
//PASS

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = new(Person)
			p.Name = "name"
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = personPool.Get().(*Person)
			p.Name = "new name"
			personPool.Put(p)
		}
	}
}
