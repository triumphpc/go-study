// Пирамидальная сортировка приоритетной очереди (кучи)
// Псевдокод
// https://habr.com/ru/companies/otus/articles/460087/
//
//  Двоичная куча — это законченное двоичное дерево, в котором элементы хранятся в особом порядке:
// значение в родительском узле больше (или меньше) значений в его двух дочерних узлах. Первый вариант называется max-heap,
// а второй — min-heap. Куча может быть представлена двоичным деревом или массивом.

// Алгоритм пирамидальной сортировки в порядке по возрастанию:
//
//
// Постройте max-heap из входных данных.
// На данном этапе самый большой элемент хранится в корне кучи. Замените его на последний элемент кучи, а затем уменьшите ее размер на 1. Наконец, преобразуйте полученное дерево в max-heap с новым корнем.
// Повторяйте вышеуказанные шаги, пока размер кучи больше 1.

package main

import "fmt"

func main() {
	// Массив, который требуется отсортировать
	arr := []int{12, 11, 13, 5, 6, 7}

	heapSort(arr)

	fmt.Println(arr)

}

func heapSort(arr []int) {
	// 1. Построение кучи из массива

	// Возьмем половину слайса
	for i := len(arr)/2 - 1; i >= 0; i-- {
		siftDown(arr, len(arr), i)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		siftDown(arr, i, 0)

	}

}

// Просеивание вниз
func siftDown(heap []int, n, idx int) {
	left := 2 * idx    // Левая
	right := 2*idx + 1 // Правый
	largestIdx := idx  // Наибольший элемент, как корень

	// Если левый дочерний элемент больше корня
	if left < n-1 && heap[left] > heap[largestIdx] {
		largestIdx = left
	}

	// Если правый дочерний элемент больше чем самый большой элмент на данный момент
	if right < n-1 && heap[left] < heap[right] {
		largestIdx = right
	}
	// Если это не корень
	if largestIdx != idx {
		heap[idx], heap[largestIdx] = heap[largestIdx], heap[idx]
		// Преобразуемв двоичную кучу затронутое дерево
		siftDown(heap, n, largestIdx)
	}
}