// Выведи диапазон
// Напишите функцию, которая будет выводить по неубыванию все ключи от L до R включительно в заданном бинарном дереве поиска.
// Ключи в дереве могут повторяться. Решение должно иметь сложность O(h+k), где h –— глубина дерева, k — число элементов в ответе.
// В данной задаче если в узле содержится ключ x, то другие ключи, равные x, могут быть как в правом, так и в левом поддереве данного узла. (Дерево строил стажёр, так что ничего страшного).
package main

import "fmt"

// do not declare Node in your submit-file

type Node struct {
	value int
	left  *Node
	right *Node
}

func printRange(root *Node, left int, right int) {
	if root == nil {
		return
	}
	if root.value >= left {
		printRange(root.left, left, right)
	}

	if root.value >= left && root.value <= right {
		fmt.Printf("%d\n", root.value)
	}

	if root.value <= right {
		printRange(root.right, left, right)
	}
}
