package main

// Бинарное дерево. Лампочки
// Гоша повесил на стену гирлянду в виде бинарного дерева, в узлах которого находятся лампочки. У каждой лампочки есть
//своя яркость. Уровень яркости лампочки соответствует числу, расположенному в узле дерева. Помогите Гоше найти самую
//яркую лампочку в гирлянде, то есть такую, у которой яркость наибольшая.

func main() {
}

type Node struct {
	value int
	left  *Node
	right *Node
}

var max = 0

func Solution(tree *Node) int {
	if tree.value > max {
		max = tree.value
	}

	if tree.left != nil {
		Solution(tree.left)
	}
	if tree.right != nil {
		Solution(tree.right)
	}

	Solution(tree)

	return max
}
