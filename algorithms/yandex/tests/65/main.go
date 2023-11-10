package main

import "fmt"

// Варианты обхода деревьев

func main() {

}

type Node struct {
	value int
	child []Node
}

type BNode struct {
	value int
	left  *BNode
	right *BNode
}

// Прямой обход
func printForward(node Node) {
	fmt.Println(node.value)

	for i := range node.child {
		printForward(node.child[i])
	}
}

// Обратный обход
func printReversed(node Node) {
	for i := range node.child {
		printForward(node.child[i])
	}

	fmt.Println(node.value)
}

// Центрированный обход для бинарного
func printLMR(node BNode) {
	if node.left != nil {
		printLMR(*node.left)
	}

	fmt.Println(node.value)

	if node.right != nil {
		printLMR(*node.right)
	}
}
