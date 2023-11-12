package main

import (
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	tree := &TNode{
		value: 8,
		left: &TNode{
			value: 4,
			left: &TNode{
				value: 2,
				left: &TNode{
					value: 1,
				},
				right: &TNode{
					value: 3,
				},
			},
			right: &TNode{
				value: 6,
				left: &TNode{
					value: 5,
				},
				right: &TNode{
					value: 7,
				},
			},
		},
		right: &TNode{
			value: 12,
			left: &TNode{
				value: 10,
				left: &TNode{
					value: 9,
					left: &TNode{
						value: 8,
						left: &TNode{
							value: 8,
						},
					},
				},
			},
			right: &TNode{
				value: 14,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		Solution1(tree)
	}
}

func BenchmarkBasic2(b *testing.B) {
	tree := &TNode{
		value: 8,
		left: &TNode{
			value: 4,
			left: &TNode{
				value: 2,
				left: &TNode{
					value: 1,
				},
				right: &TNode{
					value: 3,
				},
			},
			right: &TNode{
				value: 6,
				left: &TNode{
					value: 5,
				},
				right: &TNode{
					value: 7,
				},
			},
		},
		right: &TNode{
			value: 12,
			left: &TNode{
				value: 10,
				left: &TNode{
					value: 9,
					left: &TNode{
						value: 8,
						left: &TNode{
							value: 8,
						},
					},
				},
			},
			right: &TNode{
				value: 14,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		Solution2(tree)
	}
}
