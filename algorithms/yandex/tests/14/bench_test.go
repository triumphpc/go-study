package main

import (
	"testing"
)

// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
//BenchmarkBasic-16     	 4604036	       712.7 ns/op	    4147 B/op	       2 allocs/op
//BenchmarkBasic2-16    	 4523956	       721.7 ns/op	    4147 B/op	       2 allocs/op
//PASS

func BenchmarkBasic(b *testing.B) {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}

	for i := 0; i < b.N; i++ {
		Solution(&node0, 2)
	}

}

func BenchmarkBasic2(b *testing.B) {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}

	for i := 0; i < b.N; i++ {
		Solution2(&node0, 2)
	}

}
