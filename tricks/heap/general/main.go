package main

// Пакет heap обеспечивает операции кучи (heap) для любого типа, который реализует heap.Interface.
//Куча - это дерево со свойством того, что каждый узел является узлом с минимальным значением в своем поддереве.
//
//Минимальным элементом в дереве является корень с индексом 0.
import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))

	}
}
