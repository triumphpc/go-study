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
		a.height -= 1
		b.height += 1
	} else {
		b.height -= 2
	}

	return b
}
