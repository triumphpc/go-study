package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// Ограниченная очередь
// Астрологи объявили день очередей ограниченного размера. Тимофею нужно написать класс MyQueueSized, который принимает параметр max_size,
// означающий максимально допустимое количество элементов в очереди.
// Помогите ему —– реализуйте программу, которая будет эмулировать работу такой очереди.
// Функции, которые надо поддержать, описаны в формате ввода.

func main() {
	task(os.Stdin)

}

func task(src io.Reader) {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(src)

	// num of commands
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	max, _ := strconv.Atoi(scanner.Text())
	queue := new(queue)
	queue.values = make([]int, 0, max)

	// run commands
	for i := 0; i < num; i++ {
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "peek":
			_, _ = writer.WriteString(queue.peek() + "\n")
		case "size":
			_, _ = writer.WriteString(strconv.Itoa(len(queue.values)) + "\n")
		case "pop":
			_, _ = writer.WriteString(queue.pop() + "\n")

		default:
			if len(queue.values) == max {
				_, _ = writer.WriteString("error\n")
				continue
			}

			words := strings.Split(line, " ")
			n, _ := strconv.Atoi(words[1])

			queue.push(n)
		}
	}

	writer.Flush()
}

type queue struct {
	values []int
	head   int
	max    int
}

func (s *queue) pop() string {
	switch len(s.values) {
	case 0:
		return "None"
	case 1:
		v := s.values[s.head]
		s.values = []int{}
		s.head = 0

		return strconv.Itoa(v)
	default:
		v := s.values[s.head]
		s.values = s.values[1:]
		s.head++

		return strconv.Itoa(v)
	}
}

func (s *queue) push(val int) {
	s.values = append(s.values, val)
}

func (s *queue) peek() string {
	if len(s.values) == 0 {
		return "None"
	}

	return strconv.Itoa(s.values[s.head])
}

type roundQueue struct {
	Head     int
	Tail     int
	elements []int
	Size     int
}

func (rq *roundQueue) pop() string {
	if rq.Size == 0 {
		return "None"
	}
	val := strconv.Itoa(rq.elements[rq.Head])
	rq.Head = (rq.Head + 1) % cap(rq.elements)
	rq.Size--
	return val
}

func (rq *roundQueue) push(val int) error {
	if rq.Size == cap(rq.elements) {
		return errors.New("error")
	}
	rq.elements[rq.Tail] = val
	rq.Tail = (rq.Tail + 1) % cap(rq.elements)
	rq.Size++
	return nil
}
func (rq *roundQueue) peek() string {
	if rq.Size == 0 {
		return "None"
	}
	return strconv.Itoa(rq.elements[rq.Head])
}
func (rq *roundQueue) size() string {
	return strconv.Itoa(rq.Size)
}
func NewRoundQueue(maxSize int) *roundQueue {
	return &roundQueue{
		elements: make([]int, maxSize, maxSize),
	}
}

func task2(src io.Reader, dst io.Writer) {
	writer := bufio.NewWriter(dst)
	defer func() {
		writer.Flush()
	}()
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	commandsTotal, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	qSize, _ := strconv.Atoi(scanner.Text())
	rq := NewRoundQueue(qSize)
	for commandsTotal > 0 {
		scanner.Scan()
		commandParts := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		switch commandParts[0] {
		case "push":
			value, _ := strconv.Atoi(commandParts[1])
			if err := rq.push(value); err != nil {
				writer.WriteString(err.Error() + "\n")
			}
		case "pop":
			writer.WriteString(rq.pop() + "\n")
		case "peek":
			writer.WriteString(rq.peek() + "\n")
		case "size":
			writer.WriteString(rq.size() + "\n")
		}
		commandsTotal--
	}
}
