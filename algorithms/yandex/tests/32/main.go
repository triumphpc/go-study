package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// QuickSort быстрая сортировка

func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	size, _ := strconv.Atoi(scanner.Text())

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	digits := make([]int, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		digits[idx] = num
	}

	result := quickSort(digits)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(result))
}

func quickSort(array []int) []int {
	if len(array) < 2 { //  базовый случай, return array # массивы с 0 или 1 элементами фактически отсортированы
		return array
	}

	pivot := array[len(array)/2] // случайны элемент
	left, center, right := partition(array, pivot)

	result := quickSort(left)
	result = append(result, center...)
	result = append(result, quickSort(right)...)

	return result
}
func partition(array []int, pivot int) (left, center, right []int) {
	for idx := range array {
		if array[idx] > pivot {
			right = append(right, array[idx])
		} else if array[idx] < pivot {
			left = append(left, array[idx])
		} else {
			center = append(center, array[idx])
		}
	}

	return left, center, right

}
