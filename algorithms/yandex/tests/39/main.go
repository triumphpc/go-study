package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Большое число
// Даны числа. Нужно определить, какое самое большое число можно из них составить.
// В первой строке записано n — количество чисел. Оно не превосходит 100.
//
// Во второй строке через пробел записаны n неотрицательных чисел, каждое из которых не превосходит 1000.

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
	parts := make([]int, 0, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		part, _ := strconv.Atoi(scanner.Text())
		parts = append(parts, part)
	}

	parts = mergeSortDesc(parts)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(parts))
}

func mergeSortDesc(arr []int) []int { // Сортировка слиянием
	if len(arr) <= 1 {
		return arr
	}
	left := mergeSortDesc(arr[:len(arr)/2])
	right := mergeSortDesc(arr[len(arr)/2:])
	result := make([]int, 0, len(arr))

	l, r := 0, 0
	for l < len(left) && r < len(right) { // Сортировка O(n) переданных частей массива
		lParts := strconv.Itoa(left[l])
		rParts := strconv.Itoa(right[r])

		if lParts[0] > rParts[0] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	for l < len(left) { // Тут добавляем оставшиеся части
		result = append(result, left[l])
		l++
	}
	for r < len(right) {
		result = append(result, right[r])
		r++
	}
	return result
}

func task2(src io.Reader, dst io.Writer) { // Оптимальное решение
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")
	for i := 1; i < total; i++ {
		itemToInsert := values[i]                      // 1. Берем текущий элемент итерации
		j := i                                         // 2. Запоминаем для шага назад
		for j > 0 && less(values[j-1], itemToInsert) { // Пузырьковая сортировка - меняем местами каждый раз
			values[j] = values[j-1] // Смещаем значение
			j--
		}
		values[j] = itemToInsert // Проставляем текущее
	}
	fmt.Fprintf(dst, "%s\n", strings.Join(values, ""))
}

func less(a string, b string) bool {
	return a+b < b+a
}
