// Куча. Приоритетная очередь. Добавление и удаление элементов

package main

import "fmt"

type HeapS []int

func main() {
	h := make(HeapS, 0, 10)
	h = append(h, 50, 45, 20, 17, 14, 10, 6, 5, 4)

	heapAdd(h, 33)

	fmt.Println(h)

}

// Добавление элементов
func heapAdd(heapS HeapS, key int) {
	index := len(heapS) // 1. Определяем индекс для добавления элемента

	heapS = append(heapS, key)

	// 2. Просеивание вверх, элемент должен быть отсортирован в дереве
	siftUp(heapS, index)

}

func siftUp(heap HeapS, index int) {
	if index == 1 {
		return
	}

	// 3. Вычисляем индекс родителя
	parentIndex := index / 2 // Целочисленное деление

	fmt.Println(heap, parentIndex, index)
	// Меняем местами если значение родителя больше
	if heap[parentIndex] < heap[index] {
		heap[parentIndex], heap[index] = heap[index], heap[parentIndex]
	}

	siftUp(heap, parentIndex)
}

// Удаление максимального элемента
func popMax(heap HeapS) int {
	result := heap[0]
	heap[0] = heap[len(heap)]
	heap = heap[:len(heap)-1]

	siftDown(heap, 1)

	return result
}

func siftDown(heap HeapS, index int) {
	left := 2 * index
	right := 2*index + 1

	if len(heap) < left {
		return
	}

	indexL := 0
	if right <= len(heap) {
		indexL = right
	} else {
		indexL = left
	}

	if heap[index] < heap[indexL] {
		heap[index], heap[indexL] = heap[indexL], heap[index]
	}

	siftDown(heap, indexL)
}
