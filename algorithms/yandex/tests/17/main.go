package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// StackMaxEffective
// Реализуйте класс StackMaxEffective, поддерживающий операцию определения максимума среди элементов в стеке.
// Сложность операции должна быть O(1). Для пустого стека операция должна возвращать None. При этом push(x) и pop()
// также должны выполняться за константное время.

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
	max := 0

	// run commands
	for i := 0; i < num; i++ {
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "get_max":
			if max > 0 {
				l := strconv.Itoa(max) // Уже реализовывает O(1) сложность
				_, _ = writer.WriteString(l + "\n")
			}
		case "pop":
			if len(stack) == 0 {
				_, _ = writer.WriteString("error\n")
				continue
			}

			l := strconv.Itoa(stack[len(stack)-1])

			if max == stack[len(stack)-1] {
				sortStack := stack[:len(stack)-1]
				sort.Slice(sortStack, func(i, j int) bool {
					return i > j
				})
				max = sortStack[len(sortStack)-1]
			}

			_, _ = writer.WriteString(l + "\n")
			stack = stack[:len(stack)-1]

		default:
			words := strings.Split(line, " ")
			n, _ := strconv.Atoi(words[1])
			stack = append(stack, n)
			if max < n {
				max = n
			}
		}
	}

	//writer.Flush()
}
