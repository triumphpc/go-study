package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

// 3-SUM. Эффективный способ обхода

func main() {
	task(os.Stdin, os.Stdout)
}

func task(src io.Reader, dst io.Writer) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	numbers := make([]int, 0, 10)
	result := make([][]int, 0, 5)

	for i := 0; i < n; i++ {
		scanner.Scan()
		intVal, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, intVal)
	}

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	history := make(map[int]bool, len(numbers))

	sort.Ints(numbers) // 1. Отсортируем
	// 1 1 2 3 4 4 5 8
	// 10 - 1 -2 = 7
	// 10 - 2 - 3 = 5

	for i1 := 0; i1 < n; i1++ { // Шагаем вперед
		for j := i1 + 1; j < n && len(history) > 0; j++ {
			target := k - numbers[i1] - numbers[j] // Получаем разницу между целевой и двумя правыми

			if history[target] { // Далее смотрим, был ли слева целевой - если да - то есть в наборе 3
				result = append(result, []int{target, numbers[i1], numbers[j]})
			}
		}
		history[numbers[i1]] = true
	}

	log.Println(result)
}
