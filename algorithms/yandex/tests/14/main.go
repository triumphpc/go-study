package main

import "fmt"

// Нелюбимое дело
// Вася размышляет, что ему можно не делать из того списка дел, который он составил.
// Но, кажется, все пункты очень важные! Вася решает загадать число и удалить дело, которое идёт под этим номером.
// Список дел представлен в виде односвязного списка. Напишите функцию solution, которая принимает на вход голову списка
// и номер удаляемого дела и возвращает голову обновлённого списка.
// Внимание: в этой задаче не нужно считывать входные данные. Нужно написать только функцию, которая принимает
// на вход голову списка и номер удаляемого элемента и возвращает голову обновлённого списка. Ниже дано описание структуры,
// которая задаёт вершину списка.
// Мы рекомендуем воспользоваться заготовками кода для данной задачи, расположенными по ссылке.
// Решение надо отправлять только в виде файла с расширением, которое соответствует вашему языку.
// Иначе даже корректно написанное решение не пройдет тесты.

// go test -run none -bench . -benchmem -benchtime 3s                                                                        triumphpc@MacBook-Pro-triumphpc
//goos: darwin
//goarch: amd64
//pkg: github.com/triumphpc/go-study/algorithms/yandex/tests/14
//cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
//BenchmarkBasic-16     	1000000000	         1.769 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBasic2-16    	1000000000	         0.9520 ns/op	       0 B/op	       0 allocs/op

func main() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}

	current := Solution(&node0, 3)

	for {
		fmt.Println(current.data)

		if current.next == nil {
			break
		}

		current = current.next
	}

}

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		return head.next
	}

	var currentIdx int
	current := head
	prev := head

	for {
		if current.next == nil {
			if currentIdx == idx {
				prev.next = nil
			}

			return head
		}

		if currentIdx == idx {
			prev.next = current.next

			return head
		} else {
			currentIdx++
			prev = current
			current = current.next
		}
	}
}

func Solution2(head *ListNode, index int) *ListNode {
	if index == 0 {
		return head.next
	}
	pos := 1
	prev := head
	for {
		cur := prev.next
		if pos == index {
			prev.next = cur.next
			return head
		}
		prev = cur
		pos++
		if prev.next == nil {
			return head
		}
	}
}
