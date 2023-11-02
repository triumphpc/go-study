package main

import "fmt"

// Бинарное дерево. Бинарный поиск

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

	fmt.Println(findNode(tree, 15))
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func findNode(tree *Node, value int) *Node {
	if tree == nil {
		return nil
	}

	if tree.value == value {
		return tree
	}

	if tree.value > value {
		return findNode(tree.left, value)
	}

	return findNode(tree.right, value)
}
