// Куча. Приоритетная очередь.
// Напишите функцию, совершающую просеивание вниз в куче на максимум. Гарантируется, что порядок элементов в куче может быть нарушен только элементом, от которого запускается просеивание.
//Функция принимает в качестве аргументов массив, в котором хранятся элементы кучи, и индекс элемента, от которого надо сделать просеивание вниз. Функция должна вернуть индекс, на котором элемент оказался после просеивания. Также необходимо изменить порядок элементов в переданном в функцию массиве.
//Индексация в массиве, содержащем кучу, начинается с единицы. Таким образом, сыновья вершины на позиции v это 2v и 2v+1. Обратите внимание, что нулевой элемент в передаваемом массиве фиктивный, вершина кучи соответствует 1-му элементу.

package main

import "fmt"

func main() {
	h := make([]int, 0, 6)
	h = append(h, 12, 1, 8, 3, 4, 7)

	kk := siftDown(h, 1)

	fmt.Println(h, kk)

}

// Просеивание вниз
func siftDown(heap []int, idx int) int {
	left := 2 * idx
	right := 2*idx + 1

	if len(heap)-1 < left {
		return idx
	}
	largestIdx := left
	if right <= len(heap)-1 && heap[left] < heap[right] {
		largestIdx = right
	}
	if heap[idx] < heap[largestIdx] {
		heap[idx], heap[largestIdx] = heap[largestIdx], heap[idx]
		return siftDown(heap, largestIdx)
	}
	return idx
}

// Просеивание вверх
func siftUp(heap []int, idx int) int {
	if idx == 1 {
		return 1
	}
	parentIdx := idx / 2
	if heap[idx] > heap[parentIdx] {
		heap[idx], heap[parentIdx] = heap[parentIdx], heap[idx]
		return siftUp(heap, parentIdx)
	}
	return idx
}
