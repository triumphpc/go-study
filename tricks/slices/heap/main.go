// Обход слайса
package main

import (
	"container/heap"
	"fmt"
)

type ItHeap []int

func (h ItHeap) Len() int {
	return len(h)
}
func (h ItHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h ItHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ItHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *ItHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {

	h := &ItHeap{2, 1, 5}
	heap.Init(h)

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

}
