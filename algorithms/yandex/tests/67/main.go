// Бинарное дерево
// Вставка вершины
// Дано BST. Надо вставить узел с заданным ключом. Ключи в дереве могут повторяться.
//На вход функции подаётся корень корректного бинарного дерева поиска и ключ, который надо вставить в дерево. Осуществите вставку этого ключа. Если ключ уже есть в дереве, то его дубликаты уходят в правого сына. Таким образом вид дерева после вставки определяется однозначно. Функция должна вернуть корень дерева после вставки вершины.
//Ваше решение должно работать за O(h), где h –— высота дерева.
//На рисунках ниже даны два примера вставки вершин в дерево.
// BenchmarkBasic-16     	   65751	    380480 ns/op	      24 B/op	       1 allocs/op
//BenchmarkBasic2-16    	   57348	    366225 ns/op	      24 B/op	       1 allocs/op

package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
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

	_ = insert1(tree, 10)
	printForward(tree)

}

func insert1(root *Node, key int) *Node {
	if root.value >= key {
		if root.right == nil {
			root.right = &Node{value: key}
			return root
		}

		return insert1(root.right, key)
	}

	if root.left == nil {
		root.left = &Node{value: key}
		return root
	}

	return insert1(root.left, key)
}

func printForward(node *Node) {
	fmt.Println(node.value)

	if node.left != nil {
		printForward(node.left)
	}

	if node.right != nil {
		printForward(node.right)
	}
}

func insert2(root *Node, key int) *Node {
	var recurentInsert func(root *Node, key int)
	recurentInsert = func(root *Node, key int) {
		if key < root.value {
			if root.left == nil {
				root.left = &Node{value: key}
			} else {
				recurentInsert(root.left, key)
			}
		} else {
			if root.right == nil {
				root.right = &Node{value: key}
			} else {
				recurentInsert(root.right, key)
			}
		}
	}
	recurentInsert(root, key)
	return root
}
