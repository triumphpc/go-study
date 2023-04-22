package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// Списочная очередь
// Любимый вариант очереди Тимофея — очередь, написанная с использованием связного списка. Помогите ему с реализацией. Очередь должна поддерживать выполнение трёх команд:
//* get() — вывести элемент, находящийся в голове очереди, и удалить его. Если очередь пуста, то вывести «error».
//* put(x) — добавить число x в очередь
//* size() — вывести текущий размер очереди

func main() {
	task(os.Stdin)

}

func task(src io.Reader) {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(src)

	// num of commands
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())

	queue := new(queue)
	queue.values = make([]int, 0, 10)

	// run commands
	for i := 0; i < num; i++ {
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "size":
			_, _ = writer.WriteString(strconv.Itoa(len(queue.values)) + "\n")
		case "get":
			_, _ = writer.WriteString(queue.pop() + "\n")

		default:
			words := strings.Split(line, " ")
			n, _ := strconv.Atoi(words[1])

			queue.push(n)
		}
	}

	writer.Flush()
}

type queue struct {
	values []int
}

func (s *queue) pop() string {
	switch len(s.values) {
	case 0:
		return "None"
	case 1:
		v := s.values[0]
		s.values = []int{}

		return strconv.Itoa(v)
	default:
		v := s.values[0]
		s.values = s.values[1:]

		return strconv.Itoa(v)
	}
}

func (s *queue) push(val int) {
	s.values = append(s.values, val)
}

type lqItem struct {
	value string
	next  *lqItem
}
type linkedQueue struct {
	Head *lqItem
	Tail *lqItem
	Size int
}

func (lq *linkedQueue) get() string {
	if lq.Head == nil {
		return "error"
	}
	val := lq.Head.value
	lq.Head = lq.Head.next
	lq.Size--
	return val
}
func (lq *linkedQueue) put(val string) {
	if lq.Size == 0 {
		lq.Head = &lqItem{val, nil}
		lq.Tail = lq.Head
		lq.Size++
		return
	}
	lq.Tail.next = &lqItem{val, nil}
	lq.Tail = lq.Tail.next
	lq.Size++
}
func (lq *linkedQueue) size() string {
	return strconv.Itoa(lq.Size)
}
func task2(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(dst)
	defer func() {
		writer.Flush()
	}()
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	commandsTotal, _ := strconv.Atoi(scanner.Text())
	lq := &linkedQueue{}
	for commandsTotal > 0 {
		scanner.Scan()
		commandParts := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		switch commandParts[0] {
		case "put":
			lq.put(commandParts[1])
		case "get":
			writer.WriteString(lq.get() + "\n")
		case "size":
			writer.WriteString(lq.size() + "\n")
		}
		commandsTotal--
	}
}
