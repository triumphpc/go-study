package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

// 3-SUM
// Дан набор чисел, вывести набор суммы целевого числа
// A=10,x=[5,2,8,1,1,3,4,4]. Следующие тройки чисел дадут 10 в сумме:
// (2,3,5)
// (1,4,5)
// (2,4,4)
// (1,1,8)

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

	// Обход каждого элемента
	for i1 := 0; i1 < len(numbers); i1++ {
		for i2 := 1; i2 < len(numbers); i2++ {
			for i3 := 2; i3 < len(numbers); i3++ {
				if numbers[i1]+numbers[i2]+numbers[i3] == k {
					result = append(result, []int{numbers[i1], numbers[i2], numbers[i3]})
					// При сохранении в папу удалим дубли
					// Так же нужно сортировать при вставке
				}
			}
		}
	}

	log.Println(result)
}
