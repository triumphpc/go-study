package main

// Список дел
// Васе нужно распечатать свой список дел на сегодня. Помогите ему: напишите функцию, которая печатает все его дела.
// Известно, что дел у Васи не больше 5000. Внимание: в этой задаче не нужно считывать входные данные.
// Нужно написать только функцию, которая принимает на вход голову списка и печатает его элементы.
// Ниже дано описание структуры, которая задаёт узел списка. Используйте компилятор Make.

// go test -run none -bench . -benchmem -benchtime 3s
// BenchmarkBasic-16     	533758747	         7.059 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBasic2-16    	1000000000	         3.062 ns/op	       0 B/op	       0 allocs/op

func main() {
	test()
}

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode) {
	if head == nil {
		return
	}

	//fmt.Println(head.data)

	Solution(head.next)
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}

	Solution(&node0)
	/*
	   Output is:
	   node0
	   node1
	   node2
	   node3
	*/
}

func Solution2(head *ListNode) {
	//writer := bufio.NewWriter(os.Stdout)
	//defer writer.Flush()
	for {
		//writer.WriteString(head.data + "\n")
		if head.next == nil {
			break
		}
		head = head.next
	}
}
