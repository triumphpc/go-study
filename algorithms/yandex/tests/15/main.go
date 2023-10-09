package main

import (
	"fmt"
	"strings"
)

// Заботливая мама
// Мама Васи хочет знать, что сын планирует делать и когда. Помогите ей: напишите функцию
// solution, определяющую индекс первого вхождения передаваемого ей на вход значения
//в связном списке, если значение присутствует.

// On сложность

// go test -run none -bench . -benchmem -benchtime 3s                                                                        triumphpc@MacBook-Pro-triumphpc
// BenchmarkBasic-16     	75783692	        47.84 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBasic2-16    	322531226	        11.26 ns/op	       0 B/op	       0 allocs/op
func main() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}

	idx := Solution(&node0, "node0")

	fmt.Println(idx)

}

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, val string) int {
	if strings.EqualFold(head.data, val) {
		return 0
	}

	idx := 1
	cur := head.next

	for {
		if cur == nil {
			return -1
		}

		if strings.EqualFold(cur.data, val) {
			return idx
		}

		idx++
		cur = cur.next
	}
}

func Solution2(head *ListNode, value string) int {
	pos := 0
	for {
		if head.data == value { // Пошагово проходим по элементам
			return pos
		}
		pos++
		if head.next == nil {
			break
		}
		head = head.next // смешаем шаг
	}
	return -1
}
