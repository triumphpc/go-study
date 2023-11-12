// Сбалансированное дерево
// Гоше очень понравилось слушать рассказ Тимофея про деревья. Особенно часть про сбалансированные деревья.
// Он решил написать функцию, которая определяет, сбалансировано ли дерево.
//
// Дерево считается сбалансированным, если левое и правое поддеревья каждой вершины отличаются по высоте не больше, чем на единицу.
package main

import (
	"fmt"
	"math"
)

type TNode struct {
	value int
	left  *TNode
	right *TNode
}

func main() {
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

	fmt.Println(Solution1(tree))
}

func Solution1(root *TNode) bool {
	rl := deepCnt(root.left, 0)
	rr := deepCnt(root.right, 0)
	if math.Abs(float64(rl-rr)) > 1 {
		return false
	}

	return true
}

func deepCnt(root *TNode, deep int) int {
	deep++

	if root.left == nil && root.right == nil {
		return deep
	}

	var deepLeft, deepRight int
	if root.left != nil {
		deepLeft = deepCnt(root.left, deep)
	}

	if root.right != nil {
		deepRight = deepCnt(root.right, deep)
	}

	if deepLeft > deepRight {
		return deepLeft
	}

	return deepRight
}

func Solution2(root *TNode) bool {
	if root == nil {
		return true
	}
	var treeDeep func(*TNode, int) int
	treeDeep = func(tree *TNode, deep int) int {
		result := 0
		if tree == nil {
			return deep
		}
		deep++
		leftDeep := treeDeep(tree.left, deep)
		rightDeep := treeDeep(tree.right, deep)
		result = leftDeep
		if rightDeep > result {
			result = rightDeep
		}
		return result
	}
	return math.Abs(float64(treeDeep(root.left, 0)-treeDeep(root.right, 0))) <= 1 && Solution2(root.left) && Solution2(root.right)
}
