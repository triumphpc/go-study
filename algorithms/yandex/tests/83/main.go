package main

// Вам даны два непустых связанных списка, представляющих два неотрицательных целых числа. Цифры хранятся в обратном порядке,
//и каждый из их узлов содержит одну цифру. Сложите два числа и верните сумму в виде связанного списка.
//Вы можете предположить, что эти два числа не содержат ведущих нулей, кроме самого числа 0.
// https://leetcode.com/problems/add-two-numbers/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())

	var l1 *ListNode

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		part, _ := strconv.Atoi(scanner.Text())

		// Создаем новый узел
		temp1 := &ListNode{part, nil}

		if l1 == nil {
			l1 = temp1
		} else {
			// Ищем конец списка и вставляем в конец
			temp2 := l1
			for temp2.Next != nil {
				temp2 = temp2.Next
			}
			temp2.Next = temp1
		}
	}

	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next

	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return nil

}
