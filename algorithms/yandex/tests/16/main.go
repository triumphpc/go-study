package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Стек MAX
// В первой строке записано одно число n — количество команд, которое не превосходит 10000. В следующих n строках идут команды. Команды могут быть следующих видов:
//* push(x) — добавить число x в стек;
//* pop() — удалить число с вершины стека;
//* get_max() — напечатать максимальное число в стеке;
// Если стек пуст, при вызове команды get_max() нужно напечатать «None», для команды pop() — «error».
// Формат вывода
//
// Для каждой команды get_max() напечатайте результат её выполнения. Если стек пустой, для команды get_max() напечатайте «None».
// Если происходит удаление из пустого стека — напечатайте «error».

// go test -run none -bench . -benchmem -benchtime 3s                                1 ↵ triumphpc@MacBook-Pro-triumphpc
//BenchmarkBasic-16     	 2776242	      1350 ns/op	    8240 B/op	       3 allocs/op
//BenchmarkBasic2-16    	 2613036	      1385 ns/op	    8240 B/op	       3 allocs/op

func main() {
	task(os.Stdin)

}

func task(src io.Reader) {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(src)

	// num of commands
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	stack := make([]int, 0, num)
	max := 0 //

	// run commands
	for i := 0; i < num; i++ { // 1. Читаем по очереди команды
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "get_max": // Возвращаем максимальное
			if max > 0 {
				l := strconv.Itoa(max)
				_, _ = writer.WriteString(l + "\n")
			}
		case "pop": // удаляем число из стека
			if len(stack) == 0 {
				_, _ = writer.WriteString("error\n")
				continue
			}

			l := strconv.Itoa(stack[len(stack)-1])

			if max == stack[len(stack)-1] { // Тут определяем максимальное число в стеке. Но можно и хранить предыдыщий шаг
				sortStack := stack[:len(stack)-1]
				sort.Slice(sortStack, func(i, j int) bool {
					return i > j
				})
				max = sortStack[len(sortStack)-1]
			}

			_, _ = writer.WriteString(l + "\n")
			stack = stack[:len(stack)-1]

		default: // добавляем в стек
			words := strings.Split(line, " ")
			n, _ := strconv.Atoi(words[1])
			stack = append(stack, n)
			if max < n { // обновляем максимальное
				max = n
			}
		}
	}

	//writer.Flush()
}

func task2(src io.Reader) {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(src)

	// num of commands
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	stack := new(stack)

	// run commands
	for i := 0; i < num; i++ {
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "get_max":
			_, _ = writer.WriteString(stack.getMax() + "\n")
		case "pop":
			_ = stack.pop()

		default:
			words := strings.Split(line, " ")
			n, _ := strconv.Atoi(words[1])

			stack.push(n)
		}
	}

	//writer.Flush()
}

type stack struct {
	values []int
}

func (s *stack) pop() error {
	if len(s.values) == 0 {
		return errors.New("error")
	}
	s.values = s.values[:len(s.values)-1]
	return nil
}

func (s *stack) push(val int) {
	s.values = append(s.values, val)
}
func (s *stack) getMax() string {
	if len(s.values) == 0 {
		return "None"
	}
	max := s.values[0]
	for i := 1; i < len(s.values); i++ {
		if s.values[i] > max {
			max = s.values[i]
		}
	}

	return strconv.Itoa(max)
}
