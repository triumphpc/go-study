// Балансировка АВЛ-дерева
package main

type TNode struct {
	value  int
	left   *TNode
	right  *TNode
	height int
}

func main() {

}

// Малое левое вращение
// Балансирует АВЛ-дерево
func smallLeftRotation(a *TNode) *TNode {
	b := a.right
	C := b.left
	R := b.right

	// Переусыновляем
	a.right = C
	b.left = a

	// Корректируем высоты в зависимости от того, равны ливысоты C и R
	if C.height == R.height {
		a.height--
		b.height++
	} else {
		b.height -= 2
	}

	return b
}

// Большое левое вращение
func bigLeftRotation(a *TNode) *TNode {
	// Задаем обозначения
	b := a.right
	c := b.left
	m := c.left
	n := c.right

	// Перевыставляем вершины
	a.right = m
	b.left = n
	c.left = a
	c.right = b

	// Корректируем высоты
	a.height -= 2
	b.height--
	c.height++

	return c
}
