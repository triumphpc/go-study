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
			left: &Node{
				value: 10,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		insert1(tree, 20)
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
			left: &Node{
				value: 10,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		insert2(tree, 20)
	}
}
