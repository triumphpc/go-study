// Псевдокод подсчета размера поддерева

package main

type TNode struct {
	value int
	left  *TNode
	right *TNode
	size  int
}

func main() {

}

// Простое решение
func calcSize(root *TNode) int {
	if root == nil {
		return 0
	}

	return calcSize(root.left) + calcSize(root.right)
}

// Быстрое решение
func calcSizeBest(root *TNode, k int) int {
	// Вычисляем размер левого поддерева, с учетом того, что оно может быть пустым
	leftSize := 0
	if root.left != nil {
		leftSize = root.left.size
	}

	// Если слева k вершин,то искомая вершина - корень
	if leftSize == k {
		return root.value
	}

	// Если левое дерево слишком мало, то продолжим поиск уже в правом поддереве
	if leftSize < k {
		return calcSizeBest(root.right, k-leftSize-1)
	}

	// Иначе продолжаем поиск в левом поддереве
	return calcSizeBest(root.left, k)
}
