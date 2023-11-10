// Бинарное дерево
// Удаление узла
// Дано бинарное дерево поиска, в котором хранятся ключи. Ключи — уникальные целые числа. Найдите вершину с заданным ключом и удалите её из дерева так, чтобы дерево осталось корректным бинарным деревом поиска. Если ключа в дереве нет, то изменять дерево не надо.
// На вход вашей функции подаётся корень дерева и ключ, который надо удалить. Функция должна вернуть корень изменённого дерева. Сложность удаления узла должна составлять O(h), где h –— высота дерева.
// Создавать новые вершины (вдруг очень захочется) нельзя.
package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	tree := &Node{
		value: 8,
		left: &Node{
			value: 4,
			left: &Node{
				value: 2,
				left: &Node{
					value: 1,
				},
				right: &Node{
					value: 3,
				},
			},
			right: &Node{
				value: 6,
				left: &Node{
					value: 5,
				},
				right: &Node{
					value: 7,
				},
			},
		},
		right: &Node{
			value: 12,
			left: &Node{
				value: 10,
				left: &Node{
					value: 9,
				},
				right: &Node{
					value: 11,
				},
			},
			right: &Node{
				value: 14,
				left: &Node{
					value: 13,
				},
				right: &Node{
					value: 15,
				},
			},
		},
	}

	_ = remove1(tree, 12)
	printForward(tree)

}

func remove1(root *Node, key int) *Node {
	if root.value == key {
		// Получаем максимальное значение всех правых листьев левого ствола удаляемого узла
		max, n := getMax(root.left)
		root.left = n
		root.value = max

		return root
	}

	if root.value < key {
		if root.right == nil {
			return root
		}

		return remove1(root.right, key)
	}

	if root.left == nil {
		return root
	}

	return remove1(root.left, key)
}

func getMax(node *Node) (max int, n *Node) {
	if node.right != nil {
		max, node.right = getMax(node.right)
		return max, node
	}

	max = node.value
	node = nil

	return max, node
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
