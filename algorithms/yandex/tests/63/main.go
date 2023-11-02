package main

import "fmt"

// Дерево поиска
// Функция должна вернуть True, если дерево является деревом поиска, иначе - False.
// Набрать здесьОтправить файл
// BenchmarkBasic-16     	635281918	         5.726 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBasic2-16    	138166214	        26.30 ns/op	       0 B/op	       0 allocs/op
func main() {
	tree := &Node{
		value: 10,
		left: &Node{
			value: 5,
		},
		right: &Node{
			value: 15,
		},
	}

	fmt.Println(Solution(tree))

}

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) bool {
	if root == nil {
		return true
	}

	if root.left != nil {
		if root.value < root.left.value {
			return false
		}

		return Solution(root.left)
	}

	if root.right != nil {
		if root.value > root.right.value {
			return false
		}

		return Solution(root.right)
	}

	return true
}

func Solution2(root *Node) bool {
	var checkTree func(*Node, *int, *int) bool
	checkTree = func(tree *Node, min, max *int) bool {

		if tree == nil {
			if min == nil && max == nil {
				return false
			}
			return true
		}

		if max != nil && tree.value >= *max || min != nil && tree.value <= *min {
			return false
		}

		return checkTree(tree.left, min, &tree.value) && checkTree(tree.right, &tree.value, max)
	}

	return checkTree(root, nil, nil)
}
