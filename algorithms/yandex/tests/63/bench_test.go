package main

import (
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	tree := &Node{
		value: 10,
		left: &Node{
			value: 5,
		},
		right: &Node{
			value: 15,
		},
	}

	for i := 0; i < b.N; i++ {
		Solution(tree)
	}
}

func BenchmarkBasic2(b *testing.B) {
	tree := &Node{
		value: 10,
		left: &Node{
			value: 5,
		},
		right: &Node{
			value: 15,
		},
	}

	for i := 0; i < b.N; i++ {
		Solution2(tree)
	}
}
