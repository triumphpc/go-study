package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Рекурсивные числа фибоначи
// У Тимофея было n (0≤n≤32) стажёров. Каждый стажёр хотел быть лучше своих предшественников, поэтому i-й стажёр
// делал столько коммитов, сколько делали два предыдущих стажёра в сумме. Два первых стажёра были менее инициативными —– они сделали по одному коммиту.
// Пусть Fi —– число коммитов, сделанных i-м стажёром (стажёры нумеруются с нуля). Тогда выполняется следующее: F0=F1=1. Для всех i≥2  выполнено Fi=Fi−1+Fi−2.
// Определите, сколько кода напишет следующий стажёр –— найдите Fn.
// Решение должно быть реализовано рекурсивно.

func main() {
	task(os.Stdin)

}

var queue []int
var num int
var idx int

func task(src io.Reader) {
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		//writer.Flush()
	}()

	scanner := bufio.NewScanner(src)

	// num of commands
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())

	if num < 3 {
		writer.WriteString("1")

		return
	}

	queue = make([]int, num)
	queue[0] = 1
	queue[1] = 1
	idx = 2

	req()
}

func req() {
	queue[idx] = queue[idx-1] + queue[idx-2]
	if num == idx+1 {
		//fmt.Println(queue[idx])

		return
	}

	idx++

	req()
}

func fi(n int) int { /// Плохое решение, ростет стек вызова
	if n == 0 || n == 1 {
		return 1
	}
	return fi(n-1) + fi(n-2)
}

func task2(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(dst)
	defer func() {
		//writer.Flush()
	}()
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())

	//if num < 3 {
	//	return
	//}

	writer.WriteString(strconv.Itoa(fi(num)) + "\n")
}

func fi3(n *int) int { /// Плохое решение, ростет стек вызова
	if *n == 0 || *n == 1 {
		return 1
	}
	return fi(*n-1) + fi(*n-2)
}

func task3(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(dst)
	defer func() {
		//writer.Flush()
	}()
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())

	//if num < 3 {
	//	return
	//}

	writer.WriteString(strconv.Itoa(fi3(&num)) + "\n")
}
