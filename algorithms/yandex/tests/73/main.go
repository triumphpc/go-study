// Реализация дерева решений - балансировка с использованием левого и правого вращения
// Высота левого поддерева больше высоты правого поддерева?
// > Да
//		Используем правое вращение
//			Высота правого поддерева левого ребенка больше высоты правого ребенка корня?
//			> Нет
//				Малове правое вращение
//			> Да
//				Большое правое вращение
// > Нет
//		Используем левое вращение
//			Высота левого поддерева правого ребенка больше высоты левого ребенка?
//			> Нет
//				Малое левое вращение
//			Да
//				Большое левое вращение

package main

import "C"
import "math"

type TNode struct {
	value  int
	left   *TNode
	right  *TNode
	height int
}

func main() {
	tree := &TNode{
		height: 3,
		left: &TNode{
			height: 2,
			left: &TNode{
				height: 0,
			},
			right: &TNode{
				height: 1,
				left: &TNode{
					height: -1,
				},
				right: &TNode{
					height: 0,
				},
			},
		},
		right: &TNode{
			height: 0,
		},
	}

}

func rotate(vertex *TNode) *TNode {

	//1. Высота левого поддерева больше высоты правого поддерева меньше 2
	if math.Abs(float64(vertex.left.height)-float64(vertex.right.height)) < 2 {
		return vertex
	}

	// 2. Если высота правого поддерева больше высоты левого поддерева на 2
	// На нужны левые вращения
	if vertex.left.height-vertex.right.height == -2 {
		b := vertex.right
		r := b.right
		c := b.left

		if c.height <= r.height {
			// Нужно малое левок дерево вращения
			return smallLeftRotation(vertex)
		} else { // Нужно большое левое вращение
			return bigLeftRotation(vertex)
		}
	} else {
		b := vertex.left
		c := b.right
		l := b.left

		if c.height <= l.height {
			// Нужно правое малое дерево вращения
			return smallRightRotation(vertex)
		} else { // Нужно большое правое вращение
			return bigRightRotation(vertex)
		}
	}
}

// Малое левое вращение
// Балансирует АВЛ-дерево
func smallLeftRotation(a *TNode) *TNode {
	b := a.right
	c := b.left
	r := b.right

	// Переусыновляем
	a.right = c
	b.left = a

	// Корректируем высоты в зависимости от того, равны ливысоты C и R
	if c.height == r.height {
		a.height--
		b.height++
	} else {
		b.height -= 2
	}

	return b
}

func smallRightRotation(a *TNode) *TNode {
	b := a.left
	c := b.right
	l := b.left

	// Переусыновляем
	a.left = c
	b.right = a

	// Корректируем высоты в зависимости от того, равны ливысоты C и R
	if c.height == l.height {
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

// Большое правое вращение
func bigRightRotation(a *TNode) *TNode {
	// Задаем обозначения
	b := a.left
	c := b.right
	m := c.left
	n := c.right

	// Перевыставляем вершины
	a.left = n
	b.right = m
	c.left = b
	c.right = a

	// Корректируем высоты
	a.height -= 2
	b.height--
	c.height++

	return c
}
