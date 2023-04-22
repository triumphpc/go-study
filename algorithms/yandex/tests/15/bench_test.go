package main

import (
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}

	for i := 0; i < b.N; i++ {
		Solution(&node0, "node3")
	}

}

func BenchmarkBasic2(b *testing.B) {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}

	for i := 0; i < b.N; i++ {
		Solution2(&node0, "node3")
	}

}
